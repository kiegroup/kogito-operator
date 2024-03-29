// Copyright 2020 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package framework

import (
	"regexp"
	"strings"

	"github.com/kiegroup/kogito-operator/apis"
)

const (
	// dockerTagRegx matches a docker image name, to test it check: https://regex101.com/r/lAJKau/2.
	// this is a super relax regexp, since we accept pretty much anything see the test cases on image_test.go
	// see: https://github.com/docker/distribution/blob/main/reference/regexp.go
	dockerTagRegx = `(?P<domain>(?:[a-z0-9]?:{0,1}\.?-?_?)+\/(?:(?:[a-z0-9]|[._]|__|[-]*)+\/)?)?(?P<image>[^:]+)(?P<tag>:.+)?`
)

var (
	// DockerTagRegxCompiled is the compiled regex to verify docker tag names
	DockerTagRegxCompiled = *regexp.MustCompile(dockerTagRegx)
)

// ConvertImageTagToImage converts a plain string into an Image structure. For example, see https://regex101.com/r/1YX9rh/1.
func ConvertImageTagToImage(imageName string) api.Image {
	domain, name, tag := SplitImageTag(imageName)
	image := api.Image{
		Domain: domain,
		Name:   name,
		Tag:    tag,
	}

	return image
}

// ConvertImageToImageTag converts an Image into a plain string (domain/namespace/name:tag).
func ConvertImageToImageTag(image api.Image) string {
	imageTag := ""
	if len(image.Domain) > 0 {
		imageTag += image.Domain + "/"
	}
	imageTag += image.Name
	if len(image.Tag) > 0 {
		imageTag += ":" + image.Tag
	}
	return imageTag
}

// splitImageTag
func splitImageTag(imageTag string) (domain, name, tag string) {
	domain = ""
	name = ""
	tag = ""
	if len(imageTag) > 0 {
		if strings.HasPrefix(imageTag, ":") {
			tag = strings.Split(imageTag, ":")[1]
			return
		}

		imageMatch := DockerTagRegxCompiled.FindStringSubmatch(imageTag)
		if len(imageMatch[1]) > 0 {
			domain = strings.TrimRight(imageMatch[1], "/")
		}
		name = imageMatch[2]
		tag = strings.ReplaceAll(imageMatch[3], ":", "")
	}
	return
}

// SplitImageTag breaks into parts a given tag name, adds "latest" to the tag name if it's empty. For example, see https://regex101.com/r/1YX9rh/1.
func SplitImageTag(imageTag string) (domain, name, tag string) {
	if len(imageTag) == 0 {
		return
	}
	domain, name, tag = splitImageTag(imageTag)
	if len(tag) == 0 {
		tag = "latest"
	}
	return
}

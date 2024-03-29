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

package mappers

import (
	"fmt"

	"github.com/cucumber/messages-go/v16"
)

// TableRow represents a row of godog.Table made to a step definition
type TableRow = messages.PickleTableRow

const (
	enabledKey  = "enabled"
	disabledKey = "disabled"
)

// GetFirstColumn returns first table row column
func GetFirstColumn(row *TableRow) string {
	return row.Cells[0].Value
}

// GetSecondColumn returns second table row column
func GetSecondColumn(row *TableRow) string {
	return row.Cells[1].Value
}

// GetThirdColumn returns third table row column
func GetThirdColumn(row *TableRow) string {
	return row.Cells[2].Value
}

// MustParseEnabledDisabled parse a boolean string value
func MustParseEnabledDisabled(value string) bool {
	switch value {
	case enabledKey:
		return true
	case disabledKey:
		return false
	default:
		panic(fmt.Errorf("Unknown value for enabled/disabled: %s", value))
	}
}

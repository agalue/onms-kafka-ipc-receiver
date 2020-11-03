// @author Alejandro Galue <agalue@opennms.org>

package client

import (
	"fmt"
	"strings"
)

// EnumValue a enumaration array of strings
type EnumValue struct {
	Enum     []string
	Default  string
	selected string
}

// Set sets a value of the enum
func (e *EnumValue) Set(value string) error {
	for _, enum := range e.Enum {
		if enum == value {
			e.selected = value
			return nil
		}
	}
	return fmt.Errorf("allowed values are %s", strings.Join(e.Enum, ", "))
}

// String gets the value of the enum as string
func (e EnumValue) String() string {
	if e.selected == "" {
		return e.Default
	}
	return e.selected
}

// EnumAsString gets a CSV with all the values on the enum
func (e EnumValue) EnumAsString() string {
	return strings.Join(e.Enum, ", ")
}

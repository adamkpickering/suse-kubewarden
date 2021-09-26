package main

import (
	"testing"

	"github.com/kubewarden/gjson"
)

func TestParseSettingsWithAValueProvided(t *testing.T) {
	request := `
	{
		"request": "doesn't matter here",
		"settings": {
			"denied_names": [ "foo", "bar" ]
		}
	}
	`
	rawRequest := []byte(request)

	output, err := validateSettings(rawRequest)
	if err != nil {
		t.Errorf("Unexpected error: %w", err)
	}

	output_data := gjson.GetBytes(output, "valid")
	if valid := output_data.Bool(); valid {
		t.Errorf("Got %t, not false, in 'valid' key of output", valid)
	}
}

func TestParseSettingsWithNoValueProvided(t *testing.T) {
	request := "{}"
	rawRequest := []byte(request)

	output, err := validateSettings(rawRequest)
	if err != nil {
		t.Errorf("Unexpected error: %w", err)
	}

	output_data := gjson.GetBytes(output, "valid")
	if valid := output_data.Bool(); !valid {
		t.Errorf("Got %t, not true, in 'valid' key of output", valid)
	}
}

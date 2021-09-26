package main

import (
	"fmt"
	"strings"

	"github.com/kubewarden/gjson"
	kubewarden "github.com/kubewarden/policy-sdk-go"
)

func validate(payload []byte) ([]byte, error) {
	// ensure we got valid json
	if !gjson.ValidBytes(payload) {
		return kubewarden.RejectRequest(
			kubewarden.Message("Not a valid JSON document"),
			kubewarden.Code(400),
		)
	}

	// get labels as map
	data := gjson.GetBytes(payload, "request.object.metadata.labels")
	if !data.Exists() {
		logger.Warn("cannot read labels from metadata: accepting request")
		return kubewarden.AcceptRequest()
	}
	label_mappings := data.Map()

	// check label keys
	for key, _ := range label_mappings {
		if isPalindrome(key) {
			return kubewarden.RejectRequest(
				kubewarden.Message(
					fmt.Sprintf("'%s' is a palindrome, which is not allowed", key),
				),
				kubewarden.Code(400),
			)
		}
	}

	return kubewarden.AcceptRequest()
}


func isPalindrome(input string) bool {
	input_length := len([]rune(input))
	iterations := input_length / 2
	lowered_input := strings.ToLower(input)
	lowered_runes := []rune(lowered_input)
	for i := 0; i < iterations; i++ {
		if lowered_runes[i] != lowered_runes[input_length - (1 + i)] {
			return false
		}
	}
	return true
}

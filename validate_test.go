package main

import (
	"encoding/json"
	"testing"

	kubewarden_testing "github.com/kubewarden/policy-sdk-go/testing"
)

func TestApproval(t *testing.T) {
	settings := struct{}{}

	payload, err := kubewarden_testing.BuildValidationRequest(
		"test_data/ingress.json",
		&settings)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	var response kubewarden_testing.ValidationResponse
	if err := json.Unmarshal(responsePayload, &response); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if response.Accepted != true {
		t.Error("Unexpected rejection")
	}
}

func TestRejection(t *testing.T) {
	settings := struct{}{}

	payload, err := kubewarden_testing.BuildValidationRequest(
		"test_data/palindrome.json",
		&settings)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	var response kubewarden_testing.ValidationResponse
	if err := json.Unmarshal(responsePayload, &response); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if response.Accepted != false {
		t.Error("Unexpected approval")
	}

	expected_message := "'radar' is a palindrome, which is not allowed"
	if response.Message != expected_message {
		t.Errorf("Got '%s' instead of '%s'", response.Message, expected_message)
	}
}

type palindromeTestCase struct {
	Value string
	ShouldReturn bool
}

func TestIsPalindrome(t *testing.T) {
	test_cases := []palindromeTestCase{
		palindromeTestCase{Value: "racecar", ShouldReturn: true},
		palindromeTestCase{Value: "racecat", ShouldReturn: false},
		palindromeTestCase{Value: "racebar", ShouldReturn: false},
		palindromeTestCase{Value: "pacecar", ShouldReturn: false},
		palindromeTestCase{Value: "raccar", ShouldReturn: true},
		palindromeTestCase{Value: "racear", ShouldReturn: false},
		palindromeTestCase{Value: "naccar", ShouldReturn: false},
		palindromeTestCase{Value: "asdf", ShouldReturn: false},
		palindromeTestCase{Value: "Racecar", ShouldReturn: true},
		palindromeTestCase{Value: "1racecar1", ShouldReturn: true},
		palindromeTestCase{Value: "rac#e#car", ShouldReturn: true},
		palindromeTestCase{Value: "rαcecαr", ShouldReturn: true},
		palindromeTestCase{Value: "", ShouldReturn: true},
		palindromeTestCase{Value: "x", ShouldReturn: true},
		palindromeTestCase{Value: "xx", ShouldReturn: true},
	}

	for _, test_case := range test_cases {
		result := isPalindrome(test_case.Value)
		if test_case.ShouldReturn != result {
			t.Errorf("%s returned %t instead of %t", test_case.Value, result, test_case.ShouldReturn)
		}
	}
}

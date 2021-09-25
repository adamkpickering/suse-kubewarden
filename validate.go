package main

import (
	//"fmt"
	"strings"

	//onelog "github.com/francoispqt/onelog"
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

	// build Settings struct from payload
	//settings, err := NewSettingsFromValidationReq(payload)
	//if err != nil {
	//	return kubewarden.RejectRequest(
	//		kubewarden.Message(err.Error()),
	//		kubewarden.Code(400))
	//}

	//data := gjson.GetBytes(
	//	payload,
	//	"request.object.metadata.labels")

	//if !data.Exists() {
	//	onelog.Warn("cannot read labels from metadata: accepting request")
	//	return kubewarden.AcceptRequest()
	//}
	//name := data.String()

	//onelog.DebugWithFields("validating ingress object", func(e onelog.Entry) {
	//	namespace := gjson.GetBytes(payload, "request.object.metadata.namespace").String()
	//	e.String("name", name)
	//	e.String("namespace", namespace)
	//})

	//if settings.DeniedNames.Contains(name) {
	//	onelog.InfoWithFields("rejecting ingress object", func(e onelog.Entry) {
	//		e.String("name", name)
	//		e.String("denied_names", settings.DeniedNames.String())
	//	})

	//	return kubewarden.RejectRequest(
	//		kubewarden.Message(
	//			fmt.Sprintf("The '%s' name is on the deny list", name)),
	//		kubewarden.NoCode)
	//}

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

package main

import (
	"github.com/kubewarden/gjson"
	kubewarden "github.com/kubewarden/policy-sdk-go"
)

func validateSettings(payload []byte) ([]byte, error) {
	logger.Info("validating settings")

	// verify that we have an empty object in payload
	data := gjson.ParseBytes(payload)
	if ( ! data.IsObject() ) || ( len(data.Map()) > 0 ) {
		return kubewarden.RejectSettings(
			kubewarden.Message("Got invalid payload"),
		)
	}

	return kubewarden.AcceptSettings()
}

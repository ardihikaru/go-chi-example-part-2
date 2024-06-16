package httputil

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetJsonBody extracts json data from the request payload
func GetJsonBody(rBody io.ReadCloser, destType interface{}) (*string, int, error) {
	var errMsg string
	// extracts request body
	b, err := io.ReadAll(rBody)
	if err != nil {
		errMsg = "invalid JSON body"
		return &errMsg, http.StatusPreconditionFailed, err
	}

	// read JSON body from the request
	err = json.Unmarshal(b, &destType)
	if err != nil {
		errMsg = "failed to read JSON body from the request"
		return &errMsg, http.StatusPreconditionFailed, err
	}

	return nil, 200, nil
}

package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorResponseModel defines the error response model
type ErrorResponseModel struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// ResponseType defines what kind of response is sent for the request received
type ResponseType int

const (
	// ErrorResponse - Error response type Code
	ErrorResponse ResponseType = 1
	// SuccessResponse - Success response type Code
	SuccessResponse ResponseType = 2
)

// WriteResponse - writes the response to the http request
func WriteResponse(w http.ResponseWriter, responseType ResponseType, responseBody interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	switch responseType {
	case ErrorResponse:
		return writeErrResponse(w, responseBody)
	case SuccessResponse:
		return writeOkResponse(w, responseBody)
	default:
		return fmt.Errorf("unexpected response type")
	}
}

// writeErrResponse writes error response
func writeErrResponse(w http.ResponseWriter, errResponseBody interface{}) error {
	var err error
	var errResp *ErrResponse
	var response ErrorResponseModel

	// casts response body to ErrResponse
	errResp = errResponseBody.(*ErrResponse)
	if errResp == nil {
		// failed to decode the error response body
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("failed to decode the error response body")
	} else {
		w.WriteHeader(errResp.StatusCode)
	}

	response.Error = errResp.Error.Error()
	response.Message = errResp.Message

	responseBytes, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("failed to marshal response structure: %+v", err.Error())
	}

	_, err = w.Write(responseBytes)
	if err != nil {
		return fmt.Errorf("error in writing response: %s", err.Error())
	}

	return nil
}

// writeOkResponse writes OK response
func writeOkResponse(w http.ResponseWriter, responseBody interface{}) error {
	w.WriteHeader(http.StatusOK)

	responseBytes, err := json.Marshal(responseBody)

	if err != nil {
		return fmt.Errorf("failed to marshal response structure: %+v", err.Error())
	}

	_, err = w.Write(responseBytes)
	if err != nil {
		return fmt.Errorf("error in writing response: %s", err.Error())
	}

	return nil
}

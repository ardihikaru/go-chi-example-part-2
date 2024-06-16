package httputil

type Response struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type ErrResponse struct {
	StatusCode int    `json:"status_code"`
	Error      error  `json:"error,omitempty"`
	Message    string `json:"message,omitempty"`
}

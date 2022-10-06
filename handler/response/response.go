package response

type ResponseError struct {
	ErrorMsg *string `json:"error,omitempty"`
	DebugMsg string  `json:"debug,omitempty"`
}

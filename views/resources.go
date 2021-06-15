package views

type ErrorInfo struct {
	Scope    string `json:"scope" example:"server:rest_api"`
	Level    string `json:"level" example:"handler"`
	Position string `json:"position" example:"handlers.go: Line 41"`
	Trace    string `json:"trace" example:"..."`
}

type R struct {
	Status    string      `json:"status" example:"Success | Failure"`
	ErrorCode int         `json:"error_code" example:"0"`
	ErrorNote string      `json:"error_note" example:"Error Note"`
	Data      interface{} `json:"data"`
}

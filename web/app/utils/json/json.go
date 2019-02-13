package json

// Response transfer json data
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

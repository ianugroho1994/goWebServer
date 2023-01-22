package responseError

type ResponseError struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

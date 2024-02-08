package result

type Result[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func NewResult[T any](data T, message string, success bool) Result[T] {
	return Result[T]{Data: data, Message: message, Success: success}
}

func Success[T any](data T, message string) Result[T] {
	return NewResult(data, message, true)
}

func Error[T any](data T, message string) Result[T] {
	return NewResult(data, message, false)
}

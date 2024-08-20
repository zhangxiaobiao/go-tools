package utils

type Resp[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    T      `json:"data"`
}

func Success[T any](data T, message string) Resp[T] {
	resp := Resp[T]{
		Code:    0,
		Message: message,
		Data:    data,
	}
	return resp
}

func Error[T any](message string) Resp[T] {
	var zeroValue T
	resp := Resp[T]{
		Code:    1,
		Message: message,
		Data:    zeroValue,
	}
	return resp
}

package api_helper

import "errors"

type Response struct {
	Message string `json:"message"`
}

type ErrResponse struct {
	Message string `json:"errorMessage"`
}

var (
	ErrInvalidBody = errors.New("请检查你的请求体")
)

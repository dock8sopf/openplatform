package global

import "github.com/karldoenitz/Tigo/TigoWeb"

type JsonResponse struct {
	TigoWeb.BaseResponse
	Status  int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

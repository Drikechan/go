package api

import (
	"encoding/json"
	"test-todolist/pkg/ctl"
	"test-todolist/pkg/e"
)

func ErrorResponse(err error) *ctl.TrackErrorResponse {
	if _, ok := err.(*json.UnsupportedTypeError); ok {
		return ctl.ErrResponse(err, "JSON类型不匹配")
	}
	return ctl.ErrResponse(err, "参数错误", e.InvalidParams)
}

package ctl

import "test-todolist/pkg/e"

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type TrackErrorResponse struct {
	Response
	TrackId string `json:"trackId"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

// 错误返回
func ErrResponse(err error, data string, code ...int) *TrackErrorResponse {
	status := e.ERROR
	if code != nil {
		status = code[0]
	}
	return &TrackErrorResponse{
		Response: Response{
			Status: status,
			Msg:    e.GetMessage(status),
			Data:   data,
			Error:  err.Error(),
		},
	}
}

func RespSuccess(code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		code[0] = status
	}
	return &Response{
		Status: status,
		Data:   "操作成功",
		Msg:    e.GetMessage(status),
	}
}

func RespSuccessWithData(data interface{}, code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		code[0] = status
	}
	r := &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMessage(status),
	}
	return r
}

func RespList(items interface{}, total int64) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}

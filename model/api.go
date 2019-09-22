package model

type Response struct {
	Status bool        `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func APIResponse(status bool, objects interface{}, msg string) (response *Response) {
	response = &Response{Status: status, Data: objects, Msg: msg}
	return response
}
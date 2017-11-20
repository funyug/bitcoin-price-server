package models

type Response struct {
	Success int
	Data interface{}
}

func Success(data interface{}) Response {
	response := Response{Success:1,Data:data}
	return response
}

func Fail(data interface{}) Response {
	response := Response{Success:0,Data:data}
	return response
}
package response

import "net/http"

type Response struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Validation string      `json:"validation"`
	Data       interface{} `json:"data"`
}

func ResOK(message string, data interface{}) Response {
	res := Response{
		Status:     http.StatusOK,
		Message:    message,
		Validation: "",
		Data:       data,
	}
	return res
}

func ResErr(status int, message string) Response {
	res := Response{
		Status:     status,
		Message:    message,
		Validation: "",
		Data:       map[string]interface{}{},
	}
	return res
}

func ResErrValidate(status int, validation string) Response {
	res := Response{
		Status:     status,
		Message:    "required field",
		Validation: validation,
		Data:       map[string]interface{}{},
	}
	return res
}

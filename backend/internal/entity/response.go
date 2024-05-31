package entity

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Response struct {
	RetCode
	Data interface{} `json:"data,omitempty"`
}

func ToResponseCode(writer http.ResponseWriter, err error) {
	var resp Response
	var code RetCode
	if errors.As(err, &code) {
		resp = Response{RetCode: code}
	}
	result, _ := json.Marshal(resp)
	writer.Write(result)
}

func ToResponseData(writer http.ResponseWriter, data interface{}) {
	response := Response{
		RetCode: Success,
		Data:    data,
	}
	result, _ := json.Marshal(response)
	writer.Write(result)
}

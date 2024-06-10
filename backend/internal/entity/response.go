package entity

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hexiaopi/data-space/pkg/retcode"
)

type Response struct {
	*retcode.RetCode
	Data interface{} `json:"data,omitempty"`
}

func ToResponseCode(writer http.ResponseWriter, err error) {
	var resp Response
	var code *retcode.RetCode
	if errors.As(err, &code) {
		resp = Response{RetCode: code}
	}
	result, _ := json.Marshal(resp)
	writer.Write(result)
}

func ToResponseData(writer http.ResponseWriter, data interface{}) {
	response := Response{
		RetCode: retcode.Success,
		Data:    data,
	}
	result, _ := json.Marshal(response)
	writer.Write(result)
}

package entity

import (
	"encoding/json"
	"fmt"
)

type RetCode struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

var codes = map[string]string{}

func NewCode(code, desc string) RetCode {
	if v, ok := codes[code]; ok {
		panic(fmt.Sprintf("code:%s already exist with desc:%s", code, v))
	}
	codes[code] = desc
	return RetCode{Code: code, Desc: desc}
}

func (e RetCode) Marshal() []byte {
	data, _ := json.Marshal(e)
	return data
}

func (e RetCode) Error() string {
	return fmt.Sprintf("code:%s desc:%s", e.Code, e.Desc)
}

package retcode

import (
	"encoding/json"
	"fmt"
)

type RetCode struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

var codes = map[string]string{}

func NewCode(code, desc string) *RetCode {
	if v, ok := codes[code]; ok {
		panic(fmt.Sprintf("code:%s already exist with desc:%s", code, v))
	}
	codes[code] = desc
	return &RetCode{Code: code, Desc: desc}
}

func (e RetCode) Marshal() []byte {
	data, _ := json.Marshal(e)
	return data
}

func (e *RetCode) Error() string {
	return fmt.Sprintf("code:%s desc:%s", e.Code, e.Desc)
}

var (
	//router
	Success                = NewCode("000000", "Success")
	UnknownError           = NewCode("999999", "Unknown Error")
	RequestPathNotFound    = NewCode("000001", "Request Path Not Found")
	RequestMethodNotAllow  = NewCode("000002", "Request Method Not Allowed")
	RequestIllegal         = NewCode("000003", "Request Illegal")
	RequestUnMarshalError  = NewCode("000004", "Request Body UnMarshal Error")
	RequestNoPermission    = NewCode("000005", "Request No Permission")
	RequestAuthCheckFail   = NewCode("000006", "Request Auth Check Fail")
	JWTGenerateTokenFail   = NewCode("000007", "Generate Auth Token Fail")
	RequestTokenEmpty      = NewCode("000008", "Request Token Empty")
	RequestTokenWrong      = NewCode("000009", "Request Token Wrong")
	RequestTokenAuthFail   = NewCode("000010", "Request Token Auth Fail")
	RequestTokenAuthExpire = NewCode("000011", "Request Token Expire")
	UserGetFail            = NewCode("000012", "User Found Fail")
	UserNotExist           = NewCode("000013", "User Not Exist")
	UserPassWordWrong      = NewCode("000014", "User PassWord Wrong")
)

package entity

var (
	Success                = NewCode("000000", "Success")
	RequestPathNotFound    = NewCode("010000", "Request Path Not Found")
	RequestMethodNotAllow  = NewCode("010001", "Request Method Not Allowed")
	RequestIllegal         = NewCode("010002", "Request Illegal")
	RequestUnMarshalError  = NewCode("010003", "Request Body UnMarshal Error")
	RequestNoPermission    = NewCode("010004", "Request No Permission")
	RequestAuthCheckFail   = NewCode("020001", "Request Auth Check Fail")
	GenerateAuthTokenFail  = NewCode("020002", "Generate Auth Token Fail")
	RequestTokenEmpty      = NewCode("020003", "Request Token Empty")
	RequestTokenWrong      = NewCode("020004", "Request Token Wrong")
	RequestTokenAuthFail   = NewCode("020005", "Request Token Auth Fail")
	RequestTokenAuthExpire = NewCode("020006", "Request Token Expire")
	RequestUserGetFail     = NewCode("020007", "Get User Fail")
)

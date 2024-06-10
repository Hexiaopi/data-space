package global

import "github.com/hexiaopi/data-space/pkg/retcode"

var (
	//router
	RequestPathNotFound    = retcode.NewCode("000001", "Request Path Not Found")
	RequestMethodNotAllow  = retcode.NewCode("000002", "Request Method Not Allowed")
	RequestIllegal         = retcode.NewCode("000003", "Request Illegal")
	RequestUnMarshalError  = retcode.NewCode("000004", "Request Body UnMarshal Error")
	RequestNoPermission    = retcode.NewCode("000005", "Request No Permission")
	RequestAuthCheckFail   = retcode.NewCode("000006", "Request Auth Check Fail")
	JWTGenerateTokenFail   = retcode.NewCode("000007", "Generate Auth Token Fail")
	RequestTokenEmpty      = retcode.NewCode("000008", "Request Token Empty")
	RequestTokenWrong      = retcode.NewCode("000009", "Request Token Wrong")
	RequestTokenAuthFail   = retcode.NewCode("000010", "Request Token Auth Fail")
	RequestTokenAuthExpire = retcode.NewCode("000011", "Request Token Expire")

	UserGetFail       = retcode.NewCode("010001", "User Found Fail")
	UserNotExist      = retcode.NewCode("010002", "User Not Exist")
	UserPassWordWrong = retcode.NewCode("010003", "User PassWord Wrong")

	RoleGetFail  = retcode.NewCode("020001", "Role Found Fail")
	RoleNotExist = retcode.NewCode("020002", "Role Not Exist")
)

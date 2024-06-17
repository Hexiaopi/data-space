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

	UserGetFail       = retcode.NewCode("000101", "User Found Fail")
	UserNotExist      = retcode.NewCode("000102", "User Not Exist")
	UserPassWordWrong = retcode.NewCode("000103", "User PassWord Wrong")

	RoleListFail   = retcode.NewCode("000201", "Role List Fail")
	RoleCountFail  = retcode.NewCode("000202", "Role Count Fail")
	RoleGetFail    = retcode.NewCode("000203", "Role Found Fail")
	RoleNotExist   = retcode.NewCode("000204", "Role Not Exist")
	RoleCreateFail = retcode.NewCode("000205", "Role Create Fail")
	RoleUpdateFail = retcode.NewCode("000206", "Role Update Fail")
	RoleDeleteFail = retcode.NewCode("000207", "Role Delete Fail")
)

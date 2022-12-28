package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	// code = 1000... 用户模块的错误

	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenExist     = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	ErrorUserNoRight    = 1008

	// code = 2000... 文章模块的错误

	ErrorArtNotExist = 2001

	// code = 3000... 分类模块的错误

	ErrorCateNameUsed = 3001
	ErrorCateNotExist = 3002

	// code = 4000... token错误

	ErrorAuthCheckTokenFail    = 4001 //token 错误
	ErrorAuthCheckTokenTimeout = 4002 //token 过期
)

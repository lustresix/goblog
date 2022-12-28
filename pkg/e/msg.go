package e

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误！",

	ErrorUsernameUsed:   "用户名已存在！",
	ErrorPasswordWrong:  "密码错误！",
	ErrorUserNotExist:   "用户不存在！",
	ErrorTokenExist:     "TOKEN不存在！",
	ErrorTokenRuntime:   "TOKEN已过期！",
	ErrorTokenWrong:     "TOKEN不正确！",
	ErrorTokenTypeWrong: "TOKEN格式错误,请重新登陆！",
	ErrorUserNoRight:    "该用户无权限！",

	ErrorArtNotExist: "文章不存在",

	ErrorCateNameUsed: "该分类已存在",
	ErrorCateNotExist: "该分类不存在",

	ErrorAuthCheckTokenFail:    "无权限，token错误!",
	ErrorAuthCheckTokenTimeout: "token 过期!",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

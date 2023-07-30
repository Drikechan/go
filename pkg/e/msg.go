package e

var MsgFlags = map[int]string{
	ERROR:   "请求失败",
	SUCCESS: "操作成功",

	ErrorCheckTokenNotFound: "缺少token",
	ErrorCheckTokenFail:     "token验证失败",
	ErrorCheckTokenTimeout:  "token过期",
}

func GetMessage(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

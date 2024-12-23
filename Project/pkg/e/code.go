package e

const (
	SUCCESS                 = 200
	InvalidParams           = 400
	ErrorAuthCheckTokenFail = 500
)

func GetMsg(code int) string {
	switch code {
	case SUCCESS:
		return "成功"
	case InvalidParams:
		return "参数无效"
	case ErrorAuthCheckTokenFail:
		return "认证失败"
	default:
		return "未知错误"
	}
}

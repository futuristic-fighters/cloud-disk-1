package cfg

const (
	//1-100 公共操作码

	InvalidRequest = -1 // 请求无效

	//100-200 用户模块操作码
	UpdateUserSuccess   = 100 //更新成功
	UserNameValidateErr = 101 //用户名验证错误
)

var Lang = map[int]string{
	InvalidRequest:      "Invalid request",
	UpdateUserSuccess:   "Invalid request",
	UserNameValidateErr: "Invalid user name",
}

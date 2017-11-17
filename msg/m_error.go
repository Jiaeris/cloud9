package msg

var (
	OK            = 0
	PasswordError = 2
	NotLoggedIn   = 5

	AccountNotExist        = 1 //账户不存在
	AccountExist           = 6 //账户已存在
	AccountDetailsNotExist = 7 //账户资料不存在
)

type ResponseMessage struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

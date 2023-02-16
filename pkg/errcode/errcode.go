package errcode

// 成功响应
const (
	OK = 0
)

// user-1
const (
	// system-0
	Bcrypt_error             = 1001
	Create_user_error        = 1002
	Failed_to_generate_token = 1003

	// user-1
	Invalid_pass_parameter = 1101
	User_already_exists    = 1102
	User_does_not_exist    = 1103
	Password_error         = 1104
)

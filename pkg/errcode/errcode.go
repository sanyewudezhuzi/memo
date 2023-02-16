package errcode

// 成功响应
const (
	OK = 0
)

// middleware-0
const (
	// system-0

	// user-1
	Failed_to_get_request_header = 0101
	Failed_to_load_token         = 0102
	Token_has_expired            = 0103
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

// task-2
const (
	// system-0
	Create_task_error = 2001

	// user-1
	Failed_to_verify_identity = 2101
	Parameter_transfer_error  = 2102
)

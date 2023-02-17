package errcode

// 错误码
/*
由四位数表示
第一位：表示出错的模块
第二位：表示系统出错/用户出错
第三四位：表示错误
*/

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
	Update_task_error = 2002
	Delete_task_error = 2003

	// user-1
	Failed_to_verify_identity = 2101
	Parameter_transfer_error  = 2102
	No_title_found            = 2103
	The_title_has_been_used   = 2104
)

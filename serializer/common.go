package serializer

// 基础序列化器
type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"msg"`
	Error      string      `json:"error"`
}

// token data
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// list data
type ListData struct {
	Total int         `json:"num"`
	List  interface{} `json:"list"`
}

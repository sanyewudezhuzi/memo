package serializer

type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"msg"`
	Error      string      `json:"error"`
}

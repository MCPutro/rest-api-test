package response

type Response struct {
	Code    string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

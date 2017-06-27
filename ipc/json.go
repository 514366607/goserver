package ipc

//初始化请求格式结构体
type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

//初始化返回格式结构体
type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

//聊天信息格式结构体
type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:conetnt`
}

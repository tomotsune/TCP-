package common

const (
	LoginResType = iota
	LoginReqType
)

type Message struct {
	Type int    `json:"type"`
	Data string `json:"data"`
}
type LoginReq struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"pwd"`
}
type LoginRes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

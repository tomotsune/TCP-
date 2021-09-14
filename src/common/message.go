package common

const (
	Res = iota
	Login
	Register
	NotifyUserStatus
	SMS
)
const (
	OffLine = iota
	OnLine
)

func Converter(msg *Message, field string) interface{} {
	return msg.Data.(map[string]interface{})[field]
}

type Message struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}

type R struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
type NotifyUserStatusMes struct {
	UserMobile string `json:"userMobile"`
	Status     int    `json:"status"`
}
type SmsMsg struct {
	UserMobile string `json:"userMobile"`
	Content    string `json:"content"`
}

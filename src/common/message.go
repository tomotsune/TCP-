package common

const (
	Res = iota
	Login
	Register
)

type Message struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}

type R struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

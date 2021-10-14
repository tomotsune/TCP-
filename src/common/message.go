// @Title  通用信息协议
// @Description 定义了通信双方所要用到的所有协议类型
// @Author  haipinHu  08/10/2021 08:23
// @Update  haipinHu  08/10/2021 08:23
package common

// 定义通信类型常量:
// Res  响应信息类型
// Login 登录信息类型
// NotifyUserStatus 用户状态信息类型
// SMS 短信息类型
const (
	Res = iota
	Login
	Register
	NotifyUserStatus
	SMS
)

// 定义用户信息类型常量
const (
	OffLine = iota
	OnLine
)

// @title    消息转化
// @description   从Message协议数据单元获取信息
// @auth      hupinHu             08/9/2021 08:23
// @param     msg        *Message
// @param     field      string   所需数据条目
// @return    data[field]        interface{}
func Converter(msg *Message, field string) interface{} {
	return msg.Data.(map[string]interface{})[field]
}

// Message 信息对象, 定义了信息头和信息体
type Message struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}

// R 响应信息对象
type R struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// NotifyUserStatusMes 用户状态通知对象, 包含源地址信息和用户状态
type NotifyUserStatusMes struct {
	UserMobile string `json:"userMobile"`
	Status     int    `json:"status"`
}

// SmsMsg 端信息对象, 包含源地址信息和信息体
type SmsMsg struct {
	UserMobile string `json:"userMobile"`
	Content    string `json:"content"`
}

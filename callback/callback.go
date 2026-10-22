// Package callback 提供创蓝短信/视频短信推送回调的服务端解析模型。
//
// 这些回调由创蓝服务器主动 POST 到用户配置的回调地址，SDK 不发起调用；
// 接入方使用本包提供的结构体对推送 body 解码，按文档要求返回约定应答即可。
package callback

// Ack 表示接收回调约定应答。
type Ack struct {
	// ClCode 接收成功标识，用于 SMS 类回调应答。
	ClCode string `json:"clcode"`
	Msg    string `json:"msg"`
}

// OKACK 返回短信类回调的接收成功应答：{"clcode":"000000","msg":"接收成功"}。
func OKAck() *Ack {
	return &Ack{ClCode: "000000", Msg: "接收成功"}
}

// RCSAck 表示视频短信回调的接收成功应答（如需返回）。
const RCSAck = "ok"

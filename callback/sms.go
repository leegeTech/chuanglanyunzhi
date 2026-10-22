// Package callback 提供创蓝短信/视频短信推送回调的服务端解析模型（docs 中的推送类接口）。
package callback

import "encoding/json"

// 事件类型常量。
const (
	// EventSMSReport 短信状态回执。
	EventSMSReport = "SMS_REPORT"
	// EventSMSMO 上行报告。
	EventSMSMO = "SMS_MO"
)

// SmsPushEnvelope 短信推送回调的顶层结构。
// 对应 docs/sms/6-状态回执.md 与 7-上行报告.md 的顶层参数。
type SmsPushEnvelope struct {
	// ID 推送事件 ID，无业务含义。
	ID string `json:"ID"`
	// EventCode 事件类型：SMS_REPORT-状态回执；SMS_MO-上行报告。
	EventCode string `json:"EventCode"`
	// CreatedAt 推送时间，格式 yyyy-MM-ddTHH:mm:ss.SSS+时区。
	CreatedAt string `json:"CreatedAt"`
	// CustomTraceId 原样返回的请求追踪标识；请求未携带则不返回。
	CustomTraceId string `json:"X-Custom-TraceId"`
	// Data 业务数据（状态回执或上行）。
	Data json.RawMessage `json:"Data"`
}

// SmsReportData 状态回执业务数据。对应 docs/sms/6。
type SmsReportData struct {
	Receiver   string `json:"receiver"`
	Pswd       string `json:"pswd"`
	Mobile     string `json:"mobile"`
	Status     string `json:"status"`
	StatusDesc string `json:"statusDesc"`
	NotifyTime string `json:"notifyTime"`
	ReportTime string `json:"reportTime"`
	MsgId      string `json:"msgid"`
	Length     string `json:"length"`
	Uid        string `json:"uid"`
}

// SmsMOData 上行报告业务数据。对应 docs/sms/7。
type SmsMOData struct {
	Receiver   string `json:"receiver"`
	Pswd       string `json:"pswd"`
	Mobile     string `json:"mobile"`
	NotifyTime string `json:"notifyTime"`
	MoTime     string `json:"moTime"`
	Msg        string `json:"msg"`
	Extend     string `json:"extend"`
}

// ParseSmsReport 解析短信状态回执推送，返回 *SmsPushEnvelope（EventCode=SMS_REPORT）。
func ParseSmsReport(data []byte) (*SmsPushEnvelope, error) {
	var env SmsPushEnvelope
	if err := json.Unmarshal(data, &env); err != nil {
		return nil, err
	}
	return &env, nil
}

// ParseSmsReportData 解析状态回执的 Data 字段。
func ParseSmsReportData(raw json.RawMessage) (*SmsReportData, error) {
	var d SmsReportData
	if err := json.Unmarshal(raw, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// ParseSmsMO 解析上行报告 Data 字段。
func ParseSmsMO(raw json.RawMessage) (*SmsMOData, error) {
	var d SmsMOData
	if err := json.Unmarshal(raw, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

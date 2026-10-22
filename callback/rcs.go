package callback

import "encoding/json"

// RcsSignPush 视频短信签名审核结果推送。对应 docs/video-sms/3-签名管理.md 第 2 节。
type RcsSignPush struct {
	AccountId  string `json:"accountId"`
	SignId     string `json:"signId"`
	Status     int    `json:"status"` // 1 审核中 2 审核成功 3 审核失败
	Remark     string `json:"remark"`
	CmccStatus string `json:"cmccStatus"` // 0 不可用 1 可用
	CuccStatus string `json:"cuccStatus"`
	CtccStatus string `json:"ctccStatus"`
	PushUrl    string `json:"pushUrl"`
}

// RcsTemplatePush 视频短信模板审核结果推送。对应 docs/video-sms/4-模版管理.md 第 3 节。
type RcsTemplatePush struct {
	TemplateId string `json:"templateId"`
	Status     int    `json:"status"` // 1 平台审核中 2 运营商审核中 3 审核成功 4 审核失败
	Remark     string `json:"remark"`
	Time       string `json:"time"`
	CmccStatus int    `json:"cmccStatus"`
	CuccStatus int    `json:"cuccStatus"`
	CtccStatus int    `json:"ctccStatus"`
	ExpireFlag string `json:"expireFlag"`
}

// RcsReportPush 视频短信状态报告推送（JSON 数组）。对应 docs/video-sms/5-状态报告.md 第 1 节。
type RcsReportPush struct {
	AccountId      string `json:"accountId"`
	BatchId        string `json:"batchId"`
	ClErrorCode    string `json:"clErrorCode"`
	CreateTime     int64  `json:"createTime"`
	DeliveryStatus string `json:"deliveryStatus"`
	MsgId          string `json:"msgId"`
	Phone          string `json:"phone"`
	SubmitNo       string `json:"submitNo"`
}

// RcsReplyPush 视频短信上行消息推送（JSON 数组）。对应 docs/video-sms/6-上行消息.md 第 1 节。
type RcsReplyPush struct {
	AccountId     string `json:"accountId"`
	CodeNo        string `json:"codeNo"`
	CodeNoAddress string `json:"codeNoAddress"`
	CreateTime    int64  `json:"createTime"`
	MsgId         string `json:"msgId"`
	Phone         string `json:"phone"`
	UpContent     string `json:"upContent"`
	BatchId       string `json:"batchId"`
	SubmitNo      string `json:"submitNo"`
}

// ParseRcsSignPush 解析视频短信签名审核结果推送。
func ParseRcsSignPush(data []byte) (*RcsSignPush, error) {
	var p RcsSignPush
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// ParseRcsTemplatePush 解析视频短信模板审核结果推送。
func ParseRcsTemplatePush(data []byte) (*RcsTemplatePush, error) {
	var p RcsTemplatePush
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// ParseRcsReportPush 解析视频短信状态报告推送，返回报告列表。
func ParseRcsReportPush(data []byte) ([]*RcsReportPush, error) {
	var list []*RcsReportPush
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	return list, nil
}

// ParseRcsReplyPush 解析视频短信上行消息推送，返回上行列表。
func ParseRcsReplyPush(data []byte) ([]*RcsReplyPush, error) {
	var list []*RcsReplyPush
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	return list, nil
}

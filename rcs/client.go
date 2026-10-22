// Package rcs 提供视频短信(RCS)业务封装，共 6 个方法。
package rcs

import (
	"context"
	"encoding/json"

	"chuanglanyunzhi/cloud-sdk/internal/transport"
)

// RcsSmsTemplateAddRequest 添加视频模板请求。
type RcsSmsTemplateAddRequest struct {
	TemplateName    string `json:"templateName,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	VideoUrl        string `json:"videoUrl,omitempty"`
	SignName        string `json:"signName,omitempty"`
}

func (r *RcsSmsTemplateAddRequest) SetTemplateName(v string) *RcsSmsTemplateAddRequest {
	r.TemplateName = v
	return r
}
func (r *RcsSmsTemplateAddRequest) SetTemplateContent(v string) *RcsSmsTemplateAddRequest {
	r.TemplateContent = v
	return r
}
func (r *RcsSmsTemplateAddRequest) SetVideoUrl(v string) *RcsSmsTemplateAddRequest {
	r.VideoUrl = v
	return r
}
func (r *RcsSmsTemplateAddRequest) SetSignName(v string) *RcsSmsTemplateAddRequest {
	r.SignName = v
	return r
}

type RcsSmsTemplateAddResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		TemplateId string `json:"templateId"`
	} `json:"data"`
}

func (r *RcsSmsTemplateAddResponse) IsSuccess() bool { return r.Code == "000000" }

// RcsSmsTemplateFindRequest 查询视频模板请求。
type RcsSmsTemplateFindRequest struct {
	TemplateId string `json:"templateId,omitempty"`
}

func (r *RcsSmsTemplateFindRequest) SetTemplateId(v string) *RcsSmsTemplateFindRequest {
	r.TemplateId = v
	return r
}

type RcsSmsTemplateFindResponse struct {
	Code string              `json:"code"`
	Msg  string              `json:"msg"`
	Data *RcsSmsTemplateItem `json:"data"`
}

func (r *RcsSmsTemplateFindResponse) IsSuccess() bool { return r.Code == "000000" }

// RcsSmsTemplateItem 视频模板项。
type RcsSmsTemplateItem struct {
	TemplateId   string `json:"templateId"`
	TemplateName string `json:"templateName"`
	Content      string `json:"content"`
	VideoUrl     string `json:"videoUrl"`
	Status       string `json:"status"`
}

// RcsSmsTemplateSubmitRequest 发送视频短信请求。
type RcsSmsTemplateSubmitRequest struct {
	PhoneNumbers string `json:"phoneNumbers,omitempty"`
	TemplateCode string `json:"templateCode,omitempty"`
	Report       bool   `json:"report,omitempty"`
}

func (r *RcsSmsTemplateSubmitRequest) SetPhoneNumbers(v string) *RcsSmsTemplateSubmitRequest {
	r.PhoneNumbers = v
	return r
}
func (r *RcsSmsTemplateSubmitRequest) SetTemplateCode(v string) *RcsSmsTemplateSubmitRequest {
	r.TemplateCode = v
	return r
}
func (r *RcsSmsTemplateSubmitRequest) SetReport(v bool) *RcsSmsTemplateSubmitRequest {
	r.Report = v
	return r
}

type RcsSmsTemplateSubmitResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		MsgId string `json:"msgId"`
	} `json:"data"`
}

func (r *RcsSmsTemplateSubmitResponse) IsSuccess() bool { return r.Code == "000000" }

// RcsSmsReportPullRequest 拉取状态报告请求。
type RcsSmsReportPullRequest struct {
	Count int `json:"count,omitempty"`
}

func (r *RcsSmsReportPullRequest) SetCount(v int) *RcsSmsReportPullRequest {
	r.Count = v
	return r
}

type RcsSmsReportItem struct {
	MsgId  string `json:"msgId"`
	Phone  string `json:"phone"`
	Status string `json:"status"`
	Desc   string `json:"desc"`
}

type RcsSmsReportPullResponse struct {
	Code string              `json:"code"`
	Msg  string              `json:"msg"`
	Data []*RcsSmsReportItem `json:"data"`
}

func (r *RcsSmsReportPullResponse) IsSuccess() bool { return r.Code == "000000" }

// RcsSmsReplyPullRequest 拉取上行回复请求。
type RcsSmsReplyPullRequest struct {
	Count int `json:"count,omitempty"`
}

func (r *RcsSmsReplyPullRequest) SetCount(v int) *RcsSmsReplyPullRequest {
	r.Count = v
	return r
}

type RcsSmsReplyItem struct {
	Phone   string `json:"phone"`
	Content string `json:"content"`
	Time    string `json:"time"`
}

type RcsSmsReplyPullResponse struct {
	Code string             `json:"code"`
	Msg  string             `json:"msg"`
	Data []*RcsSmsReplyItem `json:"data"`
}

func (r *RcsSmsReplyPullResponse) IsSuccess() bool { return r.Code == "000000" }

// RcsSmsSignAddRequest 添加签名请求。
type RcsSmsSignAddRequest struct {
	SignName string `json:"signName,omitempty"`
	SignType string `json:"signType,omitempty"`
	Remark   string `json:"remark,omitempty"`
}

func (r *RcsSmsSignAddRequest) SetSignName(v string) *RcsSmsSignAddRequest {
	r.SignName = v
	return r
}
func (r *RcsSmsSignAddRequest) SetSignType(v string) *RcsSmsSignAddRequest {
	r.SignType = v
	return r
}
func (r *RcsSmsSignAddRequest) SetRemark(v string) *RcsSmsSignAddRequest {
	r.Remark = v
	return r
}

type RcsSmsSignAddResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		SignatureId string `json:"signatureId"`
	} `json:"data"`
}

func (r *RcsSmsSignAddResponse) IsSuccess() bool { return r.Code == "000000" }

// RcsClient 视频短信业务客户端。
type RcsClient struct {
	tp       *transport.Transport
	endpoint string
	auth     transport.Auth
}

// New 返回一个 RcsClient。
func New(tp *transport.Transport, endpoint string, appId, appSecret string) *RcsClient {
	return &RcsClient{
		tp:       tp,
		endpoint: endpoint,
		auth:     transport.Auth{AppId: appId, AppSecret: appSecret},
	}
}

func (c *RcsClient) do(ctx context.Context, path string, body, result any, traceId string) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	resp, err := c.tp.Do(ctx, "POST", c.endpoint, path, data, c.auth, transport.AuthSMS, traceId, false)
	if err != nil {
		return err
	}
	if resp == nil {
		return nil
	}
	return json.Unmarshal(resp, result)
}

// AddVideoTemplate 添加视频模板。
//
// 参考文档"视频短信业务 1. addVideoTemplate - 添加视频模板"。
//
//	req := &rcs.RcsSmsTemplateAddRequest{}.
//		SetTemplateName("营销视频").
//		SetTemplateContent("视频内容描述").
//		SetVideoUrl("https://example.com/video.mp4").
//		SetSignName("创蓝云")
//	resp, err := client.Rcs.AddVideoTemplate(ctx, req)
func (c *RcsClient) AddVideoTemplate(ctx context.Context, req *RcsSmsTemplateAddRequest) (*RcsSmsTemplateAddResponse, error) {
	resp := &RcsSmsTemplateAddResponse{}
	err := c.do(ctx, "/rcs/videoTemplate/add", req, resp, "")
	return resp, err
}

// AddVideoTemplateTrace 添加视频模板（带 traceId 链路追踪）。
// 参考文档"视频短信业务 1. addVideoTemplate - 添加视频模板"。
func (c *RcsClient) AddVideoTemplateTrace(ctx context.Context, req *RcsSmsTemplateAddRequest, traceId string) (*RcsSmsTemplateAddResponse, error) {
	resp := &RcsSmsTemplateAddResponse{}
	err := c.do(ctx, "/rcs/videoTemplate/add", req, resp, traceId)
	return resp, err
}

// FindVideoTemplate 查询视频模板。
//
// 参考文档"视频短信业务 2. findVideoTemplate - 查询视频模板"。
func (c *RcsClient) FindVideoTemplate(ctx context.Context, req *RcsSmsTemplateFindRequest) (*RcsSmsTemplateFindResponse, error) {
	resp := &RcsSmsTemplateFindResponse{}
	err := c.do(ctx, "/rcs/videoTemplate/find", req, resp, "")
	return resp, err
}

// FindVideoTemplateTrace 查询视频模板（带 traceId 链路追踪）。
// 参考文档"视频短信业务 2. findVideoTemplate - 查询视频模板"。
func (c *RcsClient) FindVideoTemplateTrace(ctx context.Context, req *RcsSmsTemplateFindRequest, traceId string) (*RcsSmsTemplateFindResponse, error) {
	resp := &RcsSmsTemplateFindResponse{}
	err := c.do(ctx, "/rcs/videoTemplate/find", req, resp, traceId)
	return resp, err
}

// SubmitVideoTemplate 发送视频短信。
//
// 参考文档"视频短信业务 3. submitVideoTemplate - 发送视频短信"。
//
//	req := &rcs.RcsSmsTemplateSubmitRequest{}.
//		SetPhoneNumbers("13800138000,13900139000").
//		SetTemplateCode("VT12345").
//		SetReport(true)
//	resp, err := client.Rcs.SubmitVideoTemplate(ctx, req)
func (c *RcsClient) SubmitVideoTemplate(ctx context.Context, req *RcsSmsTemplateSubmitRequest) (*RcsSmsTemplateSubmitResponse, error) {
	resp := &RcsSmsTemplateSubmitResponse{}
	err := c.do(ctx, "/rcs/videoTemplate/submit", req, resp, "")
	return resp, err
}

// SubmitVideoTemplateTrace 发送视频短信（带 traceId 链路追踪）。
// 参考文档"视频短信业务 3. submitVideoTemplate - 发送视频短信"。
func (c *RcsClient) SubmitVideoTemplateTrace(ctx context.Context, req *RcsSmsTemplateSubmitRequest, traceId string) (*RcsSmsTemplateSubmitResponse, error) {
	resp := &RcsSmsTemplateSubmitResponse{}
	err := c.do(ctx, "/rcs/videoTemplate/submit", req, resp, traceId)
	return resp, err
}

// PullReport 拉取状态报告，获取短信的发送状态。
//
// 参考文档"视频短信业务 4. pullReport - 拉取状态报告"。
func (c *RcsClient) PullReport(ctx context.Context, req *RcsSmsReportPullRequest) (*RcsSmsReportPullResponse, error) {
	resp := &RcsSmsReportPullResponse{}
	err := c.do(ctx, "/rcs/report/pull", req, resp, "")
	return resp, err
}

// PullReportTrace 拉取状态报告（带 traceId 链路追踪）。
// 参考文档"视频短信业务 4. pullReport - 拉取状态报告"。
func (c *RcsClient) PullReportTrace(ctx context.Context, req *RcsSmsReportPullRequest, traceId string) (*RcsSmsReportPullResponse, error) {
	resp := &RcsSmsReportPullResponse{}
	err := c.do(ctx, "/rcs/report/pull", req, resp, traceId)
	return resp, err
}

// PullReply 拉取上行回复，获取用户的回复内容。
//
// 参考文档"视频短信业务 5. pullReply - 拉取上行回复"。
func (c *RcsClient) PullReply(ctx context.Context, req *RcsSmsReplyPullRequest) (*RcsSmsReplyPullResponse, error) {
	resp := &RcsSmsReplyPullResponse{}
	err := c.do(ctx, "/rcs/reply/pull", req, resp, "")
	return resp, err
}

// PullReplyTrace 拉取上行回复（带 traceId 链路追踪）。
// 参考文档"视频短信业务 5. pullReply - 拉取上行回复"。
func (c *RcsClient) PullReplyTrace(ctx context.Context, req *RcsSmsReplyPullRequest, traceId string) (*RcsSmsReplyPullResponse, error) {
	resp := &RcsSmsReplyPullResponse{}
	err := c.do(ctx, "/rcs/reply/pull", req, resp, traceId)
	return resp, err
}

// AddSign 添加签名。
//
// 参考文档"视频短信业务 6. addSign - 添加签名"。
func (c *RcsClient) AddSign(ctx context.Context, req *RcsSmsSignAddRequest) (*RcsSmsSignAddResponse, error) {
	resp := &RcsSmsSignAddResponse{}
	err := c.do(ctx, "/rcs/sign/add", req, resp, "")
	return resp, err
}

// AddSignTrace 添加签名（带 traceId 链路追踪）。
// 参考文档"视频短信业务 6. addSign - 添加签名"。
func (c *RcsClient) AddSignTrace(ctx context.Context, req *RcsSmsSignAddRequest, traceId string) (*RcsSmsSignAddResponse, error) {
	resp := &RcsSmsSignAddResponse{}
	err := c.do(ctx, "/rcs/sign/add", req, resp, traceId)
	return resp, err
}

package rcs

import (
	"context"
	"encoding/json"

	"chuanglanyunzhi/cloud-sdk/internal/transport"
)

// RcsClient 视频短信(RCS)业务客户端。使用通用连接池。
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

// AddVideoTemplate 创建视频短信模板。变量模板仅文本可加变量，最多 5 个：${v1}~${v5}。
//
// 路径 /videosms/api/v2/template/add。
func (c *RcsClient) AddVideoTemplate(ctx context.Context, req *RcsTemplateAddRequest) (*RcsTemplateAddResponse, error) {
	resp := &RcsTemplateAddResponse{}
	err := c.do(ctx, "/videosms/api/v2/template/add", req, resp, "")
	return resp, err
}

// AddVideoTemplateTrace 创建视频短信模板（带 traceId 链路追踪）。
func (c *RcsClient) AddVideoTemplateTrace(ctx context.Context, req *RcsTemplateAddRequest, traceId string) (*RcsTemplateAddResponse, error) {
	resp := &RcsTemplateAddResponse{}
	err := c.do(ctx, "/videosms/api/v2/template/add", req, resp, traceId)
	return resp, err
}

// FindVideoTemplate 查看视频模板状态。
//
// 路径 /videosms/api/v2/template/getSingleTemplateInfo。
func (c *RcsClient) FindVideoTemplate(ctx context.Context, req *RcsTemplateFindRequest) (*RcsTemplateFindResponse, error) {
	resp := &RcsTemplateFindResponse{}
	err := c.do(ctx, "/videosms/api/v2/template/getSingleTemplateInfo", req, resp, "")
	return resp, err
}

// FindVideoTemplateTrace 查看视频模板状态（带 traceId 链路追踪）。
func (c *RcsClient) FindVideoTemplateTrace(ctx context.Context, req *RcsTemplateFindRequest, traceId string) (*RcsTemplateFindResponse, error) {
	resp := &RcsTemplateFindResponse{}
	err := c.do(ctx, "/videosms/api/v2/template/getSingleTemplateInfo", req, resp, traceId)
	return resp, err
}

// SubmitVideoTemplate 发送视频短信，支持静态模板（手机号列表）和动态模板（带变量）两种方式。
//
// 路径 /videosms/api/v2/template/send。
func (c *RcsClient) SubmitVideoTemplate(ctx context.Context, req *RcsTemplateSubmitRequest) (*RcsTemplateSubmitResponse, error) {
	resp := &RcsTemplateSubmitResponse{}
	err := c.do(ctx, "/videosms/api/v2/template/send", req, resp, "")
	return resp, err
}

// SubmitVideoTemplateTrace 发送视频短信（带 traceId 链路追踪）。
func (c *RcsClient) SubmitVideoTemplateTrace(ctx context.Context, req *RcsTemplateSubmitRequest, traceId string) (*RcsTemplateSubmitResponse, error) {
	resp := &RcsTemplateSubmitResponse{}
	err := c.do(ctx, "/videosms/api/v2/template/send", req, resp, traceId)
	return resp, err
}

// PullReport 拉取状态报告。Body 传空 JSON 对象。状态报告保存 72 小时，拉取成功后删除。
//
// 路径 /videosms/api/v2/report/pull。
func (c *RcsClient) PullReport(ctx context.Context, req *RcsReportPullRequest) (*RcsReportPullResponse, error) {
	resp := &RcsReportPullResponse{}
	err := c.do(ctx, "/videosms/api/v2/report/pull", req, resp, "")
	return resp, err
}

// PullReportTrace 拉取状态报告（带 traceId 链路追踪）。
func (c *RcsClient) PullReportTrace(ctx context.Context, req *RcsReportPullRequest, traceId string) (*RcsReportPullResponse, error) {
	resp := &RcsReportPullResponse{}
	err := c.do(ctx, "/videosms/api/v2/report/pull", req, resp, traceId)
	return resp, err
}

// PullReply 拉取上行消息。Body 传空 JSON 对象。上行消息保存 72 小时，拉取成功后删除。
//
// 路径 /videosms/api/v2/reply/pull。
func (c *RcsClient) PullReply(ctx context.Context, req *RcsReplyPullRequest) (*RcsReplyPullResponse, error) {
	resp := &RcsReplyPullResponse{}
	err := c.do(ctx, "/videosms/api/v2/reply/pull", req, resp, "")
	return resp, err
}

// PullReplyTrace 拉取上行消息（带 traceId 链路追踪）。
func (c *RcsClient) PullReplyTrace(ctx context.Context, req *RcsReplyPullRequest, traceId string) (*RcsReplyPullResponse, error) {
	resp := &RcsReplyPullResponse{}
	err := c.do(ctx, "/videosms/api/v2/reply/pull", req, resp, traceId)
	return resp, err
}

// AddSign 创建签名。签名内容如未带【】，系统会自动包裹。
//
// 路径 /videosms/api/v2/sign/add。
func (c *RcsClient) AddSign(ctx context.Context, req *RcsSignAddRequest) (*RcsSignAddResponse, error) {
	resp := &RcsSignAddResponse{}
	err := c.do(ctx, "/videosms/api/v2/sign/add", req, resp, "")
	return resp, err
}

// AddSignTrace 创建签名（带 traceId 链路追踪）。
func (c *RcsClient) AddSignTrace(ctx context.Context, req *RcsSignAddRequest, traceId string) (*RcsSignAddResponse, error) {
	resp := &RcsSignAddResponse{}
	err := c.do(ctx, "/videosms/api/v2/sign/add", req, resp, traceId)
	return resp, err
}

// Package sms 提供短信业务封装，共 18 个方法。
// 涵盖短信发送、签名、模板、资质管理。
package sms

import (
	"context"
	"encoding/json"

	"chuanglanyunzhi/cloud-sdk/internal/transport"
)

// SmsClient 聚合短信业务的全部 API。
// 使用独立的 SMS 连接池，避免批量发送时与其他业务竞争。
type SmsClient struct {
	tp       *transport.Transport
	endpoint string
	auth     transport.Auth
}

// New 返回一个 SmsClient。
func New(tp *transport.Transport, endpoint string, appId, appSecret string) *SmsClient {
	return &SmsClient{
		tp:       tp,
		endpoint: endpoint,
		auth:     transport.Auth{AppId: appId, AppSecret: appSecret},
	}
}

// doSMS 统一调用层：POST JSON、SMS 签名、SMS 连接池。
func (c *SmsClient) doSMS(ctx context.Context, path string, body, result any) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	resp, err := c.tp.Do(ctx, "POST", c.endpoint, path, data, c.auth, transport.AuthSMS, "", true)
	if err != nil {
		return err
	}
	if resp == nil {
		return nil
	}
	return json.Unmarshal(resp, result)
}

func (c *SmsClient) doSMSTrace(ctx context.Context, path string, body, result any, traceId string) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	resp, err := c.tp.Do(ctx, "POST", c.endpoint, path, data, c.auth, transport.AuthSMS, traceId, true)
	if err != nil {
		return err
	}
	if resp == nil {
		return nil
	}
	return json.Unmarshal(resp, result)
}

// BatchSend 批量发送短信，<br>一次请求可向多个号码发送同一条短信。
//
// 参考文档"1. batchSend - 批量发送短信"。
//
//	req := &sms.SmsBatchSendRequest{}.
//		SetProductType("notify").
//		SetPhoneNumbers("13800138000,13900139000").
//		SetTemplateCode("T12345").
//		SetSignName("创蓝云").
//		SetReport(true)
//	resp, err := client.Sms.BatchSend(ctx, req)
//	if resp.IsSuccess() {
//		fmt.Println("msgId:", resp.Data.MsgId)
//	}
func (c *SmsClient) BatchSend(ctx context.Context, req *SmsBatchSendRequest) (*SmsBatchSendResponse, error) {
	resp := &SmsBatchSendResponse{}
	err := c.doSMS(ctx, "/sms/batchSend", req, resp)
	return resp, err
}

// BatchSendTrace 批量发送短信（带 traceId 链路追踪）。
//
// traceId 会写入 HTTP 请求头 X-Trace-Id，用于分布式链路追踪与问题排查。
// 参考文档"1. batchSend - 批量发送短信"。
func (c *SmsClient) BatchSendTrace(ctx context.Context, req *SmsBatchSendRequest, traceId string) (*SmsBatchSendResponse, error) {
	resp := &SmsBatchSendResponse{}
	err := c.doSMSTrace(ctx, "/sms/batchSend", req, resp, traceId)
	return resp, err
}

// AddQualification 添加资质，用于建立可信资质信息后再申请签名/模板。
//
// 参考文档"2. addQualification - 添加资质"。
//
//	req := &sms.SmsQualificationAddRequest{}.
//		SetQualificationType("1").
//		SetCompanyName("创蓝科技有限公司").
//		SetLicenseUrl("https://example.com/license.jpg")
//	resp, err := client.Sms.AddQualification(ctx, req)
func (c *SmsClient) AddQualification(ctx context.Context, req *SmsQualificationAddRequest) (*SmsQualificationAddResponse, error) {
	resp := &SmsQualificationAddResponse{}
	err := c.doSMS(ctx, "/sms/qualification/add", req, resp)
	return resp, err
}

// AddQualificationTrace 添加资质（带 traceId 链路追踪）。
// 参考文档"2. addQualification - 添加资质"。
func (c *SmsClient) AddQualificationTrace(ctx context.Context, req *SmsQualificationAddRequest, traceId string) (*SmsQualificationAddResponse, error) {
	resp := &SmsQualificationAddResponse{}
	err := c.doSMSTrace(ctx, "/sms/qualification/add", req, resp, traceId)
	return resp, err
}

// ListQualification 查询资质列表，支持分页。
//
// 参考文档"3. listQualification - 查询资质列表"。
//
//	req := &sms.SmsQualificationListRequest{}.
//		SetPageNum(1).
//		SetPageSize(20)
//	resp, err := client.Sms.ListQualification(ctx, req)
//	for _, item := range resp.Data.List {
//		fmt.Println("资质ID:", item.QualificationId)
//	}
func (c *SmsClient) ListQualification(ctx context.Context, req *SmsQualificationListRequest) (*SmsQualificationListResponse, error) {
	resp := &SmsQualificationListResponse{}
	err := c.doSMS(ctx, "/sms/qualification/list", req, resp)
	return resp, err
}

// ListQualificationTrace 查询资质列表（带 traceId 链路追踪）。
// 参考文档"3. listQualification - 查询资质列表"。
func (c *SmsClient) ListQualificationTrace(ctx context.Context, req *SmsQualificationListRequest, traceId string) (*SmsQualificationListResponse, error) {
	resp := &SmsQualificationListResponse{}
	err := c.doSMSTrace(ctx, "/sms/qualification/list", req, resp, traceId)
	return resp, err
}

// UpdateQualification 更新资质信息，如企业名称、资质图片等。
//
// 参考文档"4. updateQualification - 更新资质"。
func (c *SmsClient) UpdateQualification(ctx context.Context, req *SmsQualificationUpdateRequest) (*SmsQualificationUpdateResponse, error) {
	resp := &SmsQualificationUpdateResponse{}
	err := c.doSMS(ctx, "/sms/qualification/update", req, resp)
	return resp, err
}

// UpdateQualificationTrace 更新资质（带 traceId 链路追踪）。
// 参考文档"4. updateQualification - 更新资质"。
func (c *SmsClient) UpdateQualificationTrace(ctx context.Context, req *SmsQualificationUpdateRequest, traceId string) (*SmsQualificationUpdateResponse, error) {
	resp := &SmsQualificationUpdateResponse{}
	err := c.doSMSTrace(ctx, "/sms/qualification/update", req, resp, traceId)
	return resp, err
}

// DeleteQualification 删除资质。
//
// 参考文档"5. deleteQualification - 删除资质"。
func (c *SmsClient) DeleteQualification(ctx context.Context, req *SmsQualificationDeleteRequest) (*SmsQualificationDeleteResponse, error) {
	resp := &SmsQualificationDeleteResponse{}
	err := c.doSMS(ctx, "/sms/qualification/delete", req, resp)
	return resp, err
}

// DeleteQualificationTrace 删除资质（带 traceId 链路追踪）。
// 参考文档"5. deleteQualification - 删除资质"。
func (c *SmsClient) DeleteQualificationTrace(ctx context.Context, req *SmsQualificationDeleteRequest, traceId string) (*SmsQualificationDeleteResponse, error) {
	resp := &SmsQualificationDeleteResponse{}
	err := c.doSMSTrace(ctx, "/sms/qualification/delete", req, resp, traceId)
	return resp, err
}

// AddSignature 添加签名。签名是短信的落款，需审核通过后用于发送。
//
// 参考文档"6. addSignature - 添加签名"。
//
//	req := &sms.SmsSignatureAddRequest{}.
//		SetSignName("创蓝云").
//		SetSignType("1").
//		SetSignPurpose("1").
//		SetRemark("官方签名")
//	resp, err := client.Sms.AddSignature(ctx, req)
func (c *SmsClient) AddSignature(ctx context.Context, req *SmsSignatureAddRequest) (*SmsSignatureAddResponse, error) {
	resp := &SmsSignatureAddResponse{}
	err := c.doSMS(ctx, "/sms/signature/add", req, resp)
	return resp, err
}

// AddSignatureTrace 添加签名（带 traceId 链路追踪）。
// 参考文档"6. addSignature - 添加签名"。
func (c *SmsClient) AddSignatureTrace(ctx context.Context, req *SmsSignatureAddRequest, traceId string) (*SmsSignatureAddResponse, error) {
	resp := &SmsSignatureAddResponse{}
	err := c.doSMSTrace(ctx, "/sms/signature/add", req, resp, traceId)
	return resp, err
}

// GetSignature 查询签名详情。
//
// 参考文档"7. getSignature - 查询签名详情"。
func (c *SmsClient) GetSignature(ctx context.Context, req *SmsSignatureGetRequest) (*SmsSignatureGetResponse, error) {
	resp := &SmsSignatureGetResponse{}
	err := c.doSMS(ctx, "/sms/signature/get", req, resp)
	return resp, err
}

// GetSignatureTrace 查询签名详情（带 traceId 链路追踪）。
// 参考文档"7. getSignature - 查询签名详情"。
func (c *SmsClient) GetSignatureTrace(ctx context.Context, req *SmsSignatureGetRequest, traceId string) (*SmsSignatureGetResponse, error) {
	resp := &SmsSignatureGetResponse{}
	err := c.doSMSTrace(ctx, "/sms/signature/get", req, resp, traceId)
	return resp, err
}

// ListSignature 查询签名列表，支持按签名名称、状态筛选并分页。
//
// 参考文档"8. listSignature - 查询签名列表"。
//
//	req := &sms.SmsSignatureListRequest{}.
//		SetSignName("创蓝云").
//		SetStatus("2").
//		SetPageNum(1).
//		SetPageSize(20)
//	resp, err := client.Sms.ListSignature(ctx, req)
func (c *SmsClient) ListSignature(ctx context.Context, req *SmsSignatureListRequest) (*SmsSignatureListResponse, error) {
	resp := &SmsSignatureListResponse{}
	err := c.doSMS(ctx, "/sms/signature/list", req, resp)
	return resp, err
}

// ListSignatureTrace 查询签名列表（带 traceId 链路追踪）。
// 参考文档"8. listSignature - 查询签名列表"。
func (c *SmsClient) ListSignatureTrace(ctx context.Context, req *SmsSignatureListRequest, traceId string) (*SmsSignatureListResponse, error) {
	resp := &SmsSignatureListResponse{}
	err := c.doSMSTrace(ctx, "/sms/signature/list", req, resp, traceId)
	return resp, err
}

// GetSignatureOperatorRejectReason 查询签名运营商驳回原因，便于了解审核被驳回的具体原因并修正。
//
// 参考文档"9. getSignatureOperatorRejectReason - 查询签名运营商驳回原因"。
func (c *SmsClient) GetSignatureOperatorRejectReason(ctx context.Context, req *SmsSignatureOperatorRejectReasonRequest) (*SmsSignatureOperatorRejectReasonResponse, error) {
	resp := &SmsSignatureOperatorRejectReasonResponse{}
	err := c.doSMS(ctx, "/sms/signature/rejectReason", req, resp)
	return resp, err
}

// GetSignatureOperatorRejectReasonTrace 查询签名运营商驳回原因（带 traceId 链路追踪）。
// 参考文档"9. getSignatureOperatorRejectReason - 查询签名运营商驳回原因"。
func (c *SmsClient) GetSignatureOperatorRejectReasonTrace(ctx context.Context, req *SmsSignatureOperatorRejectReasonRequest, traceId string) (*SmsSignatureOperatorRejectReasonResponse, error) {
	resp := &SmsSignatureOperatorRejectReasonResponse{}
	err := c.doSMSTrace(ctx, "/sms/signature/rejectReason", req, resp, traceId)
	return resp, err
}

// UpdateSignatureRealName 更新签名实名信息。
//
// 参考文档"10. updateSignatureRealName - 更新签名实名信息"。
func (c *SmsClient) UpdateSignatureRealName(ctx context.Context, req *SmsSignatureRealNameUpdateRequest) (*SmsSignatureRealNameUpdateResponse, error) {
	resp := &SmsSignatureRealNameUpdateResponse{}
	err := c.doSMS(ctx, "/sms/signature/realName", req, resp)
	return resp, err
}

// UpdateSignatureRealNameTrace 更新签名实名信息（带 traceId 链路追踪）。
// 参考文档"10. updateSignatureRealName - 更新签名实名信息"。
func (c *SmsClient) UpdateSignatureRealNameTrace(ctx context.Context, req *SmsSignatureRealNameUpdateRequest, traceId string) (*SmsSignatureRealNameUpdateResponse, error) {
	resp := &SmsSignatureRealNameUpdateResponse{}
	err := c.doSMSTrace(ctx, "/sms/signature/realName", req, resp, traceId)
	return resp, err
}

// DeleteSignature 删除签名。
//
// 参考文档"11. deleteSignature - 删除签名"。
func (c *SmsClient) DeleteSignature(ctx context.Context, req *SmsSignatureDeleteRequest) (*SmsSignatureDeleteResponse, error) {
	resp := &SmsSignatureDeleteResponse{}
	err := c.doSMS(ctx, "/sms/signature/delete", req, resp)
	return resp, err
}

// DeleteSignatureTrace 删除签名（带 traceId 链路追踪）。
// 参考文档"11. deleteSignature - 删除签名"。
func (c *SmsClient) DeleteSignatureTrace(ctx context.Context, req *SmsSignatureDeleteRequest, traceId string) (*SmsSignatureDeleteResponse, error) {
	resp := &SmsSignatureDeleteResponse{}
	err := c.doSMSTrace(ctx, "/sms/signature/delete", req, resp, traceId)
	return resp, err
}

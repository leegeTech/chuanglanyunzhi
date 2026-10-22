package sms

import "context"

// 模板相关方法。

// AddTemplate 添加模板。模板为短信内容中可复用的部分，需审核通过后用于发送。
//
// 参考文档"12. addTemplate - 添加模板"。
//
//	req := &sms.SmsTemplateAddRequest{}.
//		SetTemplateName("验证码模板").
//		SetTemplateType("1").
//		SetTemplateContent("您的验证码是{1}，{2}分钟内有效").
//		SetSignName("创蓝云")
//	resp, err := client.Sms.AddTemplate(ctx, req)
func (c *SmsClient) AddTemplate(ctx context.Context, req *SmsTemplateAddRequest) (*SmsTemplateAddResponse, error) {
	resp := &SmsTemplateAddResponse{}
	err := c.doSMS(ctx, "/sms/template/add", req, resp)
	return resp, err
}

// AddTemplateTrace 添加模板（带 traceId 链路追踪）。
// 参考文档"12. addTemplate - 添加模板"。
func (c *SmsClient) AddTemplateTrace(ctx context.Context, req *SmsTemplateAddRequest, traceId string) (*SmsTemplateAddResponse, error) {
	resp := &SmsTemplateAddResponse{}
	err := c.doSMSTrace(ctx, "/sms/template/add", req, resp, traceId)
	return resp, err
}

// QueryTemplateTypeEnum 查询模板类型枚举，用于获取可用的模板类型取值。
//
// 参考文档"13. queryTemplateTypeEnum - 查询模板类型枚举"。
func (c *SmsClient) QueryTemplateTypeEnum(ctx context.Context, req *SmsTemplateTypeEnumRequest) (*SmsTemplateTypeEnumResponse, error) {
	resp := &SmsTemplateTypeEnumResponse{}
	err := c.doSMS(ctx, "/sms/template/typeEnum", req, resp)
	return resp, err
}

// QueryTemplateTypeEnumTrace 查询模板类型枚举（带 traceId 链路追踪）。
// 参考文档"13. queryTemplateTypeEnum - 查询模板类型枚举"。
func (c *SmsClient) QueryTemplateTypeEnumTrace(ctx context.Context, req *SmsTemplateTypeEnumRequest, traceId string) (*SmsTemplateTypeEnumResponse, error) {
	resp := &SmsTemplateTypeEnumResponse{}
	err := c.doSMSTrace(ctx, "/sms/template/typeEnum", req, resp, traceId)
	return resp, err
}

// ListTemplate 查询模板列表，支持按模板名称、状态筛选并分页。
//
// 参考文档"14. listTemplate - 查询模板列表"。
func (c *SmsClient) ListTemplate(ctx context.Context, req *SmsTemplateListRequest) (*SmsTemplateListResponse, error) {
	resp := &SmsTemplateListResponse{}
	err := c.doSMS(ctx, "/sms/template/list", req, resp)
	return resp, err
}

// ListTemplateTrace 查询模板列表（带 traceId 链路追踪）。
// 参考文档"14. listTemplate - 查询模板列表"。
func (c *SmsClient) ListTemplateTrace(ctx context.Context, req *SmsTemplateListRequest, traceId string) (*SmsTemplateListResponse, error) {
	resp := &SmsTemplateListResponse{}
	err := c.doSMSTrace(ctx, "/sms/template/list", req, resp, traceId)
	return resp, err
}

// GetTemplate 查询模板详情。
//
// 参考文档"15. getTemplate - 查询模板详情"。
func (c *SmsClient) GetTemplate(ctx context.Context, req *SmsTemplateGetRequest) (*SmsTemplateGetResponse, error) {
	resp := &SmsTemplateGetResponse{}
	err := c.doSMS(ctx, "/sms/template/get", req, resp)
	return resp, err
}

// GetTemplateTrace 查询模板详情（带 traceId 链路追踪）。
// 参考文档"15. getTemplate - 查询模板详情"。
func (c *SmsClient) GetTemplateTrace(ctx context.Context, req *SmsTemplateGetRequest, traceId string) (*SmsTemplateGetResponse, error) {
	resp := &SmsTemplateGetResponse{}
	err := c.doSMSTrace(ctx, "/sms/template/get", req, resp, traceId)
	return resp, err
}

// GetTemplateOperatorRejectReason 查询模板运营商驳回原因，便于了解审核被驳回的具体原因并修正。
//
// 参考文档"16. getTemplateOperatorRejectReason - 查询模板运营商驳回原因"。
func (c *SmsClient) GetTemplateOperatorRejectReason(ctx context.Context, req *SmsTemplateOperatorRejectReasonRequest) (*SmsTemplateOperatorRejectReasonResponse, error) {
	resp := &SmsTemplateOperatorRejectReasonResponse{}
	err := c.doSMS(ctx, "/sms/template/rejectReason", req, resp)
	return resp, err
}

// GetTemplateOperatorRejectReasonTrace 查询模板运营商驳回原因（带 traceId 链路追踪）。
// 参考文档"16. getTemplateOperatorRejectReason - 查询模板运营商驳回原因"。
func (c *SmsClient) GetTemplateOperatorRejectReasonTrace(ctx context.Context, req *SmsTemplateOperatorRejectReasonRequest, traceId string) (*SmsTemplateOperatorRejectReasonResponse, error) {
	resp := &SmsTemplateOperatorRejectReasonResponse{}
	err := c.doSMSTrace(ctx, "/sms/template/rejectReason", req, resp, traceId)
	return resp, err
}

// UpdateTemplate 更新模板信息。
//
// 参考文档"17. updateTemplate - 更新模板"。
func (c *SmsClient) UpdateTemplate(ctx context.Context, req *SmsTemplateUpdateRequest) (*SmsTemplateUpdateResponse, error) {
	resp := &SmsTemplateUpdateResponse{}
	err := c.doSMS(ctx, "/sms/template/update", req, resp)
	return resp, err
}

// UpdateTemplateTrace 更新模板（带 traceId 链路追踪）。
// 参考文档"17. updateTemplate - 更新模板"。
func (c *SmsClient) UpdateTemplateTrace(ctx context.Context, req *SmsTemplateUpdateRequest, traceId string) (*SmsTemplateUpdateResponse, error) {
	resp := &SmsTemplateUpdateResponse{}
	err := c.doSMSTrace(ctx, "/sms/template/update", req, resp, traceId)
	return resp, err
}

// DeleteTemplate 删除模板。
//
// 参考文档"18. deleteTemplate - 删除模板"。
func (c *SmsClient) DeleteTemplate(ctx context.Context, req *SmsTemplateDeleteRequest) (*SmsTemplateDeleteResponse, error) {
	resp := &SmsTemplateDeleteResponse{}
	err := c.doSMS(ctx, "/sms/template/delete", req, resp)
	return resp, err
}

// DeleteTemplateTrace 删除模板（带 traceId 链路追踪）。
// 参考文档"18. deleteTemplate - 删除模板"。
func (c *SmsClient) DeleteTemplateTrace(ctx context.Context, req *SmsTemplateDeleteRequest, traceId string) (*SmsTemplateDeleteResponse, error) {
	resp := &SmsTemplateDeleteResponse{}
	err := c.doSMSTrace(ctx, "/sms/template/delete", req, resp, traceId)
	return resp, err
}

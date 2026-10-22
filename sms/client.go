// Package sms 提供短信业务封装。
// 涵盖短信发送、签名、模板、资质管理，统一使用 /sms/v2 前缀（文档见 docs/sms）。
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

// do 统一调用层：POST JSON、统一应用签名、SMS 连接池。
func (c *SmsClient) do(ctx context.Context, path string, body, result any, traceId string) error {
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

// BatchSend 批量发送短信，一次请求可向多个号码发送同一条短信。
//
// 对应 docs/sms/5-短信发送.md，路径 /sms/v2/batch-send。
//
//	req := (&sms.SmsBatchSendRequest{}).
//		SetProductType("market").
//		SetPhoneNumbers("15800000000,15300000000").
//		SetTemplateCode("1111111").
//		SetTemplateParam(`[{"param1":"张三","param2":"13"}]`).
//		SetReport(true)
//	resp, err := client.Sms.BatchSend(ctx, req)
//	if resp.IsSuccess() {
//		fmt.Println("msgId:", resp.Data.MsgId)
//	}
func (c *SmsClient) BatchSend(ctx context.Context, req *SmsBatchSendRequest) (*SmsBatchSendResponse, error) {
	resp := &SmsBatchSendResponse{}
	err := c.do(ctx, "/sms/v2/batch-send", req, resp, "")
	return resp, err
}

// BatchSendTrace 批量发送短信（带 traceId 链路追踪）。
// traceId 会写入请求头 X-Custom-TraceId（≤64 字符），用于上下游请求关联。
func (c *SmsClient) BatchSendTrace(ctx context.Context, req *SmsBatchSendRequest, traceId string) (*SmsBatchSendResponse, error) {
	resp := &SmsBatchSendResponse{}
	err := c.do(ctx, "/sms/v2/batch-send", req, resp, traceId)
	return resp, err
}

// AddQualification 创建资质，用于建立可信资质信息后再申请签名/模板。
//
// 对应 docs/sms/2-资质管理.md 第 1 节，路径 /sms/v2/qualification/add。
func (c *SmsClient) AddQualification(ctx context.Context, req *SmsQualificationAddRequest) (*SmsQualificationAddResponse, error) {
	resp := &SmsQualificationAddResponse{}
	err := c.do(ctx, "/sms/v2/qualification/add", req, resp, "")
	return resp, err
}

// AddQualificationTrace 创建资质（带 traceId 链路追踪）。
func (c *SmsClient) AddQualificationTrace(ctx context.Context, req *SmsQualificationAddRequest, traceId string) (*SmsQualificationAddResponse, error) {
	resp := &SmsQualificationAddResponse{}
	err := c.do(ctx, "/sms/v2/qualification/add", req, resp, traceId)
	return resp, err
}

// ListQualification 查询资质列表，支持企业名称模糊搜索与分页。
// 返回的文件 URL 有效期为 5 分钟，过期后需重新调用获取。
func (c *SmsClient) ListQualification(ctx context.Context, req *SmsQualificationListRequest) (*SmsQualificationListResponse, error) {
	resp := &SmsQualificationListResponse{}
	err := c.do(ctx, "/sms/v2/qualification/list", req, resp, "")
	return resp, err
}

// ListQualificationTrace 查询资质列表（带 traceId 链路追踪）。
func (c *SmsClient) ListQualificationTrace(ctx context.Context, req *SmsQualificationListRequest, traceId string) (*SmsQualificationListResponse, error) {
	resp := &SmsQualificationListResponse{}
	err := c.do(ctx, "/sms/v2/qualification/list", req, resp, traceId)
	return resp, err
}

// UpdateQualification 编辑资质信息。注意营业执照可以补充，不能修改。
func (c *SmsClient) UpdateQualification(ctx context.Context, req *SmsQualificationUpdateRequest) (*SmsQualificationUpdateResponse, error) {
	resp := &SmsQualificationUpdateResponse{}
	err := c.do(ctx, "/sms/v2/qualification/update", req, resp, "")
	return resp, err
}

// UpdateQualificationTrace 编辑资质（带 traceId 链路追踪）。
func (c *SmsClient) UpdateQualificationTrace(ctx context.Context, req *SmsQualificationUpdateRequest, traceId string) (*SmsQualificationUpdateResponse, error) {
	resp := &SmsQualificationUpdateResponse{}
	err := c.do(ctx, "/sms/v2/qualification/update", req, resp, traceId)
	return resp, err
}

// DeleteQualification 删除资质。有绑定签名的资质无法删除，需先解除签名绑定。
func (c *SmsClient) DeleteQualification(ctx context.Context, req *SmsQualificationDeleteRequest) (*SmsQualificationDeleteResponse, error) {
	resp := &SmsQualificationDeleteResponse{}
	err := c.do(ctx, "/sms/v2/qualification/delete", req, resp, "")
	return resp, err
}

// DeleteQualificationTrace 删除资质（带 traceId 链路追踪）。
func (c *SmsClient) DeleteQualificationTrace(ctx context.Context, req *SmsQualificationDeleteRequest, traceId string) (*SmsQualificationDeleteResponse, error) {
	resp := &SmsQualificationDeleteResponse{}
	err := c.do(ctx, "/sms/v2/qualification/delete", req, resp, traceId)
	return resp, err
}

// ListQualificationChangeLog 查询资质变更日志。
func (c *SmsClient) ListQualificationChangeLog(ctx context.Context, req *SmsQualificationChangeLogRequest) (*SmsQualificationChangeLogResponse, error) {
	resp := &SmsQualificationChangeLogResponse{}
	err := c.do(ctx, "/sms/v2/qualification/change-log/list", req, resp, "")
	return resp, err
}

// ListQualificationChangeLogTrace 查询资质变更日志（带 traceId 链路追踪）。
func (c *SmsClient) ListQualificationChangeLogTrace(ctx context.Context, req *SmsQualificationChangeLogRequest, traceId string) (*SmsQualificationChangeLogResponse, error) {
	resp := &SmsQualificationChangeLogResponse{}
	err := c.do(ctx, "/sms/v2/qualification/change-log/list", req, resp, traceId)
	return resp, err
}

// AddSignature 创建签名。签名是短信的落款，需审核通过后用于发送。
//
// 对应 docs/sms/3-签名管理.md 第 1 节，路径 /sms/v2/signature/create。
func (c *SmsClient) AddSignature(ctx context.Context, req *SmsSignatureAddRequest) (*SmsSignatureAddResponse, error) {
	resp := &SmsSignatureAddResponse{}
	err := c.do(ctx, "/sms/v2/signature/create", req, resp, "")
	return resp, err
}

// AddSignatureTrace 创建签名（带 traceId 链路追踪）。
func (c *SmsClient) AddSignatureTrace(ctx context.Context, req *SmsSignatureAddRequest, traceId string) (*SmsSignatureAddResponse, error) {
	resp := &SmsSignatureAddResponse{}
	err := c.do(ctx, "/sms/v2/signature/create", req, resp, traceId)
	return resp, err
}

// ListSignature 查询签名列表，按签名 ID/名称筛选并分页。
func (c *SmsClient) ListSignature(ctx context.Context, req *SmsSignatureListRequest) (*SmsSignatureListResponse, error) {
	resp := &SmsSignatureListResponse{}
	err := c.do(ctx, "/sms/v2/signature/list", req, resp, "")
	return resp, err
}

// ListSignatureTrace 查询签名列表（带 traceId 链路追踪）。
func (c *SmsClient) ListSignatureTrace(ctx context.Context, req *SmsSignatureListRequest, traceId string) (*SmsSignatureListResponse, error) {
	resp := &SmsSignatureListResponse{}
	err := c.do(ctx, "/sms/v2/signature/list", req, resp, traceId)
	return resp, err
}

// GetSignature 查询签名详情，包含各运营商审核状态与驳回原因。
func (c *SmsClient) GetSignature(ctx context.Context, req *SmsSignatureGetRequest) (*SmsSignatureGetResponse, error) {
	resp := &SmsSignatureGetResponse{}
	err := c.do(ctx, "/sms/v2/signature/detail", req, resp, "")
	return resp, err
}

// GetSignatureTrace 查询签名详情（带 traceId 链路追踪）。
func (c *SmsClient) GetSignatureTrace(ctx context.Context, req *SmsSignatureGetRequest, traceId string) (*SmsSignatureGetResponse, error) {
	resp := &SmsSignatureGetResponse{}
	err := c.do(ctx, "/sms/v2/signature/detail", req, resp, traceId)
	return resp, err
}

// GetSignatureOperatorRejectReason 查询签名指定运营商（cm/ct/cu）的审核驳回原因。
// 底层复用签名详情接口。
func (c *SmsClient) GetSignatureOperatorRejectReason(ctx context.Context, req *SmsSignatureOperatorRejectReasonRequest) (*SmsSignatureOperatorRejectReasonResponse, error) {
	detail, err := c.GetSignature(ctx, &SmsSignatureGetRequest{
		ProductType: req.ProductType,
		SignId:      req.SignId,
	})
	if err != nil {
		return nil, err
	}
	resp := &SmsSignatureOperatorRejectReasonResponse{
		Code: detail.Code, Msg: detail.Msg, RequestId: detail.RequestId,
	}
	if detail.Data != nil {
		resp.Data = pickOperatorRejectReason(req.Operator, detail.Data)
	}
	return resp, nil
}

// GetSignatureOperatorRejectReasonTrace 查询签名运营商驳回原因（带 traceId 链路追踪）。
func (c *SmsClient) GetSignatureOperatorRejectReasonTrace(ctx context.Context, req *SmsSignatureOperatorRejectReasonRequest, traceId string) (*SmsSignatureOperatorRejectReasonResponse, error) {
	detail, err := c.GetSignatureTrace(ctx, &SmsSignatureGetRequest{
		ProductType: req.ProductType,
		SignId:      req.SignId,
	}, traceId)
	if err != nil {
		return nil, err
	}
	resp := &SmsSignatureOperatorRejectReasonResponse{
		Code: detail.Code, Msg: detail.Msg, RequestId: detail.RequestId,
	}
	if detail.Data != nil {
		resp.Data = pickOperatorRejectReason(req.Operator, detail.Data)
	}
	return resp, nil
}

func pickOperatorRejectReason(operator string, info *SmsSignatureGetInfo) []*RejectReasonItem {
	switch operator {
	case "ct":
		return info.CtOperatorRejectReason
	case "cu":
		return info.CuOperatorRejectReason
	default:
		return info.CmOperatorRejectReason
	}
}

// UpdateSignatureRealName 编辑签名终端资质或应用场景。
func (c *SmsClient) UpdateSignatureRealName(ctx context.Context, req *SmsSignatureRealNameUpdateRequest) (*SmsSignatureRealNameUpdateResponse, error) {
	resp := &SmsSignatureRealNameUpdateResponse{}
	err := c.do(ctx, "/sms/v2/signature/real-name-update", req, resp, "")
	return resp, err
}

// UpdateSignatureRealNameTrace 编辑签名终端资质（带 traceId 链路追踪）。
func (c *SmsClient) UpdateSignatureRealNameTrace(ctx context.Context, req *SmsSignatureRealNameUpdateRequest, traceId string) (*SmsSignatureRealNameUpdateResponse, error) {
	resp := &SmsSignatureRealNameUpdateResponse{}
	err := c.do(ctx, "/sms/v2/signature/real-name-update", req, resp, traceId)
	return resp, err
}

// DeleteSignature 删除签名。审核通过的签名不允许删除。
func (c *SmsClient) DeleteSignature(ctx context.Context, req *SmsSignatureDeleteRequest) (*SmsSignatureDeleteResponse, error) {
	resp := &SmsSignatureDeleteResponse{}
	err := c.do(ctx, "/sms/v2/signature/delete", req, resp, "")
	return resp, err
}

// DeleteSignatureTrace 删除签名（带 traceId 链路追踪）。
func (c *SmsClient) DeleteSignatureTrace(ctx context.Context, req *SmsSignatureDeleteRequest, traceId string) (*SmsSignatureDeleteResponse, error) {
	resp := &SmsSignatureDeleteResponse{}
	err := c.do(ctx, "/sms/v2/signature/delete", req, resp, traceId)
	return resp, err
}

// AddTemplate 创建模板。模板不包含签名和退订语，变量用 {s} 标识。
func (c *SmsClient) AddTemplate(ctx context.Context, req *SmsTemplateAddRequest) (*SmsTemplateAddResponse, error) {
	resp := &SmsTemplateAddResponse{}
	err := c.do(ctx, "/sms/v2/template/add", req, resp, "")
	return resp, err
}

// AddTemplateTrace 创建模板（带 traceId 链路追踪）。
func (c *SmsClient) AddTemplateTrace(ctx context.Context, req *SmsTemplateAddRequest, traceId string) (*SmsTemplateAddResponse, error) {
	resp := &SmsTemplateAddResponse{}
	err := c.do(ctx, "/sms/v2/template/add", req, resp, traceId)
	return resp, err
}

// QueryTemplateTypeEnum 查询模板类型枚举（变量类型、号码类型、链接类型、业务类型）。
func (c *SmsClient) QueryTemplateTypeEnum(ctx context.Context, req *SmsTemplateTypeEnumRequest) (*SmsTemplateTypeEnumResponse, error) {
	resp := &SmsTemplateTypeEnumResponse{}
	err := c.do(ctx, "/sms/v2/template/query-type-enum", req, resp, "")
	return resp, err
}

// QueryTemplateTypeEnumTrace 查询模板类型枚举（带 traceId 链路追踪）。
func (c *SmsClient) QueryTemplateTypeEnumTrace(ctx context.Context, req *SmsTemplateTypeEnumRequest, traceId string) (*SmsTemplateTypeEnumResponse, error) {
	resp := &SmsTemplateTypeEnumResponse{}
	err := c.do(ctx, "/sms/v2/template/query-type-enum", req, resp, traceId)
	return resp, err
}

// ListTemplate 查询模板列表，列表查询最大长度为 100。
func (c *SmsClient) ListTemplate(ctx context.Context, req *SmsTemplateListRequest) (*SmsTemplateListResponse, error) {
	resp := &SmsTemplateListResponse{}
	err := c.do(ctx, "/sms/v2/template/list", req, resp, "")
	return resp, err
}

// ListTemplateTrace 查询模板列表（带 traceId 链路追踪）。
func (c *SmsClient) ListTemplateTrace(ctx context.Context, req *SmsTemplateListRequest, traceId string) (*SmsTemplateListResponse, error) {
	resp := &SmsTemplateListResponse{}
	err := c.do(ctx, "/sms/v2/template/list", req, resp, traceId)
	return resp, err
}

// GetTemplate 查询模板详情，含各运营商审核状态与驳回原因。
func (c *SmsClient) GetTemplate(ctx context.Context, req *SmsTemplateGetRequest) (*SmsTemplateGetResponse, error) {
	resp := &SmsTemplateGetResponse{}
	err := c.do(ctx, "/sms/v2/template/detail", req, resp, "")
	return resp, err
}

// GetTemplateTrace 查询模板详情（带 traceId 链路追踪）。
func (c *SmsClient) GetTemplateTrace(ctx context.Context, req *SmsTemplateGetRequest, traceId string) (*SmsTemplateGetResponse, error) {
	resp := &SmsTemplateGetResponse{}
	err := c.do(ctx, "/sms/v2/template/detail", req, resp, traceId)
	return resp, err
}

// GetTemplateOperatorRejectReason 查询模板指定运营商（cm/ct/cu）的审核驳回原因。
// 底层复用模板详情接口。
func (c *SmsClient) GetTemplateOperatorRejectReason(ctx context.Context, req *SmsTemplateOperatorRejectReasonRequest) (*SmsTemplateOperatorRejectReasonResponse, error) {
	detail, err := c.GetTemplate(ctx, &SmsTemplateGetRequest{
		ProductType:  req.ProductType,
		TemplateCode: req.TemplateCode,
	})
	if err != nil {
		return nil, err
	}
	resp := &SmsTemplateOperatorRejectReasonResponse{
		Code: detail.Code, Msg: detail.Msg, RequestId: detail.RequestId,
	}
	if detail.Data != nil {
		switch req.Operator {
		case "ct":
			resp.Data = detail.Data.CtOperatorRejectReason
		case "cu":
			resp.Data = detail.Data.CuOperatorRejectReason
		default:
			resp.Data = detail.Data.CmOperatorRejectReason
		}
	}
	return resp, nil
}

// GetTemplateOperatorRejectReasonTrace 查询模板运营商驳回原因（带 traceId 链路追踪）。
func (c *SmsClient) GetTemplateOperatorRejectReasonTrace(ctx context.Context, req *SmsTemplateOperatorRejectReasonRequest, traceId string) (*SmsTemplateOperatorRejectReasonResponse, error) {
	detail, err := c.GetTemplateTrace(ctx, &SmsTemplateGetRequest{
		ProductType:  req.ProductType,
		TemplateCode: req.TemplateCode,
	}, traceId)
	if err != nil {
		return nil, err
	}
	resp := &SmsTemplateOperatorRejectReasonResponse{
		Code: detail.Code, Msg: detail.Msg, RequestId: detail.RequestId,
	}
	if detail.Data != nil {
		switch req.Operator {
		case "ct":
			resp.Data = detail.Data.CtOperatorRejectReason
		case "cu":
			resp.Data = detail.Data.CuOperatorRejectReason
		default:
			resp.Data = detail.Data.CmOperatorRejectReason
		}
	}
	return resp, nil
}

// UpdateTemplate 编辑模板。模板不支持编辑时返回错误码 201659。
func (c *SmsClient) UpdateTemplate(ctx context.Context, req *SmsTemplateUpdateRequest) (*SmsTemplateUpdateResponse, error) {
	resp := &SmsTemplateUpdateResponse{}
	err := c.do(ctx, "/sms/v2/template/update", req, resp, "")
	return resp, err
}

// UpdateTemplateTrace 编辑模板（带 traceId 链路追踪）。
func (c *SmsClient) UpdateTemplateTrace(ctx context.Context, req *SmsTemplateUpdateRequest, traceId string) (*SmsTemplateUpdateResponse, error) {
	resp := &SmsTemplateUpdateResponse{}
	err := c.do(ctx, "/sms/v2/template/update", req, resp, traceId)
	return resp, err
}

// DeleteTemplate 删除模板。
func (c *SmsClient) DeleteTemplate(ctx context.Context, req *SmsTemplateDeleteRequest) (*SmsTemplateDeleteResponse, error) {
	resp := &SmsTemplateDeleteResponse{}
	err := c.do(ctx, "/sms/v2/template/delete", req, resp, "")
	return resp, err
}

// DeleteTemplateTrace 删除模板（带 traceId 链路追踪）。
func (c *SmsClient) DeleteTemplateTrace(ctx context.Context, req *SmsTemplateDeleteRequest, traceId string) (*SmsTemplateDeleteResponse, error) {
	resp := &SmsTemplateDeleteResponse{}
	err := c.do(ctx, "/sms/v2/template/delete", req, resp, traceId)
	return resp, err
}

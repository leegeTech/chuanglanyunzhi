// Package sms 提供短信业务封装，共 18 个方法。
// 涵盖短信发送、签名、模板、资质管理。
package sms

// SmsBatchSendRequest 批量发送短信请求。
// TODO: 第二阶段补充各字段含义与约束说明。
type SmsBatchSendRequest struct {
	// ProductType 业务类型，如 "notify"。
	ProductType string `json:"productType,omitempty"`
	// PhoneNumbers 接收号码，多个以英文逗号分隔。
	PhoneNumbers string `json:"phoneNumbers,omitempty"`
	// TemplateCode 模板编码。
	TemplateCode string `json:"templateCode,omitempty"`
	// SignName 签名名称。
	SignName string `json:"signName,omitempty"`
	// Message 短信内容（变量模板时传入）。
	Message string `json:"message,omitempty"`
	// Params 模板变量参数。
	Params []string `json:"params,omitempty"`
	// Report 是否返回状态报告。
	Report bool `json:"report,omitempty"`
}

func (r *SmsBatchSendRequest) SetProductType(v string) *SmsBatchSendRequest {
	r.ProductType = v
	return r
}
func (r *SmsBatchSendRequest) SetPhoneNumbers(v string) *SmsBatchSendRequest {
	r.PhoneNumbers = v
	return r
}
func (r *SmsBatchSendRequest) SetTemplateCode(v string) *SmsBatchSendRequest {
	r.TemplateCode = v
	return r
}
func (r *SmsBatchSendRequest) SetSignName(v string) *SmsBatchSendRequest {
	r.SignName = v
	return r
}
func (r *SmsBatchSendRequest) SetMessage(v string) *SmsBatchSendRequest {
	r.Message = v
	return r
}
func (r *SmsBatchSendRequest) SetParams(v []string) *SmsBatchSendRequest {
	r.Params = v
	return r
}
func (r *SmsBatchSendRequest) SetReport(v bool) *SmsBatchSendRequest {
	r.Report = v
	return r
}

// SmsBatchSendData 批量发送结果数据。
type SmsBatchSendData struct {
	// MsgId 消息 ID。
	MsgId string `json:"msgId"`
	// MsgGroupId 消息组 ID。
	MsgGroupId string `json:"msgGroupId"`
}

// SmsBatchSendResponse 批量发送短信响应。
type SmsBatchSendResponse struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data *SmsBatchSendData `json:"data"`
}

// IsSuccess 判断业务是否成功。
func (r *SmsBatchSendResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsQualificationAddRequest 添加资质请求。
type SmsQualificationAddRequest struct {
	QualificationType string `json:"qualificationType,omitempty"` // 1=企业 2=个人
	CompanyName       string `json:"companyName,omitempty"`
	LicenseUrl        string `json:"licenseUrl,omitempty"`
}

func (r *SmsQualificationAddRequest) SetQualificationType(v string) *SmsQualificationAddRequest {
	r.QualificationType = v
	return r
}
func (r *SmsQualificationAddRequest) SetCompanyName(v string) *SmsQualificationAddRequest {
	r.CompanyName = v
	return r
}
func (r *SmsQualificationAddRequest) SetLicenseUrl(v string) *SmsQualificationAddRequest {
	r.LicenseUrl = v
	return r
}

type SmsQualificationAddResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		QualificationId string `json:"qualificationId"`
	} `json:"data"`
}

func (r *SmsQualificationAddResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsQualificationListRequest 查询资质列表请求。
type SmsQualificationListRequest struct {
	PageNum  int `json:"pageNum,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

func (r *SmsQualificationListRequest) SetPageNum(v int) *SmsQualificationListRequest {
	r.PageNum = v
	return r
}
func (r *SmsQualificationListRequest) SetPageSize(v int) *SmsQualificationListRequest {
	r.PageSize = v
	return r
}

// QualificationItem 资质列表项。
type QualificationItem struct {
	QualificationId   string `json:"qualificationId"`
	QualificationType string `json:"qualificationType"`
	CompanyName       string `json:"companyName"`
	Status            string `json:"status"`
}

type SmsQualificationListResponse struct {
	Code string                    `json:"code"`
	Msg  string                    `json:"msg"`
	Data *SmsQualificationListData `json:"data"`
}

type SmsQualificationListData struct {
	List     []*QualificationItem `json:"list"`
	Total    int                  `json:"total"`
	PageNum  int                  `json:"pageNum"`
	PageSize int                  `json:"pageSize"`
}

func (r *SmsQualificationListResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsQualificationUpdateRequest 更新资质请求。
type SmsQualificationUpdateRequest struct {
	QualificationId   string `json:"qualificationId,omitempty"`
	QualificationType string `json:"qualificationType,omitempty"`
	CompanyName       string `json:"companyName,omitempty"`
	LicenseUrl        string `json:"licenseUrl,omitempty"`
}

func (r *SmsQualificationUpdateRequest) SetQualificationId(v string) *SmsQualificationUpdateRequest {
	r.QualificationId = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetQualificationType(v string) *SmsQualificationUpdateRequest {
	r.QualificationType = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetCompanyName(v string) *SmsQualificationUpdateRequest {
	r.CompanyName = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetLicenseUrl(v string) *SmsQualificationUpdateRequest {
	r.LicenseUrl = v
	return r
}

type SmsQualificationUpdateResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (r *SmsQualificationUpdateResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsQualificationDeleteRequest 删除资质请求。
type SmsQualificationDeleteRequest struct {
	QualificationId string `json:"qualificationId,omitempty"`
}

func (r *SmsQualificationDeleteRequest) SetQualificationId(v string) *SmsQualificationDeleteRequest {
	r.QualificationId = v
	return r
}

type SmsQualificationDeleteResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (r *SmsQualificationDeleteResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsSignatureAddRequest 添加签名请求。
type SmsSignatureAddRequest struct {
	SignName    string `json:"signName,omitempty"`
	SignType    string `json:"signType,omitempty"`    // 1=网站 2=APP 3=微信公众号 4=企业名称
	SignPurpose string `json:"signPurpose,omitempty"` // 1=自用 2=他用
	Remark      string `json:"remark,omitempty"`
}

func (r *SmsSignatureAddRequest) SetSignName(v string) *SmsSignatureAddRequest {
	r.SignName = v
	return r
}
func (r *SmsSignatureAddRequest) SetSignType(v string) *SmsSignatureAddRequest {
	r.SignType = v
	return r
}
func (r *SmsSignatureAddRequest) SetSignPurpose(v string) *SmsSignatureAddRequest {
	r.SignPurpose = v
	return r
}
func (r *SmsSignatureAddRequest) SetRemark(v string) *SmsSignatureAddRequest {
	r.Remark = v
	return r
}

type SmsSignatureAddResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		SignatureId string `json:"signatureId"`
	} `json:"data"`
}

func (r *SmsSignatureAddResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsSignatureGetRequest 查询签名详情请求。
type SmsSignatureGetRequest struct {
	SignatureId string `json:"signatureId,omitempty"`
}

func (r *SmsSignatureGetRequest) SetSignatureId(v string) *SmsSignatureGetRequest {
	r.SignatureId = v
	return r
}

type SmsSignatureGetResponse struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data *SignatureItem `json:"data"`
}

func (r *SmsSignatureGetResponse) IsSuccess() bool { return r.Code == "000000" }

// SignatureItem 签名详情项。
type SignatureItem struct {
	SignatureId string `json:"signatureId"`
	SignName    string `json:"signName"`
	SignType    string `json:"signType"`
	SignPurpose string `json:"signPurpose"`
	Status      string `json:"status"`
	Remark      string `json:"remark"`
}

// SmsSignatureListRequest 查询签名列表请求。
type SmsSignatureListRequest struct {
	SignName string `json:"signName,omitempty"`
	Status   string `json:"status,omitempty"` // 1=待审核 2=已通过 3=已驳回
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

func (r *SmsSignatureListRequest) SetSignName(v string) *SmsSignatureListRequest {
	r.SignName = v
	return r
}
func (r *SmsSignatureListRequest) SetStatus(v string) *SmsSignatureListRequest {
	r.Status = v
	return r
}
func (r *SmsSignatureListRequest) SetPageNum(v int) *SmsSignatureListRequest {
	r.PageNum = v
	return r
}
func (r *SmsSignatureListRequest) SetPageSize(v int) *SmsSignatureListRequest {
	r.PageSize = v
	return r
}

type SmsSignatureListResponse struct {
	Code string                `json:"code"`
	Msg  string                `json:"msg"`
	Data *SmsSignatureListData `json:"data"`
}

type SmsSignatureListData struct {
	List     []*SignatureItem `json:"list"`
	Total    int              `json:"total"`
	PageNum  int              `json:"pageNum"`
	PageSize int              `json:"pageSize"`
}

func (r *SmsSignatureListResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsSignatureOperatorRejectReasonRequest 查询签名运营商驳回原因请求。
type SmsSignatureOperatorRejectReasonRequest struct {
	SignatureId string `json:"signatureId,omitempty"`
	ReasonType  string `json:"reasonType,omitempty"`
}

func (r *SmsSignatureOperatorRejectReasonRequest) SetSignatureId(v string) *SmsSignatureOperatorRejectReasonRequest {
	r.SignatureId = v
	return r
}
func (r *SmsSignatureOperatorRejectReasonRequest) SetReasonType(v string) *SmsSignatureOperatorRejectReasonRequest {
	r.ReasonType = v
	return r
}

type SmsSignatureOperatorRejectReasonResponse struct {
	Code string              `json:"code"`
	Msg  string              `json:"msg"`
	Data []*RejectReasonItem `json:"data"`
}

func (r *SmsSignatureOperatorRejectReasonResponse) IsSuccess() bool { return r.Code == "000000" }

type RejectReasonItem struct {
	Reason string `json:"reason"`
	Remark string `json:"remark"`
}

// SmsSignatureRealNameUpdateRequest 更新签名实名信息请求。
type SmsSignatureRealNameUpdateRequest struct {
	SignatureId string `json:"signatureId,omitempty"`
	RealName    string `json:"realName,omitempty"`
	IdCardNo    string `json:"idCardNo,omitempty"`
}

func (r *SmsSignatureRealNameUpdateRequest) SetSignatureId(v string) *SmsSignatureRealNameUpdateRequest {
	r.SignatureId = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetRealName(v string) *SmsSignatureRealNameUpdateRequest {
	r.RealName = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetIdCardNo(v string) *SmsSignatureRealNameUpdateRequest {
	r.IdCardNo = v
	return r
}

type SmsSignatureRealNameUpdateResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (r *SmsSignatureRealNameUpdateResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsSignatureDeleteRequest 删除签名请求。
type SmsSignatureDeleteRequest struct {
	SignatureId string `json:"signatureId,omitempty"`
}

func (r *SmsSignatureDeleteRequest) SetSignatureId(v string) *SmsSignatureDeleteRequest {
	r.SignatureId = v
	return r
}

type SmsSignatureDeleteResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (r *SmsSignatureDeleteResponse) IsSuccess() bool { return r.Code == "000000" }

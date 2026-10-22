// Package sms 提供短信业务封装。
// 涵盖短信发送、签名、模板、资质管理，统一使用 /sms/v2 前缀（文档见 docs/sms）。
package sms

// 统一应用 API Key 模式下的请求公共字段，各请求结构体通过内嵌 ProductType 提供产品类型。
type ProductType struct {
	// ProductType 产品类型：
	// notify-通知短信；market-营销短信；verify-验证码短信；
	// 2ec-CS 短信（二类电商）；finance-DK 短信（金融）；
	// game-BK 短信（游戏）；acquisition-HK 短信（拉新）。
	ProductType string `json:"productType"`
}

func (p *ProductType) SetProductType(v string) *ProductType {
	p.ProductType = v
	return p
}

// PageRequest 分页公共字段。
type PageRequest struct {
	// PageNo 页码，默认值为 1。
	PageNo int `json:"pageNo,omitempty"`
	// PageSize 每页条数，默认值为 10（列表接口最大值为 100）。
	PageSize int `json:"pageSize,omitempty"`
}

func (p *PageRequest) SetPageNum(v int) *PageRequest { p.PageNo = v; return p }
func (p *PageRequest) SetPageSize(v int) *PageRequest {
	p.PageSize = v
	return p
}

// SmsBatchSendRequest 批量发送短信请求。
// 对应 docs/sms/5-短信发送.md。
type SmsBatchSendRequest struct {
	ProductType
	// PhoneNumbers 手机号：11 位号码，多个用英文逗号分隔，上限 1000 个。
	PhoneNumbers string `json:"phoneNumbers"`
	// TemplateCode 模板 ID。
	TemplateCode string `json:"templateCode"`
	// TemplateParam 模板变量值：JSON 数组字符串，按手机号顺序为每个手机号提供一组变量值；
	// 模板中第 1、2、3…个变量依次使用 param1、param2、param3… 传值。模板含变量时必填。
	TemplateParam string `json:"templateParam,omitempty"`
	// Report 是否需要状态回执：true-需要；false-不需要。默认 false。
	Report bool `json:"report,omitempty"`
	// CallbackUrl 状态回执推送地址：需以 http/https 开头；不传时使用 API Key 配置的回调地址。
	CallbackUrl string `json:"callbackUrl,omitempty"`
	// OutId 自定义参数：状态回执中以 uid 字段原样回传，最大 256 字符。
	OutId string `json:"outId,omitempty"`
	// SmsUpExtendCode 上行扩展码：上行报告中以 extend 字段原样回传，限 5 位以内数字。
	SmsUpExtendCode string `json:"smsUpExtendCode,omitempty"`
}

func (r *SmsBatchSendRequest) SetPhoneNumbers(v string) *SmsBatchSendRequest {
	r.PhoneNumbers = v
	return r
}
func (r *SmsBatchSendRequest) SetTemplateCode(v string) *SmsBatchSendRequest {
	r.TemplateCode = v
	return r
}
func (r *SmsBatchSendRequest) SetTemplateParam(v string) *SmsBatchSendRequest {
	r.TemplateParam = v
	return r
}
func (r *SmsBatchSendRequest) SetCallbackUrl(v string) *SmsBatchSendRequest {
	r.CallbackUrl = v
	return r
}
func (r *SmsBatchSendRequest) SetOutId(v string) *SmsBatchSendRequest {
	r.OutId = v
	return r
}
func (r *SmsBatchSendRequest) SetSmsUpExtendCode(v string) *SmsBatchSendRequest {
	r.SmsUpExtendCode = v
	return r
}
func (r *SmsBatchSendRequest) SetReport(v bool) *SmsBatchSendRequest {
	r.Report = v
	return r
}

// SmsBatchSendResult 批量发送结果数据。
type SmsBatchSendResult struct {
	// FailNum 提交失败的短信条数。
	FailNum string `json:"failNum"`
	// SuccessNum 提交成功的短信条数。
	SuccessNum string `json:"successNum"`
	// MsgId 消息 ID：32 位纯数字组成的唯一标识。
	MsgId string `json:"msgId"`
}

// SmsBatchSendResponse 批量发送短信响应。
type SmsBatchSendResponse struct {
	Code      string              `json:"code"`
	Msg       string              `json:"msg"`
	RequestId string              `json:"requestId"`
	Data      *SmsBatchSendResult `json:"data"`
}

// IsSuccess 判断业务是否成功（code 为 000000）。
func (r *SmsBatchSendResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsQualificationAddRequest 创建资质请求。对应 docs/sms/2-资质管理.md 第 1 节。
type SmsQualificationAddRequest struct {
	ProductType
	// EnterpriseName 企业名称。
	EnterpriseName string `json:"enterpriseName"`
	// UnifiedSocialCreditCode 统一社会信用代码。
	UnifiedSocialCreditCode string `json:"unifiedSocialCreditCode"`
	// LegalName 法人姓名。
	LegalName string `json:"legalName"`
	// AttentionLine 经办人姓名。
	AttentionLine string `json:"attentionLine"`
	// AttentionIdCard 经办人身份证号。
	AttentionIdCard string `json:"attentionIdCard"`
	// PhoneNum 经办人手机号。
	PhoneNum string `json:"phoneNum"`
	// IdentityType 经办人类型：1-法人；2-经办人。默认 2。
	IdentityType string `json:"identityType,omitempty"`
	// EndType 签名实名使用方：1-自用；2-他用。默认 2。
	EndType string `json:"endType,omitempty"`
	// ShareType 共享范围：0-此账号；1-所有主子账号。
	ShareType string `json:"shareType"`
	// BusinessLicenseImg 营业执照图片（Base64，PNG/JPG/JPEG，≤2MB）。
	BusinessLicenseImg string `json:"businessLicenseImg"`
	// LegalNationalEmblemIdentityImg 法人身份证国徽面（Base64，≤2MB）。
	LegalNationalEmblemIdentityImg string `json:"legalNationalEmblemIdentityImg"`
	// LegalFigureIdentityImg 法人身份证人像面（Base64，≤2MB）。
	LegalFigureIdentityImg string `json:"legalFigureIdentityImg"`
	// AgentAuthorizationLetterImg 经办人授权书图片（identityType=2 时必填，Base64，≤2MB）。
	AgentAuthorizationLetterImg string `json:"agentAuthorizationLetterImg,omitempty"`
	// NationalEmblemIdentityImg 经办人身份证国徽面（identityType=2 时必填，Base64，≤2MB）。
	NationalEmblemIdentityImg string `json:"nationalEmblemIdentityImg,omitempty"`
	// FigureIdentityImg 经办人身份证人像面（identityType=2 时必填，Base64，≤2MB）。
	FigureIdentityImg string `json:"figureIdentityImg,omitempty"`
	// IdentityOfficePhotoImg 经办人办公场景照片（Base64，≤2MB）。
	IdentityOfficePhotoImg string `json:"identityOfficePhotoImg,omitempty"`
	// ContractOtherFile 终端合同照片（endType=2 时必填，≤20 个，PDF≤5MB/图片≤2MB）。
	ContractOtherFile []string `json:"contractOtherFile,omitempty"`
	// Remark 备注。
	Remark string `json:"remark,omitempty"`
}

func (r *SmsQualificationAddRequest) SetEnterpriseName(v string) *SmsQualificationAddRequest {
	r.EnterpriseName = v
	return r
}
func (r *SmsQualificationAddRequest) SetUnifiedSocialCreditCode(v string) *SmsQualificationAddRequest {
	r.UnifiedSocialCreditCode = v
	return r
}
func (r *SmsQualificationAddRequest) SetLegalName(v string) *SmsQualificationAddRequest {
	r.LegalName = v
	return r
}
func (r *SmsQualificationAddRequest) SetAttentionLine(v string) *SmsQualificationAddRequest {
	r.AttentionLine = v
	return r
}
func (r *SmsQualificationAddRequest) SetAttentionIdCard(v string) *SmsQualificationAddRequest {
	r.AttentionIdCard = v
	return r
}
func (r *SmsQualificationAddRequest) SetPhoneNum(v string) *SmsQualificationAddRequest {
	r.PhoneNum = v
	return r
}
func (r *SmsQualificationAddRequest) SetIdentityType(v string) *SmsQualificationAddRequest {
	r.IdentityType = v
	return r
}
func (r *SmsQualificationAddRequest) SetEndType(v string) *SmsQualificationAddRequest {
	r.EndType = v
	return r
}
func (r *SmsQualificationAddRequest) SetShareType(v string) *SmsQualificationAddRequest {
	r.ShareType = v
	return r
}
func (r *SmsQualificationAddRequest) SetBusinessLicenseImg(v string) *SmsQualificationAddRequest {
	r.BusinessLicenseImg = v
	return r
}
func (r *SmsQualificationAddRequest) SetLegalNationalEmblemIdentityImg(v string) *SmsQualificationAddRequest {
	r.LegalNationalEmblemIdentityImg = v
	return r
}
func (r *SmsQualificationAddRequest) SetLegalFigureIdentityImg(v string) *SmsQualificationAddRequest {
	r.LegalFigureIdentityImg = v
	return r
}
func (r *SmsQualificationAddRequest) SetAgentAuthorizationLetterImg(v string) *SmsQualificationAddRequest {
	r.AgentAuthorizationLetterImg = v
	return r
}
func (r *SmsQualificationAddRequest) SetNationalEmblemIdentityImg(v string) *SmsQualificationAddRequest {
	r.NationalEmblemIdentityImg = v
	return r
}
func (r *SmsQualificationAddRequest) SetFigureIdentityImg(v string) *SmsQualificationAddRequest {
	r.FigureIdentityImg = v
	return r
}
func (r *SmsQualificationAddRequest) SetIdentityOfficePhotoImg(v string) *SmsQualificationAddRequest {
	r.IdentityOfficePhotoImg = v
	return r
}
func (r *SmsQualificationAddRequest) SetContractOtherFile(v []string) *SmsQualificationAddRequest {
	r.ContractOtherFile = v
	return r
}
func (r *SmsQualificationAddRequest) SetRemark(v string) *SmsQualificationAddRequest {
	r.Remark = v
	return r
}

// SmsQualificationAddResult 创建资质结果。
type SmsQualificationAddResult struct {
	// EndCustomerId 终端资质 ID。
	EndCustomerId string `json:"endCustomerId"`
}

type SmsQualificationAddResponse struct {
	Code      string                     `json:"code"`
	Msg       string                     `json:"msg"`
	RequestId string                     `json:"requestId"`
	Data      *SmsQualificationAddResult `json:"data"`
}

func (r *SmsQualificationAddResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsQualificationListRequest 资质列表查询请求。
type SmsQualificationListRequest struct {
	ProductType
	// CompanyName 企业名称，支持模糊搜索。
	CompanyName string `json:"companyName"`
	// EndType 终端类型：1-自用；2-他用。
	EndType string `json:"endType,omitempty"`
	PageRequest
}

func (r *SmsQualificationListRequest) SetCompanyName(v string) *SmsQualificationListRequest {
	r.CompanyName = v
	return r
}
func (r *SmsQualificationListRequest) SetEndType(v string) *SmsQualificationListRequest {
	r.EndType = v
	return r
}

// QualificationItem 资质列表项（文件 URL 有效期 5 分钟）。
type QualificationItem struct {
	EndCustomerId                  string   `json:"endCustomerId"`
	EnterpriseName                 string   `json:"enterpriseName"`
	UnifiedSocialCreditCode        string   `json:"unifiedSocialCreditCode"`
	AttentionLine                  string   `json:"attentionLine"`
	AttentionIdCard                string   `json:"attentionIdCard"`
	LegalName                      string   `json:"legalName"`
	PhoneNum                       string   `json:"phoneNum"`
	IdentityType                   string   `json:"identityType"`
	EndType                        string   `json:"endType"`
	OneToManyFlag                  string   `json:"oneToManyFlag"`
	AgentAuthorizationLetterUrl    string   `json:"agentAuthorizationLetterUrl"`
	BusinessLicenseUrl             string   `json:"businessLicenseUrl"`
	NationalEmblemIdentityUrl      string   `json:"nationalEmblemIdentityUrl"`
	FigureIdentityUrl              string   `json:"figureIdentityUrl"`
	LegalNationalEmblemIdentityUrl string   `json:"legalNationalEmblemIdentityUrl"`
	LegalFigureIdentityUrl         string   `json:"legalFigureIdentityUrl"`
	IdentityOfficePhotoUrl         string   `json:"identityOfficePhotoUrl"`
	ContractOtherPreviewUrl        []string `json:"contractOtherPreviewUrl"`
	Remark                         string   `json:"remark"`
}

type SmsQualificationListData struct {
	Total int                  `json:"total"`
	List  []*QualificationItem `json:"list"`
}

type SmsQualificationListResponse struct {
	Code      string                    `json:"code"`
	Msg       string                    `json:"msg"`
	RequestId string                    `json:"requestId"`
	Data      *SmsQualificationListData `json:"data"`
}

func (r *SmsQualificationListResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsQualificationUpdateRequest 编辑资请求。
type SmsQualificationUpdateRequest struct {
	ProductType
	// EndCustomerId 终端资质 ID。
	EndCustomerId string `json:"endCustomerId"`
	// IdentityType 经办人类型：1-法人；2-经办人。
	IdentityType string `json:"identityType,omitempty"`
	LegalName    string `json:"legalName,omitempty"`
	PhoneNum     string `json:"phoneNum,omitempty"`
	// BusinessLicenseImg 营业执照可以补充，不能修改。
	BusinessLicenseImg             string   `json:"businessLicenseImg,omitempty"`
	AgentAuthorizationLetterImg    string   `json:"agentAuthorizationLetterImg,omitempty"`
	NationalEmblemIdentityImg      string   `json:"nationalEmblemIdentityImg,omitempty"`
	FigureIdentityImg              string   `json:"figureIdentityImg,omitempty"`
	LegalNationalEmblemIdentityImg string   `json:"legalNationalEmblemIdentityImg,omitempty"`
	LegalFigureIdentityImg         string   `json:"legalFigureIdentityImg,omitempty"`
	IdentityOfficePhotoImg         string   `json:"identityOfficePhotoImg,omitempty"`
	ContractOtherFile              []string `json:"contractOtherFile,omitempty"`
	Remark                         string   `json:"remark,omitempty"`
}

func (r *SmsQualificationUpdateRequest) SetEndCustomerId(v string) *SmsQualificationUpdateRequest {
	r.EndCustomerId = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetIdentityType(v string) *SmsQualificationUpdateRequest {
	r.IdentityType = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetLegalName(v string) *SmsQualificationUpdateRequest {
	r.LegalName = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetPhoneNum(v string) *SmsQualificationUpdateRequest {
	r.PhoneNum = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetBusinessLicenseImg(v string) *SmsQualificationUpdateRequest {
	r.BusinessLicenseImg = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetAgentAuthorizationLetterImg(v string) *SmsQualificationUpdateRequest {
	r.AgentAuthorizationLetterImg = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetNationalEmblemIdentityImg(v string) *SmsQualificationUpdateRequest {
	r.NationalEmblemIdentityImg = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetFigureIdentityImg(v string) *SmsQualificationUpdateRequest {
	r.FigureIdentityImg = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetLegalNationalEmblemIdentityImg(v string) *SmsQualificationUpdateRequest {
	r.LegalNationalEmblemIdentityImg = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetLegalFigureIdentityImg(v string) *SmsQualificationUpdateRequest {
	r.LegalFigureIdentityImg = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetIdentityOfficePhotoImg(v string) *SmsQualificationUpdateRequest {
	r.IdentityOfficePhotoImg = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetContractOtherFile(v []string) *SmsQualificationUpdateRequest {
	r.ContractOtherFile = v
	return r
}
func (r *SmsQualificationUpdateRequest) SetRemark(v string) *SmsQualificationUpdateRequest {
	r.Remark = v
	return r
}

type SmsQualificationUpdateResponse struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

func (r *SmsQualificationUpdateResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsQualificationDeleteRequest 删除资请求（有绑定签名的资质需先解除绑定）。
type SmsQualificationDeleteRequest struct {
	ProductType
	// EndCustomerId 终端资质 ID。
	EndCustomerId string `json:"endCustomerId"`
}

func (r *SmsQualificationDeleteRequest) SetEndCustomerId(v string) *SmsQualificationDeleteRequest {
	r.EndCustomerId = v
	return r
}

type SmsQualificationDeleteResponse struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

func (r *SmsQualificationDeleteResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsQualificationChangeLogRequest 资质变更日志查询请求。
type SmsQualificationChangeLogRequest struct {
	ProductType
	EndCustomerId string `json:"endCustomerId"`
	PageRequest
}

func (r *SmsQualificationChangeLogRequest) SetEndCustomerId(v string) *SmsQualificationChangeLogRequest {
	r.EndCustomerId = v
	return r
}

// QualificationChangeLogItem 资质变更记录。
type QualificationChangeLogItem struct {
	Id           int                    `json:"id"`
	OldInfo      *QualificationSnapshot `json:"oldInfo"`
	NewInfo      *QualificationSnapshot `json:"newInfo"`
	CreateTime   string                 `json:"createTime"`
	ApplyUser    string                 `json:"applyUser"`
	AuditOpinion string                 `json:"auditOpinion"`
	ChangeType   string                 `json:"changeType"`
	AuditStatus  int                    `json:"auditStatus"`
}

// QualificationSnapshot 资质快照（变更前后）。
type QualificationSnapshot struct {
	LegalCertNo                   string `json:"legalCertNo"`
	IdentityNo                    string `json:"identityNo"`
	PhoneNum                      string `json:"phoneNum"`
	IdentityOfficePhotoPreviewUrl string `json:"identityOfficePhotoPreviewUrl"`
}

type SmsQualificationChangeLogData struct {
	Total int                           `json:"total"`
	List  []*QualificationChangeLogItem `json:"list"`
}

type SmsQualificationChangeLogResponse struct {
	Code      string                         `json:"code"`
	Msg       string                         `json:"msg"`
	RequestId string                         `json:"requestId"`
	Data      *SmsQualificationChangeLogData `json:"data"`
}

func (r *SmsQualificationChangeLogResponse) IsSuccess() bool { return r.Code == codeSuccess }

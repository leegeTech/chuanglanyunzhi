package sms

// 签名相关请求与响应。对应 docs/sms/3-签名管理.md。

// SmsSignatureAddRequest 创建签名请求。
type SmsSignatureAddRequest struct {
	ProductType
	// SignType 签名类型：1-企业名称；3-商标；4-其他（如政企）。
	SignType string `json:"signType"`
	// SignName 签名名称，长度为 2～18 个字符。
	SignName string `json:"signName"`
	// SignatureSceneType 关联资质用途：3-自用；4-他人。
	SignatureSceneType int `json:"signatureSceneType"`
	// EndCustomerId 资质 ID。
	EndCustomerId string `json:"endCustomerId"`
	// SbjScreenshotImg 国家商标局官网查询结果截图（Base64，≤2MB）。
	SbjScreenshotImg string `json:"sbjScreenshotImg,omitempty"`
	// TrademarkAuthImg 商标持有人出具的授权书（Base64，≤2MB）。
	TrademarkAuthImg string `json:"trademarkAuthImg,omitempty"`
	// InstitutionImg 组织机构代码截图（Base64，≤2MB）。
	InstitutionImg string `json:"institutionImg,omitempty"`
	// AdditionalAttachmentsImg 补充材料（Base64，≤2MB，最多 9 张）。
	AdditionalAttachmentsImg []string `json:"additionalAttachmentsImg,omitempty"`
}

func (r *SmsSignatureAddRequest) SetSignType(v string) *SmsSignatureAddRequest {
	r.SignType = v
	return r
}
func (r *SmsSignatureAddRequest) SetSignName(v string) *SmsSignatureAddRequest {
	r.SignName = v
	return r
}
func (r *SmsSignatureAddRequest) SetSignatureSceneType(v int) *SmsSignatureAddRequest {
	r.SignatureSceneType = v
	return r
}
func (r *SmsSignatureAddRequest) SetEndCustomerId(v string) *SmsSignatureAddRequest {
	r.EndCustomerId = v
	return r
}
func (r *SmsSignatureAddRequest) SetSbjScreenshotImg(v string) *SmsSignatureAddRequest {
	r.SbjScreenshotImg = v
	return r
}
func (r *SmsSignatureAddRequest) SetTrademarkAuthImg(v string) *SmsSignatureAddRequest {
	r.TrademarkAuthImg = v
	return r
}
func (r *SmsSignatureAddRequest) SetInstitutionImg(v string) *SmsSignatureAddRequest {
	r.InstitutionImg = v
	return r
}
func (r *SmsSignatureAddRequest) SetAdditionalAttachmentsImg(v []string) *SmsSignatureAddRequest {
	r.AdditionalAttachmentsImg = v
	return r
}

// SmsSignatureAddResult 创建签名结果。
type SmsSignatureAddResult struct {
	// SignId 签名 ID。
	SignId string `json:"signId"`
	// ExistedSignId 已存在的签名对应的签名 ID：签名已存在时返回。
	ExistedSignId string `json:"existedSignId"`
}

type SmsSignatureAddResponse struct {
	Code      string                 `json:"code"`
	Msg       string                 `json:"msg"`
	RequestId string                 `json:"requestId"`
	Data      *SmsSignatureAddResult `json:"data"`
}

func (r *SmsSignatureAddResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsSignatureListRequest 签名列表查询请求。
type SmsSignatureListRequest struct {
	ProductType
	// SignName 签名名称，支持模糊查询。
	SignName string `json:"signName,omitempty"`
	// SignatureStatus 平台审核签名状态：1-未审核；2-审核通过；3-审核驳回；-1-历史签名。
	SignatureStatus string `json:"signatureStatus,omitempty"`
	PageRequest
}

func (r *SmsSignatureListRequest) SetSignName(v string) *SmsSignatureListRequest {
	r.SignName = v
	return r
}
func (r *SmsSignatureListRequest) SetSignatureStatus(v string) *SmsSignatureListRequest {
	r.SignatureStatus = v
	return r
}

// SignatureItem 签名列表项。
type SignatureItem struct {
	SignId            string `json:"signId"`
	SignName          string `json:"signName"`
	SignatureStatus   string `json:"signatureStatus"`
	IsSystemSignature string `json:"isSystemSignature"`
	RealNameStatus    string `json:"realNameStatus"`
	CompanyName       string `json:"companyName"`
	CompanyLicenseNo  string `json:"companyLicenseNo"`
	LegalName         string `json:"legalName"`
	IdentityNo        string `json:"identityNo"`
	IdentityName      string `json:"identityName"`
	PhoneNum          string `json:"phoneNum"`
	EndCustomerId     string `json:"endCustomerId"`
	CmOperatorStatus  string `json:"cmOperatorStatus"`
	CtOperatorStatus  string `json:"ctOperatorStatus"`
	CuOperatorStatus  string `json:"cuOperatorStatus"`
}

type SmsSignatureListData struct {
	Total int              `json:"total"`
	List  []*SignatureItem `json:"list"`
}

type SmsSignatureListResponse struct {
	Code      string                `json:"code"`
	Msg       string                `json:"msg"`
	RequestId string                `json:"requestId"`
	Data      *SmsSignatureListData `json:"data"`
}

func (r *SmsSignatureListResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsSignatureGetRequest 签名详情查询请求。
type SmsSignatureGetRequest struct {
	ProductType
	// SignId 签名 ID。
	SignId string `json:"signId"`
}

func (r *SmsSignatureGetRequest) SetSignId(v string) *SmsSignatureGetRequest {
	r.SignId = v
	return r
}

// RejectReasonItem 运营商审核驳回原因。
type RejectReasonItem struct {
	// IssueType 问题类型：1-签名问题；2-签名材料问题；3-终端资质问题；4-其他；5-引流问题。
	IssueType int `json:"issueType"`
	// IssueDesc 问题描述。
	IssueDesc string `json:"issueDesc"`
	// OperationRecommend 操作建议。
	OperationRecommend string `json:"operationRecommend"`
	// ModifyModules 修改模块：1-新增终端资质…8-无需修改。
	ModifyModules string `json:"modifyModules"`
}

// SmsSignatureGetInfo 签名详情。
type SmsSignatureGetInfo struct {
	SignId                 string              `json:"signId"`
	SignName               string              `json:"signName"`
	SignatureStatus        string              `json:"signatureStatus"`
	IsSystemSignature      string              `json:"isSystemSignature"`
	AuditReason            string              `json:"auditReason"`
	RealNameStatus         string              `json:"realNameStatus"`
	CompanyName            string              `json:"companyName"`
	CompanyLicenseNo       string              `json:"companyLicenseNo"`
	LegalName              string              `json:"legalName"`
	IdentityNo             string              `json:"identityNo"`
	IdentityName           string              `json:"identityName"`
	PhoneNum               string              `json:"phoneNum"`
	EndCustomerId          string              `json:"endCustomerId"`
	CmOperatorStatus       string              `json:"cmOperatorStatus"`
	CtOperatorStatus       string              `json:"ctOperatorStatus"`
	CuOperatorStatus       string              `json:"cuOperatorStatus"`
	CmOperatorRejectReason []*RejectReasonItem `json:"cmOperatorRejectReason"`
	CtOperatorRejectReason []*RejectReasonItem `json:"ctOperatorRejectReason"`
	CuOperatorRejectReason []*RejectReasonItem `json:"cuOperatorRejectReason"`
}

type SmsSignatureGetResponse struct {
	Code      string               `json:"code"`
	Msg       string               `json:"msg"`
	RequestId string               `json:"requestId"`
	Data      *SmsSignatureGetInfo `json:"data"`
}

func (r *SmsSignatureGetResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsSignatureOperatorRejectReasonRequest 查询签名运营商驳回原因请求。
// 底层复用签名详情接口 detail。
type SmsSignatureOperatorRejectReasonRequest struct {
	ProductType
	SignId string `json:"signId"`
	// Operator 运营商：cm-移动；ct-电信；cu-联通。
	Operator string `json:"-"`
}

func (r *SmsSignatureOperatorRejectReasonRequest) SetSignId(v string) *SmsSignatureOperatorRejectReasonRequest {
	r.SignId = v
	return r
}
func (r *SmsSignatureOperatorRejectReasonRequest) SetOperator(v string) *SmsSignatureOperatorRejectReasonRequest {
	r.Operator = v
	return r
}

// SmsSignatureOperatorRejectReasonResponse 签名运营商驳回原因。
type SmsSignatureOperatorRejectReasonResponse struct {
	Code      string              `json:"code"`
	Msg       string              `json:"msg"`
	RequestId string              `json:"requestId"`
	Data      []*RejectReasonItem `json:"data"`
}

func (r *SmsSignatureOperatorRejectReasonResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsSignatureRealNameUpdateRequest 编辑签名终端资请求。
type SmsSignatureRealNameUpdateRequest struct {
	ProductType
	SignId        string `json:"signId"`
	EndCustomerId string `json:"endCustomerId"`
	// Type 修改类型：2-修改实名资质；3-修改应用场景。
	Type string `json:"type"`
	// SignatureSceneType 关联资质用途：3-自用；4-他人；type=3 时必填。
	SignatureSceneType         int    `json:"signatureSceneType,omitempty"`
	AppImg                     string `json:"appImg,omitempty"`
	AppOpenScreenshotImg       string `json:"appOpenScreenshotImg,omitempty"`
	SbjScreenshotImg           string `json:"sbjScreenshotImg,omitempty"`
	TrademarkImg               string `json:"trademarkImg,omitempty"`
	TrademarkAuthImg           string `json:"trademarkAuthImg,omitempty"`
	InstitutionImg             string `json:"institutionImg,omitempty"`
	CompanyUniqueScreenshotImg string `json:"companyUniqueScreenshotImg,omitempty"`
	// SignAuthImg 签名授权书（signatureSceneType=4 时必填）。
	SignAuthImg string `json:"signAuthImg,omitempty"`
	Remark      string `json:"remark,omitempty"`
}

func (r *SmsSignatureRealNameUpdateRequest) SetSignId(v string) *SmsSignatureRealNameUpdateRequest {
	r.SignId = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetEndCustomerId(v string) *SmsSignatureRealNameUpdateRequest {
	r.EndCustomerId = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetType(v string) *SmsSignatureRealNameUpdateRequest {
	r.Type = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetSignatureSceneType(v int) *SmsSignatureRealNameUpdateRequest {
	r.SignatureSceneType = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetAppImg(v string) *SmsSignatureRealNameUpdateRequest {
	r.AppImg = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetAppOpenScreenshotImg(v string) *SmsSignatureRealNameUpdateRequest {
	r.AppOpenScreenshotImg = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetSbjScreenshotImg(v string) *SmsSignatureRealNameUpdateRequest {
	r.SbjScreenshotImg = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetTrademarkImg(v string) *SmsSignatureRealNameUpdateRequest {
	r.TrademarkImg = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetTrademarkAuthImg(v string) *SmsSignatureRealNameUpdateRequest {
	r.TrademarkAuthImg = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetInstitutionImg(v string) *SmsSignatureRealNameUpdateRequest {
	r.InstitutionImg = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetCompanyUniqueScreenshotImg(v string) *SmsSignatureRealNameUpdateRequest {
	r.CompanyUniqueScreenshotImg = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetSignAuthImg(v string) *SmsSignatureRealNameUpdateRequest {
	r.SignAuthImg = v
	return r
}
func (r *SmsSignatureRealNameUpdateRequest) SetRemark(v string) *SmsSignatureRealNameUpdateRequest {
	r.Remark = v
	return r
}

type SmsSignatureRealNameUpdateResponse struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

func (r *SmsSignatureRealNameUpdateResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsSignatureDeleteRequest 删除签名请求（审核通过的签名不允许删除）。
type SmsSignatureDeleteRequest struct {
	ProductType
	SignId string `json:"signId"`
}

func (r *SmsSignatureDeleteRequest) SetSignId(v string) *SmsSignatureDeleteRequest {
	r.SignId = v
	return r
}

type SmsSignatureDeleteResponse struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

func (r *SmsSignatureDeleteResponse) IsSuccess() bool { return r.Code == codeSuccess }

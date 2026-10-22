package sms

// 模板相关请求与响应。对应 docs/sms/4-模版管理.md。

// SmsTemplateVariableParam 模板变量属性。
type SmsTemplateVariableParam struct {
	// Name 变量名称，例如变量 1、变量 2。
	Name string `json:"name,omitempty"`
	// Type 变量类型，从 QueryTemplateTypeEnum 获取。
	Type string `json:"type,omitempty"`
	// Length 变量长度，最大 30 位。
	Length string `json:"length,omitempty"`
	// Position 变量位置：第一个变量为 0，第二个为 1，依次递增。
	Position int `json:"position,omitempty"`
}

// SmsTemplateGuideUrl 模板引流链接。
type SmsTemplateGuideUrl struct {
	// GuideType 引流链接类型，从 QueryTemplateTypeEnum 获取。
	GuideType string `json:"guideType,omitempty"`
	// Content 引流链接内容，需与实际下发链接一致。
	Content string `json:"content,omitempty"`
	// DomainExample 完整链接示例，运营商审核需要。
	DomainExample string `json:"domainExample,omitempty"`
	// IcpImg 域名 ICP 备案截图（guideType=2 域名类型时必填，Base64，≤2MB）。
	IcpImg string `json:"icpImg,omitempty"`
	// OriginLinkAuthImg 原链接授权书（Base64，≤2MB）。
	OriginLinkAuthImg string `json:"originLinkAuthImg,omitempty"`
	// RedirectLinkAuthImg 跳转后链接授权书（Base64，≤2MB）。
	RedirectLinkAuthImg string `json:"redirectLinkAuthImg,omitempty"`
}

// SmsTemplateGuidePhone 模板联系方式。
type SmsTemplateGuidePhone struct {
	// GuideType 联系方式类型，从 QueryTemplateTypeEnum 获取。
	GuideType string `json:"guideType,omitempty"`
	// Content 联系方式内容（引流号码不要带空格或横杠）。
	Content string `json:"content,omitempty"`
	// PhoneMaterialImg 引流号码材料（Base64，≤2MB）。
	PhoneMaterialImg string `json:"phoneMaterialImg,omitempty"`
}

// SmsTemplateAddRequest 创建模板请求。
type SmsTemplateAddRequest struct {
	ProductType
	// Content 模板内容：不包含短信签名和退订语，长度不超过 500 字符；包含链接时需在链接前后各加一个空格；变量用 {s} 标识。
	Content string `json:"content"`
	// ContentName 模板名称，长度为 2～20 个字符。
	ContentName string `json:"contentName"`
	// SignName 关联签名名称。
	SignName string `json:"signName"`
	// BusinessCategory 业务大类，从 QueryTemplateTypeEnum 获取。
	BusinessCategory string `json:"businessCategory"`
	// BusinessSubcategory 业务细类，从 QueryTemplateTypeEnum 获取。
	BusinessSubcategory string `json:"businessSubcategory"`
	// VariableParams 变量属性：Object[] 序列化后的 JSON 字符串；运营商建议 3 个变量，最大 20 个。
	VariableParams string `json:"variableParams,omitempty"`
	// GuideUrl 引流链接：Object 序列化后的 JSON 字符串。
	GuideUrl string `json:"guideUrl,omitempty"`
	// GuidePhone 联系方式：Object[] 序列化后的 JSON 字符串。
	GuidePhone string `json:"guidePhone,omitempty"`
	// NeedUnsubscribe 退订语：0-不需要；1-需要（营销必填）。
	NeedUnsubscribe string `json:"needUnsubscribe,omitempty"`
	// Remark 备注，有利于加速审核。
	Remark string `json:"remark,omitempty"`
}

func (r *SmsTemplateAddRequest) SetContent(v string) *SmsTemplateAddRequest {
	r.Content = v
	return r
}
func (r *SmsTemplateAddRequest) SetContentName(v string) *SmsTemplateAddRequest {
	r.ContentName = v
	return r
}
func (r *SmsTemplateAddRequest) SetSignName(v string) *SmsTemplateAddRequest {
	r.SignName = v
	return r
}
func (r *SmsTemplateAddRequest) SetBusinessCategory(v string) *SmsTemplateAddRequest {
	r.BusinessCategory = v
	return r
}
func (r *SmsTemplateAddRequest) SetBusinessSubcategory(v string) *SmsTemplateAddRequest {
	r.BusinessSubcategory = v
	return r
}
func (r *SmsTemplateAddRequest) SetVariableParams(v string) *SmsTemplateAddRequest {
	r.VariableParams = v
	return r
}
func (r *SmsTemplateAddRequest) SetGuideUrl(v string) *SmsTemplateAddRequest {
	r.GuideUrl = v
	return r
}
func (r *SmsTemplateAddRequest) SetGuidePhone(v string) *SmsTemplateAddRequest {
	r.GuidePhone = v
	return r
}
func (r *SmsTemplateAddRequest) SetNeedUnsubscribe(v string) *SmsTemplateAddRequest {
	r.NeedUnsubscribe = v
	return r
}
func (r *SmsTemplateAddRequest) SetRemark(v string) *SmsTemplateAddRequest {
	r.Remark = v
	return r
}

type SmsTemplateAddResult struct {
	// TemplateCode 模板 ID。
	TemplateCode string `json:"templateCode"`
}

type SmsTemplateAddResponse struct {
	Code      string                `json:"code"`
	Msg       string                `json:"msg"`
	RequestId string                `json:"requestId"`
	Data      *SmsTemplateAddResult `json:"data"`
}

func (r *SmsTemplateAddResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsTemplateTypeEnumRequest 查询模板类型枚举请求。
type SmsTemplateTypeEnumRequest struct {
	ProductType
}

// EnumOption 枚举选项。
type EnumOption struct {
	// Id 类型 ID。
	Id string `json:"id"`
	// Type 类型名称。
	Type string `json:"type"`
}

// EnumBusinessType 业务类型（含明细）。
type EnumBusinessType struct {
	Id      string        `json:"id"`
	Type    string        `json:"type"`
	Details []*EnumOption `json:"details"`
}

// SmsTemplateTypeEnumData 模板类型枚举数据。
type SmsTemplateTypeEnumData struct {
	VariableTypes []*EnumOption       `json:"variableTypes"`
	MobileTypes   []*EnumOption       `json:"mobileTypes"`
	UrlTypes      []*EnumOption       `json:"urlTypes"`
	BusinessTypes []*EnumBusinessType `json:"businessTypes"`
}

type SmsTemplateTypeEnumResponse struct {
	Code      string                   `json:"code"`
	Msg       string                   `json:"msg"`
	RequestId string                   `json:"requestId"`
	Data      *SmsTemplateTypeEnumData `json:"data"`
}

func (r *SmsTemplateTypeEnumResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsTemplateListRequest 模板列表查询请求。
type SmsTemplateListRequest struct {
	ProductType
	// PageNo 起始值：传 0 代表从第一个查起，默认 0。
	PageNo int `json:"pageNo,omitempty"`
	// PageSize 查询条数：默认 10，最大 100。
	PageSize int `json:"pageSize,omitempty"`
	// SignName 关联签名名称，精确匹配。
	SignName string `json:"signName,omitempty"`
	// ContentName 模板名称，模糊查询。
	ContentName string `json:"contentName,omitempty"`
}

func (r *SmsTemplateListRequest) SetSignName(v string) *SmsTemplateListRequest {
	r.SignName = v
	return r
}
func (r *SmsTemplateListRequest) SetContentName(v string) *SmsTemplateListRequest {
	r.ContentName = v
	return r
}
func (r *SmsTemplateListRequest) SetPageNo(v int) *SmsTemplateListRequest {
	r.PageNo = v
	return r
}
func (r *SmsTemplateListRequest) SetPageSize(v int) *SmsTemplateListRequest {
	r.PageSize = v
	return r
}

// TemplateItem 模板列表项。
type TemplateItem struct {
	TemplateCode     string `json:"templateCode"`
	WebOrApi         string `json:"webOrApi"`
	AccountId        string `json:"accountId"`
	Type             string `json:"type"`
	ContentTask      string `json:"contentTask"`
	Content          string `json:"content"`
	NeedUnsubscribe  string `json:"needUnsubscribe"`
	Remark           string `json:"remark"`
	CreateTime       string `json:"createTime"`
	Status           string `json:"status"`
	AuditDate        string `json:"auditDate"`
	AuditReason      string `json:"auditReason"`
	SignName         string `json:"signName"`
	ContentName      string `json:"contentName"`
	CmOperatorStatus string `json:"cmOperatorStatus"`
	CtOperatorStatus string `json:"ctOperatorStatus"`
	CuOperatorStatus string `json:"cuOperatorStatus"`
}

type SmsTemplateListData struct {
	Total int             `json:"total"`
	List  []*TemplateItem `json:"list"`
}

type SmsTemplateListResponse struct {
	Code      string               `json:"code"`
	Msg       string               `json:"msg"`
	RequestId string               `json:"requestId"`
	Data      *SmsTemplateListData `json:"data"`
}

func (r *SmsTemplateListResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsTemplateGetRequest 模板详情查询请求。
type SmsTemplateGetRequest struct {
	ProductType
	// TemplateCode 模板 ID。
	TemplateCode string `json:"templateCode"`
}

func (r *SmsTemplateGetRequest) SetTemplateCode(v string) *SmsTemplateGetRequest {
	r.TemplateCode = v
	return r
}

// SmsTemplateGetInfo 模板详情。
type SmsTemplateGetInfo struct {
	TemplateCode           string                      `json:"templateCode"`
	WebOrApi               string                      `json:"webOrApi"`
	AccountId              string                      `json:"accountId"`
	AuditDate              string                      `json:"auditDate"`
	AuditReason            string                      `json:"auditReason"`
	BusinessCategory       string                      `json:"businessCategory"`
	BusinessSubcategory    string                      `json:"businessSubcategory"`
	Content                string                      `json:"content"`
	ContentName            string                      `json:"contentName"`
	ContentTask            string                      `json:"contentTask"`
	CreateTime             string                      `json:"createTime"`
	ModelId                string                      `json:"modelId"`
	NeedUnsubscribe        string                      `json:"needUnsubscribe"`
	Remark                 string                      `json:"remark"`
	RemoteId               string                      `json:"remoteId"`
	SignName               string                      `json:"signName"`
	Signature              string                      `json:"signature"`
	Status                 string                      `json:"status"`
	StatusDesc             string                      `json:"statusDesc"`
	Type                   string                      `json:"type"`
	Unsubscribe            string                      `json:"unsubscribe"`
	GuidePhone             []*SmsTemplateGuidePhone    `json:"guidePhone"`
	VariableParams         []*SmsTemplateVariableParam `json:"variableParams"`
	GuideUrl               *SmsTemplateGuideUrl        `json:"guideUrl"`
	CmOperatorStatus       string                      `json:"cmOperatorStatus"`
	CtOperatorStatus       string                      `json:"ctOperatorStatus"`
	CuOperatorStatus       string                      `json:"cuOperatorStatus"`
	CmOperatorRejectReason []*RejectReasonItem         `json:"cmOperatorRejectReason"`
	CtOperatorRejectReason []*RejectReasonItem         `json:"ctOperatorRejectReason"`
	CuOperatorRejectReason []*RejectReasonItem         `json:"cuOperatorRejectReason"`
}

type SmsTemplateGetResponse struct {
	Code      string              `json:"code"`
	Msg       string              `json:"msg"`
	RequestId string              `json:"requestId"`
	Data      *SmsTemplateGetInfo `json:"data"`
}

func (r *SmsTemplateGetResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsTemplateOperatorRejectReasonRequest 查询模板运营商驳回原因请求。
// 底层复用模板详情接口 detail。
type SmsTemplateOperatorRejectReasonRequest struct {
	ProductType
	TemplateCode string `json:"templateCode"`
	// Operator 运营商：cm-移动；ct-电信；cu-联通。
	Operator string `json:"-"`
}

func (r *SmsTemplateOperatorRejectReasonRequest) SetTemplateCode(v string) *SmsTemplateOperatorRejectReasonRequest {
	r.TemplateCode = v
	return r
}
func (r *SmsTemplateOperatorRejectReasonRequest) SetOperator(v string) *SmsTemplateOperatorRejectReasonRequest {
	r.Operator = v
	return r
}

type SmsTemplateOperatorRejectReasonResponse struct {
	Code      string              `json:"code"`
	Msg       string              `json:"msg"`
	RequestId string              `json:"requestId"`
	Data      []*RejectReasonItem `json:"data"`
}

func (r *SmsTemplateOperatorRejectReasonResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsTemplateUpdateRequest 编辑模板请求。
type SmsTemplateUpdateRequest struct {
	ProductType
	TemplateCode        string `json:"templateCode"`
	Content             string `json:"content"`
	ContentName         string `json:"contentName"`
	SignName            string `json:"signName"`
	BusinessCategory    string `json:"businessCategory"`
	BusinessSubcategory string `json:"businessSubcategory"`
	VariableParams      string `json:"variableParams,omitempty"`
	GuideUrl            string `json:"guideUrl,omitempty"`
	GuidePhone          string `json:"guidePhone,omitempty"`
	NeedUnsubscribe     string `json:"needUnsubscribe,omitempty"`
	Remark              string `json:"remark,omitempty"`
}

func (r *SmsTemplateUpdateRequest) SetTemplateCode(v string) *SmsTemplateUpdateRequest {
	r.TemplateCode = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetContent(v string) *SmsTemplateUpdateRequest {
	r.Content = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetContentName(v string) *SmsTemplateUpdateRequest {
	r.ContentName = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetSignName(v string) *SmsTemplateUpdateRequest {
	r.SignName = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetBusinessCategory(v string) *SmsTemplateUpdateRequest {
	r.BusinessCategory = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetBusinessSubcategory(v string) *SmsTemplateUpdateRequest {
	r.BusinessSubcategory = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetVariableParams(v string) *SmsTemplateUpdateRequest {
	r.VariableParams = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetGuideUrl(v string) *SmsTemplateUpdateRequest {
	r.GuideUrl = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetGuidePhone(v string) *SmsTemplateUpdateRequest {
	r.GuidePhone = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetNeedUnsubscribe(v string) *SmsTemplateUpdateRequest {
	r.NeedUnsubscribe = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetRemark(v string) *SmsTemplateUpdateRequest {
	r.Remark = v
	return r
}

type SmsTemplateUpdateResponse struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

func (r *SmsTemplateUpdateResponse) IsSuccess() bool { return r.Code == codeSuccess }

// SmsTemplateDeleteRequest 删除模板请求。
type SmsTemplateDeleteRequest struct {
	ProductType
	TemplateCode string `json:"templateCode"`
}

func (r *SmsTemplateDeleteRequest) SetTemplateCode(v string) *SmsTemplateDeleteRequest {
	r.TemplateCode = v
	return r
}

type SmsTemplateDeleteResponse struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

func (r *SmsTemplateDeleteResponse) IsSuccess() bool { return r.Code == codeSuccess }

// Package rcs 提供视频短信(RCS)业务封装。
// 统一使用 /videosms/api/v2 前缀（文档见 docs/video-sms）。
package rcs

// codeSuccess 表示 RCS 业务请求成功（与短信业务相同）。
const codeSuccess = "000000"

// RcsMaterial 素材参数（SignMaterialParam）。
type RcsMaterial struct {
	// Source 素材来源：1-base64（默认）；2-url。
	Source string `json:"source,omitempty"`
	// Type 素材格式，如 jpg、png、pdf。
	Type string `json:"type"`
	// Content 素材内容：base64 字符串或 url。
	Content string `json:"content"`
}

func (r *RcsMaterial) SetSource(v string) *RcsMaterial { r.Source = v; return r }
func (r *RcsMaterial) SetType(v string) *RcsMaterial   { r.Type = v; return r }
func (r *RcsMaterial) SetContent(v string) *RcsMaterial {
	r.Content = v
	return r
}

// RcsSignAddRequest 创建签名请求。对应 docs/video-sms/3-签名管理.md。
type RcsSignAddRequest struct {
	// Sign 签名内容（可不带【】，系统会自动添加）。
	Sign string `json:"sign"`
	// CustomerCode 所属客户编码：1 党政军 2 民生 3 医疗器械/药店 4 电商 5 沿街商铺(中小) 6 教育培训 7 房地产 8 游戏。
	CustomerCode string `json:"customerCode"`
	// IndustryCode 行业属性编码。
	IndustryCode string `json:"industryCode"`
	// SignCategory 签名分类：1 企业名称 2 app 3 商标 4 其他(如政企)。
	SignCategory string `json:"signCategory"`
	// SignType 签名类型：0 非商标类；1 商标类。
	SignType string `json:"signType"`
	// BusinessName 企业名称。
	BusinessName string `json:"businessName,omitempty"`
	// CreditCode 统一社会信用代码。
	CreditCode string `json:"creditCode,omitempty"`
	// HandledbyName 经办人姓名。
	HandledbyName string `json:"handledbyName,omitempty"`
	// RelevantIdcard 经办人身份证号。
	RelevantIdcard string `json:"relevantIdcard,omitempty"`
	// Body 签名素材体。
	Body *RcsSignBody `json:"body"`
}

func (r *RcsSignAddRequest) SetSign(v string) *RcsSignAddRequest { r.Sign = v; return r }
func (r *RcsSignAddRequest) SetCustomerCode(v string) *RcsSignAddRequest {
	r.CustomerCode = v
	return r
}
func (r *RcsSignAddRequest) SetIndustryCode(v string) *RcsSignAddRequest {
	r.IndustryCode = v
	return r
}
func (r *RcsSignAddRequest) SetSignCategory(v string) *RcsSignAddRequest {
	r.SignCategory = v
	return r
}
func (r *RcsSignAddRequest) SetSignType(v string) *RcsSignAddRequest { r.SignType = v; return r }
func (r *RcsSignAddRequest) SetBusinessName(v string) *RcsSignAddRequest {
	r.BusinessName = v
	return r
}
func (r *RcsSignAddRequest) SetCreditCode(v string) *RcsSignAddRequest { r.CreditCode = v; return r }
func (r *RcsSignAddRequest) SetHandledbyName(v string) *RcsSignAddRequest {
	r.HandledbyName = v
	return r
}
func (r *RcsSignAddRequest) SetRelevantIdcard(v string) *RcsSignAddRequest {
	r.RelevantIdcard = v
	return r
}
func (r *RcsSignAddRequest) SetBody(v *RcsSignBody) *RcsSignAddRequest { r.Body = v; return r }

// RcsSignBody 签名素材体。
type RcsSignBody struct {
	// EntrustFile 委托书文件（最多 1 个，必填）。
	EntrustFile *RcsMaterial `json:"entrustFile"`
	// JobPermitFile 营业执照文件（最多 1 个，必填）。
	JobPermitFile *RcsMaterial `json:"jobPermitFile"`
	// AuthorizeFile 商标授权文件列表（最多 5 个，选填）。
	AuthorizeFile []*RcsMaterial `json:"authorizeFile,omitempty"`
	// SoftnessFile 商标注册证文件列表（最多 5 个，signType=1 时必传）。
	SoftnessFile []*RcsMaterial `json:"softnessFile,omitempty"`
}

func (r *RcsSignBody) SetEntrustFile(v *RcsMaterial) *RcsSignBody { r.EntrustFile = v; return r }
func (r *RcsSignBody) SetJobPermitFile(v *RcsMaterial) *RcsSignBody {
	r.JobPermitFile = v
	return r
}
func (r *RcsSignBody) SetAuthorizeFile(v []*RcsMaterial) *RcsSignBody {
	r.AuthorizeFile = v
	return r
}
func (r *RcsSignBody) SetSoftnessFile(v []*RcsMaterial) *RcsSignBody {
	r.SoftnessFile = v
	return r
}

type RcsSignAddResult struct {
	SignId string `json:"signId"`
}

type RcsSignAddResponse struct {
	Code      string            `json:"code"`
	Msg       string            `json:"msg"`
	RequestId string            `json:"requestId"`
	Data      *RcsSignAddResult `json:"data"`
}

func (r *RcsSignAddResponse) IsSuccess() bool { return r.Code == codeSuccess }

// RcsTemplateBodyItem 视频模板内容项。
type RcsTemplateBodyItem struct {
	// Type 媒体类型：text / video / audio / image。
	Type string `json:"type"`
	// ExType 后缀格式，与 type 配合：text:txt video:mp4/3gp audio:mp3 image:jpg/gif/png。
	ExType string `json:"exType"`
	// Content 文本内容；文件素材时可传 base64 或 url，由 ContentType 指定。
	Content string `json:"content"`
	// ContentType content 类型（文本忽略）：1-base64；2-文件素材 url。默认 1。
	ContentType int `json:"contentType,omitempty"`
	// Sort 内容排序。
	Sort int `json:"sort"`
}

func (r *RcsTemplateBodyItem) SetType(v string) *RcsTemplateBodyItem     { r.Type = v; return r }
func (r *RcsTemplateBodyItem) SetExType(v string) *RcsTemplateBodyItem   { r.ExType = v; return r }
func (r *RcsTemplateBodyItem) SetContent(v string) *RcsTemplateBodyItem  { r.Content = v; return r }
func (r *RcsTemplateBodyItem) SetContentType(v int) *RcsTemplateBodyItem { r.ContentType = v; return r }
func (r *RcsTemplateBodyItem) SetSort(v int) *RcsTemplateBodyItem        { r.Sort = v; return r }

// RcsTemplateAddRequest 创建视频短信模板请求。
// 变量模板仅文本可添加变量，最多 5 个：${v1}~${v5}。对应 docs/video-sms/4-模版管理.md 第 1 节。
type RcsTemplateAddRequest struct {
	// TemplateName 模板名称。
	TemplateName string `json:"templateName"`
	// Sign 已在平台创建的签名。
	Sign string `json:"sign"`
	// Body 模板内容列表。
	Body []*RcsTemplateBodyItem `json:"body"`
	// IsDynamic 动态模板标记，创建动态模板时必填（1）。
	IsDynamic int `json:"isDynamic,omitempty"`
	// TemplateLink 模板链接完整示例。
	TemplateLink string `json:"templateLink,omitempty"`
}

func (r *RcsTemplateAddRequest) SetTemplateName(v string) *RcsTemplateAddRequest {
	r.TemplateName = v
	return r
}
func (r *RcsTemplateAddRequest) SetSign(v string) *RcsTemplateAddRequest { r.Sign = v; return r }
func (r *RcsTemplateAddRequest) SetBody(v []*RcsTemplateBodyItem) *RcsTemplateAddRequest {
	r.Body = v
	return r
}
func (r *RcsTemplateAddRequest) SetIsDynamic(v int) *RcsTemplateAddRequest {
	r.IsDynamic = v
	return r
}
func (r *RcsTemplateAddRequest) SetTemplateLink(v string) *RcsTemplateAddRequest {
	r.TemplateLink = v
	return r
}

type RcsTemplateAddResult struct {
	TemplateId string `json:"templateId"`
}

type RcsTemplateAddResponse struct {
	Code      string                `json:"code"`
	Msg       string                `json:"msg"`
	RequestId string                `json:"requestId"`
	Data      *RcsTemplateAddResult `json:"data"`
}

func (r *RcsTemplateAddResponse) IsSuccess() bool { return r.Code == codeSuccess }

// RcsTemplateFindRequest 查看视频模板状态请求。
type RcsTemplateFindRequest struct {
	// TemplateId 模板 id。
	TemplateId string `json:"templateId"`
}

func (r *RcsTemplateFindRequest) SetTemplateId(v string) *RcsTemplateFindRequest {
	r.TemplateId = v
	return r
}

// RcsTemplateFindInfo 视频模板状态信息。
type RcsTemplateFindInfo struct {
	// StatusName 模板状态说明。
	StatusName string `json:"statusName"`
	// Status 模板状态：1 平台审核中 2 运营商审核中 3 审核成功 4 审核失败。
	Status int `json:"status"`
	// TemplateId 模板 id。
	TemplateId string `json:"templateId"`
	// Operator 模板可发运营商标识：1 移动 2 联通 3 电信。
	Operator []string `json:"operator"`
	// RejectReason 审核驳回原因。
	RejectReason string `json:"rejectReason"`
	// ExpireFlag 模板有效性：0 失效 1 有效。
	ExpireFlag string `json:"expireFlag"`
	// IsDynamic 动参或非动参标识：0 非动参 1 动参。
	IsDynamic string `json:"isDynamic"`
}

type RcsTemplateFindResponse struct {
	Code      string               `json:"code"`
	Msg       string               `json:"msg"`
	RequestId string               `json:"requestId"`
	Data      *RcsTemplateFindInfo `json:"data"`
}

func (r *RcsTemplateFindResponse) IsSuccess() bool { return r.Code == codeSuccess }

// RcsSubmitDynamic 动态模板变量。
type RcsSubmitDynamic struct {
	// Phone 手机号。
	Phone string `json:"phone"`
	// V1~V5 变量值，对应模板中的 ${v1}~${v5}。
	V1 string `json:"v1,omitempty"`
	V2 string `json:"v2,omitempty"`
	V3 string `json:"v3,omitempty"`
	V4 string `json:"v4,omitempty"`
	V5 string `json:"v5,omitempty"`
}

func (r *RcsSubmitDynamic) SetPhone(v string) *RcsSubmitDynamic { r.Phone = v; return r }
func (r *RcsSubmitDynamic) SetV1(v string) *RcsSubmitDynamic    { r.V1 = v; return r }
func (r *RcsSubmitDynamic) SetV2(v string) *RcsSubmitDynamic    { r.V2 = v; return r }
func (r *RcsSubmitDynamic) SetV3(v string) *RcsSubmitDynamic    { r.V3 = v; return r }
func (r *RcsSubmitDynamic) SetV4(v string) *RcsSubmitDynamic    { r.V4 = v; return r }
func (r *RcsSubmitDynamic) SetV5(v string) *RcsSubmitDynamic    { r.V5 = v; return r }

// RcsTemplateSubmitRequest 发送视频短信请求。对应 docs/video-sms/2-短信发送.md。
// 静态模板传 PhoneNumbers，动态模板传 PhoneNumberJson，二者取一。
type RcsTemplateSubmitRequest struct {
	// SubmitNo 提交号，客户端提供，长度不超过 32 位。
	SubmitNo string `json:"submitNo"`
	// PhoneNumbers 手机号列表，静态模板发送时必填。
	PhoneNumbers []string `json:"phoneNumbers,omitempty"`
	// PhoneNumberJson 动态模板变量列表，动态模板发送时必填。
	PhoneNumberJson []*RcsSubmitDynamic `json:"phoneNumberJson,omitempty"`
	// TemplateId 模板 id。
	TemplateId string `json:"templateId"`
}

func (r *RcsTemplateSubmitRequest) SetSubmitNo(v string) *RcsTemplateSubmitRequest {
	r.SubmitNo = v
	return r
}
func (r *RcsTemplateSubmitRequest) SetPhoneNumbers(v []string) *RcsTemplateSubmitRequest {
	r.PhoneNumbers = v
	return r
}
func (r *RcsTemplateSubmitRequest) SetPhoneNumberJson(v []*RcsSubmitDynamic) *RcsTemplateSubmitRequest {
	r.PhoneNumberJson = v
	return r
}
func (r *RcsTemplateSubmitRequest) SetTemplateId(v string) *RcsTemplateSubmitRequest {
	r.TemplateId = v
	return r
}

// RcsTemplateSubmitResult 发送结果。
type RcsTemplateSubmitResult struct {
	// ErrPhone 提交失败的号码。
	ErrPhone []string `json:"errPhone"`
	// SubmitNo 提交号。
	SubmitNo string `json:"submitNo"`
	// BatchId 批次号，一次提交算一个。
	BatchId string `json:"batchId"`
}

type RcsTemplateSubmitResponse struct {
	Code      string                   `json:"code"`
	Msg       string                   `json:"msg"`
	RequestId string                   `json:"requestId"`
	Data      *RcsTemplateSubmitResult `json:"data"`
}

func (r *RcsTemplateSubmitResponse) IsSuccess() bool { return r.Code == codeSuccess }

// RcsReportItem 状态报告项。
type RcsReportItem struct {
	Phone          string `json:"phone"`
	BatchId        string `json:"batchId"`
	SubmitNo       string `json:"submitNo"`
	DeliveryStatus string `json:"deliveryStatus"`
	CreateTime     int64  `json:"createTime"`
	ClErrorCode    string `json:"clErrorCode"`
}

// RcsReportPullRequest 拉取状态报告请求。
// 无业务请求参数，Body 传空 JSON 对象 {}。状态报告保存 72 小时，拉取成功后删除。
type RcsReportPullRequest struct{}

type RcsReportPullResponse struct {
	Code      string           `json:"code"`
	Msg       string           `json:"msg"`
	RequestId string           `json:"requestId"`
	Data      []*RcsReportItem `json:"data"`
}

func (r *RcsReportPullResponse) IsSuccess() bool { return r.Code == codeSuccess }

// RcsReplyItem 上行消息项。
type RcsReplyItem struct {
	Phone         string `json:"phone"`
	CodeNoAddress string `json:"codeNoAddress"`
	MsgId         string `json:"msgId"`
	UpContent     string `json:"upContent"`
	CreateTime    int64  `json:"createTime"`
	BatchId       string `json:"batchId"`
	SubmitNo      string `json:"submitNo"`
}

// RcsReplyPullRequest 拉取上行消息请求。无业务请求参数，Body 传 {}。
// 上行消息保存 72 小时，拉取成功后删除。
type RcsReplyPullRequest struct{}

type RcsReplyPullResponse struct {
	Code      string          `json:"code"`
	Msg       string          `json:"msg"`
	RequestId string          `json:"requestId"`
	Data      []*RcsReplyItem `json:"data"`
}

func (r *RcsReplyPullResponse) IsSuccess() bool { return r.Code == codeSuccess }

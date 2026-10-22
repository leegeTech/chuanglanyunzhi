package sms

// 模板相关请求与响应类型。

// SmsTemplateAddRequest 添加模板请求。
type SmsTemplateAddRequest struct {
	TemplateName    string `json:"templateName,omitempty"`
	TemplateType    string `json:"templateType,omitempty"` // 1=验证码 2=通知 3=营销
	TemplateContent string `json:"templateContent,omitempty"`
	SignName        string `json:"signName,omitempty"`
}

func (r *SmsTemplateAddRequest) SetTemplateName(v string) *SmsTemplateAddRequest {
	r.TemplateName = v
	return r
}
func (r *SmsTemplateAddRequest) SetTemplateType(v string) *SmsTemplateAddRequest {
	r.TemplateType = v
	return r
}
func (r *SmsTemplateAddRequest) SetTemplateContent(v string) *SmsTemplateAddRequest {
	r.TemplateContent = v
	return r
}
func (r *SmsTemplateAddRequest) SetSignName(v string) *SmsTemplateAddRequest {
	r.SignName = v
	return r
}

type SmsTemplateAddResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		TemplateId string `json:"templateId"`
	} `json:"data"`
}

func (r *SmsTemplateAddResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsTemplateTypeEnumRequest 查询模板类型枚举请求。
type SmsTemplateTypeEnumRequest struct {
}

// EnumItem 枚举项。
type EnumItem struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type SmsTemplateTypeEnumResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data []*EnumItem `json:"data"`
}

func (r *SmsTemplateTypeEnumResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsTemplateListRequest 查询模板列表请求。
type SmsTemplateListRequest struct {
	TemplateName string `json:"templateName,omitempty"`
	Status       string `json:"status,omitempty"`
	PageNum      int    `json:"pageNum,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
}

func (r *SmsTemplateListRequest) SetTemplateName(v string) *SmsTemplateListRequest {
	r.TemplateName = v
	return r
}
func (r *SmsTemplateListRequest) SetStatus(v string) *SmsTemplateListRequest {
	r.Status = v
	return r
}
func (r *SmsTemplateListRequest) SetPageNum(v int) *SmsTemplateListRequest {
	r.PageNum = v
	return r
}
func (r *SmsTemplateListRequest) SetPageSize(v int) *SmsTemplateListRequest {
	r.PageSize = v
	return r
}

// TemplateItem 模板列表项。
type TemplateItem struct {
	TemplateId      string `json:"templateId"`
	TemplateName    string `json:"templateName"`
	TemplateType    string `json:"templateType"`
	TemplateContent string `json:"templateContent"`
	SignName        string `json:"signName"`
	Status          string `json:"status"`
}

type SmsTemplateListResponse struct {
	Code string               `json:"code"`
	Msg  string               `json:"msg"`
	Data *SmsTemplateListData `json:"data"`
}

type SmsTemplateListData struct {
	List     []*TemplateItem `json:"list"`
	Total    int             `json:"total"`
	PageNum  int             `json:"pageNum"`
	PageSize int             `json:"pageSize"`
}

func (r *SmsTemplateListResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsTemplateGetRequest 查询模板详情请求。
type SmsTemplateGetRequest struct {
	TemplateId string `json:"templateId,omitempty"`
}

func (r *SmsTemplateGetRequest) SetTemplateId(v string) *SmsTemplateGetRequest {
	r.TemplateId = v
	return r
}

type SmsTemplateGetResponse struct {
	Code string        `json:"code"`
	Msg  string        `json:"msg"`
	Data *TemplateItem `json:"data"`
}

func (r *SmsTemplateGetResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsTemplateOperatorRejectReasonRequest 查询模板运营商驳回原因请求。
type SmsTemplateOperatorRejectReasonRequest struct {
	TemplateId string `json:"templateId,omitempty"`
	ReasonType string `json:"reasonType,omitempty"`
}

func (r *SmsTemplateOperatorRejectReasonRequest) SetTemplateId(v string) *SmsTemplateOperatorRejectReasonRequest {
	r.TemplateId = v
	return r
}
func (r *SmsTemplateOperatorRejectReasonRequest) SetReasonType(v string) *SmsTemplateOperatorRejectReasonRequest {
	r.ReasonType = v
	return r
}

type SmsTemplateOperatorRejectReasonResponse struct {
	Code string              `json:"code"`
	Msg  string              `json:"msg"`
	Data []*RejectReasonItem `json:"data"`
}

func (r *SmsTemplateOperatorRejectReasonResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsTemplateUpdateRequest 更新模板请求。
type SmsTemplateUpdateRequest struct {
	TemplateId      string `json:"templateId,omitempty"`
	TemplateName    string `json:"templateName,omitempty"`
	TemplateType    string `json:"templateType,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	SignName        string `json:"signName,omitempty"`
}

func (r *SmsTemplateUpdateRequest) SetTemplateId(v string) *SmsTemplateUpdateRequest {
	r.TemplateId = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetTemplateName(v string) *SmsTemplateUpdateRequest {
	r.TemplateName = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetTemplateType(v string) *SmsTemplateUpdateRequest {
	r.TemplateType = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetTemplateContent(v string) *SmsTemplateUpdateRequest {
	r.TemplateContent = v
	return r
}
func (r *SmsTemplateUpdateRequest) SetSignName(v string) *SmsTemplateUpdateRequest {
	r.SignName = v
	return r
}

type SmsTemplateUpdateResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (r *SmsTemplateUpdateResponse) IsSuccess() bool { return r.Code == "000000" }

// SmsTemplateDeleteRequest 删除模板请求。
type SmsTemplateDeleteRequest struct {
	TemplateId string `json:"templateId,omitempty"`
}

func (r *SmsTemplateDeleteRequest) SetTemplateId(v string) *SmsTemplateDeleteRequest {
	r.TemplateId = v
	return r
}

type SmsTemplateDeleteResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (r *SmsTemplateDeleteResponse) IsSuccess() bool { return r.Code == "000000" }

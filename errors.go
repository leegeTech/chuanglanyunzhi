package cloudsdk

import "fmt"

// 成功业务错误码。
const CodeSuccess = "000000"

// 常见业务错误码。
const (
	CodeParameterMissing    = "ParameterMissing"
	CodeInvalidSignature    = "InvalidSignature"
	CodeInsufficientBalance = "InsufficientBalance"
	CodeTemplateNotFound    = "TemplateNotFound"
	CodeRateLimitExceeded   = "RateLimitExceeded"
	CodeHTTPError           = "HttpError"
)

// SdkError 对应 Java SDK 中的 CloudSdkException。
type SdkError struct {
	// Code 业务错误码，如 "000000" 表示成功。
	Code string
	// Message 错误描述。
	Message string
	// RequestID 请求 ID，用于问题排查。
	RequestID string
	// StatusCode HTTP 状态码。
	StatusCode int
	// Cause 原始错误（如有）。
	Cause error
}

func (e *SdkError) Error() string {
	return fmt.Sprintf("cloud-sdk error: code=%s message=%s requestId=%s statusCode=%d",
		e.Code, e.Message, e.RequestID, e.StatusCode)
}

// Unwrap 支持 errors.Is/As 解包原始错误。
func (e *SdkError) Unwrap() error { return e.Cause }

// NewCommonError 构造通用业务/传输错误。
func NewCommonError(code, msg string, statusCode int, cause error) *SdkError {
	return &SdkError{Code: code, Message: msg, StatusCode: statusCode, Cause: cause}
}

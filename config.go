package cloudsdk

import "time"

// 默认官方 endpoint
const (
	DefaultSmsEndpoint = "https://smssh.253.com"
	DefaultRcsEndpoint = "https://rcs.253.com"
)

// CloudApiConfig 承载 SDK 的全局配置。
type CloudApiConfig struct {
	// SmsEndpoint 短信业务 endpoint。
	SmsEndpoint string
	// RcsEndpoint 视频短信(RCS)业务 endpoint。
	RcsEndpoint string

	// ConnectTimeout 连接超时，默认 10s。
	ConnectTimeout time.Duration
	// ReadTimeout 读取超时，默认 10s。
	ReadTimeout time.Duration

	// MaxRetries 最大重试次数（不含首次请求），默认 3。
	MaxRetries int
	// RetryBaseIntervalMs 指数退避初始间隔(毫秒)，退避算法 2^n*base，默认 100。
	RetryBaseIntervalMs int
}

// NewCloudApiConfig 返回一份带默认值的配置。
func NewCloudApiConfig() *CloudApiConfig {
	return &CloudApiConfig{
		SmsEndpoint:         DefaultSmsEndpoint,
		RcsEndpoint:         DefaultRcsEndpoint,
		ConnectTimeout:      10 * time.Second,
		ReadTimeout:         10 * time.Second,
		MaxRetries:          3,
		RetryBaseIntervalMs: 100,
	}
}

// SmsEndpoint 设置短信 endpoint。
func (c *CloudApiConfig) SetSmsEndpoint(v string) *CloudApiConfig {
	c.SmsEndpoint = v
	return c
}

// RcsEndpoint 设置视频短信 endpoint。
func (c *CloudApiConfig) SetRcsEndpoint(v string) *CloudApiConfig {
	c.RcsEndpoint = v
	return c
}

// ConnectTimeout 设置连接超时。
func (c *CloudApiConfig) SetConnectTimeout(v time.Duration) *CloudApiConfig {
	c.ConnectTimeout = v
	return c
}

// ReadTimeout 设置读取超时。
func (c *CloudApiConfig) SetReadTimeout(v time.Duration) *CloudApiConfig {
	c.ReadTimeout = v
	return c
}

// MaxRetries 设置最大重试次数。
func (c *CloudApiConfig) SetMaxRetries(v int) *CloudApiConfig {
	c.MaxRetries = v
	return c
}

// RetryBaseIntervalMs 设置指数退避初始间隔(毫秒)。
func (c *CloudApiConfig) SetRetryBaseIntervalMs(v int) *CloudApiConfig {
	c.RetryBaseIntervalMs = v
	return c
}

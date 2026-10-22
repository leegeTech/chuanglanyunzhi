// Package cloudsdk 是创蓝统一应用的 Go SDK。
// 提供短信、视频短信(RCS)等业务线的统一封装：
//   - 统一入口 CloudApiClient，一次初始化全业务可用；
//   - 短信业务独立连接池，避免批量发送时与其他业务竞争；
//   - 内置指数退避重试，重试时重新生成签名；
//   - 所有方法支持可选 traceId 链路追踪。
package cloudsdk

import (
	"io"

	"chuanglanyunzhi/cloud-sdk/internal/transport"
	"chuanglanyunzhi/cloud-sdk/rcs"
	"chuanglanyunzhi/cloud-sdk/sms"
)

// CloudApiClient 是所有业务的统一入口，聚合各业务客户端。
// 内部维护两个独立的 HTTP 连接池（SMS 专用 + 通用）。
// 应创建为单例复用，使用完毕后调用 Close() 释放连接池。
type CloudApiClient struct {
	cfg *CloudApiConfig
	tr  *transport.Transport
	Sms *sms.SmsClient
	Rcs *rcs.RcsClient
}

// New 返回一个默认配置的 CloudApiClient。
func New() *CloudApiClient {
	return NewWithConfig(NewCloudApiConfig())
}

// NewWithConfig 返回一个按自定义配置初始化的 CloudApiClient。
func NewWithConfig(cfg *CloudApiConfig) *CloudApiClient {
	if cfg == nil {
		cfg = NewCloudApiConfig()
	}
	tcfg := transport.NewConfig()
	tcfg.ConnectTimeout = cfg.ConnectTimeout
	tcfg.ReadTimeout = cfg.ReadTimeout
	tcfg.MaxRetries = cfg.MaxRetries
	tcfg.RetryBaseMs = cfg.RetryBaseIntervalMs

	tr := transport.New(tcfg, transport.SignStamp)
	return &CloudApiClient{cfg: cfg, tr: tr}
}

// Clients 绑定 appId/appSecret 到各业务客户端。
// 返回 (*CloudApiClient) 以便链式调用。
func (c *CloudApiClient) Clients(appId, appSecret string) *CloudApiClient {
	c.Sms = sms.New(c.tr, c.cfg.SmsEndpoint, appId, appSecret)
	c.Rcs = rcs.New(c.tr, c.cfg.RcsEndpoint, appId, appSecret)
	return c
}

// Close 释放两个 HTTP 连接池（SMS 专用 + 通用）的空闲连接。
// 释放后的实例不可再使用。
func (c *CloudApiClient) Close() error {
	if c.tr != nil {
		c.tr.Close()
	}
	return nil
}

var _ io.Closer = (*CloudApiClient)(nil)

// Package transport 提供底层 HTTP 传输、签名与重试能力。
// 短信与视频短信业务各自使用独立的连接池，避免批量发送时相互竞争。
package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"chuanglanyunzhi/cloud-sdk/internal/retry"
)

// 连接池参数。
const (
	maxIdleConns    = 200
	maxPerHost      = 100
	smsMaxIdle      = 50
	keepAlive       = 5 * time.Minute
	defaultReadTime = 10 * time.Second
	defaultConnTime = 10 * time.Second
)

// AuthStyle 标识两种签名算法。
type AuthStyle int

const (
	// AuthSMS 短信/视频短信签名：CheckSum=MD5(appSecret+Nonce+CurTime)。
	AuthSMS AuthStyle = iota
	// AuthCommon 通用业务签名：Signature=MD5(appId+appSecret+Timestamp)。
	AuthCommon
)

// Config 是 transport 层的连接池、超时与重试配置。
type Config struct {
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	MaxRetries     int
	RetryBaseMs    int
}

// NewConfig 返回带默认超时与重试配置的 transport 配置。
func NewConfig() Config {
	return Config{
		ConnectTimeout: defaultConnTime,
		ReadTimeout:    defaultReadTime,
		MaxRetries:     3,
		RetryBaseMs:    100,
	}
}

// Auth 携带 appId/appSecret，用于生成签名。
type Auth struct {
	AppId     string
	AppSecret string
}

// Transport 持有双连接池 HTTP 传输层。
type Transport struct {
	cfg          Config
	smsClient    *http.Client
	commonClient *http.Client
	stamp        StampFunc
}

// StampFunc 在请求发出前附加鉴权头与 traceId。
type StampFunc func(r *http.Request, auth Auth, style AuthStyle, traceId string)

// Option 是 Transport 的可选参数。
type Option func(*Transport)

// WithStamp 覆盖默认签名函数。
func WithStamp(fn StampFunc) Option {
	return func(t *Transport) { t.stamp = fn }
}

// New 创建双连接池 Transport。
func New(cfg Config, stamp StampFunc, opts ...Option) *Transport {
	t := &Transport{
		cfg:          cfg,
		smsClient:    newClient(cfg, smsMaxIdle),
		commonClient: newClient(cfg, maxIdleConns),
		stamp:        stamp,
	}
	for _, o := range opts {
		o(t)
	}
	return t
}

func newClient(cfg Config, maxIdle int) *http.Client {
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   cfg.ConnectTimeout,
			KeepAlive: keepAlive,
		}).DialContext,
		MaxIdleConns:          maxIdle,
		MaxIdleConnsPerHost:   maxPerHost,
		IdleConnTimeout:       keepAlive,
		TLSHandshakeTimeout:   cfg.ConnectTimeout,
		ExpectContinueTimeout: time.Second,
	}
	return &http.Client{Transport: tr, Timeout: cfg.ConnectTimeout + cfg.ReadTimeout}
}

// Close 释放两个连接池的空闲连接。
func (t *Transport) Close() {
	closeIdle(t.smsClient)
	closeIdle(t.commonClient)
}

func closeIdle(c *http.Client) {
	if tr, ok := c.Transport.(*http.Transport); ok {
		tr.CloseIdleConnections()
	}
}

// HTTPError 表示一次可重试或已失败的 HTTP 层错误。
type HTTPError struct {
	Status int
	Body   string
	Err    error
}

func (e *HTTPError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("http error status=%d body=%s cause=%v", e.Status, e.Body, e.Err)
	}
	return fmt.Sprintf("http error status=%d body=%s", e.Status, e.Body)
}

func (e *HTTPError) Unwrap() error { return e.Err }

// IsRetryable 判断错误是否可重试(5xx 或网络/超时错误)。
func IsRetryable(err error) bool {
	if err == nil {
		return false
	}
	if he, ok := err.(*HTTPError); ok && he.Err == nil {
		return he.Status >= 500 && he.Status < 600
	}
	ne, ok := err.(net.Error)
	return ok && (ne.Timeout() || ne.Temporary())
}

// Do 执行一次带指数退避重试的 JSON 请求并返回响应体。
// smsStyle 决定使用哪套连接池；stamp 负责注入鉴权头。
func (t *Transport) Do(ctx context.Context, method, endpoint, path string,
	body []byte, auth Auth, style AuthStyle, traceId string, smsStyle bool) ([]byte, error) {
	var lastErr error
	for attempt := 0; attempt <= t.cfg.MaxRetries; attempt++ {
		data, err := t.doOnce(ctx, method, endpoint, path, body, auth, style, traceId, smsStyle)
		if err == nil {
			return data, nil
		}
		lastErr = err
		if !IsRetryable(err) || attempt >= t.cfg.MaxRetries {
			return nil, err
		}
		if !retry.SleepCtx(ctx, retry.ExponentialBackoff(attempt, t.cfg.RetryBaseMs)) {
			return nil, err
		}
	}
	return nil, lastErr
}

func (t *Transport) doOnce(ctx context.Context, method, endpoint, path string,
	body []byte, auth Auth, style AuthStyle, traceId string, smsStyle bool) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, endpoint+path, bytes.NewReader(body))
	if err != nil {
		return nil, &HTTPError{Err: err}
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	if t.stamp != nil {
		t.stamp(req, auth, style, traceId)
	}

	client := t.commonClient
	if smsStyle {
		client = t.smsClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, &HTTPError{Err: err}
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &HTTPError{Status: resp.StatusCode, Err: err}
	}
	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{Status: resp.StatusCode, Body: string(data)}
	}
	return data, nil
}

// Parse 解析响应体公共字段(code/msg/requestId)。
func Parse(data []byte) (code, msg, requestId string) {
	var p struct {
		Code      string `json:"code"`
		Msg       string `json:"msg"`
		Message   string `json:"message"`
		RequestId string `json:"requestId"`
	}
	_ = json.Unmarshal(data, &p)
	if p.Code == "" {
		p.Code = p.Message
	}
	if p.Msg == "" {
		p.Msg = p.Message
	}
	return p.Code, p.Msg, p.RequestId
}

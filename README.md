# cloudsdk (Go)

创蓝统一应用的 Go SDK，提供短信、视频短信(RCS)业务线的统一封装。

## 特性

- **统一入口** `CloudApiClient`：一次初始化绑定 appId/appSecret，短信与 RCS 全业务可用
- **性能优化**：短信业务独立连接池，避免批量发送时阻塞其他业务
- **链路追踪**：所有方法支持可选 `traceId`，写入 `X-Trace-Id` 头
- **自动重试**：内置指数退避（默认最大 3 次，100ms/200ms/400ms），重试时重新生成签名
- **资源管理**：实现 `io.Closer`，使用完毕 `Close()` 释放连接池
- **认证**：短信/RCS 使用 `MD5(appSecret+Nonce+CurTime)` 防重放签名

## 系统要求

- Go 1.25+
- 仅依赖标准库（net/http）

## 快速开始

```go
package main

import (
	"context"
	"fmt"

	"chuanglanyunzhi/cloud-sdk"
	"chuanglanyunzhi/cloud-sdk/sms"
)

func main() {
	ctx := context.Background()
	client := cloudsdk.New()
	defer client.Close()

	client.Clients("your-app-id", "your-app-secret")

	req := (&sms.SmsBatchSendRequest{}).
		SetProductType("notify").
		SetPhoneNumbers("13800138000,13900139000").
		SetTemplateCode("T12345").
		SetSignName("创蓝云").
		SetReport(true)

	resp, err := client.Sms.BatchSend(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.IsSuccess() {
		fmt.Println("发送成功 msgId:", resp.Data.MsgId)
	} else {
		fmt.Println("发送失败:", resp.Code, resp.Msg)
	}
}
```

## 自定义配置

```go
cfg := cloudsdk.NewCloudApiConfig().
	SetSmsEndpoint("https://smssh.253.com").
	SetRcsEndpoint("https://videoapi.253.com").
	SetConnectTimeout(15 * time.Second).
	SetReadTimeout(30 * time.Second).
	SetMaxRetries(3).
	SetRetryBaseIntervalMs(100)

client := cloudsdk.NewWithConfig(cfg)
defer client.Close()
client.Clients(appId, appSecret)
```

## 业务范围

| 业务客户端 | 功能 | 方法数 |
| --- | --- | --- |
| SmsClient | 短信发送、签名、模板、资质管理 | 18 |
| RcsClient | 视频短信(RCS)模板、发送、报告、上行、签名 | 6 |

> 国际短信、风控、携号转网、企业信息等业务线不在本 SDK 范围内。

## 目录结构

```
.
├── go.mod                 # module chuanglanyunzhi/cloud-sdk, go 1.25
├── client.go              # CloudApiClient 统一入口
├── config.go              # CloudApiConfig 配置
├── errors.go              # SdkError 错误类型
├── internal/
│   ├── transport/         # 双连接池 HTTP 传输、签名、错误解析
│   └── retry/             # 指数退避重试
├── sms/                   # 短信业务(18 方法)
└── rcs/                   # 视频短信(6 方法)
```
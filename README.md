# cloudsdk (Go)

创蓝统一应用的 Go SDK，提供短信、视频短信(RCS)业务线的统一封装。

## 特性

- **统一入口** `CloudApiClient`：一次初始化绑定 appId/appSecret，短信与 RCS 全业务可用
- **性能优化**：短信业务独立连接池，避免批量发送时阻塞其他业务
- **链路追踪**：所有方法支持可选 `traceId`，写入请求头 `X-Custom-TraceId`
- **自动重试**：内置指数退避（默认最大 3 次，100ms/200ms/400ms），重试时重新生成签名
- **资源管理**：实现 `io.Closer`，使用完毕 `Close()` 释放连接池
- **认证**：`SHA1(AppSecret + Nonce + CurTime)`，头 `AppID/Nonce/CurTime/CheckSum`

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

> 成功状态码：短信业务 `000000`；视频短信 `102000`（`IsSuccess()` 已内置判断）。

## 自定义配置

```go
cfg := cloudsdk.NewCloudApiConfig().
	SetSmsEndpoint("https://smssh.253.com").
	SetRcsEndpoint("https://rcs.253.com").
	SetConnectTimeout(15 * time.Second).
	SetReadTimeout(30 * time.Second).
	SetMaxRetries(3).
	SetRetryBaseIntervalMs(100)

client := cloudsdk.NewWithConfig(cfg)
defer client.Close()
client.Clients(appId, appSecret)
```

## 认证机制

统一应用 API Key 模式：

- `Nonce`：32 位字母数字随机字符串
- `CurTime`：当前秒级 Unix 时间戳（有效期 5 分钟）
- `CheckSum = SHA1(AppSecret + Nonce + CurTime)`（小写 hex）

请求头：`AppID` / `Nonce` / `CurTime` / `CheckSum`，可选 `X-Custom-TraceId`（≤64 字符）。

## 业务范围

### 短信业务 `client.Sms`（路径 `/sms/v2`）

| 方法 | 路径 | 说明 |
|---|---|---|
| BatchSend | /sms/v2/batch-send | 批量发送短信 |
| AddQualification | /sms/v2/qualification/add | 创建资质 |
| ListQualification | /sms/v2/qualification/list | 资质列表 |
| UpdateQualification | /sms/v2/qualification/update | 编辑资质 |
| DeleteQualification | /sms/v2/qualification/delete | 删除资质 |
| ListQualificationChangeLog | /sms/v2/qualification/change-log/list | 资质变更日志 |
| AddSignature | /sms/v2/signature/create | 创建签名 |
| ListSignature | /sms/v2/signature/list | 签名列表 |
| GetSignature | /sms/v2/signature/detail | 签名详情 |
| GetSignatureOperatorRejectReason | （包装 detail） | 签名运营商驳回原因 |
| UpdateSignatureRealName | /sms/v2/signature/real-name-update | 编辑签名终端资质 |
| DeleteSignature | /sms/v2/signature/delete | 删除签名 |
| AddTemplate | /sms/v2/template/add | 创建模板 |
| QueryTemplateTypeEnum | /sms/v2/template/query-type-enum | 查询模板类型枚举 |
| ListTemplate | /sms/v2/template/list | 模板列表 |
| GetTemplate | /sms/v2/template/detail | 模板详情 |
| GetTemplateOperatorRejectReason | （包装 detail） | 模板运营商驳回原因 |
| UpdateTemplate | /sms/v2/template/update | 编辑模板 |
| DeleteTemplate | /sms/v2/template/delete | 删除模板 |

### 视频短信业务 `client.Rcs`（路径 `/videosms/api/v2`）

| 方法 | 路径 | 说明 |
|---|---|---|
| AddVideoTemplate | /videosms/api/v2/template/add | 创建视频模板 |
| FindVideoTemplate | /videosms/api/v2/template/getSingleTemplateInfo | 查看模板状态 |
| SubmitVideoTemplate | /videosms/api/v2/template/send | 发送视频短信 |
| PullReport | /videosms/api/v2/report/pull | 拉取状态报告 |
| PullReply | /videosms/api/v2/reply/pull | 拉取上行消息 |
| AddSign | /videosms/api/v2/sign/add | 创建签名 |

### 推送回调解析 `callback`

服务端接收创蓝推送时使用（推送由创蓝主动调用，SDK 不发起），提供各推送结构体与 `Parse*` 方法：

- 短信：状态回执（`SMS_REPORT`）、上行报告（`SMS_MO`），应答 `{"clcode":"000000","msg":"接收成功"}`
- 视频短信：签名/模板审核结果、状态报告（数组）、上行消息（数组），应答字符串 `"ok"`

## 目录结构

```
.
├── go.mod                 # module chuanglanyunzhi/cloud-sdk, go 1.25
├── client.go              # CloudApiClient 统一入口
├── config.go              # CloudApiConfig 配置
├── errors.go              # SdkError 错误类型
├── internal/
│   ├── transport/         # 双连接池 HTTP 传输、SHA1 签名、错误解析
│   └── retry/             # 指数退避重试
├── sms/                   # 短信业务(19 方法)
├── rcs/                   # 视频短信(6 方法)
└── callback/              # 推送回调解析模型
```

## 参考文档

- `docs/sms/`：短信 API 调用说明、资质管理、签名管理、模板管理、短信发送、状态回执、上行报告、状态码
- `docs/video-sms/`：视频短信 API 调用说明、短信发送、签名管理、模板管理、状态报告、上行消息
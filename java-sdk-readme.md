---
source: https://doc.chuanglan.com/docs/SF6KRNWZHOZY860T
description: "简介 Cloud SDK 是 创蓝配合统一应用的 Go SDK，提供短信、视频短信(RCS)等业务线能力的统一封装。 核心特性 统一入口 ：CloudApiClient 聚合所有业务 API，一次初始化，全业务可用 性能优化 ：短信业务独立连接池，避免批量发送时与其他业务竞"
---

## 简介

Cloud SDK 是 创蓝配合统一应用的 Java SDK，提供短信、视频短信(RCS)等业务线能力的统一封装。

### 核心特性

-   **统一入口**：`CloudApiClient` 聚合所有业务 API，一次初始化，全业务可用
-   **性能优化**：短信业务独立连接池，避免批量发送时与其他业务竞争
-   **链路追踪**：所有方法支持可选 `traceId` 参数，便于分布式链路追踪
-   **自动重试**：内置指数退避重试机制（默认最大 3 次）
-   **Spring 集成**：提供 Spring Boot 自动装配支持，开箱即用
-   **资源管理**：实现 `AutoCloseable`，支持 try-with-resources 自动释放连接

### 系统要求

-   **JDK 8+**
-   **依赖**：OkHttp 4.9.3、Jackson 2.15.2
-   **可选**：Spring Boot 2.7.18+（用于 Spring 集成）

---

## 快速开始

### Maven 依赖

```xml
<dependency>
      <groupId>io.github.chuanglanyunzhi</groupId>
      <artifactId>cloud-sdk</artifactId>
      <version>1.0.0</version>
</dependency>
```

附件下载： [统一SDK下载](https://static2.253.com/doc-center/20260813/f14f92c9-ec19-4c8f-ac36-ec1924875abc.jar)

### 基础示例

```java
import com.chuanglan.cloudsdk.api.CloudApiClient;
import com.chuanglan.cloudsdk.api.sms.*;

public class QuickStart {
    public static void main(String[] args) {
        String appId = "your-app-id";
        String appSecret = "your-app-secret";

        // 1. 创建客户端（使用 try-with-resources 自动释放资源）
        try (CloudApiClient client = new CloudApiClient()) {

            // 2. 构建请求
            SmsBatchSendRequest request = new SmsBatchSendRequest()
                .setProductType("notify")
                .setPhoneNumbers("13800138000,13900139000")
                .setTemplateCode("T12345")
                .setSignName("创蓝云")
                .setReport(true);

            // 3. 调用 API
            SmsBatchSendResponse response = client.batchSend(appId, appSecret, request);

            // 4. 处理结果
            if (response.isSuccess()) {
                System.out.println("发送成功，msgId: " + response.getData().getMsgId());
            } else {
                System.err.println("发送失败: " + response.getMsg());
            }

        } catch (CloudSdkException e) {
            System.err.println("调用异常: " + e.getMessage());
            System.err.println("错误码: " + e.getCode());
            System.err.println("请求ID: " + e.getRequestId());
        }
    }
}
```

---

## 核心概念

### 客户端架构

Cloud SDK 采用**三层架构**设计：

```
cloud-sdk-parent
├── cloud-sdk-core      # 核心层：HTTP 传输、签名、序列化、异常处理
├── cloud-sdk-api       # API 层：封装各业务线 API
└── cloud-sdk-spring    # 集成层：Spring Boot 自动装配
```

#### 统一入口：CloudApiClient

`CloudApiClient` 是所有业务的统一入口，内部聚合了短信、视频短信(RCS)两个业务客户端：

| 业务客户端 | 功能 | 方法数 |
| --- | --- | --- |
| SmsClient | 短信发送、签名、模板、资质管理 | 18 |
| RcsSmsClient | 视频短信（RCS） | 6 |

> 注：国际短信、风控、携号转网、企业信息查询等业务线不在本 SDK 范围内，相关文档已移除。

#### 连接池隔离

为避免短信批量发送阻塞其他业务，SDK 内部维护了两个独立的 HTTP 连接池：

-   **SMS 连接池**：专用于短信业务（`SmsClient`）
-   **通用连接池**：用于其他所有业务

**连接池参数**：

-   最大空闲连接：50
-   Keep-Alive 时长：5 分钟
-   最大并发请求：200
-   单主机最大并发：100

### 认证机制

Cloud SDK 使用 **appId + appSecret** 进行身份认证，不同业务线采用不同的签名算法：

#### 短信业务

```java
// 签名参数
Nonce = 随机 UUID
CurTime = 当前时间戳（秒）
CheckSum = MD5(appSecret + Nonce + CurTime)

// HTTP 请求头
X-AppId: {appId}
X-Nonce: {Nonce}
X-CurTime: {CurTime}
X-CheckSum: {CheckSum}
Content-Type: application/json;charset=utf-8
```

#### 短信业务

```java
// 签名参数
Timestamp = 当前时间戳（毫秒）
Signature = MD5(appId + appSecret + Timestamp)

// HTTP 请求头
appId: {appId}
timestamp: {Timestamp}
sign: {Signature}
Content-Type: application/json;charset=utf-8
```

**重要说明**：

-   SDK 自动完成签名计算，开发者无需手动处理
-   重试时会重新生成签名，避免时间窗超时

### 异常处理

#### CloudSdkException

所有 SDK 异常都继承自 `CloudSdkException`，包含以下关键字段：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| code | String | 业务错误码（如 `000000` 表示成功） |
| message | String | 错误描述 |
| requestId | String | 请求 ID，用于问题排查 |
| statusCode | int | HTTP 状态码 |
| cause | Throwable | 原始异常（如有） |

#### 常见错误码

| 错误码 | 说明 | 处理建议 |
| --- | --- | --- |
| `ParameterMissing` | 必填参数缺失 | 检查请求参数 |
| `InvalidSignature` | 签名校验失败 | 检查 appId/appSecret 是否正确 |
| `InsufficientBalance` | 账户余额不足 | 充值后重试 |
| `TemplateNotFound` | 模板不存在 | 检查模板 ID 是否正确 |
| `RateLimitExceeded` | 超过频率限制 | 降低调用频率或申请提额 |
| `HttpError` | HTTP 请求失败 | 检查网络连接或重试 |

#### 异常处理示例

```java
try (CloudApiClient client = new CloudApiClient()) {
    SmsBatchSendResponse response = client.batchSend(appId, appSecret, request);

    if (response.isSuccess()) {
        // 处理成功逻辑
    } else {
        // 处理业务失败
        System.err.println("业务失败: " + response.getMsg());
    }

} catch (CloudSdkException e) {
    // 处理 SDK 异常
    System.err.println("SDK 异常: " + e.getMessage());
    System.err.println("错误码: " + e.getCode());
    System.err.println("请求ID: " + e.getRequestId());
    System.err.println("HTTP状态码: " + e.getStatusCode());

    if (e.getCause() != null) {
        System.err.println("原始异常: " + e.getCause().getMessage());
    }
}
```

---

## 配置指南

### 基础配置

#### 使用默认配置

```java
CloudApiClient client = new CloudApiClient();
```

默认配置：

-   连接超时：10 秒
-   读取超时：10 秒
-   使用官方默认 endpoint

#### 自定义配置

```java
import com.chuanglan.cloudsdk.api.CloudApiConfig;

CloudApiConfig config = new CloudApiConfig()
    // 短信 endpoint
    .setSmsEndpoint("https://smssh.253.com")
    // 视频短信 endpoint
    .setRcsSmsEndpoint("https://videoapi.253.com")
    // 连接超时（毫秒）
    .setConnectTimeout(15000)
    // 读取超时（毫秒）
    .setReadTimeout(30000);

CloudApiClient client = new CloudApiClient(config);
```

### Spring Boot 集成

#### 1\. 添加依赖

\-添加pom依赖或直接下载SDK集成使用

#### 2\. 配置文件

\-sdk不需要额外配置，不同节点在调用方法中指定调用节点endpoint即可

#### 3\. 注入使用

```java
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import com.chuanglan.cloudsdk.api.CloudApiClient;

@Service
public class SmsService {

    @Autowired
    private CloudApiClient client;

    public void sendSms(String phone, String templateCode) {
        SmsBatchSendRequest request = new SmsBatchSendRequest()
            .setProductType("notify")
            .setPhoneNumbers(phone)
            .setTemplateCode(templateCode);

        SmsBatchSendResponse response = client.batchSend(appId, appSecret, request);
        // 处理响应
    }
}
```

**注意**：

-   Spring 容器会自动管理 `CloudApiClient` 的生命周期，无需手动 `close()`
-   禁用自动装配：`cloudsdk.enabled=false`

---

## 业务 API 使用

## 短信业务

### 1\. batchSend - 批量发送短信

**方法签名**：

```java
SmsBatchSendResponse batchSend(String appId, String appSecret, SmsBatchSendRequest request)
SmsBatchSendResponse batchSend(String appId, String appSecret, SmsBatchSendRequest request, String traceId)
```

**使用示例**：

```java
SmsBatchSendRequest request = new SmsBatchSendRequest()
    .setProductType("notify")
    .setPhoneNumbers("13800138000,13900139000")
    .setTemplateCode("T12345")
    .setSignName("创蓝云")
    .setReport(true);

SmsBatchSendResponse response = client.batchSend(appId, appSecret, request);
if (response.isSuccess()) {
    System.out.println("msgId: " + response.getData().getMsgId());
}
```

---

### 2\. addQualification - 添加资质

**方法签名**：

```java
SmsQualificationAddResponse addQualification(String appId, String appSecret, SmsQualificationAddRequest request)
SmsQualificationAddResponse addQualification(String appId, String appSecret, SmsQualificationAddRequest request, String traceId)
```

**使用示例**：

```java
SmsQualificationAddRequest request = new SmsQualificationAddRequest()
    .setQualificationType("1")  // 1=企业 2=个人
    .setCompanyName("创蓝科技有限公司")
    .setLicenseUrl("https://example.com/license.jpg");

SmsQualificationAddResponse response = client.addQualification(appId, appSecret, request);
```

---

### 3\. listQualification - 查询资质列表

**方法签名**：

```java
SmsQualificationListResponse listQualification(String appId, String appSecret, SmsQualificationListRequest request)
SmsQualificationListResponse listQualification(String appId, String appSecret, SmsQualificationListRequest request, String traceId)
```

**使用示例**：

```java
SmsQualificationListRequest request = new SmsQualificationListRequest()
    .setPageNum(1)
    .setPageSize(20);

SmsQualificationListResponse response = client.listQualification(appId, appSecret, request);
for (QualificationItem item : response.getData().getList()) {
    System.out.println("资质ID: " + item.getQualificationId());
}
```

---

### 4\. updateQualification - 更新资质

**方法签名**：

```java
SmsQualificationUpdateResponse updateQualification(String appId, String appSecret, SmsQualificationUpdateRequest request)
SmsQualificationUpdateResponse updateQualification(String appId, String appSecret, SmsQualificationUpdateRequest request, String traceId)
```

---

### 5\. deleteQualification - 删除资质

**方法签名**：

```java
SmsQualificationDeleteResponse deleteQualification(String appId, String appSecret, SmsQualificationDeleteRequest request)
SmsQualificationDeleteResponse deleteQualification(String appId, String appSecret, SmsQualificationDeleteRequest request, String traceId)
```

---

### 6\. addSignature - 添加签名

**方法签名**：

```java
SmsSignatureAddResponse addSignature(String appId, String appSecret, SmsSignatureAddRequest request)
SmsSignatureAddResponse addSignature(String appId, String appSecret, SmsSignatureAddRequest request, String traceId)
```

**使用示例**：

```java
SmsSignatureAddRequest request = new SmsSignatureAddRequest()
    .setSignName("创蓝云")
    .setSignType("1")           // 1=网站 2=APP 3=微信公众号 4=企业名称
    .setSignPurpose("1")        // 1=自用 2=他用
    .setRemark("官方签名");

SmsSignatureAddResponse response = client.addSignature(appId, appSecret, request);
```

---

### 7\. getSignature - 查询签名详情

**方法签名**：

```java
SmsSignatureGetResponse getSignature(String appId, String appSecret, SmsSignatureGetRequest request)
SmsSignatureGetResponse getSignature(String appId, String appSecret, SmsSignatureGetRequest request, String traceId)
```

---

### 8\. listSignature - 查询签名列表

**方法签名**：

```java
SmsSignatureListResponse listSignature(String appId, String appSecret, SmsSignatureListRequest request)
SmsSignatureListResponse listSignature(String appId, String appSecret, SmsSignatureListRequest request, String traceId)
```

**使用示例**：

```java
SmsSignatureListRequest request = new SmsSignatureListRequest()
    .setSignName("创蓝云")
    .setStatus("2")  // 1=待审核 2=已通过 3=已驳回
    .setPageNum(1)
    .setPageSize(20);

SmsSignatureListResponse response = client.listSignature(appId, appSecret, request);
```

---

### 9\. getSignatureOperatorRejectReason - 查询签名运营商驳回原因

**方法签名**：

```java
SmsSignatureOperatorRejectReasonResponse getSignatureOperatorRejectReason(String appId, String appSecret, SmsSignatureOperatorRejectReasonRequest request)
SmsSignatureOperatorRejectReasonResponse getSignatureOperatorRejectReason(String appId, String appSecret, SmsSignatureOperatorRejectReasonRequest request, String traceId)
```

---

### 10\. updateSignatureRealName - 更新签名实名信息

**方法签名**：

```java
SmsSignatureRealNameUpdateResponse updateSignatureRealName(String appId, String appSecret, SmsSignatureRealNameUpdateRequest request)
SmsSignatureRealNameUpdateResponse updateSignatureRealName(String appId, String appSecret, SmsSignatureRealNameUpdateRequest request, String traceId)
```

---

### 11\. deleteSignature - 删除签名

**方法签名**：

```java
SmsSignatureDeleteResponse deleteSignature(String appId, String appSecret, SmsSignatureDeleteRequest request)
SmsSignatureDeleteResponse deleteSignature(String appId, String appSecret, SmsSignatureDeleteRequest request, String traceId)
```

---

### 12\. addTemplate - 添加模板

**方法签名**：

```java
SmsTemplateAddResponse addTemplate(String appId, String appSecret, SmsTemplateAddRequest request)
SmsTemplateAddResponse addTemplate(String appId, String appSecret, SmsTemplateAddRequest request, String traceId)
```

**使用示例**：

```java
SmsTemplateAddRequest request = new SmsTemplateAddRequest()
    .setTemplateName("验证码模板")
    .setTemplateType("1")       // 1=验证码 2=通知 3=营销
    .setTemplateContent("您的验证码是{1}，{2}分钟内有效")
    .setSignName("创蓝云");

SmsTemplateAddResponse response = client.addTemplate(appId, appSecret, request);
```

---

### 13\. queryTemplateTypeEnum - 查询模板类型枚举

**方法签名**：

```java
SmsTemplateTypeEnumResponse queryTemplateTypeEnum(String appId, String appSecret, SmsTemplateTypeEnumRequest request)
SmsTemplateTypeEnumResponse queryTemplateTypeEnum(String appId, String appSecret, SmsTemplateTypeEnumRequest request, String traceId)
```

---

### 14\. listTemplate - 查询模板列表

**方法签名**：

```java
SmsTemplateListResponse listTemplate(String appId, String appSecret, SmsTemplateListRequest request)
SmsTemplateListResponse listTemplate(String appId, String appSecret, SmsTemplateListRequest request, String traceId)
```

---

### 15\. getTemplate - 查询模板详情

**方法签名**：

```java
SmsTemplateGetResponse getTemplate(String appId, String appSecret, SmsTemplateGetRequest request)
SmsTemplateGetResponse getTemplate(String appId, String appSecret, SmsTemplateGetRequest request, String traceId)
```

---

### 16\. getTemplateOperatorRejectReason - 查询模板运营商驳回原因

**方法签名**：

```java
SmsTemplateOperatorRejectReasonResponse getTemplateOperatorRejectReason(String appId, String appSecret, SmsTemplateOperatorRejectReasonRequest request)
SmsTemplateOperatorRejectReasonResponse getTemplateOperatorRejectReason(String appId, String appSecret, SmsTemplateOperatorRejectReasonRequest request, String traceId)
```

---

### 17\. updateTemplate - 更新模板

**方法签名**：

```java
SmsTemplateUpdateResponse updateTemplate(String appId, String appSecret, SmsTemplateUpdateRequest request)
SmsTemplateUpdateResponse updateTemplate(String appId, String appSecret, SmsTemplateUpdateRequest request, String traceId)
```

---

### 18\. deleteTemplate - 删除模板

**方法签名**：

```java
SmsTemplateDeleteResponse deleteTemplate(String appId, String appSecret, SmsTemplateDeleteRequest request)
SmsTemplateDeleteResponse deleteTemplate(String appId, String appSecret, SmsTemplateDeleteRequest request, String traceId)
```

---

## 视频短信业务

### 1\. addVideoTemplate - 添加视频模板

**方法签名**：

```java
RcsSmsTemplateAddResponse addVideoTemplate(String appId, String appSecret, RcsSmsTemplateAddRequest request)
RcsSmsTemplateAddResponse addVideoTemplate(String appId, String appSecret, RcsSmsTemplateAddRequest request, String traceId)
```

**使用示例**：

```java
RcsSmsTemplateAddRequest request = new RcsSmsTemplateAddRequest()
    .setTemplateName("营销视频")
    .setTemplateContent("视频内容描述")
    .setVideoUrl("https://example.com/video.mp4")
    .setSignName("创蓝云");

RcsSmsTemplateAddResponse response = client.addVideoTemplate(appId, appSecret, request);
```

---

### 2\. findVideoTemplate - 查询视频模板

**方法签名**：

```java
RcsSmsTemplateFindResponse findVideoTemplate(String appId, String appSecret, RcsSmsTemplateFindRequest request)
RcsSmsTemplateFindResponse findVideoTemplate(String appId, String appSecret, RcsSmsTemplateFindRequest request, String traceId)
```

---

### 3\. submitVideoTemplate - 发送视频短信

**方法签名**：

```java
RcsSmsTemplateSubmitResponse submitVideoTemplate(String appId, String appSecret, RcsSmsTemplateSubmitRequest request)
RcsSmsTemplateSubmitResponse submitVideoTemplate(String appId, String appSecret, RcsSmsTemplateSubmitRequest request, String traceId)
```

**使用示例**：

```java
RcsSmsTemplateSubmitRequest request = new RcsSmsTemplateSubmitRequest()
    .setPhoneNumbers("13800138000,13900139000")
    .setTemplateCode("VT12345")
    .setReport(true);

RcsSmsTemplateSubmitResponse response = client.submitVideoTemplate(appId, appSecret, request);
```

---

### 4\. pullReport - 拉取状态报告

**方法签名**：

```java
RcsSmsReportPullResponse pullReport(String appId, String appSecret, RcsSmsReportPullRequest request)
RcsSmsReportPullResponse pullReport(String appId, String appSecret, RcsSmsReportPullRequest request, String traceId)
```

---

### 5\. pullReply - 拉取上行回复

**方法签名**：

```java
RcsSmsReplyPullResponse pullReply(String appId, String appSecret, RcsSmsReplyPullRequest request)
RcsSmsReplyPullResponse pullReply(String appId, String appSecret, RcsSmsReplyPullRequest request, String traceId)
```

---

### 6\. addSign - 添加签名

**方法签名**：

```java
RcsSmsSignAddResponse addSign(String appId, String appSecret, RcsSmsSignAddRequest request)
RcsSmsSignAddResponse addSign(String appId, String appSecret, RcsSmsSignAddRequest request, String traceId)
```

---

## 链路追踪支持

所有业务方法都支持链路追踪，通过传入 `traceId` 参数即可：

```java
String traceId = UUID.randomUUID().toString();

// 不带 traceId
SmsBatchSendResponse response1 = client.batchSend(appId, appSecret, request);

// 带 traceId
SmsBatchSendResponse response2 = client.batchSend(appId, appSecret, request, traceId);
```

`traceId` 会添加到 HTTP 请求头 `X-Trace-Id`，用于分布式链路追踪和问题排查。

---

## 其他方法

### 重试机制

SDK 内置指数退避重试策略，自动处理临时性网络故障：

**默认配置**：

-   最大重试次数：3 次
-   退避算法：`2^重试次数 * 100ms`
-   重试间隔：100ms、200ms、400ms

**可重试的错误**：

-   网络超时（ConnectTimeout、SocketTimeout）
-   服务端 5xx 错误
-   特定业务错误码（如 `RateLimitExceeded`）

**重要特性**：

-   每次重试会重新生成签名参数（Nonce、CurTime、Timestamp），避免时间窗校验失败
-   幂等性请求（如查询）自动重试
-   非幂等请求（如发送短信）需业务层自行判断是否重试

**自定义重试策略**（需修改源码）：

```java
// HttpTransport 构造器
HttpTransport transport = new HttpTransport(
    config,
    new ExponentialBackoffRetryPolicy(5, 200)  // 最大 5 次，初始间隔 200ms
);
```

### 资源管理

#### 使用 try-with-resources（推荐）

```java
try (CloudApiClient client = new CloudApiClient(config)) {
    // 使用 client
} // 自动调用 close()，释放连接池和线程池
```

#### 手动关闭

```java
CloudApiClient client = new CloudApiClient(config);
try {
    // 使用 client
} finally {
    client.close();  // 手动释放资源
}
```

#### Spring Boot 环境

```java
@Autowired
private CloudApiClient client;  // Spring 容器自动管理生命周期，无需手动 close()
```

**资源释放说明**：

-   `close()` 会关闭两个 HTTP 连接池（SMS 专用 + 通用）
-   关闭 OkHttp 的 ConnectionPool 和 Dispatcher 线程池
-   最多等待 5 秒让正在执行的请求完成
-   释放后的 client 实例不可再使用

### 日志安全

SDK 默认记录完整的请求响应日志，但可能包含敏感信息（手机号、身份证号等）。

#### 关闭完整日志

**Spring Boot 环境**：

```java
import com.chuanglan.cloudsdk.spring.CloudSdkSafeLog;

@CloudSdkSafeLog  // 添加到 Spring Boot 启动类
@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}
```

**非 Spring 环境**： 需修改 `HttpTransport` 构造参数（源码级别）。

#### 敏感字段遮蔽

SDK 会自动遮蔽以下敏感字段（部分显示）：

-   手机号：`138****0000`
-   身份证号：`110101********1234`
-   银行卡号：`6228****890018`
-   密码、密钥：`****`

---

## 最佳实践

### 1\. 单例客户端

`CloudApiClient` 内部维护连接池，应创建为**单例**，避免重复创建：

```java
// ❌ 错误：每次请求创建新客户端
public void sendSms() {
    CloudApiClient client = new CloudApiClient();  // 连接池未复用
    client.batchSend(appId, appSecret, request);
    client.close();
}

// ✅ 正确：单例复用
public class SmsService {
    private static final CloudApiClient CLIENT = new CloudApiClient();

    public void sendSms() {
        CLIENT.batchSend(appId, appSecret, request);  // 复用连接池
    }
}
```

### 2\. 异常处理

区分**业务失败**和**SDK 异常**：

```java
try {
    SmsBatchSendResponse response = client.batchSend(appId, appSecret, request);

    if (response.isSuccess()) {
        // 成功逻辑
    } else {
        // 业务失败（如余额不足、模板不存在）
        log.warn("短信发送失败: code={}, msg={}", response.getCode(), response.getMsg());
    }

} catch (CloudSdkException e) {
    // SDK 异常（如网络超时、序列化失败）
    log.error("SDK 异常: code={}, requestId={}", e.getCode(), e.getRequestId(), e);
}
```

### 3\. 并发调用

`CloudApiClient` 是**线程安全**的，支持多线程并发调用：

```java
ExecutorService executor = Executors.newFixedThreadPool(10);

for (String phone : phoneList) {
    executor.submit(() -> {
        SmsBatchSendRequest request = new SmsBatchSendRequest()
            .setPhoneNumbers(phone)
            .setTemplateCode("T12345");

        client.batchSend(appId, appSecret, request);  // 线程安全
    });
}
```

### 4\. 参数校验

SDK 会校验必填参数，但建议业务层提前校验：

```java
// ✅ 推荐：业务层校验
if (StringUtils.isBlank(phone) || !phone.matches("^1[3-9]\\d{9}$")) {
    throw new IllegalArgumentException("手机号格式错误");
}

SmsBatchSendRequest request = new SmsBatchSendRequest()
    .setPhoneNumbers(phone)
    .setTemplateCode(templateCode);
```

### 5\. 敏感信息管理

**不要硬编码 appId/appSecret**：

```java
// ❌ 错误：硬编码
String appId = "12345";
String appSecret = "abcdef";

// ✅ 正确：从配置文件或环境变量读取
@Value("${cloudsdk.app-id}")
private String appId;

@Value("${cloudsdk.app-secret}")
private String appSecret;
```

### 6\. 批量操作

批量发送短信时，使用逗号分隔号码，减少 API 调用次数：

```java
// ✅ 推荐：批量发送（一次请求）
String phones = "13800138000,13900139000,13700137000";
SmsBatchSendRequest request = new SmsBatchSendRequest()
    .setPhoneNumbers(phones)
    .setTemplateCode("T12345");

client.batchSend(appId, appSecret, request);

// ❌ 不推荐：逐个发送（多次请求）
for (String phone : phoneList) {
    client.batchSend(appId, appSecret, request.setPhoneNumbers(phone));
}
```

---

## 常见问题

### 1\. 如何获取 appId 和 appSecret？

登录 [创蓝云平台](https://chuanglan.com/)，在「统一应用管理」中创建应用，获取对应的 appId 和 appSecret。

### 2\. 签名校验失败怎么办？

**可能原因**：

-   appId 或 appSecret 错误
-   服务器时间不同步（误差超过 5 分钟）
-   重试时使用了旧的签名参数

**解决方法**：

-   检查 appId/appSecret 是否正确
-   同步服务器时间：`ntpdate ntp.aliyun.com`
-   SDK 会自动重新签名，无需手动处理

### 3\. 如何处理超时问题？

**调整超时时间**：

```java
CloudApiConfig config = new CloudApiConfig()
    .setConnectTimeout(30000)  // 连接超时 30 秒
    .setReadTimeout(60000);    // 读取超时 60 秒
```

**检查网络连接**：

```bash
# 测试连通性
curl -I https://smssh.253.com
```

### 4\. 如何查看请求日志？

SDK 使用 OkHttp 的日志拦截器，需配置日志框架：

**Logback 配置**：

```xml
<logger name="okhttp3" level="DEBUG"/>
<logger name="com.chuanglan.cloudsdk" level="DEBUG"/>
```

**Log4j2 配置**：

```xml
<Logger name="okhttp3" level="DEBUG"/>
<Logger name="com.chuanglan.cloudsdk" level="DEBUG"/>
```

### 5\. Spring Boot 自动装配不生效？

**检查清单**：

1.  是否添加了 `cloud-sdk` 依赖？
2.  配置文件中是否设置了 `cloudsdk.enabled=false`？
3.  是否在启动类所在包或子包下？
4.  是否存在自定义的 `CloudApiClient` Bean 覆盖了自动装配？

**调试方法**：

```bash
# 启动时查看自动装配日志
java -jar app.jar --debug | grep CloudSdk
```

### 6\. 如何实现异步调用？

SDK 本身是同步阻塞的，异步需业务层实现：

```java
// 使用 CompletableFuture
CompletableFuture.runAsync(() -> {
    client.batchSend(appId, appSecret, request);
}, executor);

// 使用 Spring @Async
@Async
public void sendSmsAsync(String phone) {
    client.batchSend(appId, appSecret, request);
}
```

### 7\. 如何实现幂等性？

短信发送等非幂等操作，建议业务层实现：

```java
// 使用分布式锁
String lockKey = "sms:send:" + phone + ":" + templateCode;
if (redisTemplate.opsForValue().setIfAbsent(lockKey, "1", 60, TimeUnit.SECONDS)) {
    try {
        client.batchSend(appId, appSecret, request);
    } finally {
        redisTemplate.delete(lockKey);
    }
} else {
    log.warn("重复发送短信: phone={}", phone);
}
```

### 8\. 如何处理大批量发送？

建议分批发送，避免单次请求过大：

```java
List<String> phoneList = ...; // 1000+ 号码
int batchSize = 100;

for (int i = 0; i < phoneList.size(); i += batchSize) {
    List<String> batch = phoneList.subList(i, Math.min(i + batchSize, phoneList.size()));
    String phones = String.join(",", batch);

    SmsBatchSendRequest request = new SmsBatchSendRequest()
        .setPhoneNumbers(phones)
        .setTemplateCode("T12345");

    client.batchSend(appId, appSecret, request);

    // 避免触发限流
    Thread.sleep(100);
}
```

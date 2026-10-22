---
source: https://doc.chuanglan.com/docs/CC1W9RJF5RZU3OJ6
description: "⚠️ 2026年7月前注册的账号不支持 统一应用 API Key 模式 ，请参考 <a href=\"https://doc.chuanglan.com/docs/FACU2YN81RTWLLRA\" target=\" self\" <strong 独立产品 API 账号模式</strong </a ✨"
---

> ⚠️ 2026年7月前注册的账号不支持**统一应用 API Key 模式**，请参考 [**独立产品 API 账号模式**](https://doc.chuanglan.com/docs/FACU2YN81RTWLLRA)
>
> ✨ **统一应用 API Key 模式** 支持通过 SDK 调用接口，详见 [统一SDK使用手册](https://doc.chuanglan.com/docs/SF6KRNWZHOZY860T)

# 验签算法

-   请求头中增加 CheckSum 参数字段，用于标识接口验签。

-   SHA1(AppSecret + Nonce + CurTime)，将该三个参数拼接的字符串进行 SHA1 哈希计算从而生成 16 进制字符（小写）。

-   出于安全性考虑，每个 CheckSum 的 **有效期** 为 **5 分钟**，即服务端接收到请求的时间与请求中的 CurTime 相差不能超过 5 分钟，建议每次请求都生成新的 CheckSum，同时 **请确认** 发起请求的服务器是与标准时间同步的。

# 参数说明

参数列表的可选属性：

-   **M**：表示必选

-   **O**：表示任意可选

-   **C**：表示条件可选

> 以下请求头适用于所有 POST 接口（推送类接口除外，推送由创蓝服务器主动调用）

| **参数名称** | **参数类型** | **是否必填** | **描述** | **示例** |
| --- | --- | --- | --- | --- |
| AppID | String | 是 | 通过控制台->开发者中心->统一应用管理->API Key 获取 | rDOMKNMLaE |
| Nonce | String | 是 | 随机字符(32位长度)，用于计算CheckSum | tewtewtwetwetwetewtewt |
| CurTime | String | 是 | 当前 UTC 时间戳，从 1970 年 1 月 1 日 0 时 0 分 0 秒开始到 **现在** 的秒数。该时间用于计算 CheckSum 的有效期，请确保与标准时间同步。 | 1624938115 |
| CheckSum | String | 是 | SHA1(AppSecret + Nonce + CurTime)，将该三个参数拼接的字符串进行 SHA1 哈希计算从而生成 16 进制字符（小写）。 | 1d6cb525296b0e7c50e5dd7e05457135b5b8eed369ce5048e6bde2d6b10a1071 |
| X-Custom-TraceId | String | 否 | 开发者自定义标识，用于上下游请求关联 |  |
| **参数名称** | **参数类型** | **描述** | **示例** |
| --- | --- | --- | --- |
| Content-Type | String | 响应体的数据类型，application/json; charset=utf-8 | application/json; charset=utf-8 |
| X-Timestamp | String | 创蓝服务器接收到该请求的时间，精确到 ms 的 UTC 时间 |  |
| X-Custom-TraceId | String | 开发者自定义标识，用于上下游请求关联 |  |

---
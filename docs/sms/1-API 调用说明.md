---
source: https://doc.chuanglan.com/docs/I58861JUE7A3QRU4
description: "<! API 调用说明 更新时间：2026 10 22 11:33 更新人：Ross ⚠️ 2026 年 7 月前注册的账号不支持 统一应用 API Key 模式 ，请对接 独立产品 API 账号模式 接口。 ✨ 统一应用 API Key 模式 支持通过 SDK 调用接口，详见 统一SDK使用手册"
---

> ⚠️ 2026 年 7 月前注册的账号不支持**统一应用 API Key 模式**，请对接**独立产品 API 账号模式**接口。
>
> ✨ **统一应用 API Key 模式** 支持通过 SDK 调用接口，详见 [统一SDK使用手册](https://doc.chuanglan.com/docs/SF6KRNWZHOZY860T)

本文档适用于统一应用（新版）账号，介绍 API 的接口规范、鉴权机制与公共参数。

### 请求规范

> -   请求方式：POST
> -   Content-Type：application/json

### 公共请求头参数

> 💡 AppID 和 AppSecret 获取路径：登录 [控制台主页](https://www.chuanglan.com/new-control/home) → `开发者中心` → `统一应用管理` → 创建 API Key（生成 AppID 和 AppSecret）

| 名称 | 类型 | 必填 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| AppID | String | 是 | 统一应用凭证标识：创建 API Key 时生成。 | \- |
| Nonce | String | 是 | 随机字符串：32 位，参与 `CheckSum` 计算。 | \- |
| CurTime | String | 是 | Unix 时间戳：秒级，有效期 5 分钟。 | \- |
| CheckSum | String | 是 | 签名值：SHA1 十六进制字符串，详见 [CheckSum 计算说明](#checksum-%E8%AE%A1%E7%AE%97%E8%AF%B4%E6%98%8E)。 | \- |
| X-Custom-TraceId | String | 否 | 自定义的请求追踪标识：不超过 64 个字符。 | `trace_2026052912514` |

### CheckSum 计算说明

#### 计算规则

CheckSum 由 `AppSecret`、`Nonce`、`CurTime` 三个字段按顺序拼接（无分隔符），再做 SHA1 计算：

```
CheckSum = SHA1(AppSecret + Nonce + CurTime)
```

#### 计算示例

以下面参数为例：

| 参数 | 值 |
| --- | --- |
| AppSecret | `a3K7mP9xQA` |
| Nonce | `36SuukDa291gfN4dTSNQ07lxFzP68KPB` |
| CurTime | `1779367802` |

**Step 1.** 拼接原始字符串：

```
a3K7mP9xQA36SuukDa291gfN4dTSNQ07lxFzP68KPB1779367802
```

**Step 2.** 对拼接字符串做 SHA1 计算：

```
SHA1("a3K7mP9xQA36SuukDa291gfN4dTSNQ07lxFzP68KPB1779367802")
= 09dedc0e189b21e774686df7055c6bb003b0a128
```

将结果作为 `CheckSum` 请求头发送即可。

### 公共响应头参数

| 名称 | 类型 | 描述 | 示例 |
| --- | --- | --- | --- |
| X-Timestamp | String | 服务器接收请求的时间：Unix 时间戳，毫秒级。 | \- |
| X-Custom-TraceId | String | 原样返回自定义的请求追踪标识；若请求未携带，则不返回此字段。 | `trace_2026052912514` |

### 代码示例

#### cURL

```bash
APP_ID="DEV_7BW8WF4UIBM"
APP_SECRET="a3K7mP9xQA"
NONCE=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | head -c 32)
CUR_TIME=$(date +%s)

# 拼接并计算 CheckSum
CHECK_SUM=$(echo -n "${APP_SECRET}${NONCE}${CUR_TIME}" | sha1sum | awk '{print $1}')

curl -X POST "https://smssh.253.com/sms/v2/batch-send" \
  -H "Content-Type: application/json" \
  -H "AppID: ${APP_ID}" \
  -H "Nonce: ${NONCE}" \
  -H "CurTime: ${CUR_TIME}" \
  -H "CheckSum: ${CHECK_SUM}" \
  -d '{
    "productType": "notify",
    "phoneNumbers": "15800000000",
    "templateCode": "1111111"
  }'
```

#### Java

```java
import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.URL;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.util.UUID;

public class SmsExample {
    private static final String APP_ID = "DEV_7BW8WF4UIBM";
    private static final String APP_SECRET = "a3K7mP9xQA";

    public static void main(String[] args) throws Exception {
        String nonce = UUID.randomUUID().toString().replace("-", "");
        String curTime = String.valueOf(System.currentTimeMillis() / 1000);
        String checkSum = sha1(APP_SECRET + nonce + curTime);

        URL url = new URL("https://smssh.253.com/sms/v2/batch-send");
        HttpURLConnection conn = (HttpURLConnection) url.openConnection();
        conn.setRequestMethod("POST");
        conn.setRequestProperty("Content-Type", "application/json");
        conn.setRequestProperty("AppID", APP_ID);
        conn.setRequestProperty("Nonce", nonce);
        conn.setRequestProperty("CurTime", curTime);
        conn.setRequestProperty("CheckSum", checkSum);
        conn.setDoOutput(true);

        String body = "{"
                + "\"productType\":\"notify\","
                + "\"phoneNumbers\":\"15800000000\","
                + "\"templateCode\":\"1111111\""
                + "}";
        try (OutputStream os = conn.getOutputStream()) {
            os.write(body.getBytes(StandardCharsets.UTF_8));
        }

        int responseCode = conn.getResponseCode();
        System.out.println("Response Code: " + responseCode);

        // 读取响应参数
        InputStream is = (responseCode >= 200 && responseCode < 300)
                ? conn.getInputStream() : conn.getErrorStream();
        try (BufferedReader br = new BufferedReader(new InputStreamReader(is, StandardCharsets.UTF_8))) {
            StringBuilder response = new StringBuilder();
            String line;
            while ((line = br.readLine()) != null) {
                response.append(line);
            }
            System.out.println("Response Body: " + response);
        }
    }

    private static String sha1(String input) throws Exception {
        MessageDigest md = MessageDigest.getInstance("SHA-1");
        byte[] digest = md.digest(input.getBytes(StandardCharsets.UTF_8));
        StringBuilder sb = new StringBuilder();
        for (byte b : digest) {
            sb.append(String.format("%02x", b));
        }
        return sb.toString();
    }
}
```

#### Python

```python
import hashlib
import time
import uuid

import requests

APP_ID = "DEV_7BW8WF4UIBM"
APP_SECRET = "a3K7mP9xQA"

nonce = uuid.uuid4().hex
cur_time = str(int(time.time()))
check_sum = hashlib.sha1((APP_SECRET + nonce + cur_time).encode()).hexdigest()

headers = {
    "Content-Type": "application/json",
    "AppID": APP_ID,
    "Nonce": nonce,
    "CurTime": cur_time,
    "CheckSum": check_sum,
}
data = {
    "productType": "notify",
    "phoneNumbers": "15800000000",
    "templateCode": "1111111",
}

resp = requests.post("https://smssh.253.com/sms/v2/batch-send", headers=headers, json=data)
print("Response Code:", resp.status_code)
print("Response Body:", resp.text)
```

#### PHP

```php
// 配置参数（请替换为实际的 AppID 和 AppSecret）
define('APP_ID', 'DEV_7BW8WF4UIBM');
define('APP_SECRET', 'a3K7mP9xQA');

// 生成随机字符串（32位）
function generateNonce($length = 32) {
    $chars = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
    $nonce = '';
    for ($i = 0; $i < $length; $i++) {
        $nonce .= $chars[random_int(0, strlen($chars) - 1)];
    }
    return $nonce;
}

$nonce = generateNonce();
$curTime = strval(time());
$checkSum = sha1(APP_SECRET . $nonce . $curTime);

$headers = [
    'Content-Type: application/json',
    'AppID: ' . APP_ID,
    'Nonce: ' . $nonce,
    'CurTime: ' . $curTime,
    'CheckSum: ' . $checkSum
];

$data = [
    'productType' => 'notify',
    'phoneNumbers' => '15800000000',
    'templateCode' => '1111111'
];

// 使用 cURL 发送请求
$ch = curl_init('https://smssh.253.com/sms/v2/batch-send');
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);

$response = curl_exec($ch);
$httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);

if (curl_errno($ch)) {
    echo 'cURL Error: ' . curl_error($ch) . PHP_EOL;
} else {
    echo 'Response Code: ' . $httpCode . PHP_EOL;
    echo 'Response Body: ' . $response . PHP_EOL;
}

curl_close($ch);
```

> 🔒 生产环境请妥善保管 AppSecret，避免泄露。
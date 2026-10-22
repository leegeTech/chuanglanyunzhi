package transport

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// SMSAuthStamp 生成短信/视频短信签名并写入请求头。
// CheckSum = MD5(appSecret + Nonce + CurTime)，Nonce 为随机 UUID，CurTime 为秒级时间戳。
func SMSAuthStamp(r *http.Request, auth Auth, _ AuthStyle, traceId string) {
	nonce := randomUUID()
	curTime := strconv.FormatInt(time.Now().Unix(), 10)
	checkSum := md5Hex(auth.AppSecret + nonce + curTime)

	if traceId != "" {
		r.Header.Set("X-Trace-Id", traceId)
	}
	r.Header.Set("X-AppId", auth.AppId)
	r.Header.Set("X-Nonce", nonce)
	r.Header.Set("X-CurTime", curTime)
	r.Header.Set("X-CheckSum", checkSum)
}

// CommonAuthStamp 生成通用业务签名并写入请求头。
// Signature = MD5(appId + appSecret + Timestamp)；Timestamp 为毫秒级时间戳。
func CommonAuthStamp(r *http.Request, auth Auth, _ AuthStyle, traceId string) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign := md5Hex(auth.AppId + auth.AppSecret + timestamp)

	if traceId != "" {
		r.Header.Set("X-Trace-Id", traceId)
	}
	r.Header.Set("appId", auth.AppId)
	r.Header.Set("timestamp", timestamp)
	r.Header.Set("sign", sign)
}

func md5Hex(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// randomUUID 生成一个 v4 风格的随机 UUID 字符串。
func randomUUID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

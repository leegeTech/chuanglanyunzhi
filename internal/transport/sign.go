package transport

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

// 统一应用 API Key 模式鉴权算法：
//
//	Nonce    = 32 位字母数字随机字符串
//	CurTime  = 当前 Unix 时间戳（秒，有效期 5 分钟）
//	CheckSum = SHA1(AppSecret + Nonce + CurTime)，小写十六进制
//
// 请求头：AppID / Nonce / CurTime / CheckSum，可选 X-Custom-TraceId。
const (
	HeaderAppID         = "AppID"
	HeaderNonce         = "Nonce"
	HeaderCurTime       = "CurTime"
	HeaderCheckSum      = "CheckSum"
	HeaderCustomTraceId = "X-Custom-TraceId"

	nonceAlphabet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	nonceLength      = 32
	maxTraceIdLength = 64
)

// SignStamp 生成统一应用签名并写入请求头。
// 每次调用都生成新的 Nonce/CurTime，因此重试时会自动重新签名，避免时间窗过期。
func SignStamp(r *http.Request, auth Auth, _ AuthStyle, traceId string) {
	nonce := RandomNonce()
	curTime := strconv.FormatInt(time.Now().Unix(), 10)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set(HeaderAppID, auth.AppId)
	r.Header.Set(HeaderNonce, nonce)
	r.Header.Set(HeaderCurTime, curTime)
	r.Header.Set(HeaderCheckSum, CheckSum(auth.AppSecret, nonce, curTime))
	if traceId != "" {
		r.Header.Set(HeaderCustomTraceId, truncateTraceId(traceId))
	}
}

// CheckSum 计算 CheckSum = SHA1(appSecret + nonce + curTime)，返回小写十六进制字符串。
func CheckSum(appSecret, nonce, curTime string) string {
	sum := sha1.Sum([]byte(appSecret + nonce + curTime))
	return hex.EncodeToString(sum[:])
}

// RandomNonce 生成 32 位字母数字随机字符串。
func RandomNonce() string {
	out := make([]byte, nonceLength)
	limit := big.NewInt(int64(len(nonceAlphabet)))
	for i := range out {
		n, err := rand.Int(rand.Reader, limit)
		if err != nil {
			out[i] = nonceAlphabet[int(time.Now().UnixNano())%len(nonceAlphabet)]
			continue
		}
		out[i] = nonceAlphabet[n.Int64()]
	}
	return string(out)
}

func truncateTraceId(traceId string) string {
	if len(traceId) > maxTraceIdLength {
		return traceId[:maxTraceIdLength]
	}
	return traceId
}

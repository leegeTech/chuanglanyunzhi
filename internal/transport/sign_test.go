package transport

import "testing"

// 使用 docs/sms/1-API 调用说明.md 中的示例断言 CheckSum。
func TestCheckSum(t *testing.T) {
	appSecret := "a3K7mP9xQA"
	nonce := "36SuukDa291gfN4dTSNQ07lxFzP68KPB"
	curTime := "1779367802"
	want := "09dedc0e189b21e774686df7055c6bb003b0a128"

	if got := CheckSum(appSecret, nonce, curTime); got != want {
		t.Fatalf("CheckSum mismatch: got %s, want %s", got, want)
	}
}

func TestRandomNonce(t *testing.T) {
	n := RandomNonce()
	if len(n) != 32 {
		t.Fatalf("nonce length = %d, want 32", len(n))
	}
	// 两次生成应不相同（几乎必然）。
	n2 := RandomNonce()
	if n == n2 {
		t.Fatalf("two generated nonces are equal: %s", n)
	}
}

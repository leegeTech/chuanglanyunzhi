// Package retry 提供指数退避重试策略。
package retry

import (
	"context"
	"time"
)

// ExponentialBackoff 计算第 attempt 次(0 基)重试的等待间隔：2^attempt * base。
func ExponentialBackoff(attempt, baseMs int) time.Duration {
	ms := 1
	for i := 0; i < attempt; i++ {
		ms *= 2
	}
	return time.Duration(ms*baseMs) * time.Millisecond
}

// SleepCtx 返回 false 表示 ctx 已取消而提前结束等待。
func SleepCtx(ctx context.Context, d time.Duration) bool {
	select {
	case <-ctx.Done():
		return false
	case <-time.After(d):
		return true
	}
}

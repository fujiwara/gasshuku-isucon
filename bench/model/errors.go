package model

import (
	"context"
	"net"

	"github.com/isucon/isucandar"
	"github.com/isucon/isucandar/failure"
)

var (
	// 予期しないクリティカルなエラー
	ErrCritical failure.StringCode = "critical"

	// リクエスト失敗
	ErrRequestFailed failure.StringCode = "request_failed"
	// ステータスコードが不正
	ErrInvalidStatusCode failure.StringCode = "invalid_status_code"
	// レスポンスボディがデコード不可
	ErrUndecodableBody failure.StringCode = "undecodable_body"
	// レスポンスボディが間違っている
	ErrInvalidBody failure.StringCode = "invalid_body"

	// タイムアウト
	ErrTimeout failure.StringCode = "timeout"
)

func IsErrCritical(err error) bool {
	return failure.IsCode(err, isucandar.ErrPrepare) ||
		failure.IsCode(err, ErrCritical)
}

func IsErrTimeout(err error) bool {
	var nErr net.Error
	if failure.As(err, &nErr) && nErr.Timeout() {
		return true
	}
	if failure.Is(err, context.DeadlineExceeded) {
		return true
	}
	return failure.IsCode(err, failure.TimeoutErrorCode)
}

package contextgo

import (
	"context"
	"time"

	"golang.org/x/xerrors"
)

var currentTimeKey = "currentTime"

func CurrentTime(ctx context.Context) (time.Time, error) {
	if currentTime := ctx.Value(&currentTimeKey); currentTime != nil {
		return currentTime.(time.Time), nil
	}
	return time.Time{}, xerrors.New("currentTime not set")
}

func SetCurrentTime(ctx context.Context) context.Context {
	return context.WithValue(ctx, &currentTimeKey, time.Now().UTC())
}

// SetMockTime - テスト用にmockしたいtime.Timeをcontext.Valueに格納する関数
func SetMockTime(ctx context.Context, mockTime time.Time) context.Context {
	return context.WithValue(ctx, &currentTimeKey, mockTime)
}

// MockTime - テスト用に現在時刻をcontextに格納してそのまま取得する
func MockTime(ctx context.Context) time.Time {
	ctx = SetCurrentTime(ctx)
	return ctx.Value(&currentTimeKey).(time.Time)
}

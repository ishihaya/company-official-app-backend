package ulidgo

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// Generate - 渡された時刻からULIDを生成する
func Generate(t time.Time) (string, error) {
	timeStamp := ulid.Timestamp(t)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	ulid, err := ulid.New(timeStamp, entropy)
	if err != nil {
		return "", err
	}
	return ulid.String(), nil
}

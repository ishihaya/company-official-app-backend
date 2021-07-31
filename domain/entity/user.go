package entity

import "time"

// User - ユーザー情報
type User struct {
	ID        string    // ID(UUID)
	AuthID    string    // 認証のID(UUID)
	NickName  string    // ニックネーム
	CreatedAt time.Time // 作成日時
	UpdatedAt time.Time // 更新日時
}

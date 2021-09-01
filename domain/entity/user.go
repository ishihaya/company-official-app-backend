package entity

import "time"

// User - ユーザー
type User struct {
	ID        AppID     // ID
	AuthID    string    // 認証ID
	Nickname  string    // ニックネーム
	CreatedAt time.Time // 作成日時
	UpdatedAt time.Time // 更新日時
}

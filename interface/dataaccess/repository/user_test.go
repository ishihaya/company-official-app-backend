package repository

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/service/apperror"
	"github.com/ishihaya/company-official-app-backend/pkg/contextgo"
	"github.com/ishihaya/company-official-app-backend/pkg/db"
	"golang.org/x/xerrors"
)

func Test_userRepository_GetByAuthID(t *testing.T) {
	mockTime := contextgo.MockTime(context.Background())
	conn := db.GetInstance()
	t.Cleanup(func() {
		CleanUpRepositoryTest(t, conn, []string{"users"})
	})

	dummyUser := map[string]interface{}{
		"id":         "id",
		"auth_id":    "auth_id",
		"nick_name":  "nick_name",
		"created_at": mockTime,
		"updated_at": mockTime,
	}
	if _, err := conn.NamedExec("INSERT INTO users(id, auth_id, nick_name, created_at, updated_at) VALUES(:id, :auth_id, :nick_name, :created_at, :updated_at)", dummyUser); err != nil {
		t.Fatal(err)
	}

	type args struct {
		authID string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr error
	}{
		{
			name: "1 / 正常系",
			args: args{
				authID: "auth_id",
			},
			want: &entity.User{
				ID:        "id",
				AuthID:    "auth_id",
				NickName:  "nick_name",
				CreatedAt: mockTime.Round(time.Microsecond),
				UpdatedAt: mockTime.Round(time.Microsecond),
			},
			wantErr: nil,
		},
		{
			name: "2 / 準正常系 / ユーザーが見つからない",
			args: args{
				authID: "not_found",
			},
			want:    nil,
			wantErr: apperror.ErrUserNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			u := NewUserRepository(conn)

			got, err := u.FindByAuthID(tt.args.authID)

			if !xerrors.Is(err, tt.wantErr) {
				t.Errorf("userRepository.GetByAuthID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("userRepository.GetByAuthID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_userRepository_Store(t *testing.T) {
	mockTime := contextgo.MockTime(context.Background())
	conn := db.GetInstance()
	t.Cleanup(func() {
		CleanUpRepositoryTest(t, conn, []string{"users"})
	})

	type args struct {
		user *entity.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "1 / 正常系",
			args: args{
				user: &entity.User{
					ID:        "id",
					AuthID:    "auth_id",
					NickName:  "nick_name",
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserRepository(conn)

			err := u.Store(tt.args.user)

			if !xerrors.Is(err, tt.wantErr) {
				t.Errorf("userRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

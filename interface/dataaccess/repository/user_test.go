package repository

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/ishihaya/company-official-app-backend/application/customerror"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/infra/db"
	"github.com/ishihaya/company-official-app-backend/pkg/context"
	"golang.org/x/xerrors"
)

func Test_userRepository_GetByAuthID(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	timeContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.SetNow(timeContext)
	mockTime, err := context.Now(timeContext)
	if err != nil {
		t.Fatal(err)
	}
	conn := db.New()
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
			name: "正常系",
			args: args{
				authID: "auth_id",
			},
			want: &entity.User{
				ID:        "id",
				AuthID:    "auth_id",
				NickName:  "nick_name",
				CreatedAt: mockTime,
				UpdatedAt: mockTime,
			},
			wantErr: nil,
		},
		{
			name: "準正常系 / ユーザーが見つからない",
			args: args{
				authID: "not_found",
			},
			want:    nil,
			wantErr: customerror.ErrUserNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserRepository(conn)
			got, err := u.GetByAuthID(tt.args.authID)
			if !xerrors.Is(err, tt.wantErr) {
				t.Errorf("userRepository.GetByAuthID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("userRepository.GetByAuthID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

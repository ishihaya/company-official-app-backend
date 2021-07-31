package usecase

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/repository/mock_repository"
	"golang.org/x/xerrors"
)

func Test_userUsecase_Get(t *testing.T) {
	type fields struct {
		userRepositoryFn func(mock *mock_repository.MockUserRepository)
	}
	type args struct {
		authID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr error
	}{
		{
			name: "正常系",
			fields: fields{
				userRepositoryFn: func(mock *mock_repository.MockUserRepository) {
					mock.EXPECT().FindByAuthID("auth_id").Return(&entity.User{
						ID:        "id",
						AuthID:    "auth_id",
						NickName:  "nick_name",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)
				},
			},
			args: args{
				authID: "auth_id",
			},
			want: &entity.User{
				ID:        "id",
				AuthID:    "auth_id",
				NickName:  "nick_name",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepository := mock_repository.NewMockUserRepository(ctrl)
			tt.fields.userRepositoryFn(mockRepository)
			u := NewUserUsecase(mockRepository)

			got, err := u.Get(tt.args.authID)

			if !xerrors.Is(err, tt.wantErr) {
				t.Errorf("userUsecase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("userUsecase.Get() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_userUsecase_Create(t *testing.T) {
	type fields struct {
		userRepositoryFn func(mock *mock_repository.MockUserRepository)
	}
	type args struct {
		id       string
		authID   string
		nickName string
		now      time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "正常系",
			fields: fields{
				userRepositoryFn: func(mock *mock_repository.MockUserRepository) {
					mock.EXPECT().Store(&entity.User{
						ID:        "id",
						AuthID:    "auth_id",
						NickName:  "nick_name",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}).Return(nil)
				},
			},
			args: args{
				id:       "id",
				authID:   "auth_id",
				nickName: "nick_name",
				now:      time.Time{},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepository := mock_repository.NewMockUserRepository(ctrl)
			tt.fields.userRepositoryFn(mockRepository)
			u := NewUserUsecase(mockRepository)

			err := u.Create(tt.args.id, tt.args.authID, tt.args.nickName, tt.args.now)

			if !xerrors.Is(err, tt.wantErr) {
				t.Errorf("userUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

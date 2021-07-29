package usecase

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/repository/mock_repository"
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
		wantErr bool
	}{
		{
			name: "正常系",
			fields: fields{
				userRepositoryFn: func(mock *mock_repository.MockUserRepository) {
					mock.EXPECT().GetByAuthID("auth_id").Return(&entity.User{
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
			wantErr: false,
		},
		{
			name: "repositoryでエラーが返された場合エラーを返す",
			fields: fields{
				userRepositoryFn: func(mock *mock_repository.MockUserRepository) {
					mock.EXPECT().GetByAuthID("failed_auth_id").Return(nil, errors.New("something wrong"))
				},
			},
			args: args{
				authID: "failed_auth_id",
			},
			want: nil,
			wantErr: true,
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
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

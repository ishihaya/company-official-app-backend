package usecase

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator/mock_operator"
	"github.com/ishihaya/company-official-app-backend/domain/repository/mock_repository"
	"golang.org/x/xerrors"
)

func Test_userUsecase_Get(t *testing.T) {
	type fields struct {
		userRepositoryFn func(mock *mock_repository.MockUser)
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
			name: "1 / 正常系",
			fields: fields{
				userRepositoryFn: func(mock *mock_repository.MockUser) {
					mock.EXPECT().FindByAuthID("auth_id").Return(&entity.User{
						ID:        "id",
						AuthID:    "auth_id",
						Nickname:  "nick_name",
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
				Nickname:  "nick_name",
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
			mockRepository := mock_repository.NewMockUser(ctrl)
			tt.fields.userRepositoryFn(mockRepository)
			u := NewUserUsecase(mockRepository, nil)

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
	var id entity.AppID = "id"
	type fields struct {
		userRepositoryFn func(mock *mock_repository.MockUser)
		appIDOperatorFn  func(mock *mock_operator.MockAppIDOperator)
	}
	type args struct {
		authID      string
		nickname    string
		currentTime time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "1 / 正常系",
			fields: fields{
				userRepositoryFn: func(mock *mock_repository.MockUser) {
					mock.EXPECT().Store(&entity.User{
						ID:        id,
						AuthID:    "auth_id",
						Nickname:  "nick_name",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}).Return(nil)
				},
				appIDOperatorFn: func(mock *mock_operator.MockAppIDOperator) {
					mock.EXPECT().Generate(time.Time{}).Return(id, nil)
				},
			},
			args: args{
				authID:      "auth_id",
				nickname:    "nick_name",
				currentTime: time.Time{},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepository := mock_repository.NewMockUser(ctrl)
			tt.fields.userRepositoryFn(mockRepository)
			mockOperator := mock_operator.NewMockAppIDOperator(ctrl)
			tt.fields.appIDOperatorFn(mockOperator)
			u := NewUserUsecase(mockRepository, mockOperator)

			err := u.Create(tt.args.authID, tt.args.nickname, tt.args.currentTime)

			if !xerrors.Is(err, tt.wantErr) {
				t.Errorf("userUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

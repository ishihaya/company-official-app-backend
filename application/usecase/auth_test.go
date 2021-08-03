package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator/mock_operator"
	"golang.org/x/xerrors"
)

func Test_authUsecase_Get(t *testing.T) {
	type fields struct {
		authOperatorFn func(mock *mock_operator.MockAuthOperator)
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Auth
		wantErr error
	}{
		{
			name: "1 / 正常系",
			fields: fields{
				authOperatorFn: func(mock *mock_operator.MockAuthOperator) {
					mock.EXPECT().FindByToken("token").Return(&entity.Auth{
						ID: "id",
					}, nil)
				},
			},
			args: args{
				token: "token",
			},
			want: &entity.Auth{
				ID: "id",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockOperator := mock_operator.NewMockAuthOperator(ctrl)
			tt.fields.authOperatorFn(mockOperator)
			a := NewAuthUsecase(mockOperator)

			got, err := a.Get(tt.args.token)

			if !xerrors.Is(err, tt.wantErr) {
				t.Errorf("authUsecase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("authUsecase.Get() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

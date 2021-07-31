package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/ishihaya/company-official-app-backend/application/customerror"
	"github.com/ishihaya/company-official-app-backend/application/usecase/mock_usecase"
	"github.com/ishihaya/company-official-app-backend/config"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/infra/logger"
	"github.com/ishihaya/company-official-app-backend/pkg/context"
	"golang.org/x/xerrors"
)

func Test_userHandler_Get(t *testing.T) {
	logger.New(config.Log())
	type fields struct {
		userUsecaseFn func(mock *mock_usecase.MockUserUsecase)
	}
	tests := []struct {
		name           string
		fields         fields
		authID         string
		isExistAuthID  bool
		wantStatusCode int
		want           string
	}{
		{
			name: "正常系",
			fields: fields{
				userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {
					mock.EXPECT().Get("auth_id").Return(&entity.User{
						ID:        "id",
						AuthID:    "auth_id",
						NickName:  "nick_name",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)
				},
			},
			authID:         "auth_id",
			isExistAuthID:  true,
			wantStatusCode: 200,
			want:           `{"nickName":"nick_name"}`,
		},
		{
			name: "準正常系 / authIDが正常に取得できない場合にBad Requestを返す",
			fields: fields{
				userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {},
			},
			isExistAuthID:  false,
			wantStatusCode: 400,
			want:           fmt.Sprintf(`"%s"`, customerror.ErrGetAuthID.Error()),
		},
		{
			name: "準正常系 / ユーザーが見つからない場合にNot Foundを返す",
			fields: fields{
				userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {
					mock.EXPECT().Get("not_found").Return(nil, customerror.ErrUserNotFound)
				},
			},
			authID:         "not_found",
			isExistAuthID:  true,
			wantStatusCode: 404,
			want:           fmt.Sprintf(`"%s"`, customerror.ErrUserNotFound.Error()),
		},
		{
			name: "異常系 / 何らかの理由でユーザー取得に失敗する場合にサーバーエラーを返す",
			fields: fields{
				userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {
					mock.EXPECT().Get("error_auth_id").Return(nil, xerrors.New("something wrong"))
				},
			},
			authID:         "error_auth_id",
			isExistAuthID:  true,
			wantStatusCode: 500,
			want:           fmt.Sprintf(`"%s"`, customerror.ErrInternalServerError.Error()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUsecase := mock_usecase.NewMockUserUsecase(ctrl)
			tt.fields.userUsecaseFn(mockUsecase)
			u := NewUserHandler(mockUsecase)

			rec := httptest.NewRecorder()
			gin.SetMode(gin.ReleaseMode)
			c, _ := gin.CreateTestContext(rec)
			req := httptest.NewRequest(http.MethodGet, "/user", nil)
			c.Request = req
			if tt.isExistAuthID {
				context.SetAuthID(c, tt.authID)
			}

			u.Get(c)

			got := rec.Body.String()
			statusCode := rec.Code

			if statusCode != tt.wantStatusCode {
				t.Errorf("userHandler.Get() statusCode = %v, wantStatusCode = %v", statusCode, tt.wantStatusCode)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("userHandler.Get() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

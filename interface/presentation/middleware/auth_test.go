package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/ishihaya/company-official-app-backend/application/usecase/mock_usecase"
	"github.com/ishihaya/company-official-app-backend/config"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/service/apperror"
	"github.com/ishihaya/company-official-app-backend/pkg/logger"
	"golang.org/x/xerrors"
)

func Test_authMiddleware_AuthAPI(t *testing.T) {
	logger.New(config.Log())
	token1 := "token"
	wantAuthID1 := "auth_id"
	wantResponseBody2 := fmt.Sprintf(`"%s"`, apperror.ErrValidation.Error())
	wantStatusCode2 := 400
	wantResponseBody3 := fmt.Sprintf(`"%s"`, apperror.ErrInternalServerError.Error())
	wantStatusCode3 := 500
	type fields struct {
		authUsecaseFn func(mock *mock_usecase.MockAuthUsecase)
	}
	tests := []struct {
		name             string
		fields           fields
		token            *string
		wantAuthID       *string
		wantResponseBody *string
		wantStatusCode   *int
	}{
		{
			name: "1 / 正常系",
			fields: fields{
				authUsecaseFn: func(mock *mock_usecase.MockAuthUsecase) {
					mock.EXPECT().Get("token").Return(&entity.Auth{
						ID: "id",
					}, nil)
				},
			},
			token:            &token1,
			wantAuthID:       &wantAuthID1,
			wantResponseBody: nil,
			wantStatusCode:   nil,
		},
		{
			name: "2 / 準正常系 / headerにtokenがセットされていない場合Bad Request",
			fields: fields{
				authUsecaseFn: func(mock *mock_usecase.MockAuthUsecase) {},
			},
			token:            nil,
			wantAuthID:       nil,
			wantResponseBody: &wantResponseBody2,
			wantStatusCode:   &wantStatusCode2,
		},
		{
			name: "3 / 異常系 / 認証情報の取得に失敗した場合Server Error",
			fields: fields{
				authUsecaseFn: func(mock *mock_usecase.MockAuthUsecase) {
					mock.EXPECT().Get("token").Return(nil, xerrors.New("something wrong"))
				},
			},
			token:            &token1,
			wantAuthID:       nil,
			wantResponseBody: &wantResponseBody3,
			wantStatusCode:   &wantStatusCode3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUsecase := mock_usecase.NewMockAuthUsecase(ctrl)
			tt.fields.authUsecaseFn(mockUsecase)
			a := NewAuthMiddleware(mockUsecase)

			rec := httptest.NewRecorder()
			gin.SetMode(gin.ReleaseMode)
			c, _ := gin.CreateTestContext(rec)
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			c.Request = req
			if tt.token != nil {
				c.Request.Header.Set("Authorization", *tt.token)
			}

			a.AuthAPI(c)

			// NOTE 明日中にこのTODOを終わらせる
			// TODO1 この下を記述する
			// TODO2 operator, pkg, di, routerを記述する
			// TODO3 ローカル環境立ち上げられるか&Swaggerの起動を確認する
			// TODO4 フロントエンドへ

			// got := rec.Body.String()
			// statusCode := rec.Code

			// if statusCode != *tt.wantStatusCode {
			// 	t.Errorf("authMiddleware.AuthAPI() statusCode = %v, wantStatusCode = %v", statusCode, tt.wantStatusCode)
			// 	return
			// }
			// if diff := cmp.Diff(tt.want, got); diff != "" {
			// 	t.Errorf("authMiddleware.AuthAPI() mismatch (-want +got):\n%s", diff)
			// }
		})
	}
}

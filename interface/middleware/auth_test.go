package middleware

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/ishihaya/company-official-app-backend/application/usecase/mock_usecase"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/service/apperror"
	"github.com/ishihaya/company-official-app-backend/pkg/contextgo"
	"golang.org/x/xerrors"
)

func Test_auth_middleware_AuthAPI(t *testing.T) {
	token1 := "Bearer token"
	wantAuthID1 := "auth_id"
	// wantResponseBody2 := fmt.Sprintf(`"%s"`, apperror.ErrValidation.Error())
	// wantStatusCode2 := 400
	wantResponseBody3 := fmt.Sprintf("\"%s\"\n", apperror.ErrInternalServerError.Error())
	wantStatusCode3 := 500
	type fields struct {
		authUsecaseFn func(mock *mock_usecase.MockAuth)
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
				authUsecaseFn: func(mock *mock_usecase.MockAuth) {
					mock.EXPECT().Get(context.Background(), "token").Return(&entity.Auth{
						ID: "auth_id",
					}, nil)
				},
			},
			token:            &token1,
			wantAuthID:       &wantAuthID1,
			wantResponseBody: nil,
			wantStatusCode:   nil,
		},
		// {
		// 	name: "2 / 準正常系 / headerにtokenがセットされていない場合Bad Request",
		// 	fields: fields{
		// 		authUsecaseFn: func(mock *mock_usecase.MockAuth) {},
		// 	},
		// 	token:            nil,
		// 	wantAuthID:       nil,
		// 	wantResponseBody: &wantResponseBody2,
		// 	wantStatusCode:   &wantStatusCode2,
		// },
		{
			name: "3 / 異常系 / 認証情報の取得に失敗した場合Server Error",
			fields: fields{
				authUsecaseFn: func(mock *mock_usecase.MockAuth) {
					mock.EXPECT().Get(context.Background(), "token").Return(nil, xerrors.New("something wrong"))
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
			mockUsecase := mock_usecase.NewMockAuth(ctrl)
			tt.fields.authUsecaseFn(mockUsecase)
			a := NewAuth(mockUsecase)

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.token != nil {
				req.Header.Set("Authorization", *tt.token)
			}
			res := httptest.NewRecorder()

			ctx := req.Context()
			testHandlerFn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// After execute AuthAPI
				ctx = r.Context()
			})
			fn := a.AuthAPI(testHandlerFn)
			fn.ServeHTTP(res, req)

			authID, err := contextgo.AuthID(ctx)
			if err != nil && tt.wantAuthID == nil {
				// エラーが出ることを期待する
				got := res.Body.String()
				status := res.Code
				if diff := cmp.Diff(*tt.wantResponseBody, got); diff != "" {
					t.Errorf("authMiddleware.AuthAPI() mismatch (-want +got):\n%s", diff)
				}
				if status != *tt.wantStatusCode {
					t.Errorf("authMiddleware.AuthAPI() statusCode = %v, wantStatusCode = %v", status, *tt.wantStatusCode)
				}
			} else if tt.wantAuthID != nil {
				if diff := cmp.Diff(*tt.wantAuthID, authID); diff != "" {
					t.Errorf("authMiddleware.AuthAPI() mismatch (-want +got):\n%s", diff)
				}
			} else {
				t.Errorf("authMiddleware.AuthAPI() authID = %v, wantAuthID = %v", authID, *tt.wantAuthID)
			}
		})
	}
}

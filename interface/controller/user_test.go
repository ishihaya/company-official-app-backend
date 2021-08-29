package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/ishihaya/company-official-app-backend/application/usecase/mock_usecase"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/service/apperror"
	"github.com/ishihaya/company-official-app-backend/pkg/contextgo"
	"golang.org/x/xerrors"
)

func Test_userHandler_Get(t *testing.T) {
	authID1 := "auth_id"
	authID3 := "not_found"
	authID4 := "error_auth_id"
	type fields struct {
		userUsecaseFn func(mock *mock_usecase.MockUserUsecase)
	}
	tests := []struct {
		name           string
		fields         fields
		authID         *string
		want           string
		wantStatusCode int
	}{
		{
			name: "1 / 正常系",
			fields: fields{
				userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {
					mock.EXPECT().Get("auth_id").Return(&entity.User{
						ID:        "id",
						AuthID:    "auth_id",
						Nickname:  "nick_name",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)
				},
			},
			authID:         &authID1,
			want:           "{\"nickname\":\"nick_name\"}\n",
			wantStatusCode: http.StatusOK,
		},
		{
			name: "2 / 準正常系 / authIDが正常に取得できない場合にBad Requestを返す",
			fields: fields{
				userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {},
			},
			authID:         nil,
			want:           fmt.Sprintf("\"%s\"\n", apperror.ErrGetAuthID.Error()),
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "3 / 準正常系 / ユーザーが見つからない場合にNot Foundを返す",
			fields: fields{
				userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {
					mock.EXPECT().Get("not_found").Return(nil, apperror.ErrUserNotFound)
				},
			},
			authID:         &authID3,
			want:           fmt.Sprintf("\"%s\"\n", apperror.ErrUserNotFound.Error()),
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "4 / 異常系 / 何らかの理由でユーザー取得に失敗する場合にサーバーエラーを返す",
			fields: fields{
				userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {
					mock.EXPECT().Get("error_auth_id").Return(nil, xerrors.New("something wrong"))
				},
			},
			authID:         &authID4,
			want:           fmt.Sprintf("\"%s\"\n", apperror.ErrInternalServerError.Error()),
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUsecase := mock_usecase.NewMockUserUsecase(ctrl)
			tt.fields.userUsecaseFn(mockUsecase)
			u := NewUserController(mockUsecase)

			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/user", nil)
			ctx := req.Context()
			if tt.authID != nil {
				ctx = contextgo.SetAuthID(ctx, *tt.authID)
			}
			req = req.WithContext(ctx)

			u.Get(res, req)

			got := res.Body.String()
			statusCode := res.Code

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

func Test_userHandler_Create(t *testing.T) {
	requestBody1 := `{"nickname":"nick_name"}`
	authID1 := "auth_id"
	ct1 := time.Now().UTC()

	type fields struct {
		userUsecaseFn func(mock *mock_usecase.MockUserUsecase)
	}
	tests := []struct {
		name           string
		fields         fields
		requestBody    *string
		currentTime    *time.Time
		authID         *string
		want           string
		wantStatusCode int
	}{
		{
			name: "1 / 正常系",
			fields: fields{
				userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {
					mock.EXPECT().Create("auth_id", "nick_name", ct1).Return(nil)
				},
			},
			requestBody:    &requestBody1,
			currentTime:    &ct1,
			authID:         &authID1,
			want:           "null\n",
			wantStatusCode: http.StatusNoContent,
		},
		// {
		// 	name: "2 / 準正常系 / リクエストボディが空の場合Bad Request",
		// 	fields: fields{
		// 		userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {},
		// 	},
		// 	requestBody:    nil,
		// 	want:           fmt.Sprintf("\"%s\"\n"`, apperror.ErrValidation.Error()),
		// 	wantStatusCode: http.StatusBadRequest,
		// },
		{
			name: "3 / 準正常系 / 現在時刻が取得できない場合Bad Request",
			fields: fields{
				userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {},
			},
			requestBody:    &requestBody1,
			currentTime:    nil,
			want:           fmt.Sprintf("\"%s\"\n", apperror.ErrGetTime.Error()),
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "4 / 準正常系 / authIDが取得できない場合Bad Request",
			fields: fields{
				userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {},
			},
			requestBody:    &requestBody1,
			currentTime:    &ct1,
			authID:         nil,
			want:           fmt.Sprintf("\"%s\"\n", apperror.ErrGetAuthID.Error()),
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "5 / 異常系 / ユーザー取得に失敗した場合Internal Server Error",
			fields: fields{
				userUsecaseFn: func(mock *mock_usecase.MockUserUsecase) {
					mock.EXPECT().Create("auth_id", "nick_name", ct1).Return(xerrors.New("something wrong"))
				},
			},
			requestBody:    &requestBody1,
			currentTime:    &ct1,
			authID:         &authID1,
			want:           fmt.Sprintf("\"%s\"\n", apperror.ErrInternalServerError.Error()),
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUsecase := mock_usecase.NewMockUserUsecase(ctrl)
			tt.fields.userUsecaseFn(mockUsecase)
			u := NewUserController(mockUsecase)

			res := httptest.NewRecorder()
			var req *http.Request
			if tt.requestBody != nil {
				req = httptest.NewRequest(http.MethodGet, "/user", strings.NewReader(*tt.requestBody))
			} else {
				req = httptest.NewRequest(http.MethodGet, "/user", nil)
			}
			ctx := req.Context()
			if tt.authID != nil {
				ctx = contextgo.SetAuthID(ctx, *tt.authID)
			}
			if tt.currentTime != nil {
				ctx = contextgo.SetMockTime(ctx, *tt.currentTime)
			}
			req = req.WithContext(ctx)

			u.Create(res, req)

			got := res.Body.String()
			statusCode := res.Code

			if statusCode != tt.wantStatusCode {
				t.Errorf("userHandler.Create() statusCode = %v, wantStatusCode = %v", statusCode, tt.wantStatusCode)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("userHandler.Create() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

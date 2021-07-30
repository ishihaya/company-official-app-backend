package repository

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/infra/db"
)

func Test_userRepository_GetByAuthID(t *testing.T) {
	conn := db.New()
	// TODO: TimeをMockする
	type args struct {
		authID string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserRepository(conn)
			got, err := u.GetByAuthID(tt.args.authID)
			if (err != nil) != tt.wantErr {
				// TODO: 
				// Repositoryのテストにおいて、customerrorを返しているかを検証したい
				// →どのエラーが返ってきているかまでを確認した方がよさそう？
				// →理由は別のエラーがたまたま返ってきているだけの可能性があるため
				t.Errorf("userRepository.GetByAuthID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("userRepository.GetByAuthID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

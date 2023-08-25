package service

import (
	"context"
	"github.com/shshimamo/knowledge-main/middlewares"
	mockrepository "github.com/shshimamo/knowledge-main/mock/repository"
	"github.com/shshimamo/knowledge-main/model"
	"go.uber.org/mock/gomock"
	"testing"

	"github.com/google/go-cmp/cmp"
	gql "github.com/shshimamo/knowledge-main/graph/model"
)

func TestCreateUser(t *testing.T) {
	username := "tester"
	type args struct {
		ctx    context.Context
		gqlNew *gql.NewUser
	}

	tests := map[string]struct {
		repoReturnUser *model.User
		args           *args
		want           *gql.User
		wantErr        bool
	}{
		"no-token": {nil, &args{ctx: context.Background(), gqlNew: &gql.NewUser{Name: "tester"}}, nil, true},
		"valid-args": {
			&model.User{ID: 1, AuthUserID: 1, Name: "tester"},
			&args{ctx: context.WithValue(context.Background(), middlewares.CurrentTokenKey{}, &model.Token{AuthUserID: 1}), gqlNew: &gql.NewUser{Name: "tester"}},
			&gql.User{ID: "1", AuthUserID: "1", Name: &username}, false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockrepository.NewMockUserRepository(ctrl)
			if tt.repoReturnUser != nil {
				inUser := *tt.repoReturnUser
				inUser.ID = 0
				mockRepo.EXPECT().CreateUser(gomock.Any(), &inUser).Return(tt.repoReturnUser, nil)
			}

			service := newUserService(mockRepo)

			got, err := service.CreateUser(tt.args.ctx, tt.args.gqlNew)
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr: %v, err: %v", tt.wantErr, err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}

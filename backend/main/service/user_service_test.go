package service

import (
	"context"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shshimamo/knowledge/main/middlewares"
	mockrepository "github.com/shshimamo/knowledge/main/mock/repository"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/repository"
	"github.com/shshimamo/knowledge/main/utils"
	"go.uber.org/mock/gomock"
	"testing"

	"github.com/google/go-cmp/cmp"
	gql "github.com/shshimamo/knowledge/main/graph/model"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		ctx    context.Context
		gqlNew *gql.NewUser
	}

	baseCtx := context.Background()
	// Context without token
	noTokenCtx := baseCtx

	// Context with token
	withTokenCtx := context.WithValue(baseCtx, middlewares.CurrentTokenKey{}, &model.Token{AuthUserID: 1})

	// Context with token and existing user
	withTokenAndUserCtx := context.WithValue(withTokenCtx, middlewares.CurrentUserKey{}, &model.User{ID: 1, AuthUserID: 1, Name: "tester"})

	// Context with no AuthUserID token
	withNoAuthUserIDTokenCtx := context.WithValue(baseCtx, middlewares.CurrentTokenKey{}, &model.Token{})

	userName := "tester"
	tests := map[string]struct {
		args           *args
		repoReturnUser *model.User
		want           *gql.User
		wantErr        bool
	}{
		"valid-args": {
			&args{ctx: withTokenCtx, gqlNew: &gql.NewUser{Name: userName}},
			&model.User{ID: 1, AuthUserID: 1, Name: userName},
			&gql.User{ID: "1", AuthUserID: "1", Name: &userName},
			false,
		},
		"no-token":            {&args{ctx: noTokenCtx, gqlNew: &gql.NewUser{Name: userName}}, nil, nil, true},
		"current-user-exists": {&args{ctx: withTokenAndUserCtx, gqlNew: &gql.NewUser{Name: userName}}, nil, nil, true},
		"no-name-user":        {&args{ctx: withTokenCtx, gqlNew: &gql.NewUser{}}, nil, nil, true},
		"invalid-token":       {&args{ctx: withNoAuthUserIDTokenCtx, gqlNew: &gql.NewUser{Name: userName}}, nil, nil, true},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			t.Cleanup(func() { ctrl.Finish() })

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

func TestCreateUserNoMock(t *testing.T) {
	type args struct {
		ctx    context.Context
		gqlNew *gql.NewUser
	}

	baseCtx := context.Background()
	// Context without token
	noTokenCtx := baseCtx

	// Context with token
	withTokenCtx := context.WithValue(baseCtx, middlewares.CurrentTokenKey{}, &model.Token{AuthUserID: 1})

	// Context with token and existing user
	withTokenAndUserCtx := context.WithValue(withTokenCtx, middlewares.CurrentUserKey{}, &model.User{ID: 1, AuthUserID: 1, Name: "tester"})

	// Context with no AuthUserID token
	withNoAuthUserIDTokenCtx := context.WithValue(baseCtx, middlewares.CurrentTokenKey{}, &model.Token{})

	userName := "tester"
	tests := map[string]struct {
		args    *args
		want    *gql.User
		wantErr bool
	}{
		"valid-args": {
			&args{ctx: withTokenCtx, gqlNew: &gql.NewUser{Name: userName}},
			&gql.User{AuthUserID: "1", Name: &userName},
			false,
		},
		"no-token":            {&args{ctx: noTokenCtx, gqlNew: &gql.NewUser{Name: userName}}, nil, true},
		"current-user-exists": {&args{ctx: withTokenAndUserCtx, gqlNew: &gql.NewUser{Name: userName}}, nil, true},
		"no-name-user":        {&args{ctx: withTokenCtx, gqlNew: &gql.NewUser{}}, nil, true},
		"invalid-token":       {&args{ctx: withNoAuthUserIDTokenCtx, gqlNew: &gql.NewUser{Name: userName}}, nil, true},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			service := newUserService(setupUserRepository(t))

			got, err := service.CreateUser(tt.args.ctx, tt.args.gqlNew)
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr: %v, err: %v", tt.wantErr, err)
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(gql.User{}, "ID")); diff != "" {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}

func setupUserRepository(t *testing.T) repository.UserRepository {
	t.Helper()
	db, err := utils.SetupDatabase(model.Test)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_ = tx.Rollback()
		_ = db.Close()
	})

	repo := repository.NewUserRepository(tx)

	return repo
}

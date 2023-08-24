package repository

import (
	"context"
	"database/sql"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/shshimamo/knowledge-main/utils"
)

func setupUserRepository(t *testing.T) (*sql.DB, *sql.Tx, UserRepository) {
	t.Helper()
	db, err := utils.SetupDatabase(model.Test)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(tx)

	return db, tx, repo
}

func TestCreateUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := map[string]struct {
		args    *args
		want    *model.User
		wantErr bool
	}{
		"AuthUserID and Name": {&args{ctx: context.Background(), user: &model.User{AuthUserID: 1, Name: "test"}}, &model.User{AuthUserID: 1, Name: "test"}, false},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, tx, repo := setupUserRepository(t)
			defer func() { _ = db.Close() }()
			defer func() { _ = tx.Rollback() }()

			got, err := repo.CreateUser(tt.args.ctx, tt.args.user)

			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr: %v, err: %v", tt.wantErr, err)
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(model.User{}, "ID")); diff != "" {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
			if got.ID == 0 {
				t.Errorf("Expected got.ID is not zero, but zero")
			}
		})
	}
}

func TestGetUserByToken(t *testing.T) {
	type args struct {
		ctx   context.Context
		token *model.Token
	}
	tests := map[string]struct {
		seed    *model.User
		args    *args
		want    *model.User
		wantErr bool
	}{
		"AuthUserID and Name": {
			&model.User{AuthUserID: 1, Name: "test"},
			&args{ctx: context.Background(), token: &model.Token{AuthUserID: 1}},
			&model.User{AuthUserID: 1, Name: "test"},
			false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, tx, repo := setupUserRepository(t)
			defer func() { _ = db.Close() }()
			defer func() { _ = tx.Rollback() }()

			// MEMO: use fixture?
			_, err := repo.CreateUser(context.Background(), tt.seed)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repo.GetUserByToken(tt.args.ctx, tt.args.token)

			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr: %v, err: %v", tt.wantErr, err)
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(model.User{}, "ID")); diff != "" {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}

package repository

import (
	"context"
	"database/sql"
	"github.com/google/go-cmp/cmp/cmpopts"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/shshimamo/knowledge-main/utils"
)

func setupUserRepository(t *testing.T) (*sql.DB, *sql.Tx, UserRepository) {
	db, err := utils.SetupDatabase(model.Test)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(tx)

	return db, tx, repo
}

func TestCreateUser(t *testing.T) {
	tests := map[string]struct {
		user *model.User
		want *model.User
	}{
		"AuthUserID and Name": {user: &model.User{AuthUserID: 1, Name: "test"}, want: &model.User{AuthUserID: 1, Name: "test"}},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, tx, repo := setupUserRepository(t)
			defer func() { _ = db.Close() }()
			defer func() { _ = tx.Rollback() }()

			got, err := repo.CreateUser(context.Background(), tt.user)
			if err != nil {
				t.Fatal(err)
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
	tests := map[string]struct {
		user  *model.User
		token *model.Token
		want  *model.User
	}{
		"AuthUserID and Name": {user: &model.User{AuthUserID: 1, Name: "test"}, token: &model.Token{AuthUserID: 1}, want: &model.User{AuthUserID: 1, Name: "test"}},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, tx, repo := setupUserRepository(t)
			defer func() { _ = db.Close() }()
			defer func() { _ = tx.Rollback() }()

			// MEMO: use fixture?
			_, err := repo.CreateUser(context.Background(), tt.user)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repo.GetUserByToken(context.Background(), tt.token)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(model.User{}, "ID")); diff != "" {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}

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
	tests := []struct {
		name string
		user *model.User
		want *model.User
	}{
		{name: "AuthUserID and Name", user: &model.User{AuthUserID: 1, Name: "test"}, want: &model.User{AuthUserID: 1, Name: "test"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			db, tx, repo := setupUserRepository(t)
			defer db.Close()
			defer tx.Rollback()

			got, err := repo.CreateUser(context.Background(), test.user)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(got, test.want, cmpopts.IgnoreFields(model.User{}, "ID")); diff != "" {
				t.Errorf("%v: want: %v, but %v", test.name, test.want, got)
			}
			if got.ID == 0 {
				t.Errorf("%v: expect got.ID is not zero, but zero", test.name)
			}
		})
	}
}

func TestGetUserByToken(t *testing.T) {
	tests := []struct {
		name  string
		user  *model.User
		token *model.Token
		want  *model.User
	}{
		{name: "AuthUserID and Name", user: &model.User{AuthUserID: 1, Name: "test"}, token: &model.Token{AuthUserID: 1}, want: &model.User{AuthUserID: 1, Name: "test"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			db, tx, repo := setupUserRepository(t)
			defer db.Close()
			defer tx.Rollback()

			// MEMO: use fixture?
			_, err := repo.CreateUser(context.Background(), test.user)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repo.GetUserByToken(context.Background(), test.token)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, test.want, cmpopts.IgnoreFields(model.User{}, "ID")); diff != "" {
				t.Errorf("%v: want: %v, but %v", test.name, test.want, got)
			}
		})
	}
}

package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/shshimamo/knowledge-main/repository"
	"github.com/shshimamo/knowledge-main/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestHandler(t *testing.T) (http.Handler, *sql.Tx) {
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

	return setupHandler(tx, model.Test), tx
}

func TestUserQuery(t *testing.T) {
	handler, tx := setupTestHandler(t)

	// Seed
	seedUser, err := repository.NewUserRepository(tx).CreateUser(context.Background(), &model.User{AuthUserID: 1, Name: "tester"})
	if err != nil {
		t.Fatal(err)
	}

	// Create Body
	query := fmt.Sprintf(`{
		user(id: "%d") {
			id
			authUserId
			name
		}
    }`, seedUser.ID)
	payload := map[string]string{
		"query": query,
	}
	jsonData, _ := json.Marshal(payload)
	body := bytes.NewReader(jsonData)

	// Request
	req := httptest.NewRequest(http.MethodPost, "/query", body)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	resp := rec.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("want: http.StatusOK, got: %v", resp.StatusCode)
	}

	var responseBody map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&responseBody)

	data, _ := responseBody["data"].(map[string]interface{})
	user, _ := data["user"].(map[string]interface{})

	if user["authUserId"] != "1" {
		t.Errorf("want: 1, got: %v", user["authUserId"])
	}
	if user["name"] != "tester" {
		t.Errorf("want: tester, got: %v", user["name"])
	}
}

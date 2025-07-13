package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/repository"
	"github.com/shshimamo/knowledge/main/utils"
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

func gqlRequest(t *testing.T, query string, handler http.Handler) map[string]interface{} {
	t.Helper()

	// Create Body
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

	// Response
	resp := rec.Result()
	
	var responseBody map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&responseBody)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want: http.StatusOK, got: %v", resp.StatusCode)
		t.Errorf("Failed to cast response body: %v", responseBody)
	}

	data, ok := responseBody["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("Failed to cast response body: %v", responseBody)
	}

	return data
}

func TestUserQuery(t *testing.T) {
	handler, tx := setupTestHandler(t)

	// Seed
	seedUser, err := repository.NewUserRepository(tx).CreateUser(context.Background(), &model.User{AuthUserID: 1, Name: "tester"})
	if err != nil {
		t.Fatal(err)
	}

	query := fmt.Sprintf(`{
		user(id: "%d") {
			id
			authUserId
			name
		}
    }`, seedUser.ID)

	data := gqlRequest(t, query, handler)
	user, _ := data["user"].(map[string]interface{})

	if user["authUserId"].(string) != "1" {
		t.Errorf("want: 1, got: %v", user["authUserId"])
	}
	if user["name"].(string) != "tester" {
		t.Errorf("want: tester, got: %v", user["name"])
	}
}

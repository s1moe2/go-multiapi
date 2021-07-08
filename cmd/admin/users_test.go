package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	domain "multiapi"
	"multiapi/mock"
	"multiapi/pkg/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsersHandler_Get(t *testing.T) {
	t.Run("expect GET /users to return 200 and a list of users", func(t *testing.T) {
		userService := &mock.UserRepository{}
		userService.GetAllFn = func(_ context.Context) ([]*domain.User, error) {
			return []*domain.User{
				{
					ID:       "1",
					Name:     "Mack",
					Email:    "mack@acme.com",
					Password: "strong",
				},
			}, nil
		}
		uh := NewUsersHandler(userService)

		r := httptest.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()
		router := test.PrepareRouter(http.MethodGet, "/users", uh.Get)
		router.ServeHTTP(w, r)

		resp := w.Result()

		test.AssertStatusCode(t, resp, http.StatusOK)
		test.AssertJSONContentType(t, resp)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("failed to read response body")
		}

		var users []*domain.User
		err = json.Unmarshal(body, &users)
		if err != nil {
			t.Fatal("failed to parse response body")
		}
	})

	t.Run("expect GET /users to return 200 and an empty list of users ", func(t *testing.T) {
		userService := &mock.UserRepository{}
		userService.GetAllFn = func(_ context.Context) ([]*domain.User, error) {
			return []*domain.User{}, nil
		}
		uh := NewUsersHandler(userService)

		r := httptest.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()
		router := test.PrepareRouter(http.MethodGet, "/users", uh.Get)
		router.ServeHTTP(w, r)

		resp := w.Result()

		test.AssertStatusCode(t, resp, http.StatusOK)
		test.AssertJSONContentType(t, resp)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("failed to read response body")
		}

		var users []*domain.User
		err = json.Unmarshal(body, &users)
		if err != nil {
			t.Fatal("failed to parse response body")
		}
		if len(users) > 0 {
			t.Fatal("response should have no records")
		}
	})
}

func TestUsersHandler_GetByID(t *testing.T) {
	t.Run("expect GET /users/{id} to return 200 and a user", func(t *testing.T) {
		userService := &mock.UserRepository{}
		userService.FindByIDFn = func(_ context.Context, ID string) (*domain.User, error) {
			return &domain.User{
				ID:       "1",
				Name:     "Mack",
				Email:    "mack@acme.com",
				Password: "strong",
			}, nil
		}
		uh := NewUsersHandler(userService)

		r := httptest.NewRequest("GET", "/users/1", nil)
		w := httptest.NewRecorder()
		router := test.PrepareRouter(http.MethodGet, "/users/{id}", uh.GetByID)
		router.ServeHTTP(w, r)

		resp := w.Result()

		test.AssertStatusCode(t, resp, http.StatusOK)
		test.AssertJSONContentType(t, resp)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("failed to read response body")
		}

		var user domain.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			t.Fatal("failed to parse response body")
		}
	})

	t.Run("expect GET /users/{id} to return 404 when user not found", func(t *testing.T) {
		userService := &mock.UserRepository{}
		userService.FindByIDFn = func(_ context.Context, ID string) (*domain.User, error) {
			return nil, nil
		}
		uh := NewUsersHandler(userService)

		r := httptest.NewRequest("GET", "/users/1", nil)
		w := httptest.NewRecorder()
		router := test.PrepareRouter(http.MethodGet, "/users/{id}", uh.GetByID)
		router.ServeHTTP(w, r)

		resp := w.Result()

		test.AssertStatusCode(t, resp, http.StatusNotFound)
	})
}

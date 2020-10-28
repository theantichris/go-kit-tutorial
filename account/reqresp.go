package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	// CreateUserRequest holds the information for the create user request.
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// CreateUserResponse holds the information for the create user response.
	CreateUserResponse struct {
		OK string `json:"ok"`
	}

	// GetUserRequest holds the information for the get user request.
	GetUserRequest struct {
		ID string `json:"id"`
	}

	// GetUserResponse holds the information for the get user response.
	GetUserResponse struct {
		Email string `json:"email"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	req := GetUserRequest{
		ID: vars["id"],
	}

	return req, nil
}

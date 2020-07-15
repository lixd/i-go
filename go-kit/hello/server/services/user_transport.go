package services

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	userIdStr := r.URL.Query().Get("userId")
	if userIdStr != "" {
		userId, _ := strconv.Atoi(userIdStr)
		return UserRequest{UserId: userId}, nil
	}
	return nil, errors.New("invalid params")
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

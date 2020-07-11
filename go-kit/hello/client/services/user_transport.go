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

func GetUserInfoReq(ctx context.Context, request *http.Request, r interface{}) error {
	userRequest := r.(UserRequest)
	q := request.URL.Query()
	q.Add("userId", strconv.Itoa(userRequest.UserId))
	request.URL.RawQuery = q.Encode()
	return nil
}
func GetUserInfoResp(ctx context.Context, resp *http.Response) (response interface{}, err error) {
	if resp.StatusCode > 400 {
		return nil, errors.New("no data")
	}
	var userResponse UserRequest
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	if err != nil {
		return nil, err
	}
	return userResponse, nil
}

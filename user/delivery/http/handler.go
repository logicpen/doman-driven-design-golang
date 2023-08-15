package http

import (
	"context"
	"ddd-golang/domain"
	"ddd-golang/middleware"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type UserHandler struct {
	userUserCase domain.UserUseCase
}

func NewUserHandler(mux *http.ServeMux, userUserCase domain.UserUseCase) {
	userHandler := UserHandler{
		userUserCase: userUserCase,
	}
	mux.Handle("/user/create", middleware.CreateMiddlewareChain(middleware.SetContentTypeToJson).Then(userHandler.CreateUser()))
	mux.Handle("/user/get/", middleware.CreateMiddlewareChain(middleware.SetContentTypeToJson).Then(userHandler.GetUserById()))
}

func (u *UserHandler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Fprintln(w, "only POST method is allowed")
			return
		}
		// get user from request body and pass it for creation
		reqBody := r.Body
		user := domain.User{}
		_ = json.NewDecoder(reqBody).Decode(&user)
		createdUser, err := u.userUserCase.CreateUser(context.TODO(), user)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		response, _ := json.Marshal(createdUser)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(string(response))
	}
}

func (u *UserHandler) GetUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			fmt.Fprintln(w, "only GET method is allowed")
			return
		}

		emailId := strings.TrimPrefix(r.URL.Path, "/user/get/")
		user, _ := u.userUserCase.GetUserByEmailId(context.TODO(), emailId)
		response, _ := json.Marshal(user)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(string(response))
	}
}

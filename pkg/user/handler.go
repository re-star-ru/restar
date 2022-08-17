package user

import (
	"net/http"
	"restar/pkg/domain"
)

type IUser interface {
	Create(user *domain.User) error
}

type HttpHandler struct {
	userUsecase IUser
}

func NewHttpHandler(userUsecase IUser) *HttpHandler {
	return &HttpHandler{
		userUsecase: userUsecase,
	}
}

func (u *HttpHandler) Create(w http.ResponseWriter, r *http.Request) {
	u.userUsecase.Create(&domain.User{
		Name: "1",
	})
}

package user

import "restar/pkg/domain"

type Usecase struct {
}

func NewUserUsecase() *Usecase {
	return &Usecase{}
}

func (u *Usecase) Create(user *domain.User) error {
	println("Hello User")
	return nil
}

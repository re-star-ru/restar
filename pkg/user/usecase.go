package user

import "restar/pkg/domain"

type Usecase struct{}

func NewUserUsecase() *Usecase {
	return &Usecase{}
}

func (u *Usecase) Create(user domain.User) (*domain.User, error) {
	panic("user create not implemented")
}

func (u *Usecase) List() (*domain.User, error) {
	panic("user list not implemented")
}

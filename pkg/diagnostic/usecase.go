package diagnostic

type Usecase struct {
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (uc *Usecase) Create() {}

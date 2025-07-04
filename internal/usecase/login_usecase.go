package usecase

type LoginUseCase struct {
}

func (uc *LoginUseCase) Execute(email, password string) {
	// err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	// return err == nil
}

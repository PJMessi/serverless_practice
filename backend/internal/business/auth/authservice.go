package auth

import (
	"sharedlambdacode/internal/excep"
	"sharedlambdacode/internal/helper/pwhelper"
)

type AuthService interface {
	PhoneSignIn(phone string, encodedPassword string) (SignInRes, error)
}

func NewAuthService() AuthService {
	return &AuthServiceImpl{}
}

type SignInRes struct {
	Token string `json:"token"`
}

type AuthServiceImpl struct {
}

func (a *AuthServiceImpl) PhoneSignIn(phone string, encodedPassword string) (SignInRes, error) {
	plainPw, err := pwhelper.DecodePw(encodedPassword)
	if err != nil || plainPw == "" {
		return SignInRes{}, excep.NewAuthExcep("invalid credentials")
	}
	if plainPw == "" {
		return SignInRes{}, excep.NewAuthExcep("invalid credentials")
	}

	if plainPw != "Password123!" {
		return SignInRes{}, excep.NewAuthExcep("invalid credentials")
	}

	token := "auth token man"

	return SignInRes{
		Token: token,
	}, nil
}

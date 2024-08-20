package service

import (
	"context"
	"github.com/Eugene-Usachev/fastbytes"
	"github.com/Eugene-Usachev/fst"
	"messager/src/internal/models"
	"messager/src/internal/repository"
	"messager/src/pkg/logger"
)

type AuthService struct {
	logger.Logger
	repository.Auth
	accessConverter  *fst.EncodedConverter
	refreshConverter *fst.EncodedConverter
}

var _ Auth = (*AuthService)(nil)

func NewAuthService(logger logger.Logger, auth repository.Auth, accessConverter *fst.EncodedConverter,
	refreshConverter *fst.EncodedConverter) *AuthService {
	return &AuthService{
		Logger:           logger,
		Auth:             auth,
		accessConverter:  accessConverter,
		refreshConverter: refreshConverter,
	}
}

func (a *AuthService) SignUp(ctx context.Context, dto models.UserSingUpDTO) (int, string, string, error) {
	id, err := a.Auth.SignUp(ctx, dto)

	if err != nil {
		return -1, "", "", err
	}

	accessToken := a.accessConverter.NewToken(fastbytes.I2B(id))
	refreshToken := a.refreshConverter.NewToken(fastbytes.S2B(dto.Password))

	return id, accessToken, refreshToken, nil
}

func (a *AuthService) SignIn(ctx context.Context, dto models.UserSignInDTO) (int, string, string, error) {
	id, err := a.Auth.SingIn(ctx, dto)

	if err != nil {
		return -1, "", "", err
	}

	accessToken := a.accessConverter.NewToken(fastbytes.I2B(id))
	refreshToken := a.refreshConverter.NewToken(fastbytes.S2B(dto.Password))

	return id, accessToken, refreshToken, nil
}

func (a *AuthService) Refresh(ctx context.Context, id int, token string) (string, string, error) {
	password, err := a.refreshConverter.ParseToken(token)

	if err != nil {
		return "", "", err
	}

	err = a.Auth.Refresh(ctx, id, fastbytes.B2S(password))
	if err != nil {
		return "", "", err
	}

	access := a.accessConverter.NewToken(fastbytes.I2B(id))
	refresh := a.refreshConverter.NewToken(password)

	return access, refresh, nil
}

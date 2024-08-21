package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	e "messager/src/internal/errors"
	"messager/src/internal/models"
	"messager/src/pkg/logger"
)

type AuthPostgres struct {
	logger.Logger
	*pgxpool.Pool
}

var _ Auth = (*AuthPostgres)(nil)

func NewAuthPostgres(pool *pgxpool.Pool, logger logger.Logger) *AuthPostgres {
	return &AuthPostgres{
		logger,
		pool,
	}
}

func (a *AuthPostgres) SignUp(ctx context.Context, dto models.UserSingUpDTO) (int, error) {
	id := -1
	err := a.QueryRow(ctx, "SELECT 1 FROM users WHERE login =  $1", dto.Login).Scan(&id)
	if err != nil {
		a.Logger.Error(err.Error())
		return -1, err
	}

	if id == 1 {
		return -1, e.LoginBusy
	}

	err = a.QueryRow(ctx, "INSERT INTO users (login,password, name, sname, phone) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		dto.Login, dto.Password, dto.Name, dto.Name, dto.Sname, dto.Phone).Scan(&id)

	if err != nil {
		a.Logger.Error(err.Error())
		return -1, err
	}

	return id, nil
}

func (a *AuthPostgres) SingIn(ctx context.Context, dto models.UserSignInDTO) (int, error) {
	var id int
	err := a.QueryRow(ctx, "SELECT 1 FROM users WHERE login=$1 AND password=$2 RETURNING id", dto.Login, dto.Password).Scan(&id)

	if err != nil {
		a.Logger.Error(err.Error())
		return -1, err
	}

	return id, nil
}

func (a *AuthPostgres) Refresh(ctx context.Context, id int, password string) error {
	var buf int

	err := a.QueryRow(ctx, "SELECT 1 FROM users WHERE id=$1 AND password=$2", id, password).Scan(&buf)
	if err != nil {
		a.Logger.Error(err.Error())
		return err
	}

	if id != 1 {
		return e.UserIsAbsent
	}

	return nil
}

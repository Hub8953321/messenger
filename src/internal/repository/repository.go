package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"messager/src/internal/models"
	"messager/src/pkg/logger"
)

type Auth interface {
	SignUp(ctx context.Context, dto models.UserSingUpDTO) (int, error)
	SingIn(ctx context.Context, dto models.UserSignInDTO) (int, error)
	Refresh(ctx context.Context, id int, password string) error
}

type User interface {
}

type Message interface {
}

type Chat interface {
	CreateChat(ctx context.Context, dto models.ChatCreateDTO) (int, error)
	AddMembers(ctx context.Context, userId int, dto models.ChatAddMemberDTO) error
	RemoveMembers(ctx context.Context, userId int, dto models.ChatRemoveMemberDTO) error
}

type Repository struct {
	Auth
	User
	Message
	Chat
}

func NewRepository(pool *pgxpool.Pool, logger logger.Logger) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(pool, logger),
		/*User:    NewUserRepository(),
		Message: NewMessageRepository(),
		Chat:    NewChatRepository(),*/
	}
}

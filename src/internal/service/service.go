package service

import (
	"context"
	"github.com/Eugene-Usachev/fst"
	"messager/src/internal/models"
	"messager/src/internal/repository"
	"messager/src/pkg/logger"
)

type Auth interface {
	SignUp(ctx context.Context, dto models.UserSingUpDTO) (int, string, string, error)
	SignIn(ctx context.Context, dto models.UserSignInDTO) (int, string, string, error)
	Refresh(ctx context.Context, id int, token string) (string, string, error)
}

type User interface {
}

type Message interface {
}

type Chat interface {
	CreateChat(ctx context.Context, token string, dto models.ChatCreateDTO) (int, error)
	AddMembers(ctx context.Context, token string, dto models.ChatAddMemberDTO) error
	RemoveMembers(ctx context.Context, token string, dto models.ChatRemoveMemberDTO) error
}

type Service struct {
	Auth
	User
	Message
	Chat
}

func NewService(logger logger.Logger, auth repository.Auth,
	accessConverter *fst.EncodedConverter, refreshConverter *fst.EncodedConverter) *Service {
	return &Service{
		Auth: NewAuthService(logger, auth, accessConverter, refreshConverter),
		/*User:    NewUserService(),
		Message: NewMessageService(),
		Chat:    NewChatService(),*/
	}
}

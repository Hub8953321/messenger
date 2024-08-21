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
	CreateChat(ctx context.Context, id int, dto models.ChatCreateDTO) (int, error)
	AddMembers(ctx context.Context, id int, dto models.ChatAddMemberDTO) error
	RemoveMembers(ctx context.Context, id int, dto models.ChatRemoveMemberDTO) error
}

type Service struct {
	Auth
	User
	Message
	Chat
}

func NewService(logger logger.Logger, repository *repository.Repository,
	accessConverter *fst.EncodedConverter, refreshConverter *fst.EncodedConverter) *Service {
	return &Service{
		Auth: NewAuthService(logger, repository.Auth, accessConverter, refreshConverter),
		Chat: NewChatService(logger, repository.Chat, accessConverter, refreshConverter),
		/*User:    NewUserService(),
		Message: NewMessageService(),
		Chat:    NewChatService(),*/
	}
}

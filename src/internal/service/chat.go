package service

import (
	"context"
	"github.com/Eugene-Usachev/fastbytes"
	"github.com/Eugene-Usachev/fst"
	e "messager/src/internal/errors"
	"messager/src/internal/models"
	"messager/src/internal/repository"
	"messager/src/pkg/logger"
)

type ChatService struct {
	logger.Logger
	repository.Chat
	accessConvertor  *fst.EncodedConverter
	refreshConvertor *fst.EncodedConverter
}

var _ Chat = (*ChatService)(nil)

func NewChatService(logger logger.Logger, chat repository.Chat,
	accessConverter *fst.EncodedConverter, refreshConverter *fst.EncodedConverter) *ChatService {
	return &ChatService{
		Logger:           logger,
		Chat:             chat,
		accessConvertor:  accessConverter,
		refreshConvertor: refreshConverter,
	}
}

func (c *ChatService) CreateChat(ctx context.Context, token string, dto models.ChatCreateDTO) (int, error) {
	buf, err := c.accessConvertor.ParseToken(token)

	if err != nil {
		return 0, e.UserUnauthorized
	}

	id := fastbytes.B2I(buf)

	if len(dto.Members) < 2 {
		return -1, e.ArrayTooShort
	}

	var isUserInMembers bool
	for _, member := range dto.Members {
		if member == id {
			isUserInMembers = true
		}
	}
	if !isUserInMembers {
		return -1, e.AccessError
	}

	return c.Chat.CreateChat(ctx, dto)
}

func (c *ChatService) AddMembers(ctx context.Context, token string, dto models.ChatAddMemberDTO) error {
	buf, err := c.accessConvertor.ParseToken(token)

	if err != nil {
		return e.UserUnauthorized
	}

	id := fastbytes.B2I(buf)

	return c.Chat.AddMembers(ctx, id, dto)
}

func (c *ChatService) RemoveMembers (ctx context.Context, token string, dto models.ChatRemoveMemberDTO) error{

	buf, err := c.accessConvertor.ParseToken(token)

	if err != nil {
		return e.UserUnauthorized
	}

	id:= fastbytes.B2I(buf)

	return c.Chat.RemoveMembers(ctx, id, dto)

}

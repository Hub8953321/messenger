package service

import (
	"context"
	"github.com/Eugene-Usachev/fst"
	e "messager/src/internal/errors"
	"messager/src/internal/models"
	"messager/src/internal/repository"
	"messager/src/pkg/logger"
	"slices"
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

func (c *ChatService) CreateChat(ctx context.Context, id int, dto models.ChatCreateDTO) (int, error) {
	if len(dto.Members) < 2 {
		return -1, e.ListTooShort
	}

	if !slices.Contains(dto.Members, id) {
		return -1, e.AccessError
	}

	dto.Admin = id

	return c.Chat.CreateChat(ctx, dto)
}

func (c *ChatService) AddMembers(ctx context.Context, id int, dto models.ChatAddMemberDTO) error {
	return c.Chat.AddMembers(ctx, id, dto)
}

func (c *ChatService) RemoveMembers(ctx context.Context, id int, dto models.ChatRemoveMemberDTO) error {
	return c.Chat.RemoveMembers(ctx, id, dto)
}

package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	e "messager/src/internal/errors"
	"messager/src/internal/models"
	"messager/src/pkg/logger"
)

type ChatPostgres struct {
	logger.Logger
	*pgxpool.Pool
}

var _ Chat = (*ChatPostgres)(nil)

func NewChatPostgres(logger logger.Logger, pool *pgxpool.Pool) *ChatPostgres {
	return &ChatPostgres{
		logger,
		pool,
	}
}

func (c *ChatPostgres) CreateChat(ctx context.Context, dto models.ChatCreateDTO) (int, error) {
	var id int
	err := c.QueryRow(
		ctx,
		"INSERT INTO chats (creator_id,users_id, name) VALUES ($1, $2, $3) RETURNING id",
		dto.Admin,
		dto.Members,
		dto.Name,
	).Scan(&id)
	if err != nil {
		c.Error(err.Error())
		return -1, err
	}

	return id, nil
}

func (c *ChatPostgres) AddMembers(ctx context.Context, userId int, dto models.ChatAddMemberDTO) error {
	_, err := c.Exec(
		ctx,
		"UPDATE chats SET users_id= user_id || $1  WHERE id=$2 AND $3=ANY(users_id)",
		dto.Members,
		dto.ChatId,
		userId,
	)
	if err != nil {
		c.Logger.Error(err.Error())
		return err
	}

	return nil

}

func (c *ChatPostgres) RemoveMembers(ctx context.Context, userId int, dto models.ChatRemoveMemberDTO) error {
	if userId == dto.UserId {
		tag, err := c.Exec(ctx, "UPDATE chats SET users_id=array_remove(users_id, $1) WHERE id=$2", dto.UserId, dto.ChatId)
		if err != nil {
			c.Logger.Error(err.Error())
		}
		if tag.RowsAffected() == 0 {
			return e.NoneRowsAffected
		}
		return err
	}

	tag, err := c.Exec(
		ctx,
		"UPDATE chats SET users_id=array_remove(users_id, $1) WHERE id=$2 AND creator_id=$3",
		dto.UserId,
		dto.ChatId,
		userId,
	)
	if err != nil {
		c.Logger.Error(err.Error())
		return err
	}

	if tag.RowsAffected() == 0 {
		return e.NoneRowsAffected
	}
	return err

}

package models

type ChatCreateDTO struct {
	Name    string `json:"name" binding:"required"`
	Members []int  `json:"members" binding:"required"`
	Creator int    `json:"creater" binding:"required"`
}

type ChatAddMemberDTO struct {
	ChatId  int   `json:"chat_id" binding:"required"`
	Members []int `json:"members" binding:"required"`
}

type ChatRemoveMemberDTO struct {
	ChatId int `json:"chat_id" binding:"required"`
	UserId int `json:"user_id" binding:"required"`
}

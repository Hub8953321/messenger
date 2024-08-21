package models

type UserSingUpDTO struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Sname    string `json:"sname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type UserSignInDTO struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

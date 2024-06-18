package repository

type Auth interface {
}

type User interface {
}

type Message interface {
}

type Chat interface {
}

type Repository struct {
	Auth
	User
	Message
	Chat
}

func NewRepository() *Repository {
	return &Repository{
		/*		Auth:    NewAuthRepository(),
				User:    NewUserRepository(),
				Message: NewMessageRepository(),
				Chat:    NewChatRepository(),*/
	}
}

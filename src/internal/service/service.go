package service

type Auth interface {
}

type User interface {
}

type Message interface {
}

type Chat interface {
}

type Service struct {
	Auth
	User
	Message
	Chat
}

func NewService() *Service {
	return &Service{
		/*		Auth:    NewAuthService(),
				User:    NewUserService(),
				Message: NewMessageService(),
				Chat:    NewChatService(),*/
	}
}

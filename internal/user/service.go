package user

import (
	"fmt"

	"github.com/lithammer/shortuuid"
	"go.uber.org/zap"
)

type UserService interface {
	Create(user *User) (*User, error)
	Get(user *User) (*User, error)
}

type Service struct {
	user   UserRepository
	logger *zap.SugaredLogger
}

func NewService(user UserRepository, logger *zap.SugaredLogger) *Service {
	return &Service{user, logger}
}

func (s *Service) Create(user *User) (*User, error) {
	user.ID = shortuuid.New()
	user, err := s.user.Create(user)
	if err != nil {
		s.logger.Errorw(fmt.Sprintf("Error creating user %s with ID %s", user.Username, user.ID))
		return nil, err
	}
	return user, nil
}

func (s *Service) Get(user *User) (*User, error) {
	user, err := s.user.Get(user.ID)
	if err != nil {
		s.logger.Errorw(fmt.Sprintf("Error fetching user %s with ID %s", user.Username, user.ID))
		return nil, err
	}
	return user, nil
}

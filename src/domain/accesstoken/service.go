package accesstoken

import (
	"strings"

	"github.com/sauravgsh16/bookstore_auth/src/utils/errors"
)

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type Respository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repo Respository
}

func NewService(r Respository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetByID(id string) (*AccessToken, *errors.RestErr) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	at, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return at, nil
}

package accesstoken

import (
	"strings"

	"github.com/sauravgsh16/bookstore_auth/src/utils/errors"
)

// Service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(*AccessToken) (*AccessToken, *errors.RestErr)
	Update(*AccessToken) *errors.RestErr
}

// Repository interface
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(*AccessToken) *errors.RestErr
	Update(*AccessToken) *errors.RestErr
}

type service struct {
	repo Repository
}

// NewService returns a new service
func NewService(r Repository) Service {
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

func (s *service) Create(at *AccessToken) (*AccessToken, *errors.RestErr) {
	if err := at.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(at); err != nil {
		return nil, err
	}

	return at, nil
}

func (s *service) Update(at *AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repo.Update(at)
}

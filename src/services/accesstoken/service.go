package accesstoken

import (
	"strings"

	"github.com/sauravgsh16/bookstore_auth/src/domain/accesstoken"
	"github.com/sauravgsh16/bookstore_auth/src/repository/rest"

	"github.com/sauravgsh16/bookstore_auth/src/repository/db"
	"github.com/sauravgsh16/bookstore_auth/src/utils/errors"
)

// Service interface
type Service interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.Request) (*accesstoken.AccessToken, *errors.RestErr)
	Update(*accesstoken.AccessToken) *errors.RestErr
}

type service struct {
	db   db.Repository
	rest rest.UserRepository
}

// NewService returns a new service
func NewService(db db.Repository, rest rest.UserRepository) Service {
	return &service{
		db:   db,
		rest: rest,
	}
}

func (s *service) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	at, err := s.db.GetByID(id)
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *service) Create(req accesstoken.Request) (*accesstoken.AccessToken, *errors.RestErr) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// TODO: support for both grant types: password and client_credentials
	user, err := s.rest.LoginUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// Generate a New access token
	at := accesstoken.NewAccessToken(user.ID)
	at.Generate()

	// Save Access Token to db
	if err := s.db.Create(at); err != nil {
		return nil, err
	}

	return at, nil
}

func (s *service) Update(at *accesstoken.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.db.Update(at)
}

package db

import (
	"github.com/sauravgsh16/bookstore_auth/src/domain/accesstoken"
	"github.com/sauravgsh16/bookstore_auth/src/utils/errors"
)

type DBRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}

func NewRepository() DBRepository {
	return &dbRepository{}
}

type dbRepository struct{}

func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("DB not defined!!")
}

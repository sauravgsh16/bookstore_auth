package db

import (
	"github.com/gocql/gocql"
	"github.com/sauravgsh16/bookstore_auth/src/clients/cassandra"
	"github.com/sauravgsh16/bookstore_auth/src/domain/accesstoken"
	"github.com/sauravgsh16/bookstore_auth/src/utils/errors"
)

const (
	selectQuery = `SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;`
	insertQuery = `INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);`
	updateQuery = `UPDATE access_tokens SET expires=? WHERE access_token=?;`
)

// DBRepository interface
type DBRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(*accesstoken.AccessToken) *errors.RestErr
	Update(*accesstoken.AccessToken) *errors.RestErr
}

// NewRepository returns a new db repo
func NewRepository() DBRepository {
	return &dbRepository{}
}

type dbRepository struct{}

func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	session := cassandra.GetSession()

	var at accesstoken.AccessToken

	if err := session.Query(selectQuery, id).Scan(&at.Token, &at.UserID, &at.ClientID, &at.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("Access token not found")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &at, nil
}

func (r *dbRepository) Create(at *accesstoken.AccessToken) *errors.RestErr {
	s := cassandra.GetSession()

	if err := s.Query(insertQuery, at.Token, at.UserID, at.ClientID, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) Update(at *accesstoken.AccessToken) *errors.RestErr {
	s := cassandra.GetSession()

	if err := s.Query(insertQuery, at.Expires, at.Token).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

package rest

import (
	"encoding/json"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"

	"github.com/sauravgsh16/bookstore-auth-api/src/domain/users"
	"github.com/sauravgsh16/bookstore-auth-api/src/utils/errors"
)

const (
	endPointLogin = "/users/login"
)

var (
	restClientUser = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
)

// UserRepository interface
type UserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

// NewRepository return a UserRepository
func NewRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct{}

func (r *userRepository) LoginUser(email, pwd string) (*users.User, *errors.RestErr) {
	req := users.UserLoginRequest{
		Email:    email,
		Password: pwd,
	}

	resp := restClientUser.Post(endPointLogin, req)
	if resp == nil || resp.Response == nil {
		return nil, errors.NewInternalServerError("invalid rest client response")
	}

	if resp.StatusCode > 299 {
		var restErr errors.RestErr
		if err := json.Unmarshal(resp.Bytes(), &restErr); err != nil {
			return nil, errors.NewInternalServerError("invalid error interface, when trying to unmarshall error from rest call")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(resp.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("invalid user interface, when trying to unmarshall user from rest call")
	}
	return &user, nil
}

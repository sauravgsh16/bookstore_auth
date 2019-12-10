package accesstoken

import (
	"strings"
	"time"

	"github.com/sauravgsh16/bookstore_auth/src/utils/errors"
)

const (
	expirationTime = 24
)

// AccessToken struct - defines field of an access token
type AccessToken struct {
	Token    string `json:"token"`
	UserID   int64  `json:"user_id"`
	ClientID int64  `json:"client_id"`
	Expires  int64  `json:"expires"`
}

// NewAccessToken returns a new access token
func NewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired checks is token is expired
func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

// Validate checks if the access token is valid
func (at *AccessToken) Validate() *errors.RestErr {
	at.Token = strings.TrimSpace(at.Token)

	if at.Token == "" {
		return errors.NewBadRequestError("invalid access token id")
	}

	if at.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}

	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}

	return nil
}

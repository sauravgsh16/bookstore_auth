package accesstoken

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/sauravgsh16/bookstore_auth/src/utils/errors"
)

const (
	expirationTime = 24

	// Grant Types
	grantTypePassword   = "password"
	grantTypeClientCred = "client_credentials"
)

// Request struct
type Request struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client-credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Validate access token request parameters
func (r *Request) Validate() *errors.RestErr {
	switch r.GrantType {
	case grantTypePassword:
		break

	case grantTypeClientCred:
		break

	default:
		return errors.NewBadRequestError("invalid grant type in request")
	}

	// TODO : Validate other request parameters
	return nil
}

// AccessToken struct - defines field of an access token
type AccessToken struct {
	Token    string `json:"token"`
	UserID   int64  `json:"user_id"`
	ClientID int64  `json:"client_id"`
	Expires  int64  `json:"expires"`
}

// NewAccessToken returns a new access token
func NewAccessToken(uid int64) *AccessToken {
	return &AccessToken{
		UserID:  uid,
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

// Generate AccessToken
func (at *AccessToken) Generate() {
	s := fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires)

	at.Token = getMd5(s)
}

func getMd5(s string) string {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(s))

	return hex.EncodeToString(hash.Sum(nil))
}

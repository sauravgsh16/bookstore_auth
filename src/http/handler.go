package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sauravgsh16/bookstore_auth/src/domain/accesstoken"
	serv "github.com/sauravgsh16/bookstore_auth/src/services/accesstoken"
	"github.com/sauravgsh16/bookstore_auth/src/utils/errors"
)

// Handler interface
type Handler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
}

type accessTokenHandler struct {
	service serv.Service
}

// NewHandler returns a new handler
func NewHandler(s serv.Service) Handler {
	return &accessTokenHandler{
		service: s,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	at, err := h.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, at)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var req accesstoken.Request

	if err := c.ShouldBindJSON(&req); err != nil {
		rstErr := errors.NewBadRequestError(fmt.Sprintf("invalid json body: %s", err.Error()))
		c.JSON(rstErr.Status, rstErr)
		return
	}

	fmt.Printf("%#v\n", req)

	token, err := h.service.Create(req)
	if err != nil {
		fmt.Printf("Breaking here: %+v\n", err)
		c.JSON(err.Status, err.Error)
		return
	}
	c.JSON(http.StatusOK, *token)
}

func (h *accessTokenHandler) Update(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not implemented",
	})
}

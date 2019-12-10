package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sauravgsh16/bookstore_auth/src/domain/accesstoken"
	"github.com/sauravgsh16/bookstore_auth/src/utils/errors"
)

type Handler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

func NewHandler(s accesstoken.Service) Handler {
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
	var at accesstoken.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		rstErr := errors.NewBadRequestError(fmt.Sprintf("invalid json body: %s", err.Error()))
		c.JSON(rstErr.Status, rstErr)
		return
	}
	token, err := h.service.Create(&at)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, *token)
}

func (h *accessTokenHandler) Update(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not implemented",
	})
}

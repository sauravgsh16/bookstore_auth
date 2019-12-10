package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sauravgsh16/bookstore_auth/src/domain/accesstoken"
)

type Handler interface {
	GetByID(c *gin.Context)
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

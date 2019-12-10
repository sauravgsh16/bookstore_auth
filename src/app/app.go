package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravgsh16/bookstore_auth/src/domain/accesstoken"
)

var (
	router = gin.Default()
)

func StartApplication() {
	db := db.NewRepository()
	serv := accesstoken.NewService(db)

	h := http.NewHandler(serv)

	router.GET("/auth/access_token/:access_token_id", h.GetByID)

	router.Run(":8080")
}

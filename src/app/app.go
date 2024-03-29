package app

import (
	"github.com/gin-gonic/gin"

	"github.com/sauravgsh16/bookstore_auth/src/http"
	"github.com/sauravgsh16/bookstore_auth/src/repository/db"
	"github.com/sauravgsh16/bookstore_auth/src/repository/rest"
	"github.com/sauravgsh16/bookstore_auth/src/services/accesstoken"
)

var (
	router = gin.Default()
)

// StartApplication starts the application server
func StartApplication() {
	db := db.NewRepository()
	rest := rest.NewRepository()
	serv := accesstoken.NewService(db, rest)

	h := http.NewHandler(serv)

	router.GET("/auth/access_token/:access_token_id", h.GetByID)
	router.POST("/auth/access_token", h.Create)
	// router.PUT("/auth/access_token", h.Update)

	router.Run(":8090")
}

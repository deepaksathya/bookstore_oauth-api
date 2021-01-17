package app

import (
	accesstoken "github.com/deepaksathya/bookstore_oauth-api/src/domain/access_token"
	"github.com/deepaksathya/bookstore_oauth-api/src/http"
	"github.com/deepaksathya/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication is called from main to start the web server
func StartApplication() {
	atHandler := http.NewHandler(accesstoken.NewService(db.NewRepository()))
	router.GET("oauth/accesstoken/:access_token_id", atHandler.GetByID)
	router.Run(":3000")
}

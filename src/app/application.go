package app

import (
	"github.com/deepaksathya/bookstore_oauth-api/src/clients/cassandra"
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

	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}

	session.Close()

	//initialize repository that stores the token
	dbrepo := db.NewRepository()

	// initialize service which will act as intermediary in getting
	// access token from the repository.
	// Service will know which repo to contact
	oauthservice := accesstoken.NewService(dbrepo)

	// initialize the Handler and tell it which service to be contacted
	atHandler := http.NewHandler(oauthservice)
	router.GET("oauth/accesstoken/:access_token_id", atHandler.GetByID)
	router.Run(":3000")
}

package http

import (
	"log"
	"net/http"

	accesstoken "github.com/deepaksathya/bookstore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByID(c *gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

func NewHandler(service accesstoken.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	accessTokenID := c.Param("access_token_id")
	accessToken, err := h.service.GetByID(accessTokenID)
	log.Println(err)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, accessToken)
}

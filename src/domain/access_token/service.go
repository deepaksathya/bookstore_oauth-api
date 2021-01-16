package accesstoken

import (
	"github.com/deepaksathya/bookstore_users-api/utils/errors"
)

type Service interface {
	GetByID(string) (*AccessToken, *errors.R)
}
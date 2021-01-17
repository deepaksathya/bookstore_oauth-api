package db

import (
	accesstoken "github.com/deepaksathya/bookstore_oauth-api/src/domain/access_token"
	"github.com/deepaksathya/bookstore_oauth-api/src/utils/errors"
)

type DbReopository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}

func NewRepository() DbReopository {
	return &dbReopository{}
}

type dbReopository struct {
}

func (r *dbReopository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}

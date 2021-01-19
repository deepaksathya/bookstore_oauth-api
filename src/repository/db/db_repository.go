package db

import (
	"fmt"

	"github.com/deepaksathya/bookstore_oauth-api/src/clients/cassandra"
	accesstoken "github.com/deepaksathya/bookstore_oauth-api/src/domain/access_token"
	"github.com/deepaksathya/bookstore_users-api/utils/errors"
)

type DbReopository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type dbReopository struct {
}

func NewRepository() DbReopository {
	return &dbReopository{}
}

func (r *dbReopository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()

	if err != nil {
		panic(err)
	}

	defer session.Closed()
	fmt.Println("cassandra connection successfully created")

	return nil, errors.NewInternalServerError("database connection not implemented yet")
}

package accesstoken

import (
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)

type Service interface {
	GetByID(string) (*AccessToken, *rest_errors.RestErr)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) GetByID(string) (*AccessToken, *rest_errors.RestErr) {
	return nil, nil
}

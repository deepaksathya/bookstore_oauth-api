package accesstoken

import (
	"strings"

	"github.com/deepaksathya/bookstore_users-api/utils/errors"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

// NewService creates a Repository.
// It takes a Repository interface implementation
// - means any DB layer object that will faciliate token retrieval and manipulation
// Returns a service layer interface implementaion that will
// act as intermediary for access token retrieval and manipulation
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

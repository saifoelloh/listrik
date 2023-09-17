package pln

import "github.com/saifoelloh/listrik/api/presenter"

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	FetchUsers() (*[]presenter.UserPLN, error)
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// FetchBooks is a service layer that helps fetch all books in BookShop
func (s *service) FetchUsers() (*[]presenter.UserPLN, error) {
	return s.repository.ReadUser()
}

//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/himanshunagar33/go-grpc-services/internal/rocket Store
package rocket

import (
	"context"
)

// Rocket should contain the defintitation of our rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// Store - defines the interface we expect our database implementation to follow
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// service - our rocket service , responsible for updating the rocket inventory
type Service struct {
	Store Store
}

// Return the new instance of our rocket service
func New() Service {
	return Service{}
}

//GetRocketby ID - retrieves a rocket based on the ID from the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// insert a new rocket in the invetonry
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// deletes a rocker from the inventory
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}

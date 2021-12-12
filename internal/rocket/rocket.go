//go:generate mockgen -destination=rocket_mock_test.go -package=rocket example.com/microservice/rocket-service/internal/rocket Store

package rocket

import "context"

// Rocket - should contain defination of rocket.
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights string
}

// Store - defines the interface we expect
// our database implementation to follow.
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service -  our rocket service,  responsible for
// updating out rocket inventory.
type Service struct {
	Store Store
}

//New - return a new instance if rocket service.
func New(Store Store) Service {
	return Service{
		Store: Store,
	}
}

// GetRocketById - retrives a rocket based on id
// from s store.
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

// InsertRockrt - insert a new rocket into the store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

// DelereRocket - deletes a rocket from inventory.
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}

	return nil
}

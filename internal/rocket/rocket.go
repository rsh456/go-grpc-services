//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/rsh456/go-grpc-services/internal/rocket Store
package rocket

import "context"

//Rocket - contains the definition of rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

//Store - defines the interface we expect
//the database implementation follow
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

//Service - rocket service, responsible for
//updating the rocket inventory
type Service struct {
	Store Store
}

//New - returns a new instance of our  rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

//GetRocketByID - retrieves a rocket based on the ID from the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, err
}

//InsertRocket - inserts a new rocket into the store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

//Delete Rocket - delete a rocket based on ID from the store
func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}

package repositories

import (
	"booking_asm/grpc/flight-grpc/repositories/flight_repo"

	"gorm.io/gorm"
)

type flightRepository struct {
	FlightRepo flight_repo.FlightRepository
}

func New(db *gorm.DB) flightRepository {
	return flightRepository{FlightRepo: flight_repo.New(db)}
}

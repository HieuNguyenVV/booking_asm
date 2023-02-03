package handler

import (
	"booking_asm/grpc/flight-grpc/handler/flight_handler"
	"booking_asm/grpc/flight-grpc/repositories/flight_repo"
)

type handler struct {
	FlightHandler flight_handler.FlightHandler
}

func New(repo flight_repo.FlightRepository) handler {
	return handler{flight_handler.New(repo)}
}

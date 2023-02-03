package handler

import (
	"booking_asm/grpc/customer-grpc/handler/customer_handler"
	"booking_asm/grpc/customer-grpc/repositories/customer_repo"
)

type handler struct {
	CustomerHandler customer_handler.CustomerHandler
}

func New(customer_repo customer_repo.CustomerRepository) handler {
	return handler{CustomerHandler: customer_handler.New(customer_repo)}
}

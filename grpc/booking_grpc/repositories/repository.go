package repositories

import (
	"booking_asm/grpc/booking_grpc/repositories/booking_repo"

	"gorm.io/gorm"
)

type repository struct {
	BookingRepo booking_repo.BookingRepository
}

func New(db *gorm.DB) repository {
	return repository{BookingRepo: booking_repo.New(db)}
}

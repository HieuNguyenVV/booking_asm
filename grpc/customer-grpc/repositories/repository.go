package repositories

import (
	"booking_asm/grpc/customer-grpc/repositories/customer_repo"

	"gorm.io/gorm"
)

type customer struct {
	CustomerRepo customer_repo.CustomerRepository
}

func New(db *gorm.DB) customer {
	return customer{CustomerRepo: customer_repo.New(db)}
}
 
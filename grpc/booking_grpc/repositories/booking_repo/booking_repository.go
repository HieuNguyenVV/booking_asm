package booking_repo

import (
	"booking_asm/grpc/booking_grpc/model"
	"context"

	"gorm.io/gorm"
)

type (
	BookingRepository interface {
		ExistBookingByCode(ctx context.Context, code string) (bool, error)
		ExistBooking(ctx context.Context, id string) (bool, error)
		CreatBooking(ctx context.Context, model *model.Booking) (*model.Booking, error)
		CancelBooking(ctx context.Context, code, status string) (*model.Booking, error)
		FindBooking(ctx context.Context, code string) (*model.Booking, error)
		FindBookingByIdCustomer(ctx context.Context, id string) (*[]model.Booking, error)
	}
	bookingRepo struct {
		db *gorm.DB
	}
)

func New(db *gorm.DB) BookingRepository {
	return &bookingRepo{
		db: db,
	}
}
func (t *bookingRepo) ExistBookingByCode(ctx context.Context, code string) (bool, error) {
	var exists bool
	if err := t.db.Model(&model.Booking{}).
		Select("count(*) > 0").
		Where("code = ?", code).
		Find(&exists).
		Error; err != nil {
		return false, err
	}
	return exists, nil
}
func (t *bookingRepo) ExistBooking(ctx context.Context, id string) (bool, error) {
	var exists bool
	if err := t.db.Model(&model.Booking{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error; err != nil {
		return false, err
	}
	return exists, nil
}
func (t *bookingRepo) CreatBooking(ctx context.Context, model *model.Booking) (*model.Booking, error) {
	if err := t.db.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}
func (t *bookingRepo) CancelBooking(ctx context.Context, code, status string) (*model.Booking, error) {
	if err := t.db.Where(&model.Booking{Code: code}).Updates(&model.Booking{Status: status}).Error; err != nil {
		return &model.Booking{}, err
	}
	var resp model.Booking
	if err := t.db.Where(&model.Booking{Code: code}).First(&resp).Error; err != nil {
		return &model.Booking{}, err
	}
	return &resp, nil
}
func (t *bookingRepo) FindBooking(ctx context.Context, code string) (*model.Booking, error) {
	var resp model.Booking
	if err := t.db.Where(&model.Booking{Code: code}).First(&resp).Error; err != nil {
		return &model.Booking{}, err
	}
	return &resp, nil
}
func (t *bookingRepo) FindBookingByIdCustomer(ctx context.Context, id string) (*[]model.Booking, error) {
	var resp []model.Booking
	if err := t.db.Where(&model.Booking{Customer_id: id}).Find(&resp).Error; err != nil {
		return &resp, err
	}
	return &resp, nil
}

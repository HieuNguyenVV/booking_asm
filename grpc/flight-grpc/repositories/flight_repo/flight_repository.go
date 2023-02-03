package flight_repo

import (
	"booking_asm/grpc/flight-grpc/model"
	"context"

	"gorm.io/gorm"
)

type (
	FlightRepository interface {
		ExistFlight(ctx context.Context, id string) (bool, error)
		CreateFlight(ctx context.Context, mode *model.Flight) (*model.Flight, error)
		UpdateFlight(ctx context.Context, mode *model.Flight) (*model.Flight, error)
		SearchFlight(ctx context.Context, id string) (*model.Flight, error)
	}
	flightRepo struct {
		db *gorm.DB
	}
)

func New(db *gorm.DB) FlightRepository {
	return &flightRepo{
		db: db,
	}
}
func (t *flightRepo) ExistFlight(ctx context.Context, id string) (bool, error) {
	var exists bool
	if err := t.db.Model(&model.Flight{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error; err != nil {
		return false, err
	}
	return exists, nil
}
func (t *flightRepo) CreateFlight(ctx context.Context, mode *model.Flight) (*model.Flight, error) {
	if err := t.db.Create(mode).Error; err != nil {
		return &model.Flight{}, err
	}
	return mode, nil
}
func (t *flightRepo) UpdateFlight(ctx context.Context, mode *model.Flight) (*model.Flight, error) {
	if err := t.db.Where(&model.Flight{ID: mode.ID}).Updates(mode).Error; err != nil {
		return &model.Flight{}, err
	}
	return mode, nil
}
func (t *flightRepo) SearchFlight(ctx context.Context, id string) (*model.Flight, error) {
	resp := model.Flight{}
	if err := t.db.Where(&model.Flight{ID: id}).First(&resp).Error; err != nil {
		return &model.Flight{}, err
	}
	return &resp, nil
}

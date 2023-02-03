package customer_repo

import (
	"booking_asm/grpc/customer-grpc/model"
	"context"

	"gorm.io/gorm"
)

type (
	CustomerRepository interface {
		ExistCustomerById(ctx context.Context, id string) (bool, error)
		GetCustomerById(ctx context.Context, id string) (*model.Customer, error)
		CreateCustomer(ctx context.Context, mode *model.Customer) (*model.Customer, error)
		UpdateCustomer(ctx context.Context, mode *model.Customer) (*model.Customer, error)
		ChangePassword(ctx context.Context, mode *model.ChangePasswordCondition) (*model.Customer, error)
	}
	customerRepo struct {
		db *gorm.DB
	}
)

func New(db *gorm.DB) CustomerRepository {
	return &customerRepo{
		db: db,
	}
}
func (t *customerRepo) ExistCustomerById(ctx context.Context, id string) (bool, error) {
	var exists bool
	if err := t.db.Model(&model.Customer{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error; err != nil {
		return false, err
	}
	return exists, nil
}
func (t *customerRepo) GetCustomerById(ctx context.Context, id string) (*model.Customer, error) {
	resp := &model.Customer{}
	if err := t.db.Where(&model.Customer{ID: id}).First(resp).Error; err != nil {
		return &model.Customer{}, err
	}
	return resp, nil
}
func (t *customerRepo) CreateCustomer(ctx context.Context, mode *model.Customer) (*model.Customer, error) {
	if err := t.db.Create(mode).Error; err != nil {
		return &model.Customer{}, err
	}
	return mode, nil
}
func (t *customerRepo) UpdateCustomer(ctx context.Context, mode *model.Customer) (*model.Customer, error) {
	if err := t.db.Where(&model.Customer{ID: mode.ID}).Updates(mode).Error; err != nil {
		return &model.Customer{}, err
	}
	return mode, nil
}
func (t *customerRepo) ChangePassword(ctx context.Context, mode *model.ChangePasswordCondition) (*model.Customer, error) {
	if err := t.db.Where(&model.Customer{ID: mode.Id}).Updates(&model.Customer{Password: mode.Password}).Error; err != nil {
		return &model.Customer{}, err
	}
	resp := &model.Customer{}
	if err := t.db.Where(&model.Customer{ID: mode.Id}).First(resp).Error; err != nil {
		return &model.Customer{}, err
	}
	return resp, nil
}

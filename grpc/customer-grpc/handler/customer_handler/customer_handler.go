package customer_handler

import (
	"booking_asm/grpc/customer-grpc/model"
	"booking_asm/grpc/customer-grpc/repositories/customer_repo"
	"booking_asm/pb"
	"context"
	"fmt"
	"time"
)

type CustomerHandler struct {
	customerRepo customer_repo.CustomerRepository
	pb.UnimplementedCustomerServiceServer
}

func New(customerRepo customer_repo.CustomerRepository) CustomerHandler {
	return CustomerHandler{customerRepo: customerRepo}
}
func (t *CustomerHandler) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	ok, err := t.customerRepo.ExistCustomerById(ctx, in.Id)
	if err != nil {
		return &pb.ChangePasswordResponse{}, err
	}
	if ok {
		resp, err := t.customerRepo.ChangePassword(ctx, &model.ChangePasswordCondition{
			Id:       in.Id,
			Password: in.Password,
		})
		if err != nil {
			return &pb.ChangePasswordResponse{}, err
		}
		return &pb.ChangePasswordResponse{
			CustomerName: resp.Name,
			Address:      resp.Address,
			Phone:        resp.Phone,
			LicenseId:    resp.License_id,
			Active:       resp.Active,
			Id:           resp.ID,
			Email:        resp.Email,
			Password:     resp.Password,
		}, nil
	}
	return &pb.ChangePasswordResponse{}, fmt.Errorf("user does not exist")
}
func (t *CustomerHandler) UpdateCustomer(ctx context.Context, in *pb.UpdateCustomerRequest) (*pb.UpdateCustomerResponse, error) {
	ok, err := t.customerRepo.ExistCustomerById(ctx, in.Id)
	if err != nil {
		return &pb.UpdateCustomerResponse{}, err
	}
	if ok {
		resp, err := t.customerRepo.GetCustomerById(ctx, in.Id)
		if err != nil {
			return &pb.UpdateCustomerResponse{}, err
		}
		req := &model.Customer{}
		if in.CustomerName != "" {
			req.Name = in.CustomerName
		} else {
			req.Name = resp.Name
		}
		if in.Address != "" {
			req.Address = in.Address
		} else {
			req.Address = resp.Address
		}
		if in.Phone != "" {
			req.Phone = in.Phone
		} else {
			req.Phone = resp.Phone
		}
		if in.LicenseId != "" {
			req.License_id = in.LicenseId
		} else {
			req.License_id = resp.License_id
		}
		if in.Id != "" {
			req.ID = in.Id
		} else {
			req.ID = resp.ID
		}
		if in.Email != "" {
			req.Email = in.Email
		} else {
			req.Email = resp.Email
		}
		if in.Password != "" {
			req.Password = in.Password
		} else {
			req.Password = resp.Password
		}
		req.UpdatedAt = time.Now()
		req.CreatedAt = resp.CreatedAt
		resp, err = t.customerRepo.UpdateCustomer(ctx, req)
		if err != nil {
			return &pb.UpdateCustomerResponse{}, err
		}
		return &pb.UpdateCustomerResponse{
			CustomerName: resp.Name,
			Address:      resp.Address,
			Phone:        resp.Phone,
			LicenseId:    resp.License_id,
			Active:       resp.Active,
			Id:           resp.ID,
			Email:        resp.Email,
			Password:     resp.Password,
		}, nil
	}
	return &pb.UpdateCustomerResponse{}, fmt.Errorf("user does not exist")
}
func (t *CustomerHandler) FindCustomer(ctx context.Context, in *pb.FindCustomerRequest) (*pb.FindCustomerResponse, error) {
	ok, err := t.customerRepo.ExistCustomerById(ctx, in.Id)
	if err != nil {
		return &pb.FindCustomerResponse{}, err
	}
	if ok {
		resp, err := t.customerRepo.GetCustomerById(ctx, in.Id)
		if err != nil {
			return &pb.FindCustomerResponse{}, err
		}
		return &pb.FindCustomerResponse{
			CustomerName: resp.Name,
			Address:      resp.Address,
			Phone:        resp.Phone,
			LicenseId:    resp.License_id,
			Active:       resp.Active,
			Id:           resp.ID,
			Email:        resp.Email,
			Password:     resp.Password,
		}, nil
	}
	return &pb.FindCustomerResponse{}, fmt.Errorf("user does not exist")
}
func (t *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	ok, err := t.customerRepo.ExistCustomerById(ctx, in.Id)
	if err != nil {
		return &pb.CreateCustomerResponse{}, err
	}
	if !ok {
		resp, err := t.customerRepo.CreateCustomer(ctx, &model.Customer{
			Name:       in.CustomerName,
			Address:    in.Address,
			Phone:      in.Phone,
			License_id: in.LicenseId,
			Active:     in.Active,
			ID:         in.Id,
			Email:      in.Email,
			Password:   in.Password,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		})
		if err != nil {
			return &pb.CreateCustomerResponse{}, err
		}
		return &pb.CreateCustomerResponse{
			CustomerName: resp.Name,
			Address:      resp.Address,
			Phone:        resp.Phone,
			LicenseId:    resp.License_id,
			Active:       resp.Active,
			Id:           resp.ID,
			Email:        resp.Email,
			Password:     resp.Password,
		}, nil
	}
	return &pb.CreateCustomerResponse{}, fmt.Errorf("user already exist")
}

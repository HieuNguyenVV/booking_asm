package booking_handler

import (
	"booking_asm/grpc/booking_grpc/model"
	"booking_asm/grpc/booking_grpc/repositories/booking_repo"
	"booking_asm/pb"
	"context"
	"fmt"
	"time"
)

type BookingHandler struct {
	bookingRepo booking_repo.BookingRepository
	pb.UnimplementedBookingServiceServer
	customerClient pb.CustomerServiceClient
	flightClient   pb.FlightServiceClient
}

func New(flightClient pb.FlightServiceClient, customerClient pb.CustomerServiceClient, bookingRepo booking_repo.BookingRepository) BookingHandler {
	return BookingHandler{
		bookingRepo:                       bookingRepo,
		UnimplementedBookingServiceServer: pb.UnimplementedBookingServiceServer{},
		customerClient:                    customerClient,
		flightClient:                      flightClient,
	}
}

func (t *BookingHandler) CreateBooking(ctx context.Context, in *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	ok, err := t.bookingRepo.ExistBooking(ctx, in.Id)
	if err != nil {
		return &pb.CreateBookingResponse{}, err
	}
	if !ok {
		resp, err := t.bookingRepo.CreatBooking(ctx, &model.Booking{
			ID:          in.Id,
			Customer_id: in.CustomerId,
			Flight_id:   in.FlightId,
			Code:        in.Code,
			Status:      in.Status,
			Booked_date: in.BookedDate.AsTime(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
		if err != nil {
			return &pb.CreateBookingResponse{}, err
		}
		return &pb.CreateBookingResponse{
			Id:         resp.ID,
			CustomerId: resp.Customer_id,
			FlightId:   resp.Flight_id,
			Code:       resp.Code,
			Status:     resp.Status,
			BookedDate: resp.Booked_date.String(),
		}, nil
	}
	return &pb.CreateBookingResponse{}, fmt.Errorf("id already exist")
}
func (t *BookingHandler) CancelBooking(ctx context.Context, in *pb.CancelBookingRequest) (*pb.BookingRespnose, error) {
	ok, err := t.bookingRepo.ExistBookingByCode(ctx, in.Code)
	if err != nil {
		return &pb.BookingRespnose{}, err
	}
	if ok {
		resp, err := t.bookingRepo.CancelBooking(ctx, in.Code, in.Status)
		if err != nil {
			return &pb.BookingRespnose{}, err
		}
		return &pb.BookingRespnose{
			Id:         resp.ID,
			CustomerId: resp.Customer_id,
			FlightId:   resp.Flight_id,
			Code:       resp.Code,
			Status:     resp.Status,
			BookedDate: resp.Booked_date.String(),
		}, nil
	}
	return &pb.BookingRespnose{}, fmt.Errorf("code does not exist")
}

func (t *BookingHandler) SearchBooking(ctx context.Context, in *pb.SearchBookingRequest) (*pb.BookingInfor, error) {
	ok, err := t.bookingRepo.ExistBookingByCode(ctx, in.Code)
	if err != nil {
		return &pb.BookingInfor{}, err
	}
	if ok {
		resp, err := t.bookingRepo.FindBooking(context.Background(), in.Code)
		if err != nil {
			return &pb.BookingInfor{}, err
		}
		customer, err := t.customerClient.FindCustomer(context.Background(), &pb.FindCustomerRequest{
			Id: resp.Customer_id,
		})
		if err != nil {
			return &pb.BookingInfor{}, err
		}
		flight, err := t.flightClient.SearchFlight(context.Background(), &pb.SearchFlightRequest{
			Id: resp.Flight_id,
		})
		if err != nil {
			return &pb.BookingInfor{}, err
		}
		return &pb.BookingInfor{
			BookingDetail: &pb.BookingRespnose{
				Id:         resp.ID,
				CustomerId: resp.Customer_id,
				FlightId:   resp.Flight_id,
				Code:       resp.Code,
				Status:     resp.Status,
				BookedDate: resp.Booked_date.String(),
			},
			FlightDetail: &pb.Flight{
				Id:            flight.Id,
				Name:          flight.Name,
				From:          flight.From,
				To:            flight.To,
				Date:          flight.Date,
				Status:        flight.Status,
				AvailableSlot: int32(flight.AvailableSlot),
			},
			CustomerDetail: &pb.Customer{
				CustomerName: customer.CustomerName,
				Address:      customer.Address,
				Phone:        customer.Phone,
				LicenseId:    customer.LicenseId,
				Active:       customer.Active,
				Id:           customer.Id,
				Email:        customer.Email,
				Password:     customer.Password,
			},
		}, nil

	}
	return &pb.BookingInfor{}, fmt.Errorf("code does not exist")
}
func (t *BookingHandler) SearchBookingId(ctx context.Context, in *pb.SearchBookingByIdRequest) (*pb.ListBooking, error) {
	booking, err := t.bookingRepo.FindBookingByIdCustomer(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	var pResponse pb.ListBooking

	//slice Booking search by ID to return
	for _, i := range *booking {
		pResponse.BookingList = append(pResponse.BookingList, &pb.BookingRespnose{
			Id:         i.ID,
			CustomerId: i.Customer_id,
			FlightId:   i.Flight_id,
			Code:       i.Code,
			Status:     i.Status,
			BookedDate: i.Booked_date.String(),
		})
	}

	return &pResponse, nil
}

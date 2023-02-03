package flight_handler

import (
	"booking_asm/grpc/flight-grpc/model"
	"booking_asm/grpc/flight-grpc/repositories/flight_repo"
	"booking_asm/pb"
	"context"
	"fmt"
	"time"
)

type FlightHandler struct {
	flightRepo flight_repo.FlightRepository
	pb.UnimplementedFlightServiceServer
}

func New(flightRepo flight_repo.FlightRepository) FlightHandler {
	return FlightHandler{flightRepo: flightRepo}
}

func (t *FlightHandler) UpdateFlight(ctx context.Context, in *pb.UpdateFlightRequest) (*pb.UpdateFlightResponse, error) {
	ok, err := t.flightRepo.ExistFlight(ctx, in.Id)
	if err != nil {
		return &pb.UpdateFlightResponse{}, err
	}
	if ok {
		resp, err := t.flightRepo.SearchFlight(ctx, in.Id)
		if err != nil {
			return &pb.UpdateFlightResponse{}, err
		}
		req := &model.Flight{}
		if in.Name != "" {
			req.Name = in.Name
		} else {
			req.Name = resp.Name
		}
		if in.Id != "" {
			req.ID = in.Id
		} else {
			req.ID = resp.ID
		}
		if in.From != "" {
			req.From = in.From
		} else {
			req.From = resp.From
		}
		if in.To != "" {
			req.To = in.To
		} else {
			req.To = resp.To
		}

		if in.Status != "" {
			req.Status = in.Status
		} else {
			req.Status = resp.Status
		}
		req.Available_slot = int(in.AvailableSlot)
		req.Date = in.Date.AsTime()
		req.UpdatedAt = time.Now()
		req.CreatedAt = resp.CreatedAt
		resp, err = t.flightRepo.UpdateFlight(ctx, req)
		if err != nil {
			return &pb.UpdateFlightResponse{}, err
		}
		return &pb.UpdateFlightResponse{
			Id:            resp.ID,
			Name:          resp.Name,
			From:          resp.From,
			To:            resp.To,
			Date:          resp.Date.String(),
			Status:        resp.Status,
			AvailableSlot: int32(resp.Available_slot),
		}, nil
	}
	return &pb.UpdateFlightResponse{}, fmt.Errorf("user does not exist")
}
func (t *FlightHandler) SearchFlight(ctx context.Context, in *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	ok, err := t.flightRepo.ExistFlight(ctx, in.Id)
	if err != nil {
		return &pb.SearchFlightResponse{}, err
	}
	if ok {
		resp, err := t.flightRepo.SearchFlight(ctx, in.Id)
		if err != nil {
			return &pb.SearchFlightResponse{}, err
		}
		return &pb.SearchFlightResponse{
			Id:            resp.ID,
			Name:          resp.Name,
			From:          resp.From,
			To:            resp.To,
			Date:          resp.Date.String(),
			Status:        resp.Status,
			AvailableSlot: int32(resp.Available_slot),
		}, nil
	}
	return &pb.SearchFlightResponse{}, fmt.Errorf("user does not exist")
}
func (t *FlightHandler) CreateFlight(ctx context.Context, in *pb.CreateFlightRequest) (*pb.CreateFlightResponse, error) {
	ok, err := t.flightRepo.ExistFlight(ctx, in.Id)
	if err != nil {
		return &pb.CreateFlightResponse{}, err
	}
	if !ok {
		resp, err := t.flightRepo.CreateFlight(ctx, &model.Flight{
			ID:             in.Id,
			Name:           in.Name,
			From:           in.From,
			To:             in.To,
			Date:           in.Date.AsTime(),
			Status:         in.Status,
			Available_slot: int(in.AvailableSlot),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		})
		if err != nil {
			return &pb.CreateFlightResponse{}, err
		}
		return &pb.CreateFlightResponse{
			Id:            resp.ID,
			Name:          resp.Name,
			From:          resp.From,
			To:            resp.To,
			Date:          resp.Date.String(),
			Status:        resp.Status,
			AvailableSlot: int32(resp.Available_slot),
		}, nil
	}
	return &pb.CreateFlightResponse{}, fmt.Errorf("user already exist")
}

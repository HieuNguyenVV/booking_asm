package handler

import (
	"booking_asm/api/flight_api/request"
	"booking_asm/api/flight_api/response"
	"booking_asm/pb"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	FlightHandler interface {
		SearchFlight(c *gin.Context)
		CreateFlight(c *gin.Context)
		UpdateFlight(c *gin.Context)
	}
	flightHandler struct {
		flightClient pb.FlightServiceClient
	}
)

func New(flightClient pb.FlightServiceClient) *flightHandler {
	return &flightHandler{flightClient: flightClient}
}

func (t *flightHandler) SearchFlight(c *gin.Context) {
	req := request.SearchFlightRequest{}
	req.ID = c.Param("id")
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	if validateErr, ok := err.(validator.ValidationErrors); ok {
	// 		errMessage := make([]string, 0)
	// 		for _, v := range validateErr {
	// 			errMessage = append(errMessage, v.Error())
	// 		}
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 			"status": http.StatusText(http.StatusBadRequest),
	// 			"error":  errMessage,
	// 		})
	// 		return
	// 	}
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"status": http.StatusText(http.StatusBadRequest),
	// 		"error":  err.Error(),
	// 	})
	// 	return
	// }
	resp, err := t.flightClient.SearchFlight(context.Background(), &pb.SearchFlightRequest{
		Id: req.ID,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	dto := response.SearchFlightResponse{
		ID:             resp.Id,
		Name:           resp.Name,
		From:           resp.From,
		To:             resp.To,
		Date:           resp.Date,
		Status:         resp.Status,
		Available_slot: int(resp.AvailableSlot),
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status": http.StatusText(http.StatusOK),
		"value":  dto,
	})
}

func (t *flightHandler) CreateFlight(c *gin.Context) {
	req := request.CreateFlightRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		if validateErr, ok := err.(validator.ValidationErrors); ok {
			errMessage := make([]string, 0)
			for _, v := range validateErr {
				errMessage = append(errMessage, v.Error())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessage,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	id := uuid.New()
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	resp, err := t.flightClient.CreateFlight(context.Background(), &pb.CreateFlightRequest{
		Id:            id.String(),
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		Date:          timestamppb.New(date),
		Status:        req.Status,
		AvailableSlot: int32(req.Available_slot),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	dto := response.CreateFlightResponse{
		ID:             resp.Id,
		Name:           resp.Name,
		From:           resp.From,
		To:             resp.To,
		Date:           resp.Date,
		Status:         resp.Status,
		Available_slot: int(resp.AvailableSlot),
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status": http.StatusText(http.StatusOK),
		"value":  dto,
	})
}
func (t *flightHandler) UpdateFlight(c *gin.Context) {
	req := request.UpdateFlightRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		if validatorErr, ok := err.(validator.ValidationErrors); ok {
			messageErr := make([]string, 0)
			for _, v := range validatorErr {
				messageErr = append(messageErr, v.Error())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  messageErr,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	resp, err := t.flightClient.UpdateFlight(context.Background(), &pb.UpdateFlightRequest{
		Id:            req.ID,
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		Date:          timestamppb.New(date),
		Status:        req.Status,
		AvailableSlot: int32(req.Available_slot),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	dto := response.UpdateFlightResponse{
		ID:             resp.Id,
		Name:           resp.Name,
		From:           resp.From,
		To:             resp.To,
		Date:           resp.Date,
		Status:         resp.Status,
		Available_slot: int(resp.AvailableSlot),
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status": http.StatusText(http.StatusOK),
		"value":  dto,
	})
}

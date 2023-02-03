package handler

import (
	"booking_asm/api/booking_api/request"
	"booking_asm/api/booking_api/response"
	cRes "booking_asm/api/customer_api/response"
	fRes "booking_asm/api/flight_api/response"
	"booking_asm/pb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	BookingHandler interface {
		CreatBooking(c *gin.Context)
		CancelBooking(c *gin.Context)
		ViewBookingByCode(c *gin.Context)
	}
	bookingHandler struct {
		bookingClient pb.BookingServiceClient
	}
)

func New(bookingClient pb.BookingServiceClient) BookingHandler {
	return &bookingHandler{
		bookingClient: bookingClient,
	}
}
func (bh *bookingHandler) CreatBooking(c *gin.Context) {
	req := request.CreatBookingRequest{}

	//parse form request
	if err := c.ShouldBindJSON(&req); err != nil {
		//validate form
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
	}

	//proto request
	pRequest := &pb.CreateBookingRequest{
		CustomerId: req.Customer_id,
		FlightId:   req.Flight_id,
		Code:       req.Code,
		Status:     req.Status,
		BookedDate: timestamppb.New(req.Booked_date),
	}

	//call creat new
	pResponse, err := bh.bookingClient.CreateBooking(c.Request.Context(), pRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	date, err := time.Parse("2006-01-02", pResponse.BookedDate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &response.BookingResponse{
		ID:          pResponse.Id,
		Customer_id: pResponse.CustomerId,
		Flight_id:   pResponse.FlightId,
		Code:        pResponse.Code,
		Status:      pResponse.Status,
		Booked_date: date,
	}

	//return to client
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (bh *bookingHandler) CancelBooking(c *gin.Context) {
	req := request.CancelBookingRequest{}

	//parse form request
	if err := c.ShouldBindJSON(&req); err != nil {
		//validate form
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
	}

	pRequest := pb.CancelBookingRequest{
		Code:   req.Code,
		Status: req.Status,
	}

	pResponse, err := bh.bookingClient.CancelBooking(c.Request.Context(), &pRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	date, err := time.Parse("2006-01-02", pResponse.BookedDate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &response.BookingResponse{
		ID:          pResponse.Id,
		Customer_id: pResponse.CustomerId,
		Flight_id:   pResponse.FlightId,
		Code:        pResponse.Code,
		Status:      pResponse.Status,
		Booked_date: date,
	}

	//return to client
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (bh *bookingHandler) ViewBookingByCode(c *gin.Context) {
	req := &request.SearchBookingRequest{}

	//parse form request
	if err := c.ShouldBindJSON(&req); err != nil {
		//validate form
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
	}

	pBRequest := &pb.SearchBookingRequest{
		Code: req.Code,
	}

	pBResponse, err := bh.bookingClient.SearchBooking(c.Request.Context(), pBRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	date, err := time.Parse("2006-01-02", pBResponse.BookingDetail.BookedDate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	
	dto := &response.SearchBookingResponse{
		BookingResponse: response.BookingResponse{
			ID:          pBResponse.BookingDetail.Id,
			Customer_id: pBResponse.BookingDetail.CustomerId,
			Flight_id:   pBResponse.FlightDetail.Id,
			Code:        pBResponse.BookingDetail.Code,
			Status:      pBResponse.BookingDetail.Status,
			Booked_date: date,
		},
		CustomerInfor: cRes.CreateCustomerResponse{
			ID:         pBResponse.BookingDetail.CustomerId,
			Name:       pBResponse.CustomerDetail.CustomerName,
			Phone:      pBResponse.CustomerDetail.Phone,
			License_id: pBResponse.CustomerDetail.LicenseId,
			Address:    pBResponse.CustomerDetail.Address,
			Email:      pBResponse.CustomerDetail.Email,
			Active:     pBResponse.CustomerDetail.Active,
		},
		FlightInfor: fRes.CreateFlightResponse{
			ID:             pBResponse.BookingDetail.FlightId,
			Name:           pBResponse.FlightDetail.Name,
			From:           pBResponse.FlightDetail.From,
			To:             pBResponse.FlightDetail.To,
			Date:           pBResponse.FlightDetail.Date,
			Status:         pBResponse.FlightDetail.Status,
			Available_slot: int(pBResponse.FlightDetail.AvailableSlot),
		},
	}

	//return to client
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

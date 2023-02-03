package response

import (
	cr "booking_asm/api/customer_api/response"
	fr "booking_asm/api/flight_api/response"
	"time"
)

type BookingResponse struct {
	ID          string    `json:"id"`
	Customer_id string    `json:"customer_id"`
	Flight_id   string    `json:"flight_id"`
	Code        string    `json:"code"`
	Status      string    `json:"status"`
	Booked_date time.Time `json:"booked_date"`
}

type SearchBookingResponse struct {
	BookingResponse
	CustomerInfor cr.CreateCustomerResponse
	FlightInfor   fr.CreateFlightResponse
}

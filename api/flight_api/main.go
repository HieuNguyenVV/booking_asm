package main

import (
	"booking_asm/api/flight_api/handler"
	"booking_asm/pb"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	customerConn, err := grpc.Dial(":2223", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	flightClient := pb.NewFlightServiceClient(customerConn)
	gin := gin.Default()
	h := handler.New(flightClient)

	gr := gin.Group("/v1/api/flight")
	gr.POST("/create", h.CreateFlight)
	gr.PUT("/update", h.UpdateFlight)
	//gr.GET("/history", h.BookingHistory)
	gr.GET("/find/:id", h.SearchFlight)
	//Listen and serve
	gin.Run(":3334")
}

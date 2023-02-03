package main

import (
	"booking_asm/api/booking_api/handler"
	"booking_asm/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	bookingConn, err := grpc.Dial(":2224", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	bookingClient := pb.NewBookingServiceClient(bookingConn)
	//Handler for GIN Gonic
	h := handler.New(bookingClient)
	g := gin.Default()

	//creat routes
	gr := g.Group("/v3/api")
	gr.POST("/creat", h.CreatBooking)
	gr.POST("/cancel", h.CancelBooking)
	gr.GET("/viewbycode", h.ViewBookingByCode)
	//Listen and serve
	g.Run(":3333")
}

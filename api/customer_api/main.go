package main

import (
	"booking_asm/api/customer_api/handler"
	"booking_asm/pb"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	customerConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	customerClient := pb.NewCustomerServiceClient(customerConn)
	gin := gin.Default()
	h := handler.New(customerClient)

	gr := gin.Group("/v1/api/")
	gr.POST("/create", h.CreateCustomer)
	gr.POST("/update", h.UpdateCustomer)
	gr.POST("/changepass", h.ChangePassword)
	//gr.GET("/history", h.BookingHistory)
	gr.GET("/find/:id", h.FindCustomer)
	//Listen and serve
	gin.Run(":3333")
}

package main

import (
	"booking_asm/db"
	booking_handler "booking_asm/grpc/booking_grpc/handler"
	"booking_asm/grpc/booking_grpc/repositories"
	"booking_asm/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	conf, err := db.AutoBindConfig()
	if err != nil {
		fmt.Println(err)
		log.Println("Bind config fail!")
		return
	}
	db, err := db.Open(conf)
	if err != nil {
		log.Println("Connect db fail!")
		return
	}
	customerConn, err := grpc.Dial(":2223", grpc.WithInsecure())
	flightConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	flightClient := pb.NewFlightServiceClient(flightConn)
	customerClienr := pb.NewCustomerServiceClient(customerConn)
	repo := repositories.New(db)
	h := booking_handler.New(flightClient, customerClienr, repo.BookingRepo)
	s := grpc.NewServer()
	l, err := net.Listen("tcp", ":2224")
	if err != nil {
		return
	}
	reflection.Register(s)
	pb.RegisterBookingServiceServer(s, &h)
	log.Println("Listen at port: 2223")
	s.Serve(l)
}

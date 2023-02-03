package main

import (
	"booking_asm/db"
	"booking_asm/grpc/flight-grpc/handler"
	"booking_asm/grpc/flight-grpc/repositories"
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
	repo := repositories.New(db)
	h := handler.New(repo.FlightRepo)
	s := grpc.NewServer()
	l, err := net.Listen("tcp", ":2223")
	if err != nil {
		return
	}
	reflection.Register(s)
	pb.RegisterFlightServiceServer(s, &h.FlightHandler)
	log.Println("Listen at port: 2223")
	s.Serve(l)
}

package main

import (
	"booking_asm/db"
	"booking_asm/grpc/customer-grpc/handler"
	"booking_asm/grpc/customer-grpc/repositories"
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
	h := handler.New(repo.CustomerRepo)
	s := grpc.NewServer()
	l, err := net.Listen("tcp", ":2222")
	if err != nil {
		return
	}
	reflection.Register(s)
	pb.RegisterCustomerServiceServer(s, &h.CustomerHandler)
	log.Println("Listen at port: 2222")
	s.Serve(l)
}

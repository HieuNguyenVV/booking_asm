package db

import (
	bk "booking_asm/grpc/booking_grpc/model"
	cs "booking_asm/grpc/customer-grpc/model"
	fl "booking_asm/grpc/flight-grpc/model"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (t Postgres) getSourceDb() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		t.Host, t.User, t.Password, t.DbName, t.Port)
}

func Open(conf Config) (*gorm.DB, error) {
	pstg := conf.Postgres
	psqlconn := pstg.getSourceDb()
	log.Println(psqlconn)
	//db, err := sql.Open("postgres", psqlconn)
	db, err := gorm.Open(postgres.Open((psqlconn)))
	if err != nil {
		log.Println(err)
		return db, err
	}
	db.AutoMigrate(&cs.Customer{},
		&fl.Flight{},
		&bk.Booking{})

	return db.Debug(), nil
}

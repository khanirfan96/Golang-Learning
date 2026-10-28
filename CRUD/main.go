package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	dns := "host=localhost user=postgres password=1234 dbname=crud port=5432 sslmode=disable"

	var err error

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connected")

	// Auto Create table
	db.AutoMigrate(&User{})

	fmt.Println("Server is running")

	http.ListenAndServe(":8008", nil)

}

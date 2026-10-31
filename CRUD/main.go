package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	dsn := "host=127.0.0.1 user=postgres password=1234 dbname=crud port=5432 sslmode=disable"

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connected")

	// Auto Create table
	db.AutoMigrate(&User{})

	http.HandleFunc("/users", userHandler)

	fmt.Println("Server is running")

	http.ListenAndServe(":8008", nil)

}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createUser(w, r)

	case http.MethodGet:
		getAllUser(w, r)

	case http.MethodDelete:
		deleteUser(w, r)

	case http.MethodPut:
		updateUser(w, r)

	default:
		fmt.Println("No Method Found")
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {

	// ResponseWriter - Iska matlab hota hai client(FrontEnd) ko jo Data send kr rhe.

	// Request - Iska matlab hota hai client(Frontend) ne kya send kiya hai or Jo frontend ne send kiya hai.

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	// r.Body mein Frontend se aaya hua JSON data hota hai.
	// Decoder ka kaam JSON ko read karna aur Go data(Struct) mein convert karna hai.

	// NewDecoder - NewDecoder() us JSON data ko read karne ke liye ek Decoder create karta hai.

	// Decode - Decode() JSON data ko Go struct mein convert karta hai.

	// &user ka matlab hai user variable ka memory address pass karna.

	result := db.Create(&user)

	// Create - Create() database mein new record insert karta hai.

	// &user user struct ka address pass karta hai.
	// GORM struct ki values lekar SQL INSERT query execute karta hai.

	if result.Error != nil {
		http.Error(w, result.Error.Error(), 500)
		return
	}

	// - w → Client ko response bhejne ke liye.
	// - result.Error.Error() → Actual error message.
	// - 500 → HTTP Internal Server Error status code.
	// return ka matlab function ko yahin stop kar do.

	json.NewEncoder(w).Encode(user)

	// w ka matlab hai response writer.
	// Encoder ka kaam Go data ko JSON mein convert karke HTTP response mein write karna hai.

	// NewEncoder - NewEncoder() ek JSON Encoder create karta hai.

	// Encode - Ye Go struct ko JSON format mein convert karke Frontend ko bhej deta hai.
}

func getAllUser(w http.ResponseWriter, r *http.Request) {

	var user []User

	result := db.Find(&user)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idStr)

	// id, _ := strconv.Atoi(idStr) == Ye line string ID ko convert kr rha int/parseint me.

	result := db.Delete(&User{}, id)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("User is Deleted")
}

func updateUser(w http.ResponseWriter, r *http.Request) {

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if user.ID == 0 {
		http.Error(w, "ID required for update", 400)
		return
	}

	result := db.Model(&User{}).Where("id=?", user.ID).Updates(user)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("User Data Updated")
}

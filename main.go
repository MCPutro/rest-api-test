package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MCPutro/rest-api-test/config"
	"github.com/MCPutro/rest-api-test/services"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseCOnnection()

	user services.UserRepository = services.UserRepository{}
)

func main() {
	//defer config.DbDisconection(db)

	//create table
	db.AutoMigrate(user)

	func() {
		defer config.DbDisconection(db)
	}()
	routing()

}

func routing() {
	fmt.Println("server running")

	myRoute := mux.NewRouter().StrictSlash(true)

	PORT := os.Getenv("PORT")
	myRoute.HandleFunc("/", chectAPI).Methods("GET")
	myRoute.HandleFunc("/check", chectAPI).Methods("GET")

	myRoute.HandleFunc("/api/user/create", user.CreateUser).Methods("POST")
	myRoute.HandleFunc("/api/user/profile", user.FindUserByEmail).Methods("POST")
	myRoute.HandleFunc("/api/user/all", user.FindAllUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+PORT, myRoute))
}

func chectAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MCPutro/rest-api-test/config"
	"github.com/MCPutro/rest-api-test/entities"
	"github.com/MCPutro/rest-api-test/services"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseCOnnection()

	user services.User = services.User{}
)

func main() {
	defer config.DbDisconection(db)

	//create table
	db.AutoMigrate(entities.SocialMedia{})
	user.Connection = db

	// db.Save(
	// 	&entities.SocialMedia{
	// 		Name: "FB",
	// Accounts: []entities.User{
	// 	{Name: "orang1", Email: "orang1@gmail.com", Password: "orang1"},
	// 	{Name: "orang2", Email: "orang2@gmail.com", Password: "orang2"},
	// },
	// 	},
	// )

	// u1 := entities.User{Name: "orang1", Email: "orang1@gmail.com", Password: "orang1"}
	// db.Save(&u1)
	//sosMed := repo.SocialMedia{connection: db}

	func() {
		//defer config.DbDisconection(db)
	}()
	routing()

}

func routing() {
	fmt.Println("server running")

	myRoute := mux.NewRouter().StrictSlash(true)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9999"
	}
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

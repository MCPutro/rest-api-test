package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MCPutro/rest-api-test/config"
	repo "github.com/MCPutro/rest-api-test/repository"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseCOnnection()

	user repo.User = repo.User{}
)

func main() {
	defer config.DbDisconection(db)

	//create table
	db.AutoMigrate(user)

	/**func() {
		defer config.DbDisconection(db)
	}()**/
	routing()

}

func routing() {
	fmt.Println("server run in port 9999")

	myRoute := mux.NewRouter().StrictSlash(true)

	myRoute.HandleFunc("/", chectAPI).Methods("GET")
	myRoute.HandleFunc("/api/user/create", user.CreateUser).Methods("POST")
	myRoute.HandleFunc("/check", chectAPI).Methods("GET")

	log.Fatal(http.ListenAndServe(":9999", myRoute))
}

func chectAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

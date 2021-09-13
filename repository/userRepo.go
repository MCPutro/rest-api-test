package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/MCPutro/rest-api-test/config"
	"github.com/MCPutro/rest-api-test/entities"
)

type User struct {
	entities.UserEntity
}

func (u User) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(requestPayload))

	//fmt.Println("header Authorization : ", r.Header.Get("Authorization"))

	//parsing json test to user model
	var tmp_user User
	json.Unmarshal(requestPayload, &tmp_user)

	//set createddate
	tmp_user.CreatedDate = time.Now()

	tmp_user.insertUser()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("oke", "haha") //set data to header resp
	w.WriteHeader(http.StatusOK)

	//w.Write(result)
}

func (u *User) insertUser() {
	db := config.SetupDatabaseCOnnection()
	db.Save(u)
}

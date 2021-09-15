package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MCPutro/rest-api-test/config"
	"github.com/MCPutro/rest-api-test/entities"
	"github.com/MCPutro/rest-api-test/repository"
	"github.com/gorilla/mux"
)

type SosMed struct {
	repository.SocialMedia
}

func (sm SosMed) CreateSosMed(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := ioutil.ReadAll(r.Body)
	fmt.Println("req masuk : ", string(requestPayload))

	json.Unmarshal(requestPayload, &sm.SocialMediaIdentity) //parsing data to variable sm

	tambahSosMed := sm.AddSosMed()
	fmt.Fprint(w, tambahSosMed.Error())
}

func (sm SosMed) FindAll(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDatabaseCOnnection()
	param := mux.Vars(r)
	param1 := param["sm"]

	defer config.DbDisconection(db)

	var sosmed []entities.SocialMedia
	if len(param1) != 0 {
		db.Preload("Accounts").Where("name = ?", param1).Find(&sosmed)
	} else {
		db.Preload("Accounts").Find(&sosmed)
	}

	respJson, _ := json.Marshal(sosmed)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(respJson)

}

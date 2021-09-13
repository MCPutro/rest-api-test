package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	repo "github.com/MCPutro/rest-api-test/repository"
	"github.com/MCPutro/rest-api-test/response"
)

type UserRepository struct {
	repo.User
}

func (ur UserRepository) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(requestPayload))

	//fmt.Println("header Authorization : ", r.Header.Get("Authorization"))

	//parsing json test to user model
	//var tmp_user UserRepository
	json.Unmarshal(requestPayload, &ur)

	//set createddate
	ur.CreatedDate = time.Now()

	resp := response.Response{}
	result := ur.InsertUser()
	if result != nil {
		resp = response.Response{Code: strconv.Itoa(http.StatusInternalServerError), Message: result.Error(), Data: ur}
	} else {
		resp = response.Response{Code: "200", Message: "Succes insert ", Data: ur}
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("oke", "haha") //set data to header resp
	w.WriteHeader(http.StatusOK)

	w.Write(respJson)
}

func (ur UserRepository) FindUserByEmail(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := ioutil.ReadAll(r.Body)
	//var tmp_user UserRepository
	json.Unmarshal(requestPayload, &ur)

	existingUser, err := ur.FindByEmail(ur.Email)
	if err != nil {
		//resp := response.Response{Code: "200", Message: "Succes", Data: existingUser}
	}
	resp := response.Response{Code: "200", Message: "Succes", Data: existingUser}

	respJson, err := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("oke", "haha") //set data to header resp
	w.WriteHeader(http.StatusOK)

	w.Write(respJson)

}

func (ur UserRepository) FindAllUser(w http.ResponseWriter, r *http.Request) {
	users, _ := ur.FindAll()
	resp := response.Response{Code: "200", Message: "Succes", Data: users}
	respJson, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("oke", "haha") //set data to header resp
	w.WriteHeader(http.StatusOK)

	w.Write(respJson)
}

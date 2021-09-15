package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	repo "github.com/MCPutro/rest-api-test/repository"
	"github.com/MCPutro/rest-api-test/response"
	"github.com/gorilla/mux"
)

type User struct {
	repo.User
}

func (ur User) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := ioutil.ReadAll(r.Body)
	fmt.Println("req masuk : ", string(requestPayload))

	//fmt.Println("header Authorization : ", r.Header.Get("Authorization"))

	//parsing json test to user model
	//var tmp_user UserRepository
	json.Unmarshal(requestPayload, &ur.UserIdentity)

	//set createddate
	ur.UserIdentity.CreatedDate = time.Now()

	fmt.Println("ur : ", ur)

	var resp = response.Response{}
	result := ur.InsertUser()
	//fmt.Println("result : ", result)
	if result != nil {
		resp = response.Response{Code: strconv.Itoa(http.StatusInternalServerError), Message: result.Error(), Data: ur.UserIdentity}
	} else {
		resp = response.Response{Code: "200", Message: "bisa insert ", Data: ur.UserIdentity}
	}

	//fmt.Println(" --- > resp : ", resp)
	respJson, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("oke", "haha") //set data to header resp
	w.WriteHeader(http.StatusOK)

	w.Write(respJson)
}

func (ur User) FindUserByEmail(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := ioutil.ReadAll(r.Body)
	//var tmp_user UserRepository
	json.Unmarshal(requestPayload, &ur.UserIdentity)

	existingUser, err := ur.FindByEmail()
	//fmt.Println("-->err : ", err)

	var resp response.Response
	if err != nil {
		resp = response.Response{Code: "500", Message: err.Error()}
	} else {
		resp = response.Response{Code: "200", Message: "Succes", Data: existingUser}
	}

	respJson, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("oke", "haha") //set data to header resp
	w.WriteHeader(http.StatusOK)

	w.Write(respJson)

}

func (ur User) FindAllUser(w http.ResponseWriter, r *http.Request) {
	users, _ := ur.FindAll()
	resp := response.Response{Code: "200", Message: "Succes", Data: users}
	respJson, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("oke", "haha") //set data to header resp
	w.WriteHeader(http.StatusOK)

	w.Write(respJson)
}

func (ur User) Subscribe(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := ioutil.ReadAll(r.Body)
	param := mux.Vars(r)
	sm := param["sm"]

	json.Unmarshal(requestPayload, &ur.UserIdentity)

	status := ur.SubscribeSosMed(strings.ToUpper(sm))

	var resp response.Response
	if status != nil {
		resp = response.Response{Code: "500", Message: status.Error()}
	} else {
		resp = response.Response{Code: "200", Message: "Subscribe Succes"}
	}

	respJson, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("oke", "haha") //set data to header resp
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

func (ur User) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := ioutil.ReadAll(r.Body)
	param := mux.Vars(r)
	sm := param["sm"]

	json.Unmarshal(requestPayload, &ur.UserIdentity)

	status := ur.UnSubscribeSosMed(strings.ToUpper(sm))

	var resp response.Response
	if status != nil {
		resp = response.Response{Code: "500", Message: status.Error()}
	} else {
		resp = response.Response{Code: "200", Message: "Unsubscribe Succes"}
	}
	respJson, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("oke", "haha") //set data to header resp
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

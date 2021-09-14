package repository

import (
	/**"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"**/

	"fmt"

	"github.com/MCPutro/rest-api-test/config"
	"github.com/MCPutro/rest-api-test/entities"
	"gorm.io/gorm"
	//"github.com/MCPutro/rest-api-test/response"
)

//
type User struct {
	entities.User
	Connection *gorm.DB
}

func (u *User) InsertUser() error {
	/**db := config.SetupDatabaseCOnnection()
	result := db.Save(u)**/
	//result := u.Connection.Save(u)
	// u1 := entities.User{Name: "orang1", Email: "orang1@gmail.com", Password: "orang1"}
	// fmt.Println("u1 : ", u1)
	fmt.Println("u : ", *u)
	result := u.Connection.Table("users").Create(u)
	if result.Error != nil {
		return result.Error
	}
	config.DbDisconection(u.Connection)
	return nil
}

func (u *User) FindByEmail(email string) (User, error) {
	tmp_user := User{}

	//db := config.SetupDatabaseCOnnection()

	res := u.Connection.Where("email = ?", email).Take(&tmp_user)

	config.DbDisconection(u.Connection)

	if res.Error != nil {
		return tmp_user, res.Error
	}
	return tmp_user, nil
}

func (u *User) FindAll() ([]User, error) {
	users := []User{}
	db := config.SetupDatabaseCOnnection()

	res := db.Find(&users)

	config.DbDisconection(db)
	if res == nil {
		return nil, res.Error
	}

	fmt.Println(users)
	return users, nil
}

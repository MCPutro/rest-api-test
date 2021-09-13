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
	//"github.com/MCPutro/rest-api-test/response"
)

type User struct {
	entities.UserEntity
}

func (u *User) InsertUser() error {
	db := config.SetupDatabaseCOnnection()
	result := db.Save(u)
	if result.Error != nil {
		return result.Error
	}
	config.DbDisconection(db)
	return nil
}

func (u *User) FindByEmail(email string) (User, error) {
	tmp_user := User{}

	db := config.SetupDatabaseCOnnection()

	res := db.Where("email = ?", email).Take(&tmp_user)

	config.DbDisconection(db)

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

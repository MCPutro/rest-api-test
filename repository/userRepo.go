package repository

import (
	/**"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"**/

	"errors"
	"fmt"

	"github.com/MCPutro/rest-api-test/config"
	"github.com/MCPutro/rest-api-test/entities"
	"gorm.io/gorm"
	//"github.com/MCPutro/rest-api-test/response"
)

//
type User struct {
	UserIdentity entities.User
	Connection   *gorm.DB
}

func (u *User) InsertUser() error {
	/**db := config.SetupDatabaseCOnnection()
	result := db.Save(u)**/
	//result := u.Connection.Save(u)
	// u1 := entities.User{Name: "orang1", Email: "orang1@gmail.com", Password: "orang1"}
	// fmt.Println("u1 : ", u1)
	//fmt.Println("u : ", u)

	_, errExistingUser := u.FindByEmail()

	fmt.Println("->>", errExistingUser)

	if errExistingUser != nil { //insert krn errExistingUser = record not found
		result := u.Connection.Create(&u.UserIdentity)
		if result.Error != nil {
			return result.Error
		}
	} else {
		return errors.New("email sudah ada gan")
	}

	return nil
}

func (u *User) FindByEmail() (entities.User, error) {
	tmp_user := entities.User{}

	//db := config.SetupDatabaseCOnnection()

	res := u.Connection.Where("email = ?", u.UserIdentity.Email).First(&tmp_user)

	//fmt.Println("--> ", res)

	//config.DbDisconection(u.Connection)

	if res.Error != nil {
		return tmp_user, res.Error
	}
	return tmp_user, nil
}

func (u *User) FindAll() ([]entities.User, error) {
	users := []entities.User{}
	db := config.SetupDatabaseCOnnection()

	res := db.Find(&users)

	config.DbDisconection(db)
	if res == nil {
		return nil, res.Error
	}

	fmt.Println(users)
	return users, nil
}

func (u *User) SubscribeSosMed(sosMedName string) error {
	var sosMed entities.SocialMedia
	db := config.SetupDatabaseCOnnection()

	get_sosmed := db.Where("name = ?", sosMedName).Find(&sosMed)
	tmp_user, _ := u.FindByEmail()

	if get_sosmed.Error != nil {
		config.DbDisconection(db)
		return errors.New("social media not found")
	} else {
		if tmp_user.Id <= 0 {
			return errors.New("user not found")
		} else {
			db.Model(&sosMed).Association("Accounts").Append(&tmp_user)
		}
	}
	config.DbDisconection(db)
	return nil
}

func (u *User) UnSubscribeSosMed(sosMedName string) error {
	var sosMed entities.SocialMedia
	db := config.SetupDatabaseCOnnection()

	get_sosmed := db.Where("name = ?", sosMedName).Find(&sosMed)
	tmp_user, _ := u.FindByEmail()

	if get_sosmed.Error != nil {
		config.DbDisconection(db)
		return errors.New("social media not found")
	} else {
		if tmp_user.Id <= 0 {
			return errors.New("user not found")
		} else {
			db.Model(&sosMed).Association("Accounts").Delete(&tmp_user) //Association("Accounts") -> param adalah nama variable
		}
	}
	config.DbDisconection(db)
	return nil
}

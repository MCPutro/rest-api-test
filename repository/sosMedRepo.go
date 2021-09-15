package repository

import (
	"errors"
	"fmt"

	"github.com/MCPutro/rest-api-test/entities"
	"gorm.io/gorm"
)

type SocialMedia struct {
	SocialMediaIdentity entities.SocialMedia
	Connection          *gorm.DB
}

func (sm *SocialMedia) AddSosMed() error {
	_, isExist := sm.findSosMed()
	//fmt.Println(isExist)
	if isExist == nil { //positif
		sm.Connection.Create(&sm.SocialMediaIdentity)
		return nil
	}
	return errors.New("already exist")

}

func (sm *SocialMedia) findSosMed() (entities.SocialMedia, error) {
	var exiting entities.SocialMedia
	hasil := sm.Connection.Where("name = ?", sm.SocialMediaIdentity.Name).Find(&exiting)
	fmt.Println("findSosMed : ", hasil.Error)
	fmt.Println("exiting : ", exiting.Id)

	if exiting.Id == 0 {
		return exiting, nil
	}

	return exiting, errors.New("already exist")
}

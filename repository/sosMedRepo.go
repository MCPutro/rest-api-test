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
	fmt.Println(isExist)
	if isExist == nil { //positif
		sm.Connection.Create(&sm.SocialMediaIdentity)
		return nil
	}
	return errors.New("error gan")

}

func (sm *SocialMedia) findSosMed() (entities.SocialMedia, error) {
	var exiting entities.SocialMedia
	hasil := sm.Connection.Where("name = ?", sm.SocialMediaIdentity.Name).Find(&exiting)

	if hasil.Error != nil {
		return exiting, hasil.Error
	}
	return exiting, nil
}

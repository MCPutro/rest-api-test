package repository

import (
	"github.com/MCPutro/rest-api-test/entities"
	"gorm.io/gorm"
)

type SocialMedia struct {
	entities.SocialMedia
	Connection *gorm.DB
}

func (sm SocialMedia) addSosMed() {

}

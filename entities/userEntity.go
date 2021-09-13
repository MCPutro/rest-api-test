package entities

import "time"

type UserEntity struct {
	Id          int    `gorm:"primary_key, AUTO_INCREMENT"`
	Name        string `gorm:"type:varchar(100)"`
	Email       string `gorm:"type:varchar(100); UNIQUE_key"`
	Password    string
	Status      bool
	CreatedDate time.Time `json:"CreatedDate"`
}

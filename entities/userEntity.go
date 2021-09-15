package entities

import "time"

//entity
type User struct {
	Id          int    `gorm:"primary_key, AUTO_INCREMENT"`
	Name        string `gorm:"type:varchar(100)"`
	Email       string `gorm:"type:varchar(100); UNIQUE"`
	Password    string
	Status      bool
	CreatedDate time.Time `json:"CreatedDate"`
	//Subscriber  string    `gorm:"-"`
	//`gorm:",omitempty"`
}

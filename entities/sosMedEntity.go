package entities

type SocialMedia struct {
	Id       int64  `gorm:"primary_key, AUTO_INCREMENT"`
	Name     string `gorm:"type:varchar(100)"`
	Accounts []User `gorm:"many2many:user_sosmed"`
}

//cara 2
/**type Orang struct {
	gorm.Model
	Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;JoinReferences:ProfileRefer"`
	Refer    uint      `gorm:"index:,unique"`
}

type Profile struct {
	gorm.Model
	Name      string
	UserRefer uint `gorm:"index:,unique"`
}**/

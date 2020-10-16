package domain

type User struct {
	Id          int    `gorm:"column:id"`
	UserName    string `gorm:"column:username"`
	Password    string `gorm:"column:password"`
	Email       string `gorm:"column:email"`
	State       int    `gorm:"column:state"`
	Nickname    string `gorm:"column:nickname"`
	Gender      string `gorm:"column:gender"`
	Province    string `gorm:"column:province"`
	Country     string `gorm:"column:country"`
	Phonenumber string `gorm:"column:phonenumber"`
}

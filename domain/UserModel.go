package domain

type User struct {
	Id          int    `gorm:"column:id" json:"id"`
	UserName    string `gorm:"column:username" json:"username"`
	Password    string `gorm:"column:password" json:"password"`
	Email       string `gorm:"column:email" json:"email"`
	State       int    `gorm:"column:state" json:"state"`
	Nickname    string `gorm:"column:nickname" json:"nickname"`
	Gender      string `gorm:"column:gender" json:"gender"`
	Province    string `gorm:"column:province" json:"province"`
	Country     string `gorm:"column:country" json:"country"`
	Phonenumber string `gorm:"column:phonenumber" json:"phonenumber"`
}

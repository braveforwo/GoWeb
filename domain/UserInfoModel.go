package domain

type Userinfo struct {
	UserId           int    `gorm:"column:userid" json:"userid" `
	Id               int    `gorm:"column:id" json:"id"`
	SelfIntroduction string `gorm:"column:selfintroduction" json:"selfintroduction"`
	Avatar           string `gorm:"column:avatar" json:"avatar"`
	Address          string `gorm:"column:address" json:"address"`
	NickName         string `gorm:"column:nickname" json:"nickname"`
}

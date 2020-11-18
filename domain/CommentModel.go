package domain

type Comment struct {
	Id        int      `gorm:"column:id" json:"id" form:"id"`
	Articleid int      `gorm:"column:articleid" json:"articleid" form:"articleid"`
	Comments  string   `gorm:"column:comment" json:"comment" form:"comment"`
	Time      string   `gorm:"column:time" json:"time" form:"time"`
	Userid    int      `gorm:"column:userid" json:"userid" form:"userid"`
	UserInfo  Userinfo `gorm:"ForeignKey:UserId;AssociationForeignKey:Userid"`
	Users     User     `gorm:"ForeignKey:Id;AssociationForeignKey:Userid""`
}

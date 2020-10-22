package domain

type Article struct {
	Id          int    `gorm:"column:id" json:"id"`
	Title       string `gorm:"column:title" form:"title" json:"title"`
	Md          string `gorm:"column:md" form:"md" json:"md"`
	Html        string `gorm:"column:html" form:"html" json:"html"`
	UserId      int    `gorm:"column:userid" json:"userid"`
	ArticleType string `gorm:"column:articletype" form:"articletype" json:"articletype"`
	Pageviews   int    `gorm:"column:pageviews" form:"pageviews" json:"pageviews"`
	Comments    int    `gorm:"column:comments" form:"comments" json:"comments"`
	Time        string `gorm:"column:time" form:"time" json:"time"`
	Users       User   `gorm:"ForeignKey:userid;"`
}

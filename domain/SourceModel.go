package domain

type Source struct {
	Id            int    `gorm:"column:id" json:"id"`
	SourcePath    string `gorm:"column:sourcepath" json:"sourcepath"`
	SourceName    string `gorm:"column:sourcename" json:"sourcename"`
	UserId        int    `gorm:"column:userid" json:"userid"`
	SourceType    string `gorm:"column:sourcetype" json:"sourcetype"`
	SourceComment string `gorm:"column:sourcecomment" json:"sourcecomment"`
	Users         User   `gorm:"column:users" json:"users"`
}

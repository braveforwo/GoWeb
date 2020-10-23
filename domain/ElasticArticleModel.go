package domain

type ElasticArticleModel struct {
	Id          int `uri:"id" form:"id"`
	Title       string
	Md          string
	ArticleType string
	Author      string
	Time        string
	Pageviews   int
	Comment     int
}

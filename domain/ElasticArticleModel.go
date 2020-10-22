package domain

type ElasticArticleModel struct {
	Id          int
	Title       string
	Md          string
	ArticleType string
	Author      string
	Time        string
	Pageviews   int
	Comment     int
}

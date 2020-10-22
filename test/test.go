package main

import (
	"GoWeb/connector"
	"GoWeb/domain"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func create(client *elastic.Client) {
	db := connector.GetDBConn()
	var articles []domain.Article
	db.Preload("Users").Find(&articles)
	//使用结构体
	for i := 0; i <= len(articles); i++ {
		var elsticArticleModel domain.ElasticArticleModel
		elsticArticleModel.Id = articles[i].Id
		elsticArticleModel.Time = articles[i].Time
		elsticArticleModel.Pageviews = articles[i].Pageviews
		elsticArticleModel.Md = articles[i].Md
		elsticArticleModel.Author = articles[i].Users.UserName
		elsticArticleModel.Title = articles[i].Title
		elsticArticleModel.Comment = articles[i].Comments
		elsticArticleModel.ArticleType = articles[i].ArticleType
		put1, err := client.Index().
			Index("articlelibrary").
			Type("article").
			Id(string(elsticArticleModel.Id)).
			BodyJson(elsticArticleModel).
			Do(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)
	}
}

func delete(client *elastic.Client) {
	res, err := client.DeleteIndex("articlelibrary").Do(context.Background())
	//fmt.Printf("delete result %s\n", res.Result)
	//res, err := client.Delete().Index("articlelibrary").
	//	Type("article").
	//	Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Acknowledged)
}
func main() {
	create(connector.GetElasticConn())
}

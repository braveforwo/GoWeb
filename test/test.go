package main

import (
	"GoWeb/connector"
	"GoWeb/domain"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
)

//从数据库获取所有的articles并上传到elasticsearch中
func create(client *elastic.Client) {
	db := connector.GetDBConn()
	var articles []domain.Article
	db.Preload("Users").Find(&articles)
	//使用结构体
	for i := 0; i <= len(articles)-1; i++ {
		var elsticArticleModel domain.ElasticArticleModel
		fmt.Println(articles[i].Id)
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
			Id(string(articles[i].Id)).
			BodyJson(elsticArticleModel).
			Do(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)
	}
}
func gets(client *elastic.Client) {
	//通过id查找
	get1, err := client.Get().Index("articlelibrary").Type("article").Id(string(1)).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
		var bb domain.ElasticArticleModel
		err := json.Unmarshal(get1.Source, &bb)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(bb.Md)
		fmt.Println(string(get1.Source))
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
func query(client *elastic.Client) {
	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("Md", "bootstrap-fileinput")
	res, err := client.Search("articlelibrary").Type("article").Query(matchPhraseQuery).Do(context.Background())
	printEmployee(res, err)
}
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ domain.ElasticArticleModel
	for key, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(domain.ElasticArticleModel)
		fmt.Println(key)
		fmt.Printf("%#v\n", t)
	}
}
func queryAll(client *elastic.Client) {
	res, err := client.Search("articlelibrary").Type("article").Size(2).From(2).Do(context.Background())
	printEmployee(res, err)
}
func main() {
	create(connector.GetElasticConn())
}

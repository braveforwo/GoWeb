package impl

import (
	"GoWeb/connector"
	"GoWeb/domain"
	"context"
	"fmt"
)

type UploadArticleServiceImpl struct {
}

func (uas UploadArticleServiceImpl) UploadArticle(article *domain.Article) error {
	db := connector.GetDBConn()
	tx := db.Begin()
	if err := tx.Create(article).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (uas UploadArticleServiceImpl) UploadArticleToElastic(article *domain.Article) error {
	client := connector.GetElasticConn()
	var elsticArticleModel domain.ElasticArticleModel
	fmt.Println(article.Id)
	elsticArticleModel.Id = article.Id
	elsticArticleModel.Time = article.Time
	elsticArticleModel.Pageviews = article.Pageviews
	elsticArticleModel.Md = article.Md
	elsticArticleModel.Author = article.Users.UserName
	elsticArticleModel.Title = article.Title
	elsticArticleModel.Comment = article.Comments
	elsticArticleModel.ArticleType = article.ArticleType
	_, err := client.Index().
		Index("articlelibrary").
		Type("article").
		Id(string(article.Id)).
		BodyJson(elsticArticleModel).
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil

}

package impl

import (
	"GoWeb/connector"
	"GoWeb/domain"
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"math"
	"reflect"
)

type SearchArticleServiceImpl struct {
}

func (uas SearchArticleServiceImpl) SearchArticleServiceFromMysql(articleSearchCondition *domain.ArticleSearchCondition) (error, []domain.Article) {
	db := connector.GetDBConn()
	var articlelist []domain.Article
	if articleSearchCondition.SearchContext == "" {
		if err := db.Preload("Users").Order("id desc").Offset(articleSearchCondition.CurrentPage * articleSearchCondition.PageSize).Limit(articleSearchCondition.PageSize).Find(&articlelist).Error; err != nil {
			return err, articlelist
		}
		return nil, articlelist
	}
	if err := db.Preload("Users").Where("md like ?", "%"+articleSearchCondition.SearchContext+"%").Order("id desc").Offset(articleSearchCondition.CurrentPage * articleSearchCondition.PageSize).Limit(articleSearchCondition.PageSize).Find(&articlelist).Error; err != nil {
		return err, articlelist
	}
	return nil, articlelist
}

func (uas SearchArticleServiceImpl) SearchArticleServiceFromElastic(articleSearchCondition *domain.ArticleSearchCondition) (error, []domain.ElasticArticleModel) {
	elasticConn := connector.GetElasticConn()
	if articleSearchCondition.SearchContext == "" {
		return queryByCurrentPage(elasticConn, articleSearchCondition)
	}
	return queryByCurrentCurrentPageAndSearchContext(elasticConn, articleSearchCondition)
}

func (uas SearchArticleServiceImpl) SearchArticleByIdFromMysql(elasticArticleModel *domain.ElasticArticleModel) (error, domain.Article) {
	var article domain.Article
	db := connector.GetDBConn()
	if err := db.Where("id = ?", elasticArticleModel.Id).Preload("Users").Find(&article).Error; err != nil {
		return err, article
	}
	return nil, article

}

func (uas SearchArticleServiceImpl) UpdateArticlePageViewsToMysql(elasticArticleModel *domain.ElasticArticleModel) (error, domain.Article) {
	db := connector.GetDBConn()
	var article domain.Article
	tx := db.Begin()
	if err := tx.Exec("update article set pageviews=pageviews+1 where id = ?", elasticArticleModel.Id).
		Where("id = ?", elasticArticleModel.Id).Preload("Users").
		Find(&article).Error; err != nil {
		tx.Rollback()
		return err, article
	}
	tx.Commit()
	return nil, article
}

func (uas SearchArticleServiceImpl) UpdateArticlePageViewsToElastic(elasticArticleModel *domain.ElasticArticleModel) error {
	client := connector.GetElasticConn()
	elasticArticleModel.Pageviews = elasticArticleModel.Pageviews + 1
	_, err := client.Update().
		Index("articlelibrary").
		Type("article").
		Id(string(elasticArticleModel.Id)).
		Doc(map[string]interface{}{"Pageviews": elasticArticleModel.Pageviews}).
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func queryByCurrentPage(client *elastic.Client, articleSearchCondition *domain.ArticleSearchCondition) (error, []domain.ElasticArticleModel) {
	articleSearchCondition.AllPageSize = int(math.Ceil(float64(queryCount(client, articleSearchCondition)) / float64(articleSearchCondition.PageSize)))
	if articleSearchCondition.CurrentPage > articleSearchCondition.AllPageSize {
		articleSearchCondition.CurrentPage = articleSearchCondition.AllPageSize
		return errors.New("没有下一页了！"), nil
	}
	if articleSearchCondition.CurrentPage < 1 {
		articleSearchCondition.CurrentPage = 1
		return errors.New("没有上一页了！"), nil
	}
	res, err := client.Search("articlelibrary").
		Type("article").
		Size(articleSearchCondition.PageSize).
		From((articleSearchCondition.CurrentPage - 1) * articleSearchCondition.PageSize).
		Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, getElasticArticleModels(res)

}

func queryByCurrentCurrentPageAndSearchContext(client *elastic.Client, articleSearchCondition *domain.ArticleSearchCondition) (error, []domain.ElasticArticleModel) {
	articleSearchCondition.AllPageSize = int(math.Ceil(float64(queryCount(client, articleSearchCondition)) / float64(articleSearchCondition.PageSize)))
	//fmt.Println(articleSearchCondition.AllPageSize)
	if articleSearchCondition.CurrentPage > articleSearchCondition.AllPageSize {
		articleSearchCondition.CurrentPage = articleSearchCondition.AllPageSize
		return errors.New("没有下一页了！"), nil
	}
	if articleSearchCondition.CurrentPage < 1 {
		articleSearchCondition.CurrentPage = 1
		return errors.New("没有上一页了！"), nil
	}
	matchPhraseQuery := elastic.NewMatchPhraseQuery("Md", articleSearchCondition.SearchContext)
	res, err := client.Search("articlelibrary").
		Type("article").
		Size(articleSearchCondition.PageSize).
		From((articleSearchCondition.CurrentPage - 1) * articleSearchCondition.PageSize).
		Query(matchPhraseQuery).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, getElasticArticleModels(res)
}

func getElasticArticleModels(res *elastic.SearchResult) []domain.ElasticArticleModel {
	var elasticArticleModels []domain.ElasticArticleModel
	var elasticArticleModel domain.ElasticArticleModel
	for _, item := range res.Each(reflect.TypeOf(elasticArticleModel)) { //从搜索结果中取数据的方法
		elasticArticleModels = append(elasticArticleModels, item.(domain.ElasticArticleModel))
	}
	return elasticArticleModels
}

func queryCount(client *elastic.Client, articleSearchCondition *domain.ArticleSearchCondition) int {
	var res *elastic.SearchResult
	var err error
	if articleSearchCondition.SearchContext == "" {
		res, err = client.Search("articlelibrary").
			Type("article").
			Do(context.Background())
	} else {
		matchPhraseQuery := elastic.NewMatchPhraseQuery("Md", articleSearchCondition.SearchContext)
		res, err = client.Search("articlelibrary").
			Type("article").
			Query(matchPhraseQuery).Do(context.Background())
	}
	if err != nil {
		return 0
	}
	return int(res.TotalHits())

}

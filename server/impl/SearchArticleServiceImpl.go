package impl

import (
	"GoWeb/connector"
	"GoWeb/domain"
)

type SearchArticleServiceImpl struct {
}

func (uas SearchArticleServiceImpl) SearchArticleService(articleSearchCondition *domain.ArticleSearchCondition) (error, []domain.Article) {
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

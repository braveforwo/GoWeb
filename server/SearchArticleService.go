package server

import "GoWeb/domain"

type SearchArticleService interface {
	SearchArticleServiceFromMysql(articleSearchCondition *domain.ArticleSearchCondition) (error, []domain.Article)
	SearchArticleServiceFromElastic(articleSearchCondition *domain.ArticleSearchCondition) (error, []domain.ElasticArticleModel)
}

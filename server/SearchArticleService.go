package server

import "GoWeb/domain"

type SearchArticleService interface {
	SearchArticleServiceFromMysql(articleSearchCondition *domain.ArticleSearchCondition) (error, []domain.Article)
	SearchArticleServiceFromElastic(articleSearchCondition *domain.ArticleSearchCondition) (error, []domain.ElasticArticleModel)
	SearchArticleByIdFromMysql(elasticArticleModel *domain.ElasticArticleModel) (error, domain.Article)
	UpdateArticlePageViewsToMysql(elasticArticleModel *domain.ElasticArticleModel) (error, domain.Article)
	UpdateArticlePageViewsToElastic(elasticArticleModel *domain.ElasticArticleModel) error
}

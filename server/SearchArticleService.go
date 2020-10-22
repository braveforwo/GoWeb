package server

import "GoWeb/domain"

type SearchArticleService interface {
	SearchArticleService(articleSearchCondition *domain.ArticleSearchCondition) error
}

package server

import "GoWeb/domain"

type UploadArticleService interface {
	UploadArticle(article *domain.Article) error
	UploadArticleToElastic(article *domain.Article) error
}

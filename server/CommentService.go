package server

import "GoWeb/domain"

type CommentService interface {
	CreateComment(comment *domain.Comment) error
	GetCommentByArticleIdAndPageFromMysql(commentPageCondition *domain.CommentPageCondition) (error, []domain.Comment)
	UpdateCommentNumToMysql(article *domain.Article) error
	UpdateCommentNumToElastic(article *domain.Article) error
}

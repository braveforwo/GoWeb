package impl

import (
	"GoWeb/connector"
	"GoWeb/domain"
	"context"
	"errors"
	"math"
)

type CommentServiceImpl struct {
}

func (csi CommentServiceImpl) CreateComment(comment *domain.Comment) error {
	db := connector.GetDBConn()
	tx := db.Begin()
	if err := tx.Create(comment).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (csi CommentServiceImpl) GetCommentByArticleIdAndPageFromMysql(commentPageCondition *domain.CommentPageCondition) (error, []domain.Comment) {
	var err error
	var comments []domain.Comment
	if err, commentPageCondition.AllPageSize = getAllPageSizeById(commentPageCondition); err != nil {
		return err, comments
	}
	if commentPageCondition.CurrentPage > commentPageCondition.AllPageSize {
		commentPageCondition.CurrentPage = commentPageCondition.AllPageSize
		return errors.New("没有下一页了！"), comments
	}
	if commentPageCondition.CurrentPage < 1 {
		commentPageCondition.CurrentPage = 1
		return errors.New("没有上一页了！"), comments
	}
	db := connector.GetDBConn()
	if err := db.Where("articleid = ?", commentPageCondition.ArticleId).
		Preload("Users").
		Offset((commentPageCondition.CurrentPage - 1) * commentPageCondition.PageSize).
		Limit(commentPageCondition.PageSize).Order("id desc").Find(&comments).Error; err != nil {
		return err, comments
	}
	return nil, comments

}
func (csi CommentServiceImpl) UpdateCommentNumToMysql(article *domain.Article) error {
	db := connector.GetDBConn()
	tx := db.Begin()
	if err := tx.Exec("update article set comments=comments+1 where id = ?", article.Id).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	db.Find(&article)
	return nil
}
func (csi CommentServiceImpl) UpdateCommentNumToElastic(article *domain.Article) error {
	client := connector.GetElasticConn()
	_, err := client.Update().
		Index("articlelibrary").
		Type("article").
		Id(string(article.Id)).
		Doc(map[string]interface{}{"Comment": article.Comments}).
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func getAllPageSizeById(condition *domain.CommentPageCondition) (error, int) {
	count := 0
	db := connector.GetDBConn()
	var comment []domain.Comment
	if err := db.Where("articleid = ?", condition.ArticleId).Find(&comment).Count(&count).Error; err != nil {
		return err, count
	}
	return nil, int(math.Ceil(float64(count) / float64(condition.PageSize)))

}

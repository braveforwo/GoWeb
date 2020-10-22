package impl

import (
	"GoWeb/connector"
	"GoWeb/domain"
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

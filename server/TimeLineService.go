package server

import "GoWeb/domain"

type TimeLineService interface {
	GetAllTimeLineByUser(user *domain.User) (error, []domain.Article)
}

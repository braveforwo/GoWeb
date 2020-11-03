package impl

import (
	"GoWeb/connector"
	"GoWeb/domain"
	"strconv"
	"time"
)

type TimeLineServiceImpl struct {
}

func (tlsi TimeLineServiceImpl) GetAllTimeLineByUser(user *domain.User) (error, []domain.Article) {
	db := connector.GetDBConn()
	var articles []domain.Article
	if err := db.Where("userid = ?", user.Id).Order("id asc").Find(&articles).Error; err != nil {
		return err, nil
	}
	return nil, articles
}

func (tlsi TimeLineServiceImpl) ParseArticlesToTimeLine(articles []domain.Article) map[string]map[string][]domain.TimeLineModel {
	var year map[string]map[string][]domain.TimeLineModel = make(map[string]map[string][]domain.TimeLineModel)
	for _, s := range articles {
		articleTime, _ := time.ParseInLocation("2006-01-02 15:04:05", s.Time, time.Local)
		yearstring := strconv.Itoa(articleTime.Year())
		month := strconv.Itoa(int(articleTime.Month())) + "月"
		var minute string
		if int(articleTime.Minute()) < 10 {
			minute = "0" + strconv.Itoa(int(articleTime.Minute()))
		} else {
			minute = strconv.Itoa(int(articleTime.Minute()))
		}
		monthanday := month + strconv.Itoa(int(articleTime.Day())) + "日 " + strconv.Itoa(int(articleTime.Hour())) + ":" + minute + ":" + strconv.Itoa(int(articleTime.Second()))
		var timeLine = domain.TimeLineModel{}
		timeLine.Time = monthanday
		timeLine.Title = s.Title
		if ok := year[yearstring]; ok != nil {
			if ok1 := year[yearstring][month]; ok1 != nil {
				year[yearstring][month] = append(year[yearstring][month], timeLine)
			} else {
				year[yearstring][month] = append(year[yearstring][month], timeLine)
			}
		} else {
			year[yearstring] = make(map[string][]domain.TimeLineModel)
			year[yearstring][month] = append(year[yearstring][month], timeLine)
		}

	}
	return year
}

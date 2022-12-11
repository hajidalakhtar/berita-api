package service

import (
	"bebasinfo/entity"
	"bebasinfo/helper"
	"bebasinfo/model"
	"bebasinfo/repository"
	"context"
	"gorm.io/gorm"
)

type NewsServiceImpl struct {
	NewsRepository repository.NewsRepository
	DB             *gorm.DB
}

func NewNewsService(newsRepository *repository.NewsRepository, db *gorm.DB) NewsService {
	return &NewsServiceImpl{
		NewsRepository: *newsRepository,
		DB:             db,
	}
}

func (service *NewsServiceImpl) FeedNewsApiOrg(ctx context.Context) interface{} {
	db := service.DB
	var newsApiOrg model.NewsApiOrgResponse
	newsApiOrg = service.NewsRepository.GetNewsApiOrg()

	var newsData []entity.News
	newsData = helper.ToNewsModelFromNewsApiOrgResponse(newsApiOrg)

	result := service.NewsRepository.Insert(ctx, db, newsData)

	response := struct {
		TotalNews int `json:"total_news"`
	}{
		TotalNews: len(result),
	}

	return response

}

func (service *NewsServiceImpl) GetAllNews(ctx context.Context, page int) []entity.News {
	db := service.DB
	result := service.NewsRepository.FindByPage(ctx, db, page)
	return result
}

func (service *NewsServiceImpl) DeleteExpirationNews(ctx context.Context) int {
	db := service.DB
	count := service.NewsRepository.DeleteExpiration(ctx, db)
	return count

}

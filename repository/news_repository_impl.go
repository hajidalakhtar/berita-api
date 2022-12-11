package repository

import (
	"bebasinfo/config"
	"bebasinfo/entity"
	"bebasinfo/exception"
	"bebasinfo/model"
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type NewsRepositoryImpl struct {
	Configuration config.Config
}

func NewNewsRepository(configuration config.Config) NewsRepository {
	return &NewsRepositoryImpl{Configuration: configuration}
}

func (repository *NewsRepositoryImpl) GetNewsApiOrg() model.NewsApiOrgResponse {

	key := repository.Configuration.Get("NEWS_API_KEY")
	url := "https://newsapi.org/v2/top-headlines?country=id&apiKey=" + key
	resp, err := http.Get(url)
	defer resp.Body.Close()
	exception.PanicIfNeeded(err)

	var response model.NewsApiOrgResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)

	exception.PanicIfNeeded(err)
	return response

}

func (repository *NewsRepositoryImpl) Insert(ctx context.Context, db *gorm.DB, news []entity.News) []entity.News {

	for _, item := range news {
		if db.Model(&item).Where("title = ?", item.Title).Updates(&item).RowsAffected == 0 {
			db.Create(&item)
		}
	}
	return news

}

func (repository *NewsRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) []entity.News {
	//TODO implement me
	panic("implement me")
}

func (repository *NewsRepositoryImpl) FindByPage(ctx context.Context, db *gorm.DB, page int) []entity.News {

	perPage := 10
	offset := (page - 1) * perPage

	var newsData []entity.News
	err := db.Offset(offset).Limit(perPage).Find(&newsData).Error
	exception.PanicIfNeeded(err)

	return newsData
}

func (repository *NewsRepositoryImpl) DeleteExpiration(ctx context.Context, db *gorm.DB) int {

	var newsData entity.News

	result := db.Unscoped().Where("created_at < ?", time.Now().Add(-(168 * time.Hour))).Delete(&newsData)
	return int(result.RowsAffected)
}

func (repository *NewsRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, newsId int) entity.News {
	//TODO implement me
	panic("implement me")
}

func (repository *NewsRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, newsId int) {

}

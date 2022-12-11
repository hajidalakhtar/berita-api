package repository

import (
	"bebasinfo/entity"
	"bebasinfo/model"
	"context"
	"gorm.io/gorm"
)

type NewsRepository interface {
	GetNewsApiOrg() model.NewsApiOrgResponse
	Insert(ctx context.Context, db *gorm.DB, news []entity.News) []entity.News
	FindByPage(ctx context.Context, db *gorm.DB, page int) []entity.News
	FindById(ctx context.Context, db *gorm.DB, newsId int) entity.News
	Delete(ctx context.Context, db *gorm.DB, newsId int)
	DeleteExpiration(ctx context.Context, db *gorm.DB) int
}

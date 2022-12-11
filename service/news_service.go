package service

import (
	"bebasinfo/entity"
	"context"
)

type NewsService interface {
	FeedNewsApiOrg(ctx context.Context) interface{}
	GetAllNews(ctx context.Context, page int) []entity.News
	DeleteExpirationNews(ctx context.Context) int
}

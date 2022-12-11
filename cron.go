package main

import (
	"bebasinfo/service"
	"fmt"
	"github.com/robfig/cron"
)

type CronJob interface {
	AddFeed() string
}

type CronImpl struct {
	NewsService service.NewsService
}

func NewCron(newsService *service.NewsService) CronJob {
	return &CronImpl{NewsService: *newsService}
}

func (cronService *CronImpl) AddFeed() string {
	c := cron.New()
	c.AddFunc("*/60 * * * *", func() {
		fmt.Println("asdas")
		//result := service.FeedNewsApiOrg(context.Background())
		//fmt.Println(result)
	})

	// Start cron with one scheduled job
	c.Start()

	return "start"
}

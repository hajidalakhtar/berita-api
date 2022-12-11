package main

import (
	"bebasinfo/config"
	"bebasinfo/controller"
	"bebasinfo/repository"
	"bebasinfo/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func main() {

	configuration := config.New()
	db := config.ConnectDB(configuration)
	newsRepository := repository.NewNewsRepository(configuration)
	newsService := service.NewNewsService(&newsRepository, db)
	//cronjob := config.NewCron(&newsService)
	//row := cronjob.AddFeed()

	//c := cron.New()
	//c.AddFunc("0 * * * *", func() {
	//	result := newsService.FeedNewsApiOrg(context.Background())
	//	fmt.Println(result)
	//
	//})
	//fmt.Println("start")
	//c.Start()

	newsController := controller.NewNewsController(&newsService)
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	newsController.Route(r)
	http.ListenAndServe(":3000", r)

}

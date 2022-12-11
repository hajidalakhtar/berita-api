package controller

import (
	"bebasinfo/helper"
	"bebasinfo/model"
	"bebasinfo/service"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type NewsController struct {
	NewsService service.NewsService
}

func NewNewsController(newsService *service.NewsService) *NewsController {
	return &NewsController{NewsService: *newsService}
}

func (controller *NewsController) Route(app *chi.Mux) {
	app.Route("/news", func(r chi.Router) {
		r.Get("/feed", controller.FeedNewsApiOrg)
		r.Get("/feed/expiration", controller.DeleteExpirationNews)
		r.Get("/", controller.GetNews)

	})
}

func (controller *NewsController) FeedNewsApiOrg(w http.ResponseWriter, r *http.Request) {

	result := controller.NewsService.FeedNewsApiOrg(r.Context())

	webResponse := model.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   result,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller *NewsController) GetNews(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	result := controller.NewsService.GetAllNews(r.Context(), pageInt)

	webResponse := model.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   result,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller *NewsController) DeleteExpirationNews(w http.ResponseWriter, r *http.Request) {

	count := controller.NewsService.DeleteExpirationNews(r.Context())

	webResponse := model.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   count,
	}

	helper.WriteToResponseBody(w, webResponse)

}

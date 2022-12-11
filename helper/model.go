package helper

import (
	"bebasinfo/entity"
	"bebasinfo/model"
)

//
//func ToNewsApiOrgFromResponse(response model.NewsApiOrgResponse) model.NewsApiOrg {
//
//	return model.NewsApiOrg{
//		Source:      nil,
//		Author:      "",
//		Title:       "",
//		Description: "",
//		Url:         "",
//		UrlToImage:  "",
//		PublishedAt: "",
//		Content:     "",
//	}
//
//}

func ToNewsModelFromNewsApiOrg(data model.NewsApiOrg) entity.News {

	return entity.News{
		Title:       data.Title,
		Description: data.Description,
		Content:     data.Content,
		Url:         data.Url,
		ImageUrl:    data.UrlToImage,
		Source:      data.Source.Name,
		PublishedAt: data.PublishedAt,
		ApiSource:   "newsapi.org",
	}
}

func ToNewsModelFromNewsApiOrgResponse(data model.NewsApiOrgResponse) []entity.News {
	var newsData []entity.News
	for _, news := range data.Articles {
		newsData = append(newsData, ToNewsModelFromNewsApiOrg(news))
	}
	return newsData
}

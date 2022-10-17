package helper

import (
	"FawziLinggo/ling-go/2.restful-api-go/model/domain"
	"FawziLinggo/ling-go/2.restful-api-go/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

}

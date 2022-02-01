package service

import (
	"log"

	"github.com/andiahmads/go-api/dto"
	"github.com/andiahmads/go-api/entity"
	"github.com/andiahmads/go-api/repository"
	"github.com/mashingan/smapping"
)

type CategoryService interface {
	Insert(category dto.CreateCategory) entity.Categories
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepo,
	}
}

func (service *categoryService) Insert(category dto.CreateCategory) entity.Categories {
	cat := entity.Categories{}
	err := smapping.FillStruct(&cat, smapping.MapFields(&category))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.categoryRepository.Insert(cat)
	return res
}

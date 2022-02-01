package controller

import (
	"net/http"

	"github.com/andiahmads/go-api/dto"
	"github.com/andiahmads/go-api/helpers"
	service "github.com/andiahmads/go-api/services"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	Insert(context *gin.Context)
}

type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(catService service.CategoryService) CategoryController {
	return &categoryController{
		categoryService: catService,
	}
}

func (c *categoryController) Insert(context *gin.Context) {
	var category dto.CreateCategory
	errDTO := context.ShouldBind(&category)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("failed insert category", errDTO.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	SaveData := c.categoryService.Insert(category)
	res := helpers.BuildResponse(true, "Succes insert category", SaveData)
	context.AbortWithStatusJSON(http.StatusCreated, res)
}

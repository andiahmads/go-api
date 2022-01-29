package helpers

import (
	"strconv"

	"github.com/andiahmads/go-api/dto"
	"github.com/gin-gonic/gin"
)

func GeneratePagination(context *gin.Context) *dto.BookPaginationMeta {
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	sort := context.DefaultQuery("sort", "created_at desc")

	return &dto.BookPaginationMeta{Limit: limit, Page: page, Sort: sort}
}

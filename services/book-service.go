package service

import (
	"fmt"
	"log"

	"github.com/andiahmads/go-api/dto"
	"github.com/andiahmads/go-api/entity"
	"github.com/andiahmads/go-api/helpers"
	"github.com/andiahmads/go-api/repository"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type BookService interface {
	Insert(b dto.BookCreateDTO) entity.Book
	Update(b dto.BookUpdateDTO) entity.Book
	Delete(b entity.Book)
	All() []entity.Book
	FindByID(bookID uint64) entity.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
	AllbookWithPagination(context *gin.Context, pagination *dto.BookPaginationMeta) *dto.ResponsePaginate
	GetAllBookWithCategory() *helpers.CustomeResponse
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (service *bookService) Insert(b dto.BookCreateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bookRepository.InsertBook(book)
	return res
}

func (service *bookService) Update(b dto.BookUpdateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.bookRepository.UpdateBook(book)
	return res
}

func (service *bookService) Delete(b entity.Book) {
	service.bookRepository.DeleteBook(b)
}

func (service *bookService) All() []entity.Book {
	return service.bookRepository.AllBook()
}

func (service *bookService) FindByID(bookID uint64) entity.Book {
	return service.bookRepository.FindBookByID(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	b := service.bookRepository.FindBookByID(bookID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}

func (service *bookService) AllbookWithPagination(context *gin.Context, pagination *dto.BookPaginationMeta) *dto.ResponsePaginate {
	operationResult, totalPages := service.bookRepository.AllbookWithPagination(pagination)

	if operationResult.Error != nil {
		return &dto.ResponsePaginate{Success: false, Message: operationResult.Error.Error()}
	}
	var data = operationResult.Result.(*dto.BookPaginationMeta)

	urlPath := context.Request.URL.Path

	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 0, pagination.Sort)
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort)

	if data.Page > 0 {
		// set previus page query parameter
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page-1, pagination.Sort)
	}
	if data.Page < totalPages {
		//set next page pagination response
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort)
	}

	if data.Page > totalPages {
		//reset previus page
		data.PreviousPage = ""
	}

	return &dto.ResponsePaginate{Success: true, Data: data}

}

func (service *bookService) GetAllBookWithCategory() *helpers.CustomeResponse {
	res := service.bookRepository.GetBookWithInnerJoin()
	return &helpers.CustomeResponse{Success: true, Message: "success", Data: res}

}

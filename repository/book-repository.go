package repository

import (
	"fmt"
	"math"

	"github.com/andiahmads/go-api/dto"
	"github.com/andiahmads/go-api/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book)
	AllBook() []entity.Book
	FindBookByID(bookID uint64) entity.Book
	AllbookWithPagination(pagination *dto.BookPaginationMeta) (RepositoryResult, int)
	GetBookWithInnerJoin() interface{}
}
type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConn,
	}
}
func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) DeleteBook(b entity.Book) {
	db.connection.Delete(&b)
}

func (db *bookConnection) FindBookByID(bookID uint64) entity.Book {
	var book entity.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}

func (db *bookConnection) AllBook() []entity.Book {
	var books []entity.Book
	db.connection.Preload("User").Find(&books)
	return books
}

func (db *bookConnection) AllbookWithPagination(pagination *dto.BookPaginationMeta) (RepositoryResult, int) {
	var books []entity.Book

	totalRows, totalPages, fromRow, toRow := 0, 0, 0, 0
	fmt.Println(totalRows)

	var count int64
	// count = int64(totalRows)

	offset := pagination.Page * pagination.Limit

	errFind := db.connection.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Joins("User").Find(&books).Error

	if errFind != nil {
		// panic(errFind.Error())
		return RepositoryResult{Error: errFind}, totalPages
	}
	pagination.Rows = books

	errCount := db.connection.Model(&entity.Book{}).Count(&count).Error

	if errCount != nil {
		panic(errCount)
	}
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		//set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > totalRows {
		//set to row with total rows
		toRow = totalRows
	}
	pagination.FromRow = fromRow
	pagination.ToRow = toRow
	return RepositoryResult{Result: pagination}, totalPages
}

func (db *bookConnection) GetBookWithInnerJoin() interface{} {

	var result []dto.GetBookWithCategory
	// query := `SELECT books.id,title,description,categories.category FROM books INNER JOIN categories ON categories.id = books.category_id`

	// db.connection.Raw(query).Scan(&result)
	db.connection.Table("books").
		Select("books.id,books.title,books.description,categories.category").
		Joins("INNER JOIN categories ON categories.id = books.category_id").Scan(&result)

	return result

}

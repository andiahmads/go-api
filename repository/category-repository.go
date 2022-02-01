package repository

import (
	"github.com/andiahmads/go-api/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Insert(c entity.Categories) entity.Categories
}

type categoryConnection struct {
	connection *gorm.DB
}

func NewCategoryRepository(dbcon *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: dbcon,
	}
}

func (db *categoryConnection) Insert(c entity.Categories) entity.Categories {
	db.connection.Save(&c)
	return c
}

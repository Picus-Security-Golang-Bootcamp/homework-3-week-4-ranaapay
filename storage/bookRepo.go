package storage

import (
	"PicusHomework3/src/storage/helper"
	"PicusHomework3/src/storage/storageType"
	"fmt"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}



func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&storageType.Book{})
}

func (b *BookRepository) InsertSampleData() error {
	//jsonFile := os.Getenv("BOOK_JSON")
	books, err := helper.ReadJSONForBook("book.json")
	if err != nil {
		return err
	}
	fmt.Println(books[1])
	b.db.Create(&books)
	return nil
}
package cmd

import (
	postgres "PicusHomework3/src/pkg/db"
	"PicusHomework3/src/storage"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

func Execute() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init ", err)
	}
	log.Println("Postgres connected")

	BookRepoFunc(db)
	AuthorRepoFunc(db)
}

//BookRepoFunc BookRepo functions are called. return types are ignored.
func BookRepoFunc(db *gorm.DB) {
	bookRepo := storage.NewBookRepository(db)
	bookRepo.Migrations()
	if err := bookRepo.InsertSampleData(); err != nil {
		fmt.Println(err)
	}
	bookRepo.GetById(12)
	bookRepo.BuyBookByItsId(1, 2)
	bookRepo.GetAllBooks()
	bookRepo.FindBookByItsStockCode("abc")
	bookRepo.FindByName("asdf")
	bookRepo.FindBooksWithLimitOffset(5, 5)
}

//AuthorRepoFunc AuthorRepo functions are called. return types are ignored.
func AuthorRepoFunc(db *gorm.DB) {
	authorRepo := storage.NewAuthorRepository(db)
	authorRepo.Migrations()
	fmt.Println(authorRepo.GetAuthorByIdWithBookInfos(5))
	if err := authorRepo.InsertSampleData(); err != nil {
		fmt.Println(err)
	}
	authorRepo.GetById(10)
	authorRepo.GetAllAuthorsWithBookInfo()
	authorRepo.GetAuthorByIdWithBookInfos(3)
}

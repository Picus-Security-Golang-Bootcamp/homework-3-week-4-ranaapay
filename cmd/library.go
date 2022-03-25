package cmd

import (
	postgres "PicusHomework3/src/pkg/db"
	"PicusHomework3/src/storage"
	"fmt"
	"github.com/joho/godotenv"
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
	bookRepo := storage.NewBookRepository(db)
	bookRepo.Migrations()
	if err  = bookRepo.InsertSampleData(); err != nil {
		fmt.Println(err)
	}

	authorRepo := storage.NewAuthorRepository(db)
	authorRepo.Migrations()
	if err  = bookRepo.InsertSampleData(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(authorRepo.GetAllAuthorsWithBookInfo())
}

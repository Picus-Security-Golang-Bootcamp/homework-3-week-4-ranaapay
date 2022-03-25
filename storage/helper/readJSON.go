package helper

import (
	"PicusHomework3/src/storage/storageType"
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJSONForBook(filename string) ([]storageType.Book,error) {
	var books []storageType.Book
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err = json.Unmarshal(byteValue, &books); err != nil {
		return nil, err
	}
	return books, nil
}

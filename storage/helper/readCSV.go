package helper

import (
	"PicusHomework3/src/storage/storageType"
	"encoding/csv"
	"os"
)

func ReadCSVForAuthor(filename string) ([]storageType.Author, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	var authors []storageType.Author
	for _, line := range records[1:] {
		authors = append(authors, storageType.Author{
			Name:     line[0],
			LastName: line[1],
			Email:    line[2],
			Gender:   line[3],
		})
	}
	return authors, nil
}

package database

import (
	"fmt"
	"os"
)

func CreateFile(zipcode string) error {
	file, err := os.Create("zipcode.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprint(file, zipcode)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile() string {
	zipcode, err := os.ReadFile("zipcode.txt")
	if err != nil {
		fmt.Print(err)
		return ""
	}

	return string(zipcode)
}

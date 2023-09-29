package librarymanager

import (
	"fmt"
	"os"
	"strconv"
)

func SaveLibrary(library Library) {
	file, err := os.Create("library.txt")
	if err != nil {
		fmt.Println(err)
	}
	for _, ch := range library.Data {
		file.WriteString(string(ch.Title) + ", ")
		file.WriteString(string(ch.Author) + ", ")
		file.WriteString(strconv.Itoa(ch.PublicationDate) + ", ")
		file.WriteString(string(ch.Gender) + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	file.Sync()
}

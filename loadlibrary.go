package librarymanager

import (
	"bufio"
	"fmt"
	"os"
)

func LoadLibrary(fileName string) Library {
	var library Library
	var arrayValues [][]string
	var arrayBook []Book

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringBook := (scanner.Text())
		arrayValues = append(arrayValues, Split(stringBook, ","))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for _, ch := range arrayValues {
		Title := ch[0]
		Author := ch[1]
		PublicationDate := Atoi(ch[2])
		Gender := ch[3]
		book := Book{Title, Author, PublicationDate, Gender}
		arrayBook = append(arrayBook, book)
	}
	library = Library{arrayBook}

	return library
}

package librarymanager

import (
	"bufio"
	"fmt"
	"os"
)

func LoadLibrary() {
	// var library Library
	var arrayDictionnaire [][]string
	var arrayValue [][][]string
	// var book Book

	arguments := os.Args[1]

	file, err := os.Open(arguments)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringBook := (scanner.Text())
		arrayDictionnaire = append(arrayDictionnaire, Split(stringBook, ","))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	for _, ch := range arrayDictionnaire {
		for _, carac := range ch {
			arrayValue = append(arrayValue, Split(carac, ":")[1:])
		}
	}
}

func Split(s, sep string) []string {
	var tabStringSeparate []string
	lastIndice := 0
	for indice, val := range s {
		if val == rune(sep[0]) && s[indice+len(sep)-1] == sep[len(sep)-1] {
			if lastIndice == indice {
				lastIndice = indice + len(sep)
			} else {
				tabStringSeparate = append(tabStringSeparate, s[lastIndice:indice])
				lastIndice = indice + len(sep)
			}
		}
		if indice == len(s)-1 {
			tabStringSeparate = append(tabStringSeparate, s[lastIndice:])
		}
	}
	if tabStringSeparate[len(tabStringSeparate)-1] == "" {
		return tabStringSeparate[:len(tabStringSeparate)-1]
	}
	return tabStringSeparate
}

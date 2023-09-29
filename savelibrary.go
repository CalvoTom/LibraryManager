package librarymanager

import (
	"fmt"
	"os"
)

func SaveLibrary(library Library) {
	file, err := os.Create("library.txt")
	if err != nil {
		fmt.Println(err)
	}
	for _, ch := range library.Data {
		if err != nil {
			fmt.Println(err)
		}
	}
	file.Sync()
}

// file.WriteString(string(jsonBytes) + "\n")

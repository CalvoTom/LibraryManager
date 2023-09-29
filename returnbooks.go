package librarymanager

import "sort"

func ReturnBooks(library Library) []Book {
	sort.Slice(library.Data, func(i, j int) bool {
		return library.Data[i].Title < library.Data[j].Title
	})
	return library.Data
}

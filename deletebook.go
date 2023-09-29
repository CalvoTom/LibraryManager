package librarymanager

func DeleteBook(book Book, library Library) Library {
	for index, ch := range library.Data {
		if ch == book {
			library.Data = append(library.Data[:index], library.Data[(index+1):]...)
		}
	}
	return library
}

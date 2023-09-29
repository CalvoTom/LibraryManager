package librarymanager

func AddBook(book Book, library Library) Library {
	library.Data = append(library.Data, book)
	return library
}

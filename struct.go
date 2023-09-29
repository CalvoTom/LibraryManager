package librarymanager

type Book struct {
	Title           string
	Author          string
	PublicationDate int
	Gender          string
}

type Library struct {
	Data []Book
}

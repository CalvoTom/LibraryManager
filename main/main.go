package main

import (
	librarymanager "LibraryManager"
	"fmt"
)

func main() {
	var tab []librarymanager.Book
	var arrayValues [][]string

	book := librarymanager.Book{"Décaméron", "Boccace", 1349, "Romanesque"}
	tab = append(tab, book)
	a := librarymanager.Book{"Contes", "Hans Christian Andersen", 1935, "Romanesque"}
	tab = append(tab, a)
	library := librarymanager.Library{tab}
	fmt.Println(librarymanager.LoadLibrary("library.txt"))

	fmt.Println("Wich features do you want tu use : Add, Delete or Read ?")
	var feature string
	fmt.Scanln(&feature)

	if feature == "Add" || feature == "add" {
		fmt.Println("Please enter the book information like so, Contes,Hans_Christian_Andersen,1935,Romanesque :")
		var values string
		fmt.Scanln(&values)

		arrayValues = append(arrayValues, librarymanager.Split(values, ","))
		Title := librarymanager.Join(librarymanager.Split(arrayValues[0][0], "_"), " ")
		Author := librarymanager.Join(librarymanager.Split(arrayValues[0][1], "_"), " ")
		PublicationDate := librarymanager.Atoi(arrayValues[0][2])
		Gender := librarymanager.Join(librarymanager.Split(arrayValues[0][3], "_"), " ")

		book := librarymanager.Book{Title, Author, PublicationDate, Gender}
		librarymanager.AddBook(book, library)
	}

	if feature == "Delete" || feature == "delete" || feature == "del" {
		fmt.Println("Please enter the book information like so, Contes,Hans_Christian_Andersen,1935,Romanesque :")
		var values string
		fmt.Scanln(&values)

		arrayValues = append(arrayValues, librarymanager.Split(values, ","))
		Title := librarymanager.Join(librarymanager.Split(arrayValues[0][0], "_"), " ")
		Author := librarymanager.Join(librarymanager.Split(arrayValues[0][1], "_"), " ")
		PublicationDate := librarymanager.Atoi(arrayValues[0][2])
		Gender := librarymanager.Join(librarymanager.Split(arrayValues[0][3], "_"), " ")

		book := librarymanager.Book{Title, Author, PublicationDate, Gender}
		librarymanager.DeleteBook(book, library)
	}

	if feature == "Read" || feature == "read" {
		fmt.Println("Please enter the library file name you want to read :")
		var fileName string
		fmt.Scanln(&fileName)

		librarymanager.LoadLibrary(fileName)
	}
}

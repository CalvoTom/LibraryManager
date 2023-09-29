package main

import (
	librarymanager "LibraryManager"
)

func main() {
	var tab []librarymanager.Book
	book := librarymanager.Book{"Décaméron", "Boccace", 1349, "Romanesque"}
	tab = append(tab, book)
	a := librarymanager.Book{"Contes", "Hans Christian Andersen", 1935, "Romanesque"}
	tab = append(tab, a)
	// library := librarymanager.Library{tab}

	librarymanager.LoadLibrary()

}

// var tab []librarymanager.Book
// 	book := librarymanager.Book{"Décaméron", "Boccace", 1349, "Romanesque"}
// 	tab = append(tab, book)
// 	a := librarymanager.Book{"Contes", "Hans Christian Andersen", 1935, "Romanesque"}
// 	tab = append(tab, a)
// 	library := librarymanager.Library{tab}

// 	fmt.Println("Wich features do you want tu use : Add, Delete or Read ?")
// 	var feature string
// 	fmt.Scanln(&feature)
// 	if feature == "Add" || feature == "add" {
// 		fmt.Println("Please enter the book information")
// 		fmt.Println("Title :")
// 		var titleRune []rune
// 		fmt.Scanln(&titleRune)

// 		fmt.Println("Author :")
// 		var authorRune []rune
// 		fmt.Scanln(&authorRune)

// 		fmt.Println("Publication date (years only) :")
// 		var PublicationDate int
// 		fmt.Scanln(&PublicationDate)

// 		fmt.Println("Gender :")
// 		var genderRune []rune
// 		fmt.Scanln(&genderRune)

// 		for i := 0; i < len(titleRune); i++ {
// 			if titleRune[i] == '_' {
// 				titleRune[i] = ' '
// 			}
// 		}
// 		Title := string(titleRune)
// 		for i := 0; i < len(authorRune); i++ {
// 			if authorRune[i] == '_' {
// 				authorRune[i] = ' '
// 			}
// 		}
// 		Author := string(authorRune)
// 		for i := 0; i < len(genderRune); i++ {
// 			if genderRune[i] == '_' {
// 				genderRune[i] = ' '
// 			}
// 		}
// 		Gender := string(genderRune)

// 		book := librarymanager.Book{Title, Author, PublicationDate, Gender}
// 		fmt.Println(librarymanager.AddBook(book, library))

package main

import (
    "fmt"
    "library-management/library"
)

func main() {
    lib := library.NewLibrary()

    // isso é a forma temporaria de acrescentar.
    lib.AddBook(library.Book{Title: "Book 1", Author: "Author 1", Code: "123456"})
    lib.AddBook(library.Book{Title: "Book 2", Author: "Author 2", Code: "789012"})

    // isso vai buscar todos os livros, tenho que ver como acrescentar mais.
    books := lib.GetAllBooks()
    for _, book := range books {
        fmt.Printf("Title: %s, Author: %s, Code: %s\n", book.Title, book.Author, book.Code)
    }
}
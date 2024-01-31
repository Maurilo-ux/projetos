// library.go é o local responsavel por armazerna e adicionar livros.

package library

type Book struct {
    Title  string
    Author string
    Code   string
}

type Library struct {
    books []Book
}


// func basicas para gerencia a biblioteca.

// cria a biblioteca
func NewLibrary() *Library {
    return &Library{}
}

// adicionar os livro (ainda não sei resolver)
func (lib *Library) AddBook(book Book) {
    lib.books = append(lib.books, book)
}

// buscar tudo dentro da struct Libraby
func (lib *Library) GetAllBooks() []Book {
    return lib.books
}

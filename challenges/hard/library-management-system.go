package challenges

// Create a Library struct with fields Name, Books, and Members. Implement methods to add books, lend books, return books, and display book availability.

import "fmt"

type Book struct {
	Title  string
	Author string
}

type Member struct {
	Name  string
	Books []Book
}

type Library struct {
	Name    string
	Books   []Book
	Members []Member
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
	fmt.Println("Book added to the library.")
}

func (l *Library) LendBook(bookTitle string, memberName string) {
	for i, book := range l.Books {
		if book.Title == bookTitle {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			for j, member := range l.Members {
				if member.Name == memberName {
					l.Members[j].Books = append(l.Members[j].Books, book)
					fmt.Printf("Book '%s' lent to %s.\n", book.Title, memberName)
					return
				}
			}
			break
		}
	}
	fmt.Println("Book not available or member not found.")
}

func (l *Library) ReturnBook(bookTitle string, memberName string) {
	for i, member := range l.Members {
		if member.Name == memberName {
			for j, book := range member.Books {
				if book.Title == bookTitle {
					l.Members[i].Books = append(l.Members[i].Books[:j], l.Members[i].Books[j+1:]...)
					l.Books = append(l.Books, book)
					fmt.Printf("Book '%s' returned by %s.\n", book.Title, memberName)
					return
				}
			}
			break
		}
	}
	fmt.Println("Book not found or member not found.")
}

func (l Library) DisplayBookAvailability() {
	fmt.Println("Book availability:")
	for _, book := range l.Books {
		fmt.Println(book.Title)
	}
}

// func main() {
// 	library := Library{
// 		Name: "My Library",
// 	}

// 	library.AddBook(Book{Title: "Book 1", Author: "Author 1"})
// 	library.AddBook(Book{Title: "Book 2", Author: "Author 2"})

// 	library.DisplayBookAvailability()

// 	library.LendBook("Book 1", "John")
// 	library.LendBook("Book 3", "Jane")

// 	library.ReturnBook("Book 1", "John")

// 	library.DisplayBookAvailability()
// }

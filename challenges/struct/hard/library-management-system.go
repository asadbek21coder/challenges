package hard

import "fmt"

// Create a Library struct with fields Name, Books, and
// Members. Implement methods to add books, lend books,
// return books, and display book availability.

// Kutubxona - Nomi, Kitoblar[Kitob]
// Kitob - Nomi, Avtori

// QarzgaOlish()
// QaytaribBerish()
// MavjudliginiKorsatish()

type Library struct {
	Name  string
	Books []Book
}

type Book struct {
	Name     string
	Author   string
	Quantity int
}

func (l *Library) LendBook(book Book) {
	index := l.checkBookAvailability(book)
	if index == -1 {
		fmt.Println("We do not have this book!")
	} else if index == -2 {
		fmt.Println("Currently unavailable: ", book.Name)
	} else {
		l.Books[index].Quantity -= 1

	}

}

func (l Library) checkBookAvailability(book Book) int {
	for i, v := range l.Books {
		if v.Name == book.Name && v.Author == book.Author {
			if v.Quantity > 0 {
				return i
			} else {
				return -2
			}
		}
	}
	return -1
}

func (l *Library) Return(book Book) {
	index := l.checkBookAvailability(book)
	if index == -1 {
		fmt.Println("This book does not belong to us")
	} else {
		l.Books[index].Quantity += 1

	}
}

func (l *Library) AddBook(book Book) {
	index := l.checkBookAvailability(book)
	if index == -1 {
		l.Books = append(l.Books, book)
	} else if index == -2 {
		return
	} else {
		l.Books[index].Quantity += book.Quantity
	}

}

func (l Library) Display() {
	fmt.Println("Kutubxona nomi: ", l.Name)
	fmt.Println("Kitoblar ro'yxati: ")
	for _, v := range l.Books {
		fmt.Println("\tName: ", v.Name)
		fmt.Println("\tAuthor: ", v.Author)
		fmt.Println("\tQuantity: ", v.Quantity)
		fmt.Println("\t------------------------------------------")
	}
}

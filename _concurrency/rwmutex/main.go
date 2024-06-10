package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("program dimulai...")
	// flow:
	// 1. create bookstore
	bookstore := NewBookStore()

	// 2. simualate multiple readers
	for i := 0; i < 5; i++ {
		go func() {
			book := bookstore.getBookDetail(1)
			if book != nil {
				fmt.Printf("reader ke: %d, buku dengan id: %d ditemukan, quantity: %d\n", i, 1, book.Quantity)
			} else {
				fmt.Printf("reader ke %d, buku tidak ditemukan\n", i)
			}
		}()
	}

	// 3. simulate write
	go func() {
		fmt.Println("writer: mengupdate jumlah buku...")
		bookstore.updateBookQuantity(1, -1)
		fmt.Println("writer: berhasil mengupdate jumlah buku")
		time.Sleep(100 * time.Millisecond)
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("program selesai.")
}

type Book struct {
	Title    string
	Quantity int
}

type BookStore struct {
	books   map[int]*Book
	rwmutex sync.RWMutex
}

func NewBookStore() *BookStore {
	return &BookStore{
		books: map[int]*Book{
			1: {
				Title:    "Go Programming",
				Quantity: 10,
			},
			2: {
				Title:    "Concurrency in Go",
				Quantity: 5,
			},
		},
	}
}

func (bs *BookStore) getBookDetail(id int) *Book {
	bs.rwmutex.RLock()
	defer bs.rwmutex.RUnlock()

	book, found := bs.books[id]
	if !found {
		fmt.Printf("book with id: %d not found\n", id)
		return nil
	}

	return &Book{
		Title:    book.Title,
		Quantity: book.Quantity,
	}
}

func (bs *BookStore) updateBookQuantity(id int, change int) {
	bs.rwmutex.Lock()
	defer bs.rwmutex.Unlock()

	book, found := bs.books[id]
	if !found {
		fmt.Printf("book with id: %d not found\n", id)
		return
	}

	book.Quantity += change
}

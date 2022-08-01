//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book string
type Person string

type Lease struct {
	checkOut time.Time
	checkIn  time.Time
}
type Store struct {
	total int
	taken int
}

type Member struct {
	member     Person
	booksTaken map[Book]Lease
}

type Library struct {
	members map[Person]Member
	shelf   map[Book]Store
}

//  - Check out a book
func (l *Library) CheckOut(b Book, p Person) (ok bool, err string) {
	if _, found := l.members[p]; !found {
		return false, "member does not exist"
	}
	if _, found := l.shelf[b]; !found {
		return false, "book does not exist"
	}
	if l.shelf[b].total-l.shelf[b].taken > 0 {
		bookLease := l.members[p].booksTaken[b]
		bookLease.checkOut = time.Now()
		bookLease.checkIn = time.Time{} // set to 0
		bookShelf := l.shelf[b]
		bookShelf.taken++
	} else {
		return false, "no available books"
	}

	return true, ""
}

//  - Check in a book
func (l *Library) AddBook(b Book, count int) {
	book, found := l.shelf[b]
	if !found {
		l.shelf[b] = Store{total: count, taken: 0}
	} else {
		book.total += count
	}
}

func (l *Library) AddMember(n string) (res bool, err string) {
	p := Person(n)
	_, exist := l.members[p]
	if exist {
		res, err = false, string("member already exist")
		return
	}
	l.members[p] = Member{member: p, booksTaken: map[Book]Lease{}}
	res, err = true, ""
	return
}

func (l *Library) Print() {
	fmt.Println("Following books are taken: ")
	booksOut := []string{}
	booksIn := []string{}
	subscribers := []Person{}
	for _, m := range l.members {
		subscribers = append(subscribers, "\n"+m.member)
		fmt.Print(m.member, ":")
		for name, book := range m.booksTaken {
			if book.checkIn.IsZero() {
				booksOut = append(booksOut, string("\t"+name+"checked out:"+Book(book.checkOut.String())))
			} else {
				booksIn = append(booksIn, string("\t"+name+"returned at:"+Book(book.checkIn.String())))
			}
		}
	}

	for b, l := range l.shelf {
		fmt.Println(b, "total of", l.total, "books")
	}

	fmt.Println("subscirbers:", subscribers)
	fmt.Println("returned books", booksIn)
	fmt.Println("takenbooks", booksOut)
}

func main() {
	//  - Add at least 4 books and at least 3 members to the library

	library := Library{
		members: map[Person]Member{},
		shelf:   map[Book]Store{},
	}

	library.AddBook("Blade runner", 4)
	library.AddBook("Snowflake", 2)
	library.AddBook("A little prince", 3)
	library.AddBook("Cracking coding interview", 5)
	library.AddBook("Head first: javascript", 1)
	// library.Print()

	library.AddMember("John")
	library.AddMember("Mike")
	library.AddMember("Andrew")
	library.AddMember("kate")
	// library.Print()

	library.CheckOut(Book("Blade runner"), "John")
	library.CheckOut(Book("Snowflake"), "John")
	library.CheckOut(Book("Head first: javascript"), "kate")
	library.Print()
	//  - Check out a book
	//  - Check in a book
	//  - Print out initial library information, and after each change
}

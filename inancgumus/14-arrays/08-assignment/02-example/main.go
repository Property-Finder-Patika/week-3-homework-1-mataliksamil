// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	prev := [3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	// You can't do this:
	// books = prev // since the books variable needs to be declared first

	var books [4]string // this is the declaration

	for i, b := range prev {
		books[i] += b + " 2nd Ed." // assignment on books does not affect prev
	}

	fmt.Printf("books[3] value = %v -> which means not assigned yet \n \n", books[3])

	books[3] = "Awesomeness" // again ddoes not affect prev

	fmt.Printf("last year:\n%#v\n", prev)
	fmt.Printf("\nthis year:\n%#v\n", books)
}

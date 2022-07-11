// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Moodly
//
//   1. Get username from command-line
//
//   2. Display the usage if the username is missing
//
//   3. Create an array
//     1. Add three positive mood messages
//     2. Add three negative mood messages
//
//   4. Randomly select and print one of the mood messages
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name]
//
//   go run main.go Socrates
//     Socrates feels good 👍
//
//   go run main.go Socrates
//     Socrates feels bad 👎
//
//   go run main.go Socrates
//     Socrates feels sad 😞
//
//   go run main.go Socrates
//     Socrates feels happy 😀
//
//   go run main.go Socrates
//     Socrates feels awesome 😎
//
//   go run main.go Socrates
//     Socrates feels terrible 😩
// ---------------------------------------------------------

func main() {
	var reader string
	mood := [...]string{
		"feels good 👍",
		"feels bad 👎",
		"feels sad 😞",
		"feels happy 😀",
		"feels awesome 😎",
		"feels terrible 😩",
	}

	if len(os.Args) > 1 {
		reader = os.Args[1]
		//fmt.Printf(" your input is :  %v, and the type of the input is : %T \n", reader, reader)
	} else {
		fmt.Printf(" ERROR There is no username specified!!! \n")
		fmt.Printf("go run main.go [your name]")
		return
	}
	//fmt.Printf("mood array : %q \n", mood)
	rand.Seed(time.Now().UnixNano()) // time.Now().UnixNano can give an arbitrary(like) number as the value is in 'one thousand-millionth of a second'
	fmt.Printf("%s %s \n", reader, mood[rand.Intn(1000)%5])
	/* 	for i := 0; i < 50; i++ {
		fmt.Printf("%d - ", rand.Intn(1000)%5)
		if i :=
	} */
}

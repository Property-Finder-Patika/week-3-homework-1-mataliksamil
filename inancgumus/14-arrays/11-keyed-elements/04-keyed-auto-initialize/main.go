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
	rates := [3]float64{
		// index 0 empty
		// index 1 empty
		// just index2 assigned not with the order but with the key
		2: 1.5, // index: 2
	}

	fmt.Println(rates)

	// above array literal equals to this:
	//
	// rates := [3]float64{
	// 	0.,
	// 	0.,
	// 	1.5,
	// }
}

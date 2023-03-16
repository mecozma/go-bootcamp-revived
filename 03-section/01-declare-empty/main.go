// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare empty arrays
//
//  1. Declare and print the following arrays with their types:
//
//    1. The names of your three best friends (names array)
//
//    2. The distances to five different locations (distances array)
//
//    3. A data buffer with five bytes of capacity (data array)
//
//    4. Currency exchange ratios only for a single currency (ratios array)
//
//    5. Up/Down status of four different web servers (alives array)
//
//    6. A byte array that doesn't occupy memory space (zero array)
//
//  2. Print only the types of the same arrays.
//
//  3. Print only the elements of the same arrays.
//
// HINT
//   When printing the elements of an array, you can use the usual Printf verbs.
//
//   For example:
//     When printing a string array, you can use "%q" verb as usual.
//
// EXPECTED OUTPUT
//  names    : [3]string{"", "", ""}
//  distances: [5]int{0, 0, 0, 0, 0}
//  data     : [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
//  ratios   : [1]float64{0}
//  alives   : [4]bool{false, false, false, false}
//  zero     : [0]uint8{}
//
//  names    : [3]string
//  distances: [5]int
//  data     : [5]uint8
//  ratios   : [1]float64
//  alives   : [4]bool
//  zero     : [0]uint8
//
//  names    : ["" "" ""]
//  distances: [0 0 0 0 0]
//  data     : [0 0 0 0 0]
//  ratios   : [0.00]
//  alives   : [false false false false]
//  zero     : []
// ---------------------------------------------------------

func main() {
	names := [3]string{}
	distances := [5]int{}
	data := [5]uint8{}
	ratios := [1]float64{}
	alives := [4]bool{}
	zero := [0]uint8{}

	// print the arrays with their types
	fmt.Println()
	fmt.Println("========== print the arrays with their types ==========")
	fmt.Println()
	fmt.Printf("name: %#v\n", names)
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("data: %#v\n", data)
	fmt.Printf("ratios: %#v\n", ratios)
	fmt.Printf("alives: %#v\n", alives)
	fmt.Printf("zero: %#v\n", zero)

	// print only the types of the arrays
	fmt.Println()
	fmt.Println("========== print only the types of the arrays ==========")
	fmt.Println()
	fmt.Printf("name: %T\n", names)
	fmt.Printf("distances: %T\n", distances)
	fmt.Printf("data: %T\n", data)
	fmt.Printf("ratios: %T\n", ratios)
	fmt.Printf("alives: %T\n", alives)
	fmt.Printf("zero: %T\n", zero)

	// print onlye the elements of the arrays
	fmt.Println()
	fmt.Println("========== print onlye the elements of the arrays ==========")
	fmt.Println()
	fmt.Printf("name: %q\n", names)
	fmt.Printf("distances: %v\n", distances)
	fmt.Printf("data: %v\n", data)
	fmt.Printf("ratios: %v\n", ratios)
	fmt.Printf("alives: %v\n", alives)
	fmt.Printf("zero: %v\n", zero)

}

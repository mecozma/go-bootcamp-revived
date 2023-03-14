package main

import "fmt"

func main() {
	var books [4]string
	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] = books[0] + " 2nd edition"
	fmt.Printf("boooks : %q\n", books)
}

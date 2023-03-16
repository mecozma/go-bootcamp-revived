package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	var books = [4]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		// "Kafka's Revenge 2nd edition",
	}

	fmt.Printf("boooks : %q\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)
	wBooks[0] = books[0]
	// sBooks[0] = books[1]
	// sBooks[1] = books[2]
	// sBooks[2] = books[3]

	for book := range sBooks {
		sBooks[book] = books[book+1]
	}

	fmt.Printf("winter : %#v\n", wBooks)
	fmt.Printf("summer : %#v\n", sBooks)

	var published [len(books)]bool
	published[0] = true
	published[len(books)-1] = true

	fmt.Println("\nPublished Books:")
	for i, ok := range published {
		if ok {
			fmt.Printf("+ %s\n", books[i])
		}
	}

}

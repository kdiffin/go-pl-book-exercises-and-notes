package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// more efficient than the range
	fmt.Println(strings.Join(os.Args[1:], " "))

	// just let the formatter format it for you in the form of a slice
	fmt.Println(os.Args[1:])
}

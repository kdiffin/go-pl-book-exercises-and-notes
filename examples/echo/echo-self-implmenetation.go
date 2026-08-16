package main

import (
	"fmt"
	"os"
)

// my implementation
func main() {
	for i, arg := range os.Args {
		// skip initial arg
		if i == 0 {
			continue
		}

		// handle the case where the final arg prints a pointless space
		if i == (len(os.Args) - 1) {
			fmt.Print(arg)
		} else {
			fmt.Print(arg + " ")
		}
	}
}

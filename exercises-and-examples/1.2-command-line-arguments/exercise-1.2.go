// real simple, just change the index lol, didnt bother doing the sep stuff
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, " ", arg)
	}
}

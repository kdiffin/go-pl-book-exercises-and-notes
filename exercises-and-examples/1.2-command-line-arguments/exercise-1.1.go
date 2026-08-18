// real simple, just change the index lol, didnt bother doing the sep stuff
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}

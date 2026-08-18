package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// inefficient w/ for range loop
	fmt.Println("inefficient program: implemented using range loop")
	executed_at_inefficient := time.Now()
	for _, arg := range os.Args {
		fmt.Print(arg, " ")
	}
	finished_at_inefficient := time.Now()

	fmt.Println("\nefficient program: implemented using strings.join")
	executed_at_efficient := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	finished_at_efficient := time.Now()

	execution_time_inefficient := finished_at_inefficient.Sub(executed_at_inefficient)
	execution_time_efficient := finished_at_efficient.Sub(executed_at_efficient)
	execution_time_difference_seconds := execution_time_efficient - execution_time_inefficient

	fmt.Printf("the efficient program is %s faster than the inefficient program", execution_time_difference_seconds)
}

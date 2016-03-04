// _Switch statements_ express conditionals across many
// branches.

package main

import "fmt"
import "time"

func main() {
	fmt.Println("start " + time.Now().String())
	t := time.Now().Weekday()
	switch t {
	case time.Monday:
		fmt.Println(" This is Monday")
		break
	case time.Saturday:
		fmt.Println(" This is Saturday")
		break
	default:
		fmt.Println("This is other day")
		break
	}
}

// todo: type switches

package tricks

import (
	"fmt"
)

func format() {
    fmt.Println()
    fmt.Println("--------------------------------------------------------------------------------")
    fmt.Println()
}

func Format(title string) {

    first := "----------------------------------------"
    fmt.Println()
    fmt.Printf("%s%s\n", first,title)
    fmt.Println()
}
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

func FormatTwo(title string, value any) {

    first := "----------------------------------------"
    fmt.Println()
    fmt.Printf("%s%s %d\n", first, title, value)
    fmt.Println()
}
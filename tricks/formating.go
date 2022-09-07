package tricks

import (
	"fmt"
    "strings"
)

func init() {
    fmt.Println("file formating.go thuoc package /tricks")
}

func format() {
    fmt.Println()
    fmt.Println("--------------------------------------------------------------------------------")
    fmt.Println()
}

func Format(title string) {

    first := fmt.Sprint(strings.Repeat("-", 25)) // "*" * 25
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
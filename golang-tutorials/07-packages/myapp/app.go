package main

import (
	"fmt"
	"rootpath/numbers"
	"rootpath/strings"
	"rootpath/strings/greeting"
	str "strings" // Package Alias
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("callicoder"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}

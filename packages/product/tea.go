package product

import (
	"fmt"
)

func init() {
	fmt.Println("tao la init tea.go")
}

func Tea(name string) {
	fmt.Printf("func Tea() thuoc product/ duoc called from %20s/\n", name)
}
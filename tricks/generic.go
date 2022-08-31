package tricks

import (
	// "fmt"
)

// STOP chưa nên xài
// func init() {
//     fmt.Println("file generic.go thuoc package /tricks")
// }

type SuperNumber interface {
	int | int8 | int16 | int32 | int64 | float64
}


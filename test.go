// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

type Human struct{} 
type myInt int

func main() {
	var name any = "chi thong"
	var age int = 30
	var obj Human = struct{}{}
	obj2 := struct{}{}

	fmt.Println(name, age, obj, obj2)
	fmt.Println()

	val, ok := name.(string)
	fmt.Println(val, ok)
	fmt.Println(reflect.TypeOf(obj))
	fmt.Println()
	fmt.Println()
}

package main
 
import (
	"fmt"
	"errors"

	"phong/tricks"
 )


// Tạo struct
type MyError struct{}
 
// implement method Error() of type ""
func (myErr *MyError) Error() string {
	return "Error() run gogo"
}


func main() {
	tricks.Format("struct implement method Error() of type error")
	// tạo error
	myStructErr := &MyError{}
	// in ra thông báo lỗi; run Error()
	fmt.Println(myStructErr)

	tricks.Format("errors.New(s string) error")
	myErr := errors.New("tao la func New(s string) error , thuoc package errors ")
	fmt.Println(myErr)

	fmt.Printf("Type of myErr is %T \n", myErr)
	fmt.Printf("Value of myErr is %#v \n", myErr)

}

const NOTE string = `


bên in package: errors 
type errorString struct {
    s string
}

// fmt.Println(obj); se run Erro()
func (e *errorString) Error() string {
    return e.s
}


func New(text string) error {
    return &errorString{text}
}

`
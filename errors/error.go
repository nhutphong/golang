package main
 
import (
	"fmt"
	"errors"

	"phong/tricks"
 )


// Tạo struct
type MyError struct{Name string}
 
// implement method Error() of type ""
func (myErr *MyError) Error() string {
	return "Error() run gogo"
}


func main() {
	tricks.Format("struct implement method Error() of type error")
	// tạo error
	myStructErr := &MyError{Name: "dung"}
	// in ra thông báo lỗi; run Error()
	fmt.Println(*myStructErr)
	fmt.Println(myStructErr)
	fmt.Println(myStructErr.Error())

	tricks.Format("errors.New(s string) error")
	myErr := errors.New("tao la func New(s string) error , thuoc package errors ")
	fmt.Println(myErr)

	fmt.Printf("Type of myErr is %T \n", myErr)
	fmt.Printf("Value of myErr is %#v \n", myErr)


}

const NOTE string = `


// bên in package: errors 
type errorString struct {
    s string
}

// fmt.Println(obj); se run Error()
func (e *errorString) Error() string {
    return e.s
}


func New(text string) error {
    return &errorString{text}
}

`
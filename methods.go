package main

import (
    "fmt"
)

// Định nghĩa struct Vertex với 2 thuộc tính X và Y
type Human struct {
    name string
    age int
}

// Tạo phương thức Abs() cho struct Vertex (receiver argument)
func (self Human) get_name() string {
    return self.name
}

func (self Human) get_age() int {
    return self.age
}

func main() {
    // Khởi tạo struct
    human := Human{name: "chithong", age:28}
    fmt.Println("name: ", human.get_name())
    fmt.Println("age: ", human.get_age())
    
}
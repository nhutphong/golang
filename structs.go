package main

import "fmt"

// Định nghĩa một struct với từ khóa type
type Student struct {
    name string
    age int
}

func main() {
    // Khởi tạo biến s1 có giá trị là struct Student
    student1 := Student{"Robin", 30}   // {"Robin", 30}

    // Khởi tạo biến s2 có giá trị là struct Student với 1 field là name
    // Field còn lại sẽ có giá trị mặc định (zero value)
    student2 := Student{name: "Robin"}   // {"Robin", 0}

    // Khởi tạo biến s3 có giá trị là struct Student và không khai báo giá trị cho trường nào
    student3 := Student{}   // {"", 0}

    // Thay đổi giá trị field trong struct
    student3.name = "Robert"
    student3.age = 25

    fmt.Println(student3)   // s3 = {"Robert", 25}
}
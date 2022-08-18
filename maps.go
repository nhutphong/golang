package main

import (
	"fmt"
)

func main() {
	/* 
		var demoMap map[string]int //cach1 
		khong the add del, change key value in map
		syntax thuong dung la argument cho func

		// use make()
		dict := make(map[string]int) //cach2 co the add dell change key value in map
	*/

	dict := make(map[string]int) //cach2

	dict["old"] = 27
	value, ok := dict["old"] // ok is True if dict['old'] ton tai
	fmt.Println("value: ", value, "ok: ", ok)

	delete(dict, "old")
	fmt.Println(dict["old"])

	// Tiếp theo ví dụ trên
	v, k := dict["old"]
	fmt.Println("Giá trị của phần tử là: ", v)
	fmt.Println("Kiểm tra phần tử có tồn tại hay không: ", k)

	fmt.Println()

	var personMobileNo = map[string]string{
		"John":  "+33-8273658526",
		"Steve": "+1-8579822345",
		"David": "+44-9462834443",
	}

	mobileNo := personMobileNo["Steve"]
	fmt.Println("Steve's Mobile No : ", mobileNo)

	// If a key doesn't exist in the map, we get the zero value of the value type
	mobileNo = personMobileNo["Jack"]
	fmt.Printf("\nJack's PhoneNo : %v, type = %T", mobileNo, mobileNo)

	fmt.Println()

	var employees = map[int]string{
		1001: "Rajeev",
		1002: "Sachin",
		1003: "James",
	}

	printEmployee(employees, 1001)
	printEmployee(employees, 1010)

	if isEmployeeExists(employees, 1002) {
		fmt.Println("EmployeeId 1002 found")
	}

	fmt.Println()

	var fileExtensions = map[string]string{
		"Python": ".py",
		"C++":    ".cpp",
		"Java":   ".java",
		"Golang": ".go",
		"Kotlin": ".kt",
	}

	fmt.Println(fileExtensions)

	delete(fileExtensions, "Kotlin")

	// delete function doesn't do anything if the key doesn't exist
	delete(fileExtensions, "Javascript")

	fmt.Println(fileExtensions)


	fmt.Println()
	/*
		chi co array la value type = tham tri
		slice, map[typekey]typevalue reference type = tham chieu giong python
	*/
	var m1 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	var m2 = m1
	fmt.Println("m1 = ", m1)
	fmt.Println("m2 = ", m2)

	m2["ten"] = 10
	fmt.Println("\nm1 = ", m1)
	fmt.Println("m2 = ", m2)

	fmt.Println()
	
	var personAge = map[string]int{
		"Rajeev": 25,
		"James":  32,
		"Sarah":  29,
	}

	for key, value := range personAge {
		fmt.Println(key, value)
	}
}


func printEmployee(employees map[int]string, employeeId int) {
	if name, ok := employees[employeeId]; ok {
		fmt.Printf("name = %s, ok = %v\n", name, ok)
	} else {
		fmt.Printf("EmployeeId %d not found\n", employeeId)
	}
}

func isEmployeeExists(employees map[int]string, employeeId int) bool {
	_, ok := employees[employeeId]
	return ok
}

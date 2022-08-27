package main

import (
	"fmt"
	"phong/tricks"
)


func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("I am a string and my value is %s\n", i.(string))
    case int:
        fmt.Printf("I am an int and my value is %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }
}

func defer_test(end int) {
	fmt.Println("counting")

    for i := 0; i < end; i++ {
        defer fmt.Println("defer ", i)
    }
    /*
    	defer se run khi func chứa nó kết thúc 
    	i = 0 vao truoc, se run sau cùng 
    */

    fmt.Println("done")
}

func main() {
	// If Statement
	var x = 25
	if x%5 == 0 {
		fmt.Printf("%d is a multiple of 5\n", x)
	}

	tricks.Format()
	// Parentheses are Optional
	var y = -1
	if y < 0 {
		fmt.Printf("%d is negative\n", y)
	}

	tricks.Format()
	// If with a condition consisting of short circuit operators
	var age = 21
	if age >= 17 && age <= 30 {
		fmt.Println("My Age is between 17 and 30")
	}

	tricks.Format()
	// If with a short statement
	if n := 10; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	}

	tricks.Format()
	var old = 18
	if old >= 18 {
		fmt.Println("You're eligible to vote!")
	} else {
		fmt.Println("You're not eligible to vote!")
	}

	tricks.Format()
	var BMI = 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight")
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal")
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}

	tricks.Format()
	var dayOfWeek = 6
	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		{
			fmt.Println("Saturday")
			fmt.Println("Weekend. Yaay!")
		}
	case 7:
		{
			fmt.Println("Sunday")
			fmt.Println("Weekend. Yaay!")
		}
	default:
		fmt.Println("Invalid day")
	}

	tricks.Format()
	switch day_of_week := 5; day_of_week {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid Day")
	}

	tricks.Format()
	var AMD = 21.0
	switch {
	case AMD < 18.5:
		fmt.Println("You're underweight")
	case AMD >= 18.5 && AMD < 25.0:
		fmt.Println("Your weight is normal")
	case AMD >= 25.0 && AMD < 30.0:
		fmt.Println("You're overweight")
	default:
		fmt.Println("You're obese")
	}

	tricks.Format()
	for one := 0; one < 10; one++ {
		fmt.Printf("%d ", one)
	}

	tricks.Format()
	two := 2
	for ; two <= 10; two += 2 {
		fmt.Printf("%d ", two)
	}

	tricks.Format()
	three := 2
	for three <= 20 {
		// giong nhu while
		fmt.Printf("%d ", three)
		three *= 2
	}


	tricks.Format()
	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Printf("First positive number divisible by both 3 and 5 is %d\n", num)
			break
		}
	}

	tricks.Format()
	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Printf("%d ", num)
	}

	tricks.Format()

	findType("Naveen")
    findType(77)
    findType(89.98)

    tricks.Format()

    defer_test(10)

    tricks.Format()

    name:="chi thong"
	for index, char := range name {
		/*
			phai dung %c để print char
		*/
		fmt.Printf("index %d char %c\n", index, char)
	}

	tricks.Format()
}

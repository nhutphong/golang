package main

import "fmt"


func format() {
    fmt.Println()
    fmt.Println("--------------------------------------------------------------------------------")
    fmt.Println()
}

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

func main() {
	// If Statement
	var x = 25
	if x%5 == 0 {
		fmt.Printf("%d is a multiple of 5\n", x)
	}

	format()
	// Parentheses are Optional
	var y = -1
	if y < 0 {
		fmt.Printf("%d is negative\n", y)
	}

	format()
	// If with a condition consisting of short circuit operators
	var age = 21
	if age >= 17 && age <= 30 {
		fmt.Println("My Age is between 17 and 30")
	}

	format()
	// If with a short statement
	if n := 10; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	}

	format()
	var old = 18
	if old >= 18 {
		fmt.Println("You're eligible to vote!")
	} else {
		fmt.Println("You're not eligible to vote!")
	}

	format()
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

	format()
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

	format()
	switch day_of_week := 5; day_of_week {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid Day")
	}

	format()
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

	format()
	for one := 0; one < 10; one++ {
		fmt.Printf("%d ", one)
	}

	format()
	two := 2
	for ; two <= 10; two += 2 {
		fmt.Printf("%d ", two)
	}

	format()
	three := 2
	for three <= 20 {
		fmt.Printf("%d ", three)
		three *= 2
	}

	// Infinite Loop
	// for {
	// }


	format()
	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Printf("First positive number divisible by both 3 and 5 is %d\n", num)
			break
		}
	}

	format()
	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Printf("%d ", num)
	}

	format()

	findType("Naveen")
    findType(77)
    findType(89.98)

    format()

}

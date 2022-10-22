package main

import (
	"fmt"
	"phong/tricks"
)



func defer_test(end int) {
	fmt.Println("counting")

    for i := 0; i < end; i++ {
        defer fmt.Println("defer ", i)
    }
    /*
    	chi run sau return of func chua no
    	i = 0 vao truoc, se run sau cùng 
    */

    fmt.Println("done")
}

func main() {
	tricks.Format("if")
	var x = 25
	if x%5 == 0 {
		fmt.Printf("%d is a multiple of 5\n", x)
	}

	
	var y = -1
	if y < 0 {
		fmt.Printf("%d is negative\n", y)
	}

	var age = 21
	if age >= 17 && age <= 30 {
		fmt.Println("My Age is between 17 and 30")
	}

	tricks.Format("if n := 10; n%2 == 0 {}")
	// If with a short statement
	if n := 10; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	}

	tricks.Format("if {} else {}")
	var old = 18
	if old >= 18 {
		fmt.Println("You're eligible to vote!")
	} else {
		fmt.Println("You're not eligible to vote!")
	}

	tricks.Format("if {} else if {} else {}")
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

	// switch-case golang, khong can keyword break in moi case:
	tricks.Format("switch expression { case : } ")
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


	tricks.Format("switch case")
	switch day_of_week := 5; day_of_week {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid Day")
	}


	tricks.Format("switch case")
	var AMD = 21.0
	switch {
	case AMD < 18.5:
		fmt.Println("You're underweight")
	case AMD >= 18.5 && AMD < 25.0:
		fmt.Println("Your weight is normal")
		//fallthrough //run them 1 case ben duoi khong qua tam dieu kien
	case AMD >= 25.0 && AMD < 30.0:
		fmt.Println("You're overweight")
	default:
		fmt.Println("You're obese")
	}

	tricks.Format("for one := 0; one < 10; one++")
	for one := 0; one < 10; one++ {
		if one == 4 {
			continue
		} else if one >= 8 {
			break
		}

		fmt.Printf("%d ", one)
	}

	tricks.Format("for ; two <= 10; two += 2 {}")
	two := 2
	for ; two <= 10; two += 2 {
		fmt.Printf("%d ", two)
	}

	tricks.Format("for three <= 20 {}")
	three := 2
	for three <= 20 {
		// giong nhu while
		fmt.Printf("%d ", three)
		three *= 2
	}


	tricks.Format("*")
	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Printf("First positive number divisible by both 3 and 5 is %d\n", num)
			break
		}
	}

	tricks.Format("*")
	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Printf("%d ", num)
	}

    tricks.Format("defer")

    defer_test(10)

    tricks.Format("for index, char := range name {}")

    name:="chi thong"
	for index, char := range name {
		/*
			phai dung %c để print char
		*/
		fmt.Printf("index %d char %c\n", index, char)
	}


	tricks.Format("check type use syntax: switch v := value.(type)")
	// chi dung duoc cho type phai = interface{}
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		// Here, v is declared once, but it denotes
		// different variables in different branches.
		switch v := x.(type) {
		case []int: // a type literal
			// The type of v is "[]int" in this branch.
			fmt.Println("int slice:", v)
		case string: // one type name
			// The type of v is "string" in this branch.
			fmt.Println("string:", v)
		case int, float64, int32: // multiple type names
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("number:", v)
		case nil:
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println(v)
		default:
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("others:", v)
		}
		// Note, each variable denoted by v in the
		// last three branches is a copy of x.
	}

	fmt.Println("***********************")

	for _, x := range values {
		if v, ok := x.([]int); ok {
			fmt.Println("int slice:", v)
		} else if v, ok := x.(string); ok {
			fmt.Println("string:", v)
		} else if x == nil {
			v := x
			fmt.Println(v)
		} else {
			_, isInt := x.(int)
			_, isFloat64 := x.(float64)
			_, isInt32 := x.(int32)
			if isInt || isFloat64 || isInt32 {
				v := x
				fmt.Println("number:", v)
			} else {
				v := x
				fmt.Println("others:", v)
			}
		}
	}


}

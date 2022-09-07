package main

import (
	"fmt"
	"phong/tricks"
)

func main() {
	/*
		array := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
		slice = array[2:7] // REF-TYPE

		arr1 := [5]int{1:10,2:40} // [0, 10, 40, 0, 0]
		// index1 = 10, index2 = 40
		
	*/

	w := []byte("ba")

	z1 := append(w, 'd')
	z2 := append(w, 'g')

	fmt.Println(string(z1)) // bag
	fmt.Println(string(z2)) // bag


	tricks.Format("slice REF-TYPE to array dung: array[:]")
	var four = []int{3, 5, 7, 9, 11, 13, 17} // Creates an array, and returns a slice reference to the array

	// Short hand declaration
	t := []int{2, 4, 8, 16, 32, 64}

	fmt.Println("s = ", four)
	fmt.Println("t = ", t)

	tricks.Format("*")
	a := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	// syntax slice gan la tham chieu= REF_TYPE, tuc slice1 or slice2 thay doi thi a cung thay doi
	slice1 := a[1:]
	slice2 := a[3:]

	// syntax gan binh thuong, la tham tri = VALUE_TYPE, slice1 or slice2 thay doi, khong lam a thay doi
	// slice1 = a 
	// slice2 = a


	fmt.Println("------- Before Modifications -------")
	fmt.Println("a  = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	slice1[0] = "TUE"
	slice1[1] = "WED"
	slice1[2] = "THU"

	slice2[1] = "FRIDAY"

	fmt.Println("\n-------- After Modifications --------")
	fmt.Println("a  = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)


	/*
		The length of the slice is the number of elements in the slice.
		slice=cap la do dai tu start_index den end cua array no tham tri=value
	*/
	a1 := []int{10, 20, 30, 40, 50, 60, 80, 100, 120, 130, 150}
	s1 := a1[1:6]    // {20, 30, 40, 50, 60}
	last := s1[2:]   // {40,50, 60}

	// cap(slice) tinh = slice[0] xuat hien dau tien trong array, dem tu do den het array
	//cap(s1) = 10(tuc la tinh tu index=1 cua a1 toi het a1) REF_TYPE to a1
	//cap(last) = 8(tuc la tinh tu index=2 cua s1 toi het a1) REF_TYPE to a1


	tricks.Format("*")
	space := ""
	fmt.Printf("a1 = %v, len = %d, cap = %d\n", a1, len(a1), cap(a1))
	fmt.Printf("s1 = %v, %20s, len = %d, cap = %d\n", s1, space, len(s1), cap(s1))
	fmt.Printf("last = %v, %24s, len = %d, cap = %d\n", last, space, len(last), cap(last))


	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original Slice")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[1:5]
	fmt.Println("\nAfter slicing from index s[1:5]")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[:8]
	fmt.Println("\nAfter extending the length s[:8]")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[2:]
	fmt.Println("\nAfter dropping the first two elements s[2:]")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))


	tricks.Format("make([]int, lenght, cap)")
	// make([]int, lenght, cap)
	slice := make([]int, 5, 10)
	fmt.Printf("slice = %v, len = %d, cap = %d\n", slice, len(slice), cap(slice))


	/*
		The capacity parameter in the make() function is optional. When omitted, it defaults to the specified length
	*/
	// Creates an array of size 5, and returns a slice reference to it
	var s3 = make([]int, 5)
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s3, len(s3), cap(s3))


	tricks.Format("nil")
	// The zero value of a slice is nil. A nil slice doesn’t have any underlying array, and has a length and capacity of 0
	var six []int // = nil = None python
	fmt.Printf("six = %v, len = %d, cap = %d\n", six, len(six), cap(six))

	if six == nil {
		fmt.Println("six is nil")
	}

	
	tricks.Format("numElementsCopied := copy(dest, src)")
	/*
		The copy() function copies elements from one slice to another
		func copy(dst, src []T) int
	*/
	src := []string{"Sublime", "VSCode", "IntelliJ", "Eclipse"}
	dest := make([]string, 2)

	numElementsCopied := copy(dest, src)

	fmt.Println("src = ", src)
	fmt.Println("dest = ", dest)
	fmt.Println("Number of elements copied from src to dest = ", numElementsCopied)


	tricks.Format(`new_slice := append(slice3, "Python", "Ruby", "Go")`)
	slice3 := []string{"C", "C++", "Java"}
	// slice5 va slice3 rieng biet khong REF_TYPE nhau
	slice5 := append(slice3, "Python", "Ruby", "Go")

	fmt.Printf("slice3 = %v, \t\t\tlen = %d, cap = %d\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice5 = %v, \tlen = %d, cap = %d\n", slice5, len(slice5), cap(slice5))

	slice3[0] = "HoHo"
	fmt.Println("\nslice3 = ", slice3)
	fmt.Println("slice5 = ", slice5)


	tricks.Format("add item to last || first for slice")
	sliceInt := []int{4,5,6,7}
	item := 10
	fmt.Println("add item to last of slice", append(sliceInt, item))
	fmt.Println("add item to first of slice", append([]int{item}, sliceInt...))
	fmt.Println("test", append(sliceInt[:2], 8,9,10))
	fmt.Println("test", append(sliceInt[:2], sliceInt[3:]...))


	// syntax 3dot: "strings"... append vao []byte("old ")
	tricks.Format(`sliceByte := append([]byte("old "), "falcon"...) `)
	bytes := append([]byte("old "), "falcon"...)
    fmt.Println(bytes)
    fmt.Println(string(bytes))

	// slice1 := make([]string, 3, 10)
	// count := copy(slice1, []string{"C", "C++", "Java"})
	// slice2 := append(slice1, "Python", "Ruby", "Go")

	// slice1 := []string{"Jack", "John", "Peter"}
	// slice2 := []string{"Bill", "Mark", "Steve"}
	// unpackae array... 3dot
	// slice3 := append(slice1, slice2...)


	tricks.Format("array 2 chieu")
	two_chieu := [][]string{
		{"India", "China"},
		{"USA", "Canada"},
		{"Switzerland", "Germany"},
	}

	fmt.Println("Slice two_chieu = ", two_chieu)
	fmt.Println("length = ", len(two_chieu))
	fmt.Println("capacity = ", cap(two_chieu))


	tricks.Format("for index, value := range array {}")
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	for index, number := range primeNumbers {
		fmt.Printf("PrimeNumber(%d) = %d\n", index+1, number)
	}


	tricks.Format("*")
	numbers := []float64{3.5, 7.4, 9.2, 5.4}
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Total Sum = %.2f\n", sum)
}


func insert(orig []int, index int, value int) ([]int, error) {
    if index < 0 {
        return nil, errors.New("Index cannot be less than 0")
    }

    if index >= len(orig) {
        return append(orig, value), nil
    }

    //  orig[1 2 3 4 5] ; index=2; [1 2 3] + [3 4 5] = [1 2 3 3 4 5]
    orig = append(orig[:index+1], orig[index:]...)
    // orid[2] = 10 ; // [1 2 10 3 4 5]
    orig[index] = value

    return orig, nil
}

func delete(orig []int, index int) ([]int, error) {
    if index < 0 || index >= len(orig) {
        return nil, errors.New("Index cannot be less than 0")
    }

    orig = append(orig[:index], orig[index+1:]...)

    return orig, nil
}
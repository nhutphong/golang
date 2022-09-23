package main

import (
	"fmt"
	// "reflect"
	// "sort"
	"time"
)


/*
	khi dung for cho string=chars char đã được convert ngầm về rune=int32
	vì thế phải dùng char = string(char) về string
*/
func reverseString(chars string) (result string) {

	for _ , rune := range chars {
		result = string(rune) + result
	}
	return
}

func loop(chars string) {
	for _, rune := range chars {
		fmt.Println("rune=int32", rune)
		time.Sleep(time.Millisecond * 100)
	}
}

func loopCharsUpper() {
	for char := 'A'; char <= 'Z'; char++{
		fmt.Printf("char '%s' rune %d\n", string(char), char)
		time.Sleep(time.Millisecond * 100)
	}
}

func loopCharsLower() {
	for char := 'a'; char <= 'z'; char++{
		fmt.Printf("char '%s' rune %d\n", string(char), char)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	name := "nguyen dung thong"
	fmt.Println(name)
	nameRev := reverseString(name)
	fmt.Println(nameRev)

	// loop(name)
	fmt.Println(string([]rune{'n', 'g', 'u', 'y', 'e', 'n'}))
	fmt.Println(byte('n'))

	loopCharsUpper()
	loopCharsLower()

}

const NOTE = `

func Find(n int, cmp func(int) int) (i int, found bool)
func SliceIsSorted(x any, less func(i, j int) bool) bool


func Strings(x []string)
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)

func Ints(x []int)
func Float64s(x []float64)
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(s)
	fmt.Println(s) //[-3.8 -1.3 0.7 2.6 5.2]

	s = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // unsorted
	sort.Float64s(s)
	fmt.Println(s) //[NaN -Inf 0 +Inf]


func StringsAreSorted(x []string) bool // string co tang dan hay khong
func IntsAreSorted(x []int) bool // int co tang dan hay khong
func Float64sAreSorted(x []float64) bool // // float64 co tang dan hay khong
	s := []float64{0.7, 1.3, 2.6, 3.8, 5.2} // true = tăng  sorted ascending
	fmt.Println(sort.Float64sAreSorted(s))

	s = []float64{5.2, 3.8, 2.6, 1.3, 0.7} // false = giảm sorted descending
	fmt.Println(sort.Float64sAreSorted(s))

	s = []float64{5.2, 1.3, 0.7, 3.8, 2.6} // false = unsorted
	fmt.Println(sort.Float64sAreSorted(s))



func Search(n int, f func(int) bool) int

func SearchInts(a []int, x int) int
func SearchFloat64s(a []float64, x float64) int
	slice := []float64{1.0, 2.0, 3.3, 4.6, 6.1, 7.2, 8.0}
	item := 2.0
	index := sort.SearchFloat64s(slice, x) //index=1
	fmt.Printf("found %g at index %d in %v\n", item, index, slice)

	item = 0.5
	index = sort.SearchFloat64s(slice, item)
	fmt.Printf("%g not found, can be inserted at index %d in %v\n", item, index, slice)

func SearchStrings(a []string, x string) int


func Slice(x any, less func(i, j int) bool)
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)

	By name: [{Alice 55} {Bob 75} {Gopher 7} {Vera 24}]
	By age: [{Gopher 7} {Vera 24} {Alice 55} {Bob 75}]
func SliceStable(x any, less func(i, j int) bool) // các item=nhau không bị đảo vị trí 

*/

/*
func IsSorted(data Interface) bool
func Sort(data Interface)
func Stable(data Interface)

	sliceInt := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(s))) //[6 5 4 3 2 1]
	fmt.Println(s)


	sort.Sort(sort.Reverse(sort.Float64Slice(s)))
	sort.Sort(sort.Reverse(sort.StringSlice(s)))

`
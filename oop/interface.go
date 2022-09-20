package main

import (  
    "fmt"
    // "log"
    "strings"
    "phong/tricks"
)
/*
	interface nơi chứa methods
	khi các struct implement= tức là code các methods trùng tên với methods co trong interface
	thì khi đó các struct đều có type la interface đó 

	ứng dụng: interface type Animal là type đại diện cho các structs đã implement nó, hay dùng làm
    type arguments cho func, hay khi declared varName
    khi Human, Dog đều implement interface Animal
	có human1 human2, dog1, dog2 đều có type Animal 
	arr = []Animal{human1, human2, dog1, dog2} //ok

    var meo Animal
    meo = Cat("pass args vao") //OK
    var cho Anamal
    cho = Dog("pass args vao") //OK


    func action(meo, cho Animal) {} //

    action(meo, cho) //OK


*/

type SalaryCalculator interface {  
    CalculateSalary() int
}

type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId  int
    basicpay int
}

//tiền lương của nhân viên permanent bằng tổng của basic pay và pf
// implement SalaryCalculator interface  (ngầm, tức tụ hiểu)
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

//tiền lương của nhân viên contract chỉ là basic pay
func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

/*
tổng chi phí được tính bằng cách duyệt qua từng phần tử của slice SalaryCalculator
và tính tổng mức lương của từng nhân viên
*/
func totalExpense(list []SalaryCalculator) {  
    expense := 0
    for _, v := range list {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}


type Describer interface {  
    Describe()
}
type Person2 struct {  
    name string
    age  int
}

func (p Person2) Describe() {  
    fmt.Printf("%s is %d years old\n", p.name, p.age)
}

//go1.18, (i any) == (i interface{}) 
func findType(i interface{}) {  
    switch v := i.(type) {
    case Describer:
        v.Describe()
    default:
        fmt.Printf("unknown type\n")
    }
}


func main() {  

	tricks.Format("interface")
    pemp1 := Permanent{1, 5000, 20}
    pemp2 := Permanent{2, 6000, 30}
    cemp1 := Contract{3, 3000}
    // 
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)


    tricks.Format("varName.(type)")
   	findType("Naveen")

    p := Person2{
        name: "Naveen R",
        age:  25,
    }
    findType(p)

    var repeat = strings.Repeat("-", 15)
    fmt.Println(repeat, "implement interface; use struct access", repeat)

    var x Animal
    x = Animal{name: "Chi Thong"}
 
    fmt.Println("x", x) // Output: Hello
    fmt.Println("&x", &x)
    fmt.Println("x.name", x.name) // Output: sam


     var x2 Animal
    x2 = Animal{name: "Chi Thong"}
 
    fmt.Println("x2", x2) // Output: Hello
    fmt.Println("&x2", &x2)
    fmt.Println("x2.name", x2.name) // Output: sam




    fmt.Println(repeat, "implement interface; use interface access co receiver pointer", repeat)
    var str Stringer 
    str = new(Animal) //use type Stringer access
    fmt.Println(str)
    fmt.Println(str.(*Animal).name)
    str.(*Animal).GetName()

    animal, ok := str.(*Animal)
    fmt.Println(animal, ok)
    animal.NewName("CHI THONG")
    fmt.Println("animal", animal)
    fmt.Println("animal.name", animal.name)


    fmt.Println(repeat, "implement interface; use interface access, ko receiver pointer", repeat)
    var human Human
    human = &Person{name:"Thanh Dung"}
    fmt.Println("human", human) // run String()
    person := human.(*Person)
    person.GetName()
    person.NewName("ekekekekek")
    person.GetName()


}


type Stringer interface {
    String() string
}

type Animal struct {
    name string
}
// A implements Stringer interface
func (*Animal) String() string {
    return "Tao la Stringer"
}

func(a Animal) GetName() {
    fmt.Println(a.name)
}

func(a *Animal) NewName(name string) {
    a.name = name
}


type Human interface {
    String() string
}

type Person struct {
    name string
}
// A implements Stringer interface
func (*Person) String() string {
    return "Tao la Human"
}

func(a Person) GetName() {
    fmt.Println(a.name)
}

func(a *Person) NewName(name string) {
    a.name = name
}
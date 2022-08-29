package main

import (  
    "fmt"
    "log"
    "phong/tricks"
)
/*
	interface nơi chứa methods
	khi các struct implement= tức là code các methods trùng tên với methods co trong interface
	thì khi đó các struct đều có type la interface đó 

	ứng dụng: khi Human, Dog đều implement interface Animal
	có human1 human2, dog1, dog2 đều có type Animal 
	arr = []Animal{human1, human2, dog1, dog2} //ok 


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


type Tester interface {  
    Test()
}

// non-struct
/*
	one := MyFloat(80.5)
	suy ra: m=80.5
	one.Test() // run method print ra 80.5
*/
type MyFloat float64

func (m MyFloat) Test() {  
    fmt.Println(m)
}

func describe(t Tester) {  
    log.Printf("Interface type %T value %v\n", t, t)
}


type Describer interface {  
    Describe()
}
type Person struct {  
    name string
    age  int
}

func (p Person) Describe() {  
    fmt.Printf("%s is %d years old", p.name, p.age)
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


    tricks.Format("non-struct")
    var t Tester
    f := MyFloat(89.7) //m=89.7 
    t = f
    describe(t)
    t.Test()


    tricks.Format("varName.(type)")
   	findType("Naveen")

    p := Person{
        name: "Naveen R",
        age:  25,
    }
    findType(p)

}
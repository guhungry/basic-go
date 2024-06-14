package methods

import (
	"bee-playground/money"
	"bee-playground/utils"
	"fmt"
)

type Bee struct {
	Name string
	Age  int
}

func (m Bee) isAdult() bool {
	return m.Age > 20
}

func Methods() {
	bee := Bee{Name: "Bee", Age: 35}

	utils.Dump(bee)
	fmt.Printf("isAdult: %t\n", bee.isAdult())
}

func Workshop() {
	result := Employee{FirstName: "Bee", LastName: "Haha", Salary: 55555555}

	utils.Dump(result)
	result.FullName()
	result.PrintSalary()
}

type Employee struct {
	FirstName string
	LastName  string
	Branch    string
	Salary    float64
}

func (v Employee) FullName() {
	fmt.Printf("Full Name: %s %s\n", v.FirstName, v.LastName)
}

func (v Employee) PrintSalary() {
	fmt.Printf("Salary: $%s\n", money.FormatMoney(v.Salary))
}

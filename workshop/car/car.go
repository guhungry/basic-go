package car

import (
	"bee-playground/money"
	"fmt"
)

type Car struct {
	Name  string
	Model string
	Price float64
}

func Workshop() {
	car := Car{Name: "BMW", Model: "i8", Price: 13000000}

	PrintDetails(car.Name, car.Model, car.Price)
	PrintDetailWithCustomType(car)
	PrettyPrintPrice(car)
}

func PrintDetails(name string, model string, price float64) {
	fmt.Printf("Name: %s\nModel: %s\nPrice: $%f\n", name, model, price)
}

func PrintDetailWithCustomType(car Car) {
	fmt.Printf("%+v", car)
}

func PrettyPrintPrice(car Car) {
	fmt.Printf("Name: %s\nModel: %s\nPrice: $%s\n", car.Name, car.Model, money.FormatMoney(car.Price))
}

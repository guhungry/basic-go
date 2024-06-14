package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"bee-playground/foo"
	"bee-playground/money"

	"github.com/shopspring/decimal"
)

const gconst = "Global Constant"

var gvar = "Global Variable"

func init() {
	fmt.Println("init function")
}

func main() {
	hello()
	line()
	workshopDataType()
	line()
	conversion()
	line()
	maptype()
	line()
	typedecimal()
	line()
	customtype()
	line()
	workshopAddress()
	line()
	arraySlice()
	line()
	workshopArraySlice()
	line()
	workshopIfElse()
	line()
	workshopFor()
	line()
	workshopCar()
	line()
	workshopPackage()
	line()
	methods()
	line()
	workshopMethods()
	line()
	workshopGenerics()
	line()
	workshopHttp()
}

func hello() {
	var inline string
	inline = "Inline"

	gvar = "Change"

	fmt.Printf("Hello, Bee (%s, %s, %s)\n", gconst, gvar, inline)
}

func line() {
	fmt.Println("======================")
}

func dump(val interface{}) {
	fmt.Printf("%T: %+v\n", val, val)
}

func workshopDataType() {
	golang := "golang is a simple programming language"
	number := 1024
	float := float64(100.35)
	boolean := true
	bite := []byte(`simple byte string`)

	dump(golang)
	dump(number)
	dump(float)
	dump(boolean)
	dump(bite)
}

func conversion() {
	str := "10.14"
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("parse float error:", str, err)
		return
	}
	dump(float)

	inter := interface{}("14")
	val, ok := inter.(string)
	if !ok {
		fmt.Println("parse interface error:", inter, ok)
		return
	}
	dump(val)
}

func maptype() {
	mstring := map[string]string{"1": "one", "2": "two"}
	mint := map[int]string{1: "one", 2: "two"}

	fmt.Println(mstring)
	fmt.Println(mint)

	for k, v := range mint {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}

	key := "1"
	val, ok := mstring[key]
	if !ok {
		fmt.Println("get value error:", key, ok)
		return
	}
	fmt.Printf("value: %s\n", val)
}

func typedecimal() {
	one := decimal.NewFromFloat(1)
	two := decimal.NewFromFloat(2)

	fmt.Println(one.Add(two))
}

func customtype() {
	type CustomType struct {
		firstName string
		lastName  string
		age       int
	}

	a := CustomType{firstName: "A", lastName: "B", age: 10}

	dump(a)

	var b CustomType
	b.firstName = "B"
	b.lastName = "C"
	b.age = 15

	dump(b)
}

func workshopAddress() {
	type Address struct {
		HomeNo     string
		StreetName string
		Province   string
		ZipCode    int
	}

	a := Address{HomeNo: "15", StreetName: "My Street", Province: "Nonthaburi", ZipCode: 11120}

	dump(a)
}

func arraySlice() {
	arr := [2]string{"foo", "bar"}
	dump(arr)

	sli := []string{"foo", "bar"}
	sli = append(sli, "append")
	sli = append(sli, "extra")
	dump(sli)

	fmt.Printf("%v\n", sli[2])
	fmt.Printf("%v\n", sli[1:])
	fmt.Printf("%v\n", sli[:2])
	fmt.Printf("%v\n", sli[1:3])
}

func workshopArraySlice() {
	id := "1234567890123"
	dump(id)

	last4Index := len(id) - 4
	dump(id[last4Index:])
}

func workshopIfElse() {
	score := 90
	grade := getGrage(score)

	fmt.Printf("Score: %d, grade: %s\n", score, grade)
}

func getGrage(score int) string {
	switch {
	case score > 90:
		return "A"
	case score > 80:
		return "B"
	case score > 70:
		return "C"
	case score > 60:
		return "D"
	default:
		return "Fail"
	}
}

func workshopFor() {
	sumFor()
	oddFor()
}

func sumFor() {
	sum := 0
	for ; sum <= 100; sum += 5 {
	}
	fmt.Printf("Sum: %d\n", sum)
}

func oddFor() {
	numbers := [100]int{}
	for i := range numbers {
		numbers[i] = i + 1
	}
	dump(numbers)

	//	for i := 1; i <= 100; i++ {
	for _, i := range numbers {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

type Car struct {
	Name  string
	Model string
	Price float64
}

func workshopCar() {
	car := Car{Name: "BMW", Model: "i8", Price: 13000000}

	CarPrintDetails(car.Name, car.Model, car.Price)
	CarPrintDetailWithCustomType(car)
	CarPrettyPrintPrice(car)
}

func CarPrintDetails(name string, model string, price float64) {
	fmt.Printf("Name: %s\nModel: %s\nPrice: $%f\n", name, model, price)
}

func CarPrintDetailWithCustomType(car Car) {
	fmt.Printf("%+v", car)
}

func CarPrettyPrintPrice(car Car) {
	fmt.Printf("Name: %s\nModel: %s\nPrice: $%0.2f\n", car.Name, car.Model, money.FormatMoney(car.Price))
}

func workshopPackage() {
	foo.PrintFoo()
}

type Bee struct {
	Name string
	Age  int
}

func (m Bee) isAdult() bool {
	return m.Age > 20
}

func methods() {
	bee := Bee{Name: "Bee", Age: 35}

	dump(bee)
	fmt.Printf("isAdult: %t\n", bee.isAdult())
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

func workshopMethods() {
	result := Employee{FirstName: "Bee", LastName: "Haha", Salary: 55555555}

	dump(result)
	result.FullName()
	result.PrintSalary()
}

func workshopGenerics() {
	ints := []int{1, 2, 3, 4, 5}
	float32s := []float32{4.7, 5.9, 1.2, 8.6}
	float64s := []float64{1.23, 6.33, 12.6}

	dump(sumNumber(ints))
	dump(sumNumber(float32s))
	dump(sumNumber(float64s))
}

func sumNumber[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](numbers []T) T {
	result := T(0)

	for _, v := range numbers {
		result += v
	}
	return result
}

type Users struct {
	Code int `json:"code"`
	Data []struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Gender string `json:"gender"`
		Status string `json:"status"`
	} `json:"data"`
}

func workshopHttp() {
	url := "https://gorest.co.in/public-api/users"

	response, err := requestJson[Users](url, nil)
	if err != nil {
		log.Println(err)
	}
	dump(response)
}

func requestJson[T any](url string, body any) (T, any) {
	var result T

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("http request error:", err)
		return result, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("client do error:", err)
		return result, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Println("parse json error:", err)
		return result, err
	}
	return result, nil
}

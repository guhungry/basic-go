package main

import (
	"fmt"
	"strconv"

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
}

func hello() {
	var inline string
	inline = "Inline"

	gvar = "Change"

	fmt.Println("Hello, Bee (%s, %s, %s)", gconst, gvar, inline)
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

	fmt.Printf("%T: %s\n%T: %d\n%T: %g\n%T: %t\n%T: %v\n", golang, golang, number, number, float, float, boolean, boolean, bite, bite)
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

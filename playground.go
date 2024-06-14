package main

import (
	"fmt"
	"strconv"

	"bee-playground/utils"
	"bee-playground/workshop/arraySlice"
	"bee-playground/workshop/batch"
	"bee-playground/workshop/car"
	"bee-playground/workshop/channel"
	"bee-playground/workshop/csv"
	"bee-playground/workshop/fors"
	"bee-playground/workshop/generics"
	"bee-playground/workshop/http"
	"bee-playground/workshop/ifelse"
	"bee-playground/workshop/methods"
	"bee-playground/workshop/options"
	"bee-playground/workshop/packages"
	"bee-playground/workshop/parallel"
	"bee-playground/workshop/postgresql"
	"bee-playground/workshop/test"

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
	arraySlice.ArraySlice()
	line()
	arraySlice.Workshop()
	line()
	ifelse.Workshop()
	line()
	fors.Workshop()
	line()
	car.Workshop()
	line()
	packages.Workshop()
	line()
	methods.Methods()
	line()
	methods.Workshop()
	line()
	generics.Workshop()
	line()
	http.Workshop()
	line()
	csv.Workshop()
	line()
	batch.Workshop()
	line()
	options.Workshop()
	line()
	postgresql.Workshop()
	line()
	parallel.Workshop()
	line()
	channel.Workshop()
	line()
	test.Workshop()
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

func workshopDataType() {
	golang := "golang is a simple programming language"
	number := 1024
	float := float64(100.35)
	boolean := true
	bite := []byte(`simple byte string`)

	utils.Dump(golang)
	utils.Dump(number)
	utils.Dump(float)
	utils.Dump(boolean)
	utils.Dump(bite)
}

func conversion() {
	str := "10.14"
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("parse float error:", str, err)
		return
	}
	utils.Dump(float)

	inter := interface{}("14")
	val, ok := inter.(string)
	if !ok {
		fmt.Println("parse interface error:", inter, ok)
		return
	}
	utils.Dump(val)
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

	utils.Dump(a)

	var b CustomType
	b.firstName = "B"
	b.lastName = "C"
	b.age = 15

	utils.Dump(b)
}

func workshopAddress() {
	type Address struct {
		HomeNo     string
		StreetName string
		Province   string
		ZipCode    int
	}

	a := Address{HomeNo: "15", StreetName: "My Street", Province: "Nonthaburi", ZipCode: 11120}

	utils.Dump(a)
}

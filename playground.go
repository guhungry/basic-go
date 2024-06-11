package main

import (
	"fmt"
	"strconv"
)

const gconst = "Global Constant"

var gvar = "Global Variable"

func init() {
	fmt.Println("init function")
}

func main() {
	hello()
	line()
	datatype()
	line()
	conversion()
	line()
	maptype()
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

func datatype() {
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
	fmt.Printf("%T: %v\n", float, float)
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

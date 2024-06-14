package utils

import "fmt"

func Dump(val interface{}) {
	fmt.Printf("%T: %+v\n", val, val)
}

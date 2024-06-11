package foo

import (
	"fmt"
)

func PrintFoo() {
	fmt.Println("Print Foo")
	privateBar()
}

func privateBar() {
	fmt.Println("Print Bar")
}

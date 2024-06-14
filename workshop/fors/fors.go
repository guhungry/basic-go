package fors

import (
	"bee-playground/utils"
	"fmt"
)

func Workshop() {
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
	utils.Dump(numbers)

	//	for i := 1; i <= 100; i++ {
	for _, i := range numbers {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

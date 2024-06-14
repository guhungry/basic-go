package arraySlice

import (
	"bee-playground/utils"
	"fmt"
)

func ArraySlice() {
	arr := [2]string{"foo", "bar"}
	utils.Dump(arr)

	sli := []string{"foo", "bar"}
	sli = append(sli, "append")
	sli = append(sli, "extra")
	utils.Dump(sli)

	fmt.Printf("%v\n", sli[2])
	fmt.Printf("%v\n", sli[1:])
	fmt.Printf("%v\n", sli[:2])
	fmt.Printf("%v\n", sli[1:3])
}

func Workshop() {
	id := "1234567890123"
	utils.Dump(id)

	last4Index := len(id) - 4
	utils.Dump(id[last4Index:])
}

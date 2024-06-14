package generics

import "bee-playground/utils"

func Workshop() {
	ints := []int{1, 2, 3, 4, 5}
	float32s := []float32{4.7, 5.9, 1.2, 8.6}
	float64s := []float64{1.23, 6.33, 12.6}

	utils.Dump(sumNumber(ints))
	utils.Dump(sumNumber(float32s))
	utils.Dump(sumNumber(float64s))
}

func sumNumber[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](numbers []T) T {
	result := T(0)

	for _, v := range numbers {
		result += v
	}
	return result
}

package test

import (
	"bee-playground/utils"
	"strings"
)

func Workshop() {
	utils.Dump(TransformText(" 123 AF 2w  "))
}

func TransformText(input string) string {
	// Replace number to word, remove whitespaces
	var replacers = strings.NewReplacer(
		"1", "one",
		"2", "two",
		"3", "three",
		"4", "four",
		"5", "five",
		"6", "six",
		"7", "seven",
		"8", "eight",
		"9", "nine",
		"0", "zero",
		" ", "",
	)
	return replacers.Replace(strings.ToLower(input))
}

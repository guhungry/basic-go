package ifelse

import "fmt"

func Workshop() {
	score := 90
	grade := getGrade(score)

	fmt.Printf("Score: %d, grade: %s\n", score, grade)
}

func getGrade(score int) string {
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

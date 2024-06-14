package parallel

import (
	"fmt"
	"time"
)

func Workshop() {
	go print("foo")
	print("bar")
}

func print(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
}

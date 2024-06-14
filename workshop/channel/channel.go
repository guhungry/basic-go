package channel

import (
	"bee-playground/utils"
	"bee-playground/workshop/http"
	"log"
	"sync"
)

var wg sync.WaitGroup

func Workshop() {
	result := make(chan string, 3)
	wg.Add(3)

	go callUrl("https://www.google.com", result)
	go callUrl("https://www.twitch.tv", result)
	go callUrl("https://www.kasikornbank.com", result)

	wg.Wait()
	for i := 0; i < 3; i++ {
		utils.Dump(len(<-result))
	}
}

func callUrl(url string, result chan string) {
	defer wg.Done()

	response, err := http.RequestText(url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	result <- response
}

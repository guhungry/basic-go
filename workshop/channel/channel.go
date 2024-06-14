package channel

import (
	"bee-playground/utils"
	"bee-playground/workshop/http"
	"log"
	"sync"
)

var wg sync.WaitGroup

func Workshop() {
	wg.Add(3)

	go callUrl("https://www.google.com")
	go callUrl("https://www.twitch.tv")
	go callUrl("https://www.kasikornbank.com")

	wg.Wait()
}

func callUrl(url string) {
	defer wg.Done()

	response, err := http.RequestText(url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	utils.Dump(response)
}

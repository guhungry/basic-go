package batch

import (
	"bee-playground/utils"
	"bee-playground/workshop/http"
	"flag"
	"fmt"
	"log"
	"time"
)

const url = "https://gorest.co.in/public-api/"

func Workshop() {
	api := flag.String("api", "", "api to call only 'comments' or 'todos'")
	flag.Parse()

	utils.Dump(*api)
	if *api == "comments" {
		handleComments(*api)
	} else if *api == "todos" {
		handleTodos(*api)
	} else {
		fmt.Println("wrong")
	}
}

type Todos struct {
	Code int `json:"code"`
	Data []struct {
		Id     int       `json:"id"`
		UserId int       `json:"user_id"`
		Title  string    `json:"title"`
		DueOn  time.Time `json:"due_on"`
		Status string    `json:"status"`
	} `json:"data"`
}

type Comments struct {
	Code int `json:"code"`
	Data []struct {
		Id     int    `json:"id"`
		PostId int    `json:"post_id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Body   string `json:"body"`
	} `json:"data"`
}

func handleTodos(api string) {
	todos, err := http.RequestJson[Todos](url+api, nil)
	if err != nil {
		log.Println(err)
		return
	}

	utils.Dump(todos)
}

func handleComments(api string) {
	todos, err := http.RequestJson[Comments](url+api, nil)
	if err != nil {
		log.Println(err)
		return
	}

	utils.Dump(todos)

}

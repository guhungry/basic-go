package http

import (
	"bee-playground/utils"
	"encoding/json"
	"log"
	"net/http"
)

type Users struct {
	Code int `json:"code"`
	Data []struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Gender string `json:"gender"`
		Status string `json:"status"`
	} `json:"data"`
}

func Workshop() {
	url := "https://gorest.co.in/public-api/users"

	response, err := requestJson[Users](url, nil)
	if err != nil {
		log.Println(err)
	}
	utils.Dump(response)
}

func requestJson[T any](url string, body any) (T, any) {
	var result T

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("http request error:", err)
		return result, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("client do error:", err)
		return result, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Println("parse json error:", err)
		return result, err
	}
	return result, nil
}

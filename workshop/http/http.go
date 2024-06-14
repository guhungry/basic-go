package http

import (
	"bee-playground/utils"
	"bytes"
	"encoding/json"
	"io"
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

	response, err := GetJson[Users](url)
	if err != nil {
		log.Println(err)
	}
	utils.Dump(response)
}

func GetJson[T any](url string) (T, any) {
	return requestJson[T](http.MethodGet, url, nil)
}

func PostJson[T any](url string, body any) (T, any) {
	return requestJson[T](http.MethodGet, url, body)
}

func requestJson[T any](method string, url string, body any) (T, any) {
	var result T

	bodyReader, err := toJsonBody(body)
	if err != nil {
		log.Println("parse request error:", err)
		return result, err
	}
	req, err := makeHttpRequest(method, url, bodyReader)
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

func makeHttpRequest(method string, url string, bodyReader *bytes.Reader) (*http.Request, error) {
	if bodyReader == nil {
		return http.NewRequest(method, url, nil)
	}
	return http.NewRequest(method, url, bodyReader)
}

func toJsonBody(body any) (*bytes.Reader, error) {
	if body == nil {
		return nil, nil
	}
	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(requestBody), nil
}

func RequestText(url string, body any) (string, any) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("http request error:", err)
		return "", err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("client do error:", err)
		return "", err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyBytes), nil
}

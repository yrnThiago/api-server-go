package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Session *http.Client
}

var (
	client    Client
	LOCALHOST string = "http://localhost:3333"
)

func ClientInit() {
	client = Client{
		&http.Client{},
	}
}

func MakeRequest(url string, method string, body []byte) *http.Request {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Erro criar requesição")
	}

	req.Header.Set("Content-Type", "application/json")
	return req
}

func DoRequest(req *http.Request) []byte {
	request, err := client.Session.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer request.Body.Close()

	response, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Erro no parser")
	}

	return response
}

func Checkout() {
	data := map[string]any{
		"ID":   "12345",
		"Date": "2024-03-11",
		"Product": map[string]any{
			"ID":    "98765",
			"Name":  "Smartphone",
			"Price": 1499.99,
		},
	}

	// Convertendo para JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}
	req := MakeRequest(LOCALHOST+"/checkout", http.MethodPost, jsonData)
	DoRequest(req)
}

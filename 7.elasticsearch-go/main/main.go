package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	baseURL string
}

func newClient(host string) *Client {
	return &Client{host}
}

func (c *Client) CheckHealth() error {
	response, err := http.Get(c.baseURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed check Elasticsearch health : %v ", err)
	}

	log.Println("debug health check response : ", string(responseBody))

	return nil
}

type Employee struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Salary  float64 `json:"salary"`
}

func (c *Client) InsertData(e *Employee) error {
	body, _ := json.Marshal(e)
	http.NewRequest("PUT", c.baseURL)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type server struct {
	name string
	url  string
}

type response struct {
	Count int `json:"count"`
}

func main() {
	// Set up the servers to poll
	servers := []server{
		{"maria.ru", "http://maria.ru/api/count"},
		{"rose.ru", "http://rose.ru/api/count"},
		{"sina.ru", "http://sina.ru/api/count"},
	}

	// Poll the servers every minute
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		for _, s := range servers {
			// Send a GET request to the server
			resp, err := http.Get(s.url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			defer resp.Body.Close()

			// Read the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// Unmarshal the response into a struct
			var r response
			err = json.Unmarshal(body, &r)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// Print the current time, server name, and count
			now := time.Now().Format("2006-01-02 15:04:05")
			fmt.Printf("%s %s %d\n", now, s.name, r.Count)
		}
	}
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type StatusChecker interface {
	Check(ctx context.Context, name string) (status bool, err error)
}

type httpChecker struct{}

func (h httpChecker) Check(ctx context.Context, name string) (status bool, err error) {
	resp, err := http.Get("http://" + name)
	if err != nil || resp.StatusCode != 200 {
		return false, err
	}
	return true, nil
}

var websiteMap = make(map[string]string)
var mu sync.Mutex
var checker StatusChecker = httpChecker{}

func websitesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}

	var websites []string
	err := json.NewDecoder(r.Body).Decode(&websites)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	mu.Lock()
	for _, website := range websites {
		websiteMap[website] = ""
	}
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
	res := map[string]string{"message": "Websites received successfully"}
	json.NewEncoder(w).Encode(res)
}

func getWebsiteStatuses(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}
	var websMap = make(map[string]string)

	website := r.URL.Query().Get("name")

	if website != "" {
		if status, exists := websiteMap[website]; exists {
			websMap[website] = status
		} else {
			http.Error(w, "Invalid parameter", http.StatusBadRequest)
		}

	} else {
		websMap = websiteMap
	}
	w.Header().Set("Content-Type", "application/json")
	if len(websMap) != 0 {
		json.NewEncoder(w).Encode(websMap)
	}
}

func monitorWebsites() {
	for {
		time.Sleep(1 * time.Minute)
		mu.Lock()
		for website := range websiteMap {
			status, _ := checker.Check(context.Background(), website)
			if status {
				websiteMap[website] = "UP"
				fmt.Println(website, status)
			} else {
				websiteMap[website] = "DOWN"
				fmt.Println(website, status)
			}
		}
		mu.Unlock()
	}
}

func main() {
	go monitorWebsites()
	http.HandleFunc("/websites", websitesHandler)
	http.HandleFunc("/website-status", getWebsiteStatuses)
	fmt.Println("Server started...")
	http.ListenAndServe(":3000", nil)
}

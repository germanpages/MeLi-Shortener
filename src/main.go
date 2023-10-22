package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var urlMap = make(map[string]string)
var mutex = &sync.Mutex{}
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateShortURL() string {
	const caracteres = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := make([]byte, 6)
	for i := range shortURL {
		shortURL[i] = caracteres[rng.Intn(len(caracteres))]
	}
	return string(shortURL)
}

func createShortURL(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("url")
	if longURL == "" {
		http.Error(w, "Missing URL parameter", http.StatusBadRequest)
		return
	}
	mutex.Lock()
	shortURL := generateShortURL()
	urlMap[shortURL] = longURL
	mutex.Unlock()
	fmt.Fprintf(w, "Short URL: http://localhost:8080/%s\n", shortURL)
}

func resolveURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	mutex.Lock()
	longURL, ok := urlMap[shortURL]
	mutex.Unlock()
	if !ok {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/create", createShortURL)
	http.HandleFunc("/", resolveURL)
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

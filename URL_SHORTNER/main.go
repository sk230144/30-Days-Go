package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalUrl  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

//in memory database

var urlDB = make(map[string]URL)

func generateShortUrl(OriginalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalUrl)) //convert url in to bytes
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	fmt.Println("hasher data:", hash[:8])

	return hash[:8] //return only starting 8 chars
}

func createUrl(OriginalUrl string) string {
	shortUrl := generateShortUrl(OriginalUrl)
	id := shortUrl
	urlDB[id] = URL{ //save in db using that short url
		ID:           id,
		OriginalUrl:  OriginalUrl,
		ShortURL:     shortUrl,
		CreationDate: time.Now(),
	}
	return shortUrl
}

func getUrl(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("Url not fount")
	}
	return url, nil
}
func shortUrlHandler(w http.ResponseWriter, r *http.Request){
	var data struct{
		URL string `json:"url"`
	}
	err:=json.NewDecoder(r.Body).Decode(&data)
    if(err != nil){
		http.Error(w,"Invalid Request Body", http.StatusBadRequest)
		return
	}
	shortUrl := createUrl(data.URL)
	response := struct{
		ShortUrl string `json:"short_url"`
	}{ShortUrl: shortUrl}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectUrlHandler(w http.ResponseWriter, r *http.Request){
	id := r.URL.Path[len("/redirect/"):]
	url, err := getUrl(id)
	if(err != nil){
		http.Error(w, "Invalid request", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}







func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Menthod")
	fmt.Fprintf(w, "Hello World")  //where we access response, this fprintf print there
}

func main() {
	//Register always handler function to handle all request to root url ("/")
	http.HandleFunc("/", handler)
    http.HandleFunc("/shorten", shortUrlHandler)
	http.HandleFunc("/redirect/", redirectUrlHandler)  


	//creating local server on http
	fmt.Println("Starting server on 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error in starting server")
	}
}

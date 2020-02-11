package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//URL var
//var URL = "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api_key="

//KEY const
var KEY = os.Getenv("KEY")

//PATH var
//var PATH = URL + KEY

/*
//URLHandler will make http req to site
func URLHandler(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("We got a request")
	//fmt.Fprintf(w, "Thanks for working server with HTTP method '%v'", r.Method)
	resp, err := http.Get("http://example.com/")

	if err != nil {
		fmt.Println("Error in handler")
	}
	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
	fmt.Println(string(body))
}
*/

//URLHandler l
func URLHandler(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var URL = "https://track.onestepgps.com/v3/api/public/device?latest_point=true"
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + KEY

	// Create a new request using http
	req, err := http.NewRequest("GET", URL, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
	log.Println(string([]byte(body)))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", URLHandler).Methods("GET", "OPTIONS")
	//http.Handle("/", router)

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	fmt.Println("Listening on localhost: 8080")
	log.Fatal(http.ListenAndServe(":"+PORT, router))

}

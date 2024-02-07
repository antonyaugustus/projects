package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	details "examples/go-microservices/details"

	geometry "examples/go-microservices/geometry"

	"github.com/gorilla/mux"
)

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	title := vars["title"]
// 	page := vars["page"]

// 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
// }

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := details.GetHostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(hostname)

	ip, err := details.GetIp()
	if err != nil {
		panic(err)
	}
	fmt.Println(ip)

	response := map[string]string{
		"hostname": hostname,
		"ip":       ip,
	}
	json.NewEncoder(w).Encode(response)
}

func geometryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Calculating...")
	length, _ := strconv.ParseFloat(r.URL.Query().Get("length"), 64)
	width, _ := strconv.ParseFloat(r.URL.Query().Get("width"), 64)

	response := map[string]float64{
		"length":   length,
		"width":    width,
		"area":     geometry.Area(length, width),
		"diagonal": geometry.Diagonal(length, width),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {

	// fmt.Println(quote.Go())

	// Area := geometry.Area(2, 3)
	// Diagonal := geometry.Diagonal(2, 3)
	// fmt.Println("Area = ", Area)
	// fmt.Println("Diagonal = ", Diagonal)

	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/health", healthHandler)
	//r.HandleFunc("/books/{title}/page/{page}", rootHandler)
	r.HandleFunc("/details", detailsHandler)
	r.HandleFunc("/geometry", geometryHandler)
	log.Println("Server has started. ")
	log.Fatal(http.ListenAndServe(":80", r))

}

package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/drone/routes"
)

type API struct {
	Message string "json:message"
}

// using routing with Gorilla muxers
func Hello (w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	name := urlParams["user"]
	HelloMessage := "Hello, " + name

	message := API{HelloMessage}
	output, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	fmt.Fprintf(w, string(output) )
}

// using routing using Routes from drone.io

func sayHello (w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()
	name := urlParams.Get(":name")
	HelloMessage := "Hello " + name
	message := API{HelloMessage}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Something went wrong!")
	}

	fmt.Fprintf(w, string(output) )
}

func main() {
/*	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request ) {
		message := API{"Hello World"}
		output, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Something went wrong!")
		}
		fmt.Fprintf(w, string(output) )
	})
*/

/*
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", Hello)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
*/
	mux := routes.New()
	mux.Get("/api/:name", sayHello)
	http.Handle("/", mux)
	http.ListenAndServe(":8080", nil)
}
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Water int
	Wind  int
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/get-data", getData)
	fmt.Println("Listening to port:8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	p := randomizeInt()
	t, _ := template.ParseFiles("status.html")
	t.Execute(w, p)
}

func getData(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(randomizeInt())
	fmt.Fprintln(w, string(response))
}

func randomizeInt() Status {
	var status Status
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	status.Water = rand.Intn(max-min+1) + min
	status.Wind = rand.Intn(max-min+1) + min
	return status
}

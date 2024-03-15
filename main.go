package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting DiffDeNoiser")
	http.ListenAndServe(":5000", muxroutes())
	RunDiff()
}

func muxroutes() *http.ServeMux {
	fileServer := http.FileServer(http.Dir("./Static/"))

	//Using a serveMux is good practise because we can define all routes here instead of having many http handlefuncs
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	//mux.HandleFunc("/", Homepage)

	return mux
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hallo homepage")
}

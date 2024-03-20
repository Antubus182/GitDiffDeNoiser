package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type DiffData struct {
	Sha1 string
	Sha2 string
	Dir  string
}

func main() {
	port := flag.String("addr", ":5000", "HTTP network port")
	fmt.Printf("starting DiffDeNoiser at port %s", *port)
	http.ListenAndServe(*port, muxroutes())
}

func muxroutes() *http.ServeMux {
	homepage := http.FileServer(http.Dir("./Static/"))

	//Using a serveMux is good practise because we can define all routes here instead of having many http handlefuncs
	mux := http.NewServeMux()
	mux.Handle("/", homepage)
	mux.HandleFunc("/diff", DiffPage)
	return mux
}

func DiffPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed) //405= Method Not Allowed
		w.Write([]byte("Method not Allowed"))
		return
	}
	var d DiffData

	fmt.Println("Diffje doen")

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("the following came in: ")
	sha1 := d.Sha1
	sha2 := d.Sha2
	dir := d.Dir
	fmt.Printf("Sha1: %s, Sha2: %s, Directory: %s\n", sha1, sha2, dir)
	if !VerifyInputs(d) {
		w.Write([]byte("Invalid sha codes, please provide two 40 char long SHA-1 code"))
		return
	}
	diffoutput := RunDiff()
	w.Write([]byte(diffoutput))
}

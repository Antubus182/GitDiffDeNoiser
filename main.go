package main

import (
	"flag"
	"fmt"
	"net/http"
)

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
	fmt.Println("Diffje doen")
	diffoutput := RunDiff()
	w.Write([]byte(diffoutput))
}

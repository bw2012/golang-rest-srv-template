package main

import (
	"fmt"
	"net/http"
	//"test/testpkg"
	"encoding/json"
	"strings"
)

type RestResponse struct {
	Name string
	Body string
	Time int64
}

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	//decoder := json.NewDecoder(r.Body)
	//var t test_struct
	// err := decoder.Decode(&t)
	//if err != nil {
	//    panic()
	// }
	//log.Println(t.Test)

	ip := r.RemoteAddr
	header := r.Header

	fmt.Println("ip --- %s", ip)
	fmt.Println("header --- %s", header)

	var url string = r.URL.Path[1:]
	par := strings.Split(url, "/")

	var a string = par[0]
	var b string = par[1]

	m := RestResponse{a, b, 1111}

	out, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	fmt.Println("output --- %s", string(out))
	//fmt.Println("test")

	//fmt.Fprintf(w, "test --- %s!", r.URL.Path[1:])
	fmt.Fprintf(w, string(out))
	//fmt.Fprintf(w, "test2 --- %s", b)
}

func main() {
	http.HandleFunc("/restapi/", defaultHandler)
	http.ListenAndServe(":8080", nil)

	//testpkg.Test();
}

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

	fmt.Println("remote ip --- ", ip)
	fmt.Println("header    --- ", header)

	// parse url
	var url string = r.URL.Path[1:]
	par := strings.Split(url, "/")

	// get parameters
	var a string = par[0]
	var b string = par[1]

	// create response
	m := RestResponse{a, b, 1111}

	// response to json
	out, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	fmt.Println("output --- ", string(out))
	
	// put response to http
	fmt.Fprintf(w, string(out))
}

func main() {

	fmt.Println("Start example rest server at port 8080")

	http.HandleFunc("/restapi/", defaultHandler)
	http.ListenAndServe(":8080", nil)
	
	//testpkg.Test();
}

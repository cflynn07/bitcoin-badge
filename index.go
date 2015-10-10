// https://golang.org/doc/articles/wiki/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Equivalent of console.log?
	log.Print(r.URL.Path)

	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Missing address")
		return
	}

	// https://golang.org/pkg/net/http/
	// https://toshi.io/docs/#unconfirmed-transactions

	// https://golang.org/pkg/net/http/#Request
	// https://golang.org/pkg/net/url/
	// http.Request Struct contains URL key
	//   http.Request.URL
	resp, err := http.Get("https://bitcoin.toshi.io/api/v0/addresses" + r.URL.Path)
	if err != nil {
		fmt.Fprintf(w, "Can not fetch recent blocks from bitcoin.toshi.io")
	}
	// Pushes function call onto a list, this list of saved calls is
	// executed after surrounding function returns
	defer resp.Body.Close() // to be done after handler returns
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "error")
	}

	//Extracting JSON
	// To pluck specific keys, use predefined Struct type with
	// specific keys
	type AddressValues struct {
		hash     string
		balance  string
		received string
		sent     string
	}
	var av AddressValues
	err2 := json.Unmarshal(body, av)
	if err2 != nil {
		fmt.Fprintf(w, "error2")
		return
	}

	// How to log structs to console
	// http://stackoverflow.com/a/24512194/480807
	fmt.Printf("%+v\n", av)

	s := bytes.NewBuffer(body).String()

	// body is a byte array, must convert to string
	// http://stackoverflow.com/questions/14230145/what-is-the-best-way-to-convert-byte-array-to-string
	fmt.Fprintf(w, s)
}

/**
 *
j*/
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

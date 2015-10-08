// https://golang.org/doc/articles/wiki/
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// https://golang.org/pkg/net/http/
	// https://toshi.io/docs/#unconfirmed-transactions
	resp, err := http.Get("https://bitcoin.toshi.io/api/v0/addresses/16jbpwYG79qN5f9kLVhXMLLRAoeEEy24KM")
	if err != nil {
		fmt.Fprintf(w, "Can not fetch recent blocks from bitcoin.toshi.io")
	}
	// Pushes function call onto a list, this list of saved calls is
	// executed after surrounding function returns
	defer resp.Body.Close() // to be done after handler returns
	body, err := ioutil.ReadAll(resp.Body)

	s := bytes.NewBuffer(body).String()

	// body is a byte array, must convert to string
	// http://stackoverflow.com/questions/14230145/what-is-the-best-way-to-convert-byte-array-to-string
	fmt.Fprintf(w, s)
}

/**
 *
 */
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

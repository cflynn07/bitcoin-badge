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
	/*
		{
			"hash":"1FfmbHfnpaZjKFvyi1okTjJJusN455paPH",
			"balance":3416228,
			"received":14434157021163,
			"sent":14434153604935,
			"unconfirmed_received":0,
			"unconfirmed_sent":0,
			"unconfirmed_balance":0
		}
	*/
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

	var av interface{}
	err2 := json.Unmarshal(body, &av)
	if err2 != nil {
		fmt.Fprintf(w, "error2")
		log.Print(err2)
		return
	}
	// How to log structs to console
	// http://stackoverflow.com/a/24512194/480807
	// NOTE: output appears as: (why?)
	//   {hash: balance: received: sent:}
	fmt.Printf("%+v\n", av)

	// To access, use type assertion
	m := av.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv) //No output
		}
	}

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

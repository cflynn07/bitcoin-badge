/*
Every Go package should have a block comment at the top of the file. godoc
will extract for documentation.

This is a sandbox project for me to play with Go. Eventually I would like
to create a server that returns an image "badge" for a given Bitcoin
address that contains information such as balance, recieved, sent, etc.
*/
package main

// https://golang.org/doc/articles/wiki/
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

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("a is an array of length ", len(a))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// https://gobyexample.com/slices
	// Slices
	slice := make([]string, 3)
	fmt.Println("empty slice:", slice)
	// builtin "append" for slices, benefit over arrays
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	slice = append(slice, "d")
	slice = append(slice, "e")
	fmt.Println("slice with data: ", slice)
	fmt.Println("slice length", len(slice))

	//Copy a slice
	slice2 := make([]string, 5)
	copy(slice2, slice)
	fmt.Println("copied slice", slice2)

	fmt.Println("output sub-slice", slice2[2:4])
	fmt.Println("output slice from 2 to end", slice[2:])

	// single line delcaration and initialization
	slice3 := []string{"x", "y", "z"}
	fmt.Println("slice 3!", slice3)

	// Trivial example of multi-dimensional, variable length slice
	twoD2 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD2[i] = make([]int, innerLen)
		for k := 0; k < innerLen; k++ {
			twoD2[i][k] = i + k
		}
	}
	fmt.Println("nested slice structure example", twoD2)

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

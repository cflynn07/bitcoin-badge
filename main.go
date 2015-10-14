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
	"fmt"
	"github.com/cflynn07/bitcoin-badge/db"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	// body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "error")
	}

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, m, nil)
	if err != nil {
		log.Print("jpeg.Encode error", err)
	}
	buffBytes := buf.Bytes()
	err = ioutil.WriteFile("./img.jpg", buffBytes, 0777)
	if err != nil {
		log.Print("ioutil.WriteFile error", err)
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffBytes)))
	w.Write(buffBytes)
	//s := bytes.NewBuffer(body).String()

	// body is a byte array, must convert to string
	// http://stackoverflow.com/questions/14230145/what-is-the-best-way-to-convert-byte-array-to-stringp
	//fmt.Fprintf(w, s)
}
func main() {

	item := db.LoadItem(5)
	log.Print("item ", item)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

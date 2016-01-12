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
	"encoding/json"
	"fmt"
	"github.com/cflynn07/bitcoin-badge/guestbook"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	a := &guestbook.GuestBookEntry{1, "email@gmail.com", "Title1", "Content1"}
	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(out))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

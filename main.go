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
	"github.com/gorilla/mux"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	a := &guestbook.GuestBookEntry{1, "email@gmail.com", "Title1", "Content1"}
	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(out))
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "hello "+vars["firstName"])
}

func PrefixHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "prefixed route handler, should match /prefix/prefix2")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test", TestHandler)
	r.HandleFunc("/name/{firstName}", NameHandler)

	sub := r.PathPrefix("/prefix").Subrouter()
	sub.HandleFunc("/prefix2", PrefixHandler)

	http.ListenAndServe(":8080", r)
}

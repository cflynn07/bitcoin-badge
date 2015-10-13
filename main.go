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
	"os"
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

	//Go compiler will count elements when you init an array
	sampleArr := [...]string{"One", "two", "three"}
	fmt.Println("sampleArr", sampleArr)

	s := bytes.NewBuffer(body).String()

	//Define a function in Go with unknown number of variables?
	// http://stackoverflow.com/questions/19238143/does-golang-support-arguments
	appendByte := func(slice []byte, data ...byte) []byte {
		m := len(slice)
		n := m + len(data)
		if n > cap(slice) {
			newSlice := make([]byte, (n+1)*2) // n+1 exactly what we need, *2 to double it
			copy(newSlice, slice)
			slice = newSlice
		}
		slice = slice[0:n]
		copy(slice[m:n], data)
		return slice
	}

	p := []byte{2, 3, 5}
	p = appendByte(p, 7, 9, 11)
	fmt.Println("output of appendByte", p)

	// body is a byte array, must convert to string
	// http://stackoverflow.com/questions/14230145/what-is-the-best-way-to-convert-byte-array-to-stringp
	fmt.Fprintf(w, s)
}

func nonExportReturnValFunction() (int, int) {
	return 5, 6
}

// Go not object-oriented. No objects or inheritance. No polymorphism or overloading
// composition over inheritance
type Cray struct {
	Name string
	Age  int
}

// Associating a method with a structure (16)
func (c *Cray) incAge() {
	c.Age = c.Age + 1
}

func (c *Cray) makeMary() {
	c.Name = "Mary"
}

func aFactoryMethod() *Cray {
	return &Cray{
		"Johnny",
		5,
	}
}

/**
 *
 */
func main() {
	if len(os.Args) > 1 {
		log.Print(os.Args[1])
	}
	// use the blank identifier, value not actually assigned
	_, a := nonExportReturnValFunction()
	log.Print("a: ", a)

	k := Cray{
		Name: "Thomas",
		Age:  100,
	}

	//declare new instance without field name declarations, rely on order of delared fields
	k = Cray{"Turle", 33}

	log.Print("Cray Name ", k.Name)
	log.Print("Cray Age ", k.Age)

	// Go always passes by values not reference (opposite of JavaScript)
	copyFn := func(v Cray) {
		// Only mutates a copy, not original
		v.Name = "COPY"
	}

	refFn := func(v *Cray) {
		//Mutates
		v.Name = "Reference!"
	}

	copyFn(k)
	log.Print("Has k.Name changed? (no) ", k.Name)

	refFn(&k) // & is the "address of" operator
	log.Print("Has k.Name changed? (yes) ", k.Name)

	log.Print("How old is k? ", k.Age)

	k.incAge()
	log.Print("How old is k now? ", k.Age)

	log.Print("what is k's name? ", k.Name)
	k.makeMary()
	log.Print("what is k's name now? ", k.Name)

	k2 := aFactoryMethod()
	log.Print("k2:")
	log.Print(k2.Name)
	log.Print(k2.Age)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

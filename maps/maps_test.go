// https://golang.org/doc/code.html#Testing
package main

import "testing"
import "fmt"

func TestExportFn(t *testing.T) {
	x := ExportFn()
	fmt.Println("test output")
	if x != 5 {
		t.Errorf("ExportFn() == %q, want %q", x, 5)
	}
}

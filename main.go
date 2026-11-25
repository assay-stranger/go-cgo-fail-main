// Command gocgofail cannot build: its cgo preamble includes a header that does not exist.
package main

/*
#include <coldrun_missing_header.h>
*/
import "C"

import "fmt"

func main() {
	fmt.Println(C.coldrun_missing())
}

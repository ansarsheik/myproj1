package main

import (
	"log"

	readdata "github.com/ansarsheik/goproj/readdata"
)

func main() {

	// You can get individual args with normal indexing.
	//arg := os.Args[1]
	//fmt.Println(arg)

	a := readdata.Data()

	log.Println(a)

	//wordPtr := flag.String("app", "foo", "a string")
	//fmt.Println(*wordPtr)
}

// #Go-PHP-Parser

// Attempts to compile PHP code to Go. Stupidly.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileargs := os.Args[1:]
	files := make(map[string][]byte)
	var err error
	for _, f := range fileargs {
		files[f], err = ioutil.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(files[f])
		fmt.Println(string(files[f]))
	}
}

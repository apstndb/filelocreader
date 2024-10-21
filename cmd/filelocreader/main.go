package main

import (
	"fmt"
	"github.com/apstndb/filelocreader"
	"log"
	"os"
)

func main() {
	for _, s := range os.Args[1:] {
		loc, err := filelocreader.ParseLoc(s)
		if err != nil {
			log.Fatalln(err)
		}

		result, err := filelocreader.ExtractLoc(loc)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(result))
	}
}

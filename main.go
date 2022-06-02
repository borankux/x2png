package main

import (
	"log"
	"os"
)

func main() {
	err := CreateApp().Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

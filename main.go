package main

import (
	"flag"
	"log"
)

var instaName = flag.String("n", "", "Instangram user name such as: 'pchchv'")

func main() {
	flag.Parse()
	if *instaName == "" {
		log.Fatal("You need to input -n=name")
	}
	input := *instaName
}

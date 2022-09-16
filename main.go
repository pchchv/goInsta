package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

var instaName = flag.String("n", "", "Instangram user name such as: 'pchchv'")

func main() {
	flag.Parse()
	if *instaName == "" {
		log.Fatal("You need to input -n=name")
	}
	input := *instaName
	// Set the folder for saving photos
	baseDir, err := filepath.Abs("../photo")
	if err != nil {
		panic(err)
	}
	fmt.Println(baseDir)
}

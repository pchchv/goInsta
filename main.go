package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/gedex/go-instagram/instagram"
)

var instaName = flag.String("n", "", "Instangram user name such as: 'pchchv'")
var client *instagram.Client

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

	// Get User info
	client = instagram.NewClient(nil)
}

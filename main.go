package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/gedex/go-instagram/instagram"
	"github.com/joho/godotenv"
)

var instaName = flag.String("n", "", "Instangram user name such as: 'pchchv'")
var client *instagram.Client
var ClientID string

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value. Outputs a panic if the value is missing.
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

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
	ClientID = getEnvValue("INSTAGRAMID")
	// Get User info
	client = instagram.NewClient(nil)
	client.ClientID = ClientID
}

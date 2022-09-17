package main

import (
	"flag"
	"fmt"
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

func FindPhotos(ownerName string, albumName string, userId string, baseDir string) {
	//Create folder
	dir := fmt.Sprintf("%v/%v", baseDir, ownerName)
	os.MkdirAll(dir, 0755)
}

func main() {
	var userId string
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
	// Search Users
	searchUsers, _, err := client.Users.Search(input, nil)
	for _, user := range searchUsers {
		if user.Username == input {
			userId = user.ID
		}
	}
	if userId == "" {
		log.Fatalln("Can not address user name: ", input, err)
	}
	userFolderName := fmt.Sprintf("[%s]%s", userId, input)
	fmt.Println("Starting download [", userId, "]", input)
	FindPhotos(userFolderName, input, userId, baseDir)
}

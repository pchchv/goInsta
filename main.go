package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gedex/go-instagram/instagram"
	"github.com/joho/godotenv"
)

var instaName = flag.String("n", "", "Instangram user name such as: 'pchchv'")
var client *instagram.Client
var FileIndex int = 0
var ClientID string
var m sync.Mutex

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

func GetFileIndex() (ret int) {
	m.Lock()
	ret = FileIndex
	FileIndex = FileIndex + 1
	m.Unlock()
	return ret
}

func DownloadWorker(destDir string, linkChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for target := range linkChan {
		var imageType string
		if strings.Contains(target, ".png") {
			imageType = ".png"
		} else {
			imageType = ".jpg"
		}
		resp, err := http.Get(target)
		if err != nil {
			log.Println("Http.Get\nerror: " + err.Error() + "\ntarget: " + target)
			continue
		}
		defer resp.Body.Close()
		m, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Println("image.Decode\nerror: " + err.Error() + "\ntarget: " + target)
			continue
		}
		// Ignore small images
		bounds := m.Bounds()
		if bounds.Size().X > 300 && bounds.Size().Y > 300 {
			imgInfo := fmt.Sprintf("pic%04d", GetFileIndex())
			out, err := os.Create(destDir + "/" + imgInfo + imageType)
			if err != nil {
				log.Printf("os.Create\nerror: %s", err)
				continue
			}
			defer out.Close()
			if imageType == ".png" {
				png.Encode(out, m)
			} else {
				jpeg.Encode(out, m, nil)
			}

			if FileIndex%30 == 0 {
				fmt.Println(FileIndex, " photos downloaded.")
			}
		}
	}
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

package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	filenamePtr := flag.String("filename", "", "a string")
	contentPtr := flag.String("content", "", "a string")
	tokenPtr := flag.String("token", "", "a string")
	flag.Parse()
	if *filenamePtr == "" {
		print("Filename is required in payload.")
		os.Exit(1)
	}
	if *contentPtr == "" {
		print("Content is required in payload.")
		os.Exit(1)
	}
	if *tokenPtr == "" {
		print("Token is required in payload.")
		os.Exit(1)
	}
	App(*filenamePtr, *contentPtr, *tokenPtr)
}

func App(name string, content string, token string) string {
	if !validatePayload(name, content, token) {
		return "Payload is invalid."
	}
	createDataDir()
	dataDir := "data/"
	f, err := os.Create(dataDir + name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(content)
	print("Success")
	return "Success"
}

func createDataDir() {
	dir := "data"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
}

func validatePayload(name string, content string, token string) bool {
	if name == "" || content == "" || token != "Bearer auth-token-123" {
		return false
	}
	return true
}

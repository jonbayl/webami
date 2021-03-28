package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const version string = "v1.2"

func help() {
	fmt.Println("webami - A whoami for the internet.")
	fmt.Println("Easily retrieve your public IP address from the command-line. webami utilises Ipify to retrieve your public IP address and return it within the command-line.")
	fmt.Println("\nUsage: simply run the webami executable to retrieve your public IP address.")
	fmt.Println("\nOther arguments:")
	fmt.Println("version: prints the current version of webami.")

	os.Exit(0)
}

func getPublicIp() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.ipify.org", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "webami/"+version)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", ip)

	defer resp.Body.Close()
	os.Exit(0)
}

func main() {
	if len(os.Args) == 1 {
		getPublicIp()
	}
	switch os.Args[1] {
	case "version":
		fmt.Println("webami version: " + version)
	default:
		help()
	}

}

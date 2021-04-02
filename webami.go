package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

const version string = "v1.3"

func help() {
	fmt.Println("webami - A whoami for the internet.")
	fmt.Println("Easily retrieve your public IP address from the command-line. webami utilises Ipify to retrieve your public IP address and return it within the command-line.")
	fmt.Println("\nUsage: simply run the webami executable to retrieve your public IP address.")
	fmt.Println("\nOther arguments:")
	fmt.Println("version: prints the current version of webami.")
	fmt.Println("use: specify the IP retrieveal service to use. Takes one argument which must be a valid URL, all other arguments are ignored. Example: webami use https://api.ipify.org")
}

func validateIp(ip []byte) error {
	reg, _ := regexp.Match("^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$", ip)
	if reg != true {
		return errors.New("an invalid response was returned by the queried service, it was not an IP address")
	}

	return nil
}

func prepareRequest(srv []string) (*http.Client, *http.Request) {
	var srvString string = srv[0]
	_, err := url.ParseRequestURI(srvString)
	if err != nil {
		log.Fatal("an invalid url was provided to be used as the IP retrieval service")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", srvString, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "webami/"+version)

	return client, req
}

func getPublicIp(srv []string) []byte {
	client, req := prepareRequest(srv)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = validateIp(ip)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	return ip
}

func printIp(ip []byte) {
	fmt.Printf("%s\n", ip)
	os.Exit(0)
}

func main() {
	if len(os.Args) == 1 {
		ip := getPublicIp([]string{"https://api.ipify.org"})
		printIp(ip)
	}
	switch os.Args[1] {
	case "version":
		fmt.Println("webami version: " + version)
	case "use":
		if len(os.Args) < 3 {
			log.Fatal("you did not provide use with enough arguments")
		}
		ip := getPublicIp(os.Args[2:])
		printIp(ip)
	default:
		help()
	}

}

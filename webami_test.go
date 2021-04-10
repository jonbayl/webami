package main

import (
	"net/http"
	"reflect"
	"testing"
)

// func TestHelp(t *testing.T) {
// 	if os.Getenv("TEST_HELP") == "1" {
// 		help()
// 		return
// 	}
// 	cmd := exec.Command(os.Args[0], "-test.run=TestHelp")
// 	fmt.Printf("%s", os.Args[0])
// 	cmd.Env = append(os.Environ(), "TEST_HELP=1")
// 	err := cmd.Run()

// 	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
// 		t.Errorf("Help returned %v when exit status 0 was expected.", err)
// 	}
// 	help()
// }

func Example_help() {
	help()
	// Output:
	// webami - A whoami for the internet.
	// Easily retrieve your public IP address from the command-line. webami utilises Ipify to retrieve your public IP address and return it within the command-line.
	//
	// Usage: simply run the webami executable to retrieve your public IP address.
	//
	// Other arguments:
	// version: prints the current version of webami.
	// use: specify the IP retrieveal service to use. Takes one argument which must be a valid URL, all other arguments are ignored. Example: webami use https://api.ipify.org
}
func TestValidateIp(t *testing.T) {
	validIp := []byte("10.10.10.10")
	invalidIp := []byte("this is invalid input")

	err := validateIp(validIp)
	if err != nil {
		t.Fatalf("Valid input: %s resulted in error: %s.", validIp, err)
	}

	err = validateIp(invalidIp)
	if err == nil {
		t.Fatalf("Invalid input: %s did not result in error.", invalidIp)
	}
}

func TestPrepareRequest(t *testing.T) {
	srv := []string{"https://api.ipify.org"}
	client, req := prepareRequest(srv)

	validClient := &http.Client{}
	if reflect.TypeOf(client) != reflect.TypeOf(validClient) {
		t.Errorf("Expected client of type %s, got: %s", reflect.TypeOf(validClient), reflect.TypeOf(client))
	}

	if req.Method != "GET" {
		t.Errorf("Expected request to have method GET, got: %s", req.Method)
	}

	expect := "webami"
	if req.Header.Get("User-Agent") != "webami/"+version {
		t.Errorf("Expected user agent: %s, got: %s", expect, req.Header.Get("User-Agent"))
	}

}

func TestGetPublicIp(t *testing.T) {
	ip := getPublicIp([]string{"https://api.ipify.org"})
	err := validateIp(ip)
	if err != nil {
		t.Errorf("getPublicIp returned %s, which is not valid input.", ip)
	}
}

package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestHelp(t *testing.T) {
	if os.Getenv("TEST_HELP") == "1" {
		help()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestHelp")
	cmd.Env = append(os.Environ(), "TEST_HELP=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		t.Errorf("Help returned %v when exit status 0 was expected.", err)
	}
}
func TestValidateIp(t *testing.T) {
	validIp := []byte("10.10.10.10")
	invalidIp := []byte("this is invalid input")

	err := validateIp(validIp)
	if err != nil {
		t.Errorf("Valid input: %s resulted in error: %s.", validIp, err)
	}

	err = validateIp(invalidIp)
	if err == nil {
		t.Errorf("Invalid input: %s did not result in error.", invalidIp)
	}
}

func TestGetPublicIp(t *testing.T) {
	//This test may need some evaluation. Is just testing for exit code enough?
	if os.Getenv("TEST_PUBLIC_IP") == "1" {
		getPublicIp()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestGetPublicIp")
	cmd.Env = append(os.Environ(), "TEST_PUBLIC_IP=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		t.Errorf("Help returned %v when exit status 0 was expected.", err)
	}
}

func TestGetPublicIpAltService(t *testing.T) {
	//This test may need some evaluation. Is just testing for exit code enough?
	if os.Getenv("TEST_PUBLIC_IP_ALT") == "1" {
		getPublicIp()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestGetPublicIp")
	cmd.Env = append(os.Environ(), "TEST_PUBLIC_IP_ALT=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		t.Errorf("Help returned %v when exit status 0 was expected.", err)
	}
}

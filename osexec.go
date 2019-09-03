package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	s := "ls -lsah"

	args := strings.Split(s, " ")

	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput() //wait for return
	if err != nil {
		log.Printf("Command failed with error: %v", err)
	}

	fmt.Printf("%s\n", b)
}

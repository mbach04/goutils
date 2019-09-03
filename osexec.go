package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	waitWithCombinedOut("head -c 1G </dev/urandom > tempfile")
	waitWithCombinedOut("ls -lsah")

}

func waitWithCombinedOut(s string) {

	cmd := exec.Command("sh", "-c", s)

	b, err := cmd.CombinedOutput() //wait for return
	if err != nil {
		log.Printf("Command failed with error: %v", err)
	}

	fmt.Printf("%s\n", b)
}

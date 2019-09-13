package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

var fileSize = "10M"
var tmpFile = "/run/media/mbach/USB-Gadget/.t10M"
var count = 1000

func main() {
	cmd := fmt.Sprintf("head -c %s </dev/urandom > %s", fileSize, tmpFile)
	// fmt.Println(cmd)
	waitWithCombinedOut(cmd)
	b, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	start := time.Now()
	writeFileNTimes(tmpFile, &b, count)
	elapsed := time.Since(start)
	deleteFile(tmpFile)
	fmt.Println("Finished", count, "cycles of", fileSize)
	fmt.Println("Time Elapsed:", elapsed)

}

func writeFileNTimes(path string, b *[]byte, n int) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for i := 1; i <= n; i++ {
		fmt.Println(i, "of", count)
		_, err = file.Write(*b)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		// save changes
		err = file.Sync()
		if err != nil {
			fmt.Println("Error syncing file before close:", err)
			return
		}
	}

	fmt.Println("Finished writing to tmp file:", tmpFile)
}

func deleteFile(path string) {
	var err = os.Remove(path)
	if err != nil {
		fmt.Println("Error deleting file: %v", err)
		return
	}
}

func waitWithCombinedOut(s string) {
	cmd := exec.Command("sh", "-c", s)

	_, err := cmd.CombinedOutput() //wait for return
	if err != nil {
		log.Printf("Command failed with error: %v", err)
	}

	// fmt.Println(b)
}

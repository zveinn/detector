package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func watchShells() {
	var lastOut []byte
	for {
		time.Sleep(1 * time.Second)

		out, err := exec.Command("bash", "-c", "ps aux | awk '{print $7}' | grep -v ?").Output()
		if err != nil {
			log.Fatal(err)
		}

		if string(lastOut) != string(out) {
			log.Println("CHANGE 111")
			fmt.Println(string(out))
		}

		lastOut = out

	}
}

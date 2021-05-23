package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func scanBusses() {

	var lastOut []byte
	for {
		time.Sleep(1 * time.Second)

		out, err := exec.Command("lsusb", "-t").Output()
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

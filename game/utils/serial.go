package main

import (
	"bufio"
	"log"
	"strings"
	"time"

	"go.bug.st/serial"
)

func main() {
	PORT := "/dev/ttyUSB0"
	RATE := 9600

	mode := &serial.Mode{
		BaudRate: RATE,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(PORT, mode)

	if err != nil {
		panic(err)
	}
	defer port.Close()

	log.Printf("Opened serial: PORT = %s RATE = %d\n", PORT, RATE)

	reader := bufio.NewReader(port)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error in reading: %v\n", err)
			time.Sleep(5 * time.Second)
		}

		line = strings.TrimSpace(line)
		if line != "" {
			log.Printf("Received: %s\n", line)
		}

	}

}

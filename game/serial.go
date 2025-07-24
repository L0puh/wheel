package main

import (
	"bufio"
	"log"
	"strings"
	"time"

	"go.bug.st/serial"
)

type Serial struct {
	port   string
	rate   int
	mode   *serial.Mode
	reader *bufio.Reader
	socket serial.Port
}

func open_serial() Serial {
	var s Serial
	s.port = "/dev/ttyUSB0"
	s.rate = 9600

	s.mode = &serial.Mode{
		BaudRate: s.rate,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(s.port, s.mode)

	if err != nil {
		panic(err)
	}

	log.Printf("Opened serial: PORT = %s RATE = %d\n", s.port, s.rate)

	s.reader = bufio.NewReader(port)
	s.socket = port

	return s
}

func receive_from_serial(s Serial) string {

	// for {
	line, err := s.reader.ReadString('\n')
	if err != nil {
		log.Printf("Error in reading: %v\n", err)
		time.Sleep(5 * time.Second)
	}

	line = strings.TrimSpace(line)
	if line != "" && len(line) > 10 {
		// log.Printf("Received: %s\n", line)
		return line
	}

	return ""

	// }

}

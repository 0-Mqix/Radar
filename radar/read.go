package radar

import (
	"log"
	"strconv"

	"github.com/tarm/serial"
)

type Data struct {
	Angle    int
	Distance int
}

func Read(s *serial.Port, c chan Data) {
	last := byte(0)
	angle := 0
	distance := 0
	afterSplit := false
	data := ""

	for {
		buf := make([]byte, 1)
		_, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		char := buf[0]

		switch char {
		case '\r':
			continue

		case '\n':
			if data != "" {
				distance, _ = strconv.Atoi(data)
				c <- Data{Distance: distance, Angle: angle}
			}

			last = byte(0)
			angle = 0
			afterSplit = false
			data = ""
		}

		switch last {
		case '\n':
			angle = int(char)

		case ':':
			afterSplit = true
		}

		if afterSplit {
			data += string(char)
		}

		last = char
	}
}

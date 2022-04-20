package radar

import (
	"strconv"

	"github.com/tarm/serial"
)

type Data struct {
	Angle    int
	Distance int
}

/*
Read()
s: serial port were the direct data from the audrino comes in
c: channel to output the correct data

it looks at the current and the last byte wat to do with the
data and were to store it and when its ready
*/
func Read(s *serial.Port, c chan Data, name string) {
	last := byte(0)
	angle := 0
	distance := 0
	afterSplit := false
	data := ""

	for {
		buf := make([]byte, 1)
		_, err := s.Read(buf)
		if err != nil {
			newSerial := Reconnect(name)
			s = &newSerial
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

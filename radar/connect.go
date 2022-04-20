package radar

import (
	"fmt"

	"github.com/tarm/serial"
)

//a recursive function to force the user to enter a working a port and it returns a serial.Port and the name of the port
func TryPort() (serial.Port, string) {
	fmt.Println("enter serial port:")

	name := ""
	fmt.Scanln(&name)

	c := &serial.Config{Name: name, Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
		return TryPort()
	}
	return *s, name
}

//a its a blocking function inside the Read() function, to reconnect the adruino if its lost
func Reconnect(name string) serial.Port {
	for {
		c := &serial.Config{Name: name, Baud: 115200}
		s, err := serial.OpenPort(c)
		if err == nil {
			return *s
		}
	}
}

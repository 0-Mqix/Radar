package main

import (
	"log"

	"github.com/MqixSchool/MonkeSockets"
	"github.com/MqixSchool/radar/radar"
	"github.com/labstack/echo/v4"
	"github.com/tarm/serial"
)

func main() {
	e := echo.New()
	ws := MonkeSockets.NewRoom()

	e.Static("/static", "./web/public")
	e.File("/", "./web/index.html")

	c := &serial.Config{Name: "COM5", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	out := make(chan radar.Data)
	go radar.Read(s, out)
	go radar.Print(out, ws)

	go ws.Run()

	e.GET("/ws", ws.WebSocket)

	e.Logger.Fatal(e.Start(":8080"))
}

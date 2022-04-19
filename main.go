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

	c := &serial.Config{Name: "COM3", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	/*
		channel is basicly a tunnel between diffrent processes
		so if i pass it to radar.Read() i can use "channel<-data" to send data to that "tunnel"
		and if i pass it to radar.CalcCircle() i can use "data := <-channel" to use the the data
	*/

	out := make(chan radar.Data)

	go radar.Read(s, out)
	// go radar.PrintCircle(out, ws)
	// go radar.PrintSqaure(out, ws)
	go radar.PrintBoth(out, ws)
	// ws.On("join:", func(r *MonkeSockets.Room, c *MonkeSockets.Client, b []byte) {
	// 	for i := 0; i < 72; i++ {
	// 		radar.Debug(i*5, 100, ws)
	// 	}
	// })

	go ws.Run()

	e.GET("/ws", ws.WebSocket)

	e.Logger.Fatal(e.Start(":8080"))
}

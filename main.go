package main

import (
	"github.com/MqixSchool/MonkeSockets"
	"github.com/MqixSchool/radar/radar"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()              //new web application object
	ws := MonkeSockets.NewRoom() //new websocket room

	e.Static("/static", "./web/public")
	e.File("/", "./web/index.html")

	s, name := radar.TryPort()

	// channel is basicly a tunnel between diffrent processes
	// so if i pass it to radar.Read() i can use "channel<-data" to send data to that "tunnel"
	// and if i pass it to radar.CalcCircle() i can use "data := <-channel" to use the the data
	out := make(chan radar.Data)

	go radar.Read(&s, out, name)
	// go radar.PrintCircle(out, ws)
	// go radar.PrintSqaure(out, ws)
	go radar.PrintBoth(out, ws)

	go ws.Run()

	e.GET("/ws", ws.WebSocket)

	e.Logger.Fatal(e.Start(":8080"))
}

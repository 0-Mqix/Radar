package radar

import (
	"encoding/json"
	"math"

	"github.com/MqixSchool/MonkeSockets"
	"github.com/MqixSchool/radar/functions"
)

type Point struct {
	Angle int     `json:"angle"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}

/*
PrintCircle()
out: channel for serial data
ws: MonkeSocket room object

uses the out channel and turns the angle and the distance into cordinates
for a circle and broadcasts it to the MonkeSocket Room
*/
func PrintCircle(out chan Data, ws *MonkeSockets.Room) {
	for {
		data := <-out
		l := float64(data.Distance)
		t := (float64(data.Angle) - 90) * math.Pi / 180

		x := math.Round(l * math.Cos(t))
		y := math.Round(l * math.Sin(t))

		out, _ := json.Marshal(Point{Angle: data.Angle, X: x, Y: y})
		ws.Broadcast("pointCircle:", out)
	}
}

/*
PrintSqaure()
out: channel for serial data
ws: MonkeSocket room object

uses the out channel and turns the angle and the distance into cordinates
for a square and broadcasts it to the MonkeSocket Room
*/
func PrintSqaure(out chan Data, ws *MonkeSockets.Room) {
	for {
		data := <-out
		r := float64(data.Distance)
		t := (float64(data.Angle) - 90) * math.Pi / 180

		x := r * math.Cos(t)
		y := r * math.Sin(t)

		xS, yS := CircleToSquare(data.Angle, x, y, r)

		out, _ := json.Marshal(Point{Angle: data.Angle, X: xS, Y: yS})
		ws.Broadcast("pointSquare:", out)
	}
}

/*
PrintBoth()
out: channel for serial data
ws: MonkeSocket room object

uses the out channel and turns the angle and the distance into cordinates
for a cricle and a square and broadcasts it to the MonkeSocket Room
*/
func PrintBoth(out chan Data, ws *MonkeSockets.Room) {
	for {
		data := <-out
		r := float64(data.Distance)
		t := (float64(data.Angle) - 90) * math.Pi / 180

		x := r * math.Cos(t)
		y := r * math.Sin(t)

		out, _ := json.Marshal(Point{Angle: data.Angle, X: x, Y: y})
		ws.Broadcast("pointCircle:", out)

		xS, yS := CircleToSquare(data.Angle, x, y, r)

		outS, _ := json.Marshal(Point{Angle: data.Angle, X: xS, Y: yS})
		ws.Broadcast("pointSquare:", outS)
	}
}

/*
CircleToSquare()
angle: the angle of the radar
x: x cordinates
y: y cordinates
r: radius

returns the transformed x, y

it transforms cordinates in a circle to cordinates in a square
by looking at the angle and using the corect formula
*/
func CircleToSquare(angle int, x, y, r float64) (float64, float64) {
	xS, yS := float64(0), float64(0)

	if functions.IsBetween(angle, 0, 45) {
		xS = x * -r / y
		yS = -r
	} else if functions.IsBetween(angle, 45, 135) {
		xS = r
		yS = y * r / x
	} else if functions.IsBetween(angle, 135, 225) {
		xS = x * r / y
		yS = r
	} else if functions.IsBetween(angle, 225, 315) {
		xS = -r
		yS = y * -r / x
	} else if functions.IsBetween(angle, 315, 360) {
		xS = x * -r / y
		yS = -r
	}

	return math.Round(xS), math.Round(yS)
}

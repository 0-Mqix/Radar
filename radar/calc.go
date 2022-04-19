package radar

import (
	"encoding/json"
	"math"

	"github.com/MqixSchool/MonkeSockets"
)

type Point struct {
	Angle int     `json:"angle"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}

func Print(out chan Data, ws *MonkeSockets.Room) {
	for {
		data := <-out
		l := float64(data.Distance)
		t := (float64(data.Angle) - 90) * math.Pi / 180

		x := math.Round(l * math.Cos(t))
		y := math.Round(l * math.Sin(t))

		out, _ := json.Marshal(Point{Angle: data.Angle, X: x, Y: y})
		ws.Broadcast("point:", out)
	}
}

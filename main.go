package main

import (
	"math"
	"time"

	r "github.com/go-vgo/robotgo"
)

const (
	// Use some sort of tool (in my case xdotool)
	// to get the position of the white, center dot
	// on the screen.
	X   = 1268
	Y   = 758
	RAD = 600
	INC = 100 // 3 gets the highest score, makes a triangle, makes no sense
)

func main() {
	r.Move(X+RAD, Y)
	time.Sleep(time.Second * 1)
	r.Toggle("down")
	r.Toggle("up")
	r.Toggle("down")

	for i := 0; i <= INC; i++ {
		theta := -float64(i) / float64(INC) * 2.0000 * math.Pi
		x := X + int(float64(RAD)*math.Cos(theta))
		y := Y + int(float64(RAD)*math.Sin(theta))
		r.Move(x, y)
		if x == 0 {
			r.Toggle("down")
		}
		time.Sleep(50 * time.Millisecond)
	}
}

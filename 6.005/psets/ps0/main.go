// https://github.com/holizz/terrapin
// 小学的时候机房小乌龟画图, 画一个四方形
package main

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type Terrapin struct {
	Image       *image.RGBA
	Pos         Position
	Orientation float64
	Color       color.Color
	Pen         bool
}

type Position struct {
	X, Y float64
}

func NewTerrapin(i *image.RGBA, starting Position) (t *Terrapin) {
	t = &Terrapin{
		Image:       i,
		Pos:         starting,
		Orientation: 0.0,
		Color:       color.White,
		Pen:         true,
	}

	return
}

func (t *Terrapin) Forward(dist float64) {
	for i := 0; i < int(dist); i++ {
		if t.Pen {
			t.Image.Set(int(t.Pos.X), int(t.Pos.Y), t.Color)
		}

		x := 1.0 * math.Sin(t.Orientation)
		y := 1.0 * -math.Cos(t.Orientation)

		t.Pos = Position{t.Pos.X + x, t.Pos.Y + y}
	}
}

func (t *Terrapin) Backward(dist float64) {
	t.Forward(-dist)
}

func (t *Terrapin) Right(radians float64) {
	t.Orientation += radians
}

func (t *Terrapin) Left(radians float64) {
	t.Right(-radians)
}

func (t *Terrapin) PenUp() {
	t.Pen = false
}

func (t *Terrapin) PenDown() {
	t.Pen = true
}

func main() {
	i := image.NewRGBA(image.Rect(0, 0, 300, 300))
	t := NewTerrapin(i, Position{150.0, 150.0})

	// t.Forward(100)
	// t.Right(math.Pi * 1 / 2)
	// t.Forward(100)
	// t.Right(math.Pi * 3 / 4)
	// t.Forward(math.Hypot(100, 100))

	// drawSquare
	t.Forward(100)
	t.Right(math.Pi * 1 / 2)
	t.Forward(100)
	t.Right(math.Pi * 1 / 2)
	t.Forward(100)
	t.Right(math.Pi * 1 / 2)
	t.Forward(100)

	outFile, _ := os.Create("p1.png")
	defer outFile.Close()

	buff := bufio.NewWriter(outFile)
	png.Encode(buff, i)
	buff.Flush()
}

package main

import (
	"log"
	"os"
	"math"
	"image"
	"image/color"
	"image/gif"
)

const (
	fileName = "rotatingCube.gif"
	width, height = 640, 640
	offset        = height / 2
)

func rotateCube(angleX, angleY float64) {
	sinX := math.Sin(angleX)
	cosX := math.Cos(angleX)
	sinY := math.Sin(angleY)
	cosY := math.Cos(angleY)
	for _, node := range nodes {
		x := node[0]
		y := node[1]
		z := node[2]
		node[0] = x*cosX - z*sinX
		node[2] = z*cosX + x*sinX
		z = node[2]
		node[1] = y*cosY - z*sinY
		node[2] = z*cosY + y*sinY
	}
}

var nodes = [][]float64{{-100, -100, -100}, {-100, -100, 100}, {-100, 100, -100}, {-100, 100, 100},
	{100, -100, -100}, {100, -100, 100}, {100, 100, -100}, {100, 100, 100}}
var edges = [][]int{{0, 1}, {1, 3}, {3, 2}, {2, 0}, {4, 5}, {5, 7}, {7, 6},
	{6, 4}, {0, 4}, {1, 5}, {2, 6}, {3, 7}}

func drawLine(x0, y0, x1, y1 int, img *image.Paletted, col color.RGBA) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	var sx, sy int = -1, -1
	if x0 < x1 {
		sx = 1
	}
	if y0 < y1 {
		sy = 1
	}
	err := dx - dy
	for {
		img.Set(x0, y0, col)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fgCol := color.RGBA{0xff, 0x00, 0xff, 0xff}
	var palette = []color.Color{color.RGBA{0x00, 0x00, 0x00, 0xff}, fgCol}
	var images []*image.Paletted
	var delays []int

	imgFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()
}
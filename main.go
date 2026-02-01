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

var nodes = [][]float64{{-100, -100, -100}, {-100, -100, 100}, {-100, 100, -100}, {-100, 100, 100},
	{100, -100, -100}, {100, -100, 100}, {100, 100, -100}, {100, 100, 100}}
var edges = [][]int{{0, 1}, {1, 3}, {3, 2}, {2, 0}, {4, 5}, {5, 7}, {7, 6},
	{6, 4}, {0, 4}, {1, 5}, {2, 6}, {3, 7}}


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
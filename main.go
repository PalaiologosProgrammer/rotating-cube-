package main

import (
	"log"
	"os"
)

const (
	fileName = "rotatingCube.gif"
)

var nodes = [][]float64{{-100, -100, -100}, {-100, -100, 100}, {-100, 100, -100}, {-100, 100, 100},
	{100, -100, -100}, {100, -100, 100}, {100, 100, -100}, {100, 100, 100}}
var edges = [][]int{{0, 1}, {1, 3}, {3, 2}, {2, 0}, {4, 5}, {5, 7}, {7, 6},
	{6, 4}, {0, 4}, {1, 5}, {2, 6}, {3, 7}}


func main() {
	imgFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()
}
package main

import (
	"log"
	"os"
)

const (
	fileName = "rotatingCube.gif"
)

func main() {
	imgFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()
}
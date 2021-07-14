package gonesis

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
)

func modLikePython(number, module int) int {
	res := number % module
	if (res < 0 && module > 0) || (res > 0 && module < 0) {
		return res + module
	}
	return res
}

func randomIntBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func drawFrame(terrain *Terrain, frameNumber int) {
	out, err := os.Create(fmt.Sprintf("./frames/%d.png", frameNumber))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cellzoom := 50
	m := image.NewRGBA(image.Rect(0, 0, terrain.LatitudesCount*cellzoom, terrain.LongitudesCount*cellzoom))
	organiccolor := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	//agentcolor := color.RGBA{R: 0, G: 0, B: 255, A: 255}
	emptycolor := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	draw.Draw(m, image.Rect(0, 0, terrain.LatitudesCount*cellzoom+cellzoom, terrain.LongitudesCount*cellzoom+cellzoom), &image.Uniform{C: organiccolor}, image.Point{}, draw.Src)

	for currentLatitude := 0; currentLatitude < terrain.LatitudesCount; currentLatitude++ {
		for currentLongitude := 0; currentLongitude < terrain.LongitudesCount; currentLongitude++ {
			if currentCell := terrain.GetCell(currentLatitude, currentLongitude); currentCell.CellType == Empty {
				draw.Draw(m, image.Rect(currentLatitude*cellzoom, currentLongitude*cellzoom, currentLatitude*cellzoom+cellzoom, currentLongitude*cellzoom+cellzoom), &image.Uniform{C: emptycolor}, image.Point{}, draw.Src)
			} else if currentCell.CellType == Locked {
				draw.Draw(m, image.Rect(currentLatitude*cellzoom, currentLongitude*cellzoom, currentLatitude*cellzoom+cellzoom, currentLongitude*cellzoom+cellzoom), &image.Uniform{C: color.RGBA{R: 0, G: 0, B: 255, A: uint8(modLikePython(255+currentCell.Agent.Generation*10, 255))}}, image.Point{}, draw.Src)
			} else if currentCell.CellType == Organic {
				draw.Draw(m, image.Rect(currentLatitude*cellzoom, currentLongitude*cellzoom, currentLatitude*cellzoom+cellzoom, currentLongitude*cellzoom+cellzoom), &image.Uniform{C: organiccolor}, image.Point{}, draw.Src)
			}
		}
	}

	err = png.Encode(out, m)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package gui

import (
	"fmt"
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/utils"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func DrawFrame(terrain contracts.ITerrain, frameNumber int) {

	out, err := os.Create(fmt.Sprintf("./frames/%d.png", frameNumber))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cellzoom := 100
	m := image.NewRGBA(image.Rect(0, 0, terrain.GetWidth()*cellzoom, terrain.GetHeight()*cellzoom))
	organiccolor := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	//agentcolor := color.RGBA{R: 0, G: 0, B: 255, A: 255}
	emptycolor := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	draw.Draw(m, image.Rect(0, 0, terrain.GetWidth()*cellzoom+cellzoom, terrain.GetHeight()*cellzoom+cellzoom), &image.Uniform{C: organiccolor}, image.Point{}, draw.Src)

	for currentLatitude := 0; currentLatitude < terrain.GetWidth(); currentLatitude++ {
		for currentLongitude := 0; currentLongitude < terrain.GetHeight(); currentLongitude++ {
			if currentCell := terrain.GetCell(currentLatitude, currentLongitude); currentCell.GetCellType() == contracts.EmptyCell {
				draw.Draw(m, image.Rect(currentLatitude*cellzoom, currentLongitude*cellzoom, currentLatitude*cellzoom+cellzoom, currentLongitude*cellzoom+cellzoom), &image.Uniform{C: emptycolor}, image.Point{}, draw.Src)
			} else if currentCell.GetCellType() == contracts.LockedCell {
				draw.Draw(m, image.Rect(currentLatitude*cellzoom, currentLongitude*cellzoom, currentLatitude*cellzoom+cellzoom, currentLongitude*cellzoom+cellzoom), &image.Uniform{C: color.RGBA{R: 0, G: 0, B: 255, A: uint8(utils.ModLikePython(255+currentCell.GetAgent().GetGeneration()*10, 255))}}, image.Point{}, draw.Src)
			} else if currentCell.GetCellType() == contracts.OrganicCell {
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

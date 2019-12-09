package sif

import (
	"fmt"
	"strings"
)

//Image represents a SIF image
type Image struct {
	Width, Height int
	Code          []rune
}

//GetLayers returns the list of individual layers in a SIF image
func (img Image) GetLayers() [][]rune {
	layerSize := img.Width * img.Height
	layers := make([][]rune, len(img.Code)/layerSize)
	for i := 0; i < len(layers); i++ {
		layers[i] = img.Code[i*layerSize : i*layerSize+layerSize]
	}
	return layers
}

//CombineLayers returns the single image when all layers have been overlayed
func (img Image) CombineLayers() []rune {
	layerSize := img.Width * img.Height
	combined := make([]rune, layerSize)
	layers := img.GetLayers()
	copy(combined, layers[0])

	//loop through the items in 'combined'
	for i, v := range combined {
		if v == rune('2') {
			//for every two, go through the layers until you find a value !=2
			for _, layer := range layers {
				if layer[i] != rune('2') {
					combined[i] = layer[i]
					break
				}
			}
		}
	}

	return combined
}

//Show prints the image to the screen
func (img Image) Show() {
	combined := img.CombineLayers()
	for i := 0; i < img.Height; i++ {
		fmt.Println(strings.ReplaceAll(string(combined[i*img.Width:i*img.Width+img.Width]), "0", " "))
	}
}

package sif

import (
	"testing"
)

func Test_BasicSif_GetLayers_IsAwesome(t *testing.T) {
	img := Image{Width: 3, Height: 2, Code: []rune("123456654321")}

	layers := img.GetLayers()

	if len(layers) != 2 {
		t.Errorf("Expected 2 layers, got %d", len(layers))
	}
}

func Test_BasicSif_CombineLayers_IsAwesome(t *testing.T) {
	img := Image{Width: 2, Height: 2, Code: []rune("0222112222120000")}
	combined := string(img.CombineLayers())

	if combined != "0110" {
		t.Errorf("Expected 0110, got %v", combined)
	}
}

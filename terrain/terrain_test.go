package terrain

import (
	"fmt"
	"testing"
)

func TestGetHeights(t *testing.T) {
	fmt.Println("Starting")
	heightsJPG := GetHeights(2, 2, 1, "helpers/heightmap16x16.jpg")
	heightsBMP := GetHeights(4, 4, 10, "helpers/heightmap16x16.bmp")
	if len(heightsJPG) != 3 || len(heightsBMP) != 5 {
		t.Error("Incorrect height array length")
	}

	//the top right pixel is white
	if heightsJPG[2][2] != 1 || heightsBMP[4][4] != 10 {
		t.Error("Heights incorrectly assigned")
	}
}

package terrain

import (
	"fmt"
	"testing"
)

func TestBuildTerrain(t *testing.T) {
	terrainWithoutHeights := BuildTerrain(2, 2, 1, nil)
	//We're going to eventually want the terrain to be built of triangle strips instead
	//or better yet... Give the user the option to chose how the terrain is built
	if len(terrainWithoutHeights.VertexIndices) != 24 {
		t.Error("Incorrect number of vertex indices")
	}
	if len(terrainWithoutHeights.VertexPositions) != 9 {
		t.Error("Incorrect number of vertices")
	}
}

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

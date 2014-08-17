package terrain

import (
	"fmt"
	"testing"
)

//This test assumes that the terrain is built with each square having its own unique 4 vertices
//This might not be the most efficient since it is possible to share vertices.
//Must look into how sharing vertices effects normals, textures, and colors
func TestBuildTerrain(t *testing.T) {
	terrainWithoutHeights := BuildTerrain(2, 2, 1, nil)

	if len(terrainWithoutHeights.VertexIndices) != 4*6 {
		t.Error("Incorrect number of vertex indices")
	}
	if len(terrainWithoutHeights.VertexPositions) != 16*3 {
		t.Error("Incorrect number of vertices")
	}
}

//Test whether we are able to generate a height array using an image
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

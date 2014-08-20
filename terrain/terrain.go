package terrain

import (
	"code.google.com/p/go.image/bmp"
	"encoding/json"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
)

type Terrain struct {
	VertexPositions []float32
	VertexIndices   []uint32
	//VertexColors    []float32
	//VertexNormals []float32
	//TextureCoords   []float32
}

//Generates a Terrain struct
func BuildTerrain(MAP_X int, MAP_Z int, tileSize float32, heights [][]float32) Terrain {
	vertexPositions := make([]float32, 12*MAP_X*MAP_Z)
	vertexIndices := make([]uint32, 6*MAP_X*MAP_Z)
	tileNum := 0
	//generate terrain vertex data
	//Refactor this to call external functions and be more readable
	for z := 0; z < MAP_Z; z++ {
		for x := 0; x < MAP_X; x++ {
			start := tileNum * 12
			vertexPositions[start], vertexPositions[start+9] = float32(x), float32(x)
			vertexPositions[start+1] = nilOrHeight(heights, x, z)
			vertexPositions[start+2], vertexPositions[start+5] = float32(-z), float32(-z)
			vertexPositions[start+3], vertexPositions[start+6] = float32(x+1), float32(x+1)
			vertexPositions[start+4] = nilOrHeight(heights, x+1, z)
			vertexPositions[start+7] = nilOrHeight(heights, x+1, z+1)
			vertexPositions[start+8], vertexPositions[start+11] = float32(-z-1), float32(-z-1)
			vertexPositions[start+10] = nilOrHeight(heights, x, z+1)

			start = tileNum * 6
			startIndex := uint32(tileNum * 4)
			vertexIndices[start], vertexIndices[start+3] = startIndex, startIndex
			vertexIndices[start+1] = startIndex + 1
			vertexIndices[start+2], vertexIndices[start+4] = startIndex+2, startIndex+2
			vertexIndices[start+5] = startIndex + 3

			tileNum++
		}
	}

	terrain := Terrain{vertexPositions, vertexIndices}

	return terrain
}

//Generate a height array from a source heightmap image
func GetHeights(MAP_X int, MAP_Z int, scale float32, filename string) [][]float32 {
	//load the heightmap image
	hmImageFile, err := os.Open(filename)
	if err != nil {
		panic("Error loading heightmap image file")
	}
	defer hmImageFile.Close()

	//decode the image data into the hmImage var. Using image.Decode in order to infer file type
	hmImage, _, err := image.Decode(hmImageFile)
	//circumvent "imported and not used".. There's probably a better way to do this ..
	_, _ = jpeg.Decode(hmImageFile)
	_, _ = bmp.Decode(hmImageFile)
	if err != nil {
		panic(err)
	}

	//determine image width and height
	bounds := hmImage.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	//Used to store pixel colors
	var r, g, b, a uint32
	var R, G, B, A float32

	heights := make([][]float32, MAP_X+1)

	//populate heights array
	for x := 0; x < MAP_X+1; x++ {
		heights[x] = make([]float32, MAP_Z+1)
		for y := 0; y < MAP_Z+1; y++ {
			color := hmImage.At(x*(width-1)/MAP_X, y*(height-1)/MAP_Z)
			r, g, b, a = color.RGBA()
			R, G, B, A = float32(r), float32(g), float32(b), float32(a)
			//we average the RGB then make it range from 0-1
			//alpha should always be the maximum value for any bit depth
			heights[x][y] = scale * (R + G + B) / 3 / A
		}
	}

	return heights
}

//used to test if a heightmap array was passed or not
func nilOrHeight(heights [][]float32, x int, z int) float32 {
	if heights == nil {
		return 0.0
	} else {
		return heights[x][z]
	}
}

//SaveTerrainFile saves a json representation of the terrain to disk
func (terrain *Terrain) SaveToFile(filename string) {
	terrainJSON, _ := json.Marshal(terrain)

	//append '.json' to the filename if not present
	if filename[len(filename)-5:] != ".json" {
		filename = filename + ".json"
	}

	ioutil.WriteFile(filename, terrainJSON, 0644)
}

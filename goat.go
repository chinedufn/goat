package main

import (
	"code.google.com/p/go.image/bmp"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strconv"
)

type Terrain struct {
	Heights         [][]float32
	VertexPositions []float32
	VertexColors    []float32
	VertexIndices   []uint32
	TextureCoords   []float32
}

//args -> inputFile-bitmap, desired-tile-width, desired-tile-height
func main() {
	if os.Args[1] == "terrain" {
		fmt.Println("Terrain command")
	} else {
		fmt.Println("Starting...")
		//load the heightmap image file
		hmImageFile, err := os.Open(os.Args[1])
		if err != nil {
			panic("Goodbye")
		}
		defer hmImageFile.Close()

		//decode the image data into hmImage
		hmImage, _, err := image.Decode(hmImageFile)
		_, _ = bmp.Decode(hmImageFile)
		if err != nil {
			panic(err)
		}

		//The width of the terrain in tiles
		MAP_X, err := strconv.Atoi(os.Args[2])
		//The height of the terrain in tiles
		MAP_Y, err := strconv.Atoi(os.Args[3])

		//determine image width and height
		bounds := hmImage.Bounds()
		width, height := bounds.Max.X, bounds.Max.Y

		//used to store pixel colors at different points on the heightmap image
		var r uint32
		var R float32

		heights := make([][]float32, MAP_Y+1)

		//generate heightmap array using the image
		for x := 0; x <= MAP_X; x++ {
			heights[x] = make([]float32, MAP_X+1)
			for y := 0; y <= MAP_Y; y++ {
				color := hmImage.At(x*width/MAP_X, y*height/MAP_Y)
				r, _, _, _ = color.RGBA()
				R = float32(r) / 65535
				heights[x][y] = 9 * R
			}
		}

		vertexPositions := make([]float32, 12*MAP_X*MAP_Y)
		vertexColors := make([]float32, 16*MAP_X*MAP_Y)
		vertexIndices := make([]uint32, 6*MAP_X*MAP_Y)
		textureCoords := make([]float32, 8*MAP_X*MAP_Y)
		tileNum := 0
		//generate terrain vertex data
		for z := 0; z < MAP_X; z++ {
			for x := 0; x < MAP_Y; x++ {
				start := tileNum * 12
				vertexPositions[start], vertexPositions[start+9] = float32(x), float32(x)
				vertexPositions[start+1] = heights[x][z]
				vertexPositions[start+2], vertexPositions[start+5] = float32(-z), float32(-z)
				vertexPositions[start+3], vertexPositions[start+6] = float32(x+1), float32(x+1)
				vertexPositions[start+4] = heights[x+1][z]
				vertexPositions[start+7] = heights[x+1][z+1]
				vertexPositions[start+8], vertexPositions[start+11] = float32(-z-1), float32(-z-1)
				vertexPositions[start+10] = heights[x][z+1]

				start = tileNum * 16
				for j := 0; j < 16; j++ {
					vertexColors[start+j] = 1
				}

				start = tileNum * 6
				startIndex := uint32(tileNum * 4)
				vertexIndices[start], vertexIndices[start+3] = startIndex, startIndex
				vertexIndices[start+1] = startIndex + 1
				vertexIndices[start+2], vertexIndices[start+4] = startIndex+2, startIndex+2
				vertexIndices[start+5] = startIndex + 3

				//texture coordinates
				start = tileNum * 8
				textureCoords[start], textureCoords[start+1] = 0.0, 0.0
				textureCoords[start+2], textureCoords[start+3] = 1.0, 0.0
				textureCoords[start+4], textureCoords[start+5] = 1.0, 1.0
				textureCoords[start+6], textureCoords[start+7] = 0.0, 1.0

				tileNum++
			}
		}

		harr := Terrain{heights, vertexPositions, vertexColors, vertexIndices, textureCoords}
		jsonHarr, _ := json.Marshal(harr)

		ioutil.WriteFile("out.json", jsonHarr, 0644)
	}
}

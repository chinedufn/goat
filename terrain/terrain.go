package terrain

import (
	"code.google.com/p/go.image/bmp"
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

type Terrain struct {
	Heights         [][]float32
	VertexPositions []float32
	VertexIndices   []int
	//VertexColors    []float32
	//VertexNormals []float32
	//TextureCoords   []float32
}

func BuildTerrain(MAP_X int, MAP_Z int, tileSize float32, heights [][]float32) *Terrain {
	t := &Terrain{}
	fmt.Println("building some terrain!")
	return t
}

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

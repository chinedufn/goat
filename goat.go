package main

import (
	"bufio"
	"fmt"
	"github.com/chinedufn/goat/convert"
	"github.com/chinedufn/goat/terrain"
	"os"
	"strconv"
)

func main() {
	if os.Args[1] == "terrain" {
		HandleTerrainArg()
	} else if os.Args[1] == "convert" {
		HandleConvertArg()
	}
}

//Handles the terrain argument
func HandleTerrainArg() {
	//if the user is specifying the arguments inline
	if len(os.Args) > 2 {
		//The width of the terrain in tiles
		//MAP_X, err := strconv.Atoi(os.Args[2])
		//The height of the terrain in tiles
		//MAP_Y, err := strconv.Atoi(os.Args[3])
	} else {
		MAP_X, _ := strconv.Atoi(GetInput("Terrain width in tiles [int] : "))
		MAP_Z, _ := strconv.Atoi(GetInput("Terrain depth in tiles [int] : "))
		heightmapFile := GetInput("Heightmap file name [string] : ")
		jsonFile := GetInput("JSON output file name [string] : ")

		var heights [][]float32
		if heightmapFile != "" {
			heights = terrain.GetHeights(MAP_X, MAP_Z, 1, heightmapFile)
		}
		ter := terrain.BuildTerrain(MAP_X, MAP_Z, 1, heights)
		ter.SaveToFile(jsonFile)
	}
}

func HandleConvertArg() {
	if len(os.Args) > 2 {
		//parameters are inline
	} else {
		inputFile := GetInput("3d model filename [string] : ")
		outputFile := GetInput("JSON output file name [string] : ")
		object := convert.LoadFromFile(inputFile)
		object.SaveToFile(outputFile)
	}
}

func GetInput(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s", message)
	scanner.Scan()
	return scanner.Text()
}

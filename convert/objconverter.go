package convert

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Object struct {
	VertexPositions []float32
	VertexIndices   []int
	//VertexNormals []float32
	//VertexColors []float32
	//TextureCoords []float32
}

//Loads a .obj file and stores the data in an Object struct
func LoadFromFile(filename string) *Object {
	//initialize data arrays
	VertexPositions := make([]float32, 0)
	VertexIndices := make([]int, 0)

	//load the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	//scan the file line by line and parse appropriately
	for scanner.Scan() {
		line = scanner.Text()
		fields := strings.Fields(line)
		//vertex positions data
		if fields[0] == "v" {
			for i := 1; i < 4; i++ {
				vPos, _ := strconv.ParseFloat(fields[i], 32)
				VertexPositions = append(VertexPositions, float32(vPos))
			}
		}
		//vertex indices data
		vIndices := make([]int, 6)
		if fields[0] == "f" {
			for i := 0; i < 4; i++ {
				//subtracting one from each index in order to zero index
				vIndices[i], _ = strconv.Atoi(fields[i+1])
				vIndices[i]--
			}
			vIndices[5] = vIndices[3]
			vIndices[3], vIndices[4] = vIndices[0], vIndices[2]
			VertexIndices = append(VertexIndices, vIndices...)
		}
	}

	obj := &Object{VertexPositions, VertexIndices}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return obj
}

//Saves an object to a json file
func (o *Object) SaveToFile(filename string) {
	objJson, _ := json.Marshal(o)
	ioutil.WriteFile(filename, objJson, 0644)
}

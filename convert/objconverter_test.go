package convert

import (
	"testing"
)

func TestLoadFromFile(t *testing.T) {
	obj := Object{}
	loadedObj := obj.LoadFromFile("helpers/cube.obj")

	//8 vertices, 3 data points per vertex
	if len(loadedObj.VertexPositions) != 24 {
		t.Error("Incorrect number of vertex positions")
	}
	if len(loadedObj.VertexIndices) != 36 {
		t.Error("Incorrect number of vertex indices")
	}
}

##Goat

Goat is an opinionated command line tool for webgl development. The api is currently very unstable.

#### Features

- Generate terrain from a heightmap image
- Convert wavefront (.obj) files to json for easy vertex parsing [feature in progress]
- To be continued...

#### To Install

Using Go

```
$ go get github.com/chinedufn/goat
$ cd $GOPATH/src/github.com/chinedufn/goat
$ go build
$ go install
```

#### Terrain

Use goat to generate your terrain from a heightmap

```
//provide a heightmap file, width in squares, height in squares
goat heightmap.bmp 64 64
```

Later in your javascript file...

```
//after loading the JSON file into an object...
var vertexPositions = loadedData.VertexPositions
var vertexIndicies = loadedData.VertexIndices
```

#### Note

The goat api is currently unstable and will see unpredictable changes as I use it more and familiarize myself with golang.

## License

MIT

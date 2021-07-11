package main

import (
	"fmt"
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/gfx"
)

func fatalUsage() {
	log.Fatalf("Usage: %s file.problem\n", os.Args[0])
}

func main() {
	log.Println("Hello ICFP Contest!")
	if len(os.Args) < 2 {
		fatalUsage()
	}

	pb := data.ParseProblem(os.Args[1])
	figCopy := pb.Figure.Copy()

	vis := gfx.NewVisualizer(pixelgl.WindowConfig{
		Title:  "Hello ICFP Contest!",
		Bounds: pixel.R(0, 0, 1000, 800),
	}, pb)

	vis.PushFigure(&figCopy, true, 2, false)
	vis.PushFigure(pb.Figure, false, 2, true)

	vis.OnVertexDrag = func(v *data.Vertex, mousePos pixel.Vec) {
		v.X = mousePos.X
		v.Y = mousePos.Y
	}

	vis.Start()
	fmt.Print(pb)
}

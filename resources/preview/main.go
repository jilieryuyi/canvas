package main

import (
	"github.com/jilieryuyi/canvas"
	"github.com/jilieryuyi/canvas/renderers"
)

func main() {
	c := canvas.New(200, 100)
	ctx := canvas.NewContext(c)
	if err := canvas.DrawPreview(ctx); err != nil {
		panic(err)
	}
	c.WriteFile("preview.png", renderers.PNG(canvas.DPMM(3.2)))
}

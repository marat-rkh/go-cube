package main

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
	"github.com/octomarat/go-cube/cube"
	"log"
)

type cubeApp struct {
	*application.Application
	cube cube.Cube
}

func create(ops application.Options) (*cubeApp, error) {
	app, err := application.Create(ops)
	if err != nil {
		return nil, err
	}
	return &cubeApp{Application: app}, nil
}

func (a *cubeApp) renderCube() {
	// TODO see mesh.SetMaterial
}

func main() {
	app, _ := create(application.Options{
		Title:  "Go Cube",
		Width:  800,
		Height: 600,
	})
	app.cube.AddToScene(app.Application)
	drawPlus(app.cube)
	app.Subscribe(application.OnBeforeRender, func(eventName string, event interface{}) {
		app.renderCube()
	})
	ambientLight := light.NewAmbient(&math32.Color{R: 1.0, G: 1.0, B: 1.0}, 0.8)
	app.Scene().Add(ambientLight)
	axis := graphic.NewAxisHelper(35)
	app.Scene().Add(axis)

	app.CameraPersp().SetPosition(0, 0, -200)
	app.CameraPersp().SetRotationY(math32.Pi)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func drawPlus(cb cube.Cube) {
	color := math32.NewColor("YellowGreen")
	cb.SetColor(1, 1, 0, color)
	cb.SetColor(1, 0, 1, color)
	cb.SetColor(0, 1, 1, color)
	cb.SetColor(1, 1, 1, color)
	cb.SetColor(2, 1, 1, color)
	cb.SetColor(1, 2, 1, color)
	cb.SetColor(1, 1, 2, color)
}

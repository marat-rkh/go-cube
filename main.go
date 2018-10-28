package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
	"log"
)

type cubeApp struct {
	*application.Application
	cube   *graphic.Mesh
	isBlue bool
}

func create(ops application.Options) (*cubeApp, error) {
	app, err := application.Create(ops)
	if err != nil {
		return nil, err
	}
	return &cubeApp{Application: app}, nil
}

func (a *cubeApp) addCubeToScene() {
	geom := geometry.NewCube(2)
	mat := material.NewPhong(math32.NewColor("DarkBlue"))
	a.cube = graphic.NewMesh(geom, mat)
	a.Application.Scene().Add(a.cube)
}

func (a *cubeApp) renderCube() {
	color := "DarkBlue"
	if a.isBlue {
		color = "Red"
		a.isBlue = false
	} else {
		a.isBlue = true
	}

	mat := material.NewPhong(math32.NewColor(color))
	a.cube.SetMaterial(mat)
}

func main() {
	app, _ := create(application.Options{
		Title:  "Go Cube",
		Width:  800,
		Height: 600,
	})
	app.addCubeToScene()
	app.Subscribe(application.OnBeforeRender, func(eventName string, event interface{}) {
		app.renderCube()
	})
	// Add lights to the scene
	ambientLight := light.NewAmbient(&math32.Color{R: 1.0, G: 1.0, B: 1.0}, 0.8)
	app.Scene().Add(ambientLight)
	pointLight := light.NewPoint(&math32.Color{R: 1, G: 1, B: 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	app.Scene().Add(pointLight)

	// Add an axis helper to the scene
	axis := graphic.NewAxisHelper(0.5)
	app.Scene().Add(axis)

	app.CameraPersp().SetPosition(0, 0, 3)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

package cube

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
)

const (
	cubeSize            = 3
	cubeSizePx  float32 = 100
	gapPx               = cubeSizePx / (cubeSize - 1)
	startPos            = -cubeSizePx / 2
	pixelSizePx         = 3
)

type Cube struct {
	pixels [][][]*graphic.Mesh
}

func (c Cube) SetColor(x, y, z int, color *math32.Color) {
	c.pixels[x][y][z].SetMaterial(material.NewPhong(color))
}

func (c Cube) AddToScene(app *application.Application) {
	c.pixels = make([][][]*graphic.Mesh, cubeSize)
	for i := 0; i < cubeSize; i++ {
		c.pixels[i] = make([][]*graphic.Mesh, cubeSize)
		xPos := startPos + float32(i)*gapPx
		for j := 0; j < cubeSize; j++ {
			c.pixels[i][j] = make([]*graphic.Mesh, cubeSize)
			yPos := startPos + float32(j)*gapPx
			for k := 0; k < cubeSize; k++ {
				zPos := startPos + float32(k)*gapPx
				geom := geometry.NewSphere(pixelSizePx, 32, 32, 0, 2*math32.Pi, 0, 2*math32.Pi)
				mat := material.NewPhong(math32.NewColor("DarkBlue"))
				mesh := graphic.NewMesh(geom, mat)
				mesh.SetPosition(xPos, yPos, zPos)
				app.Scene().Add(mesh)
				c.pixels[i][j][k] = mesh
			}
		}
	}
}

package cube

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
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

func (c *Cube) SetColor(x, y, z int, color *math32.Color) {
	c.pixels[x][y][z].SetMaterial(material.NewPhong(color))
}

func (c *Cube) AddToScene(app *application.Application) {
	c.addNodes(app)
	addEdges(app)
}

func (c *Cube) addNodes(app *application.Application) {
	c.pixels = make([][][]*graphic.Mesh, cubeSize)
	for i := 0; i < cubeSize; i++ {
		c.pixels[i] = make([][]*graphic.Mesh, cubeSize)
		xPos := startPos + float32(i)*gapPx
		for j := 0; j < cubeSize; j++ {
			c.pixels[i][j] = make([]*graphic.Mesh, cubeSize)
			yPos := startPos + float32(j)*gapPx
			for k := 0; k < cubeSize; k++ {
				zPos := startPos + float32(k)*gapPx
				sphere := geometry.NewSphere(pixelSizePx, 32, 32, 0, 2*math32.Pi, 0, 2*math32.Pi)
				mat := material.NewPhong(math32.NewColor("DimGrey"))
				mesh := graphic.NewMesh(sphere, mat)
				mesh.SetPosition(xPos, yPos, zPos)
				app.Scene().Add(mesh)
				c.pixels[i][j][k] = mesh
			}
		}
	}
}

func addEdges(app *application.Application) {
	vertices := math32.NewArrayF32(0, cubeSize*cubeSize*2)
	for i := 0; i < cubeSize; i++ {
		xPos := startPos + float32(i)*gapPx
		for j := 0; j < cubeSize; j++ {
			yPos := startPos + float32(j)*gapPx
			for k := 0; k < cubeSize; k++ {
				zPos := startPos + float32(k)*gapPx
				if i == 0 {
					vertices.AppendVector3(&math32.Vector3{X: xPos, Y: yPos, Z: zPos})
					vertices.AppendVector3(&math32.Vector3{X: xPos + cubeSizePx, Y: yPos, Z: zPos})
				}
				if j == 0 {
					vertices.AppendVector3(&math32.Vector3{X: xPos, Y: yPos, Z: zPos})
					vertices.AppendVector3(&math32.Vector3{X: xPos, Y: yPos + cubeSizePx, Z: zPos})
				}
				if k == 0 {
					vertices.AppendVector3(&math32.Vector3{X: xPos, Y: yPos, Z: zPos})
					vertices.AppendVector3(&math32.Vector3{X: xPos, Y: yPos, Z: zPos + cubeSizePx})
				}
			}
		}
	}
	geom := geometry.NewGeometry()
	geom.AddVBO(gls.NewVBO(vertices).AddAttrib(gls.VertexPosition))

	mat := material.NewStandard(math32.NewColor("Black"))
	lines := graphic.NewLines(geom, mat)
	app.Scene().Add(lines)
}

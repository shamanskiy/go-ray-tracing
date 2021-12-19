package render

import (
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/chewxy/math32"
	"github.com/schollz/progressbar/v3"
)

type Camera struct {
	Origin            core.Vec3
	Upper_left_corner core.Vec3
	Horizontal        core.Vec3
	Vertical          core.Vec3
}

/*func NewCamera() camera {
	return camera{
		origin:            Vec3{0.0, 0.0, 0.0},
		upper_left_corner: Vec3{-2.0, 1.0, -1.0},
		horizontal:        Vec3{4.0, 0.0, 0.0},
		vertical:          Vec3{0.0, -2.0, 0.0}}
}*/

func (c *Camera) Render(scene *Scene) *image.RGBA {
	width := 800
	height := 400
	sampling := 4

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	bar := progressbar.Default(int64(width), "rendering")

	start := time.Now()
	for x := 0; x < width; x++ {
		bar.Add(1)
		for y := 0; y < height; y++ {
			var clr core.Color
			for s := 0; s < sampling; s++ {
				u := (core.Real(x) + rand.Float32()) / core.Real(width)
				v := (core.Real(y) + rand.Float32()) / core.Real(height)
				ray := c.GetRay(u, v)
				c := scene.TestRay(ray)
				clr = clr.Add(c)
			}
			clr = clr.Mul(1.0 / core.Real(sampling))
			img.Set(x, y, toRGBA(clr))
		}
	}
	elapsed := time.Since(start)
	log.Printf("Rendering took %s", elapsed)

	return img
}

func toRGBA(c core.Color) color.RGBA {
	return color.RGBA{toZero255(c.X()), toZero255(c.Y()), toZero255(c.Z()), 0xff}
}

func toZero255(x core.Real) uint8 {
	return uint8(math32.Floor(255.99 * math32.Sqrt(x)))
}

func (c *Camera) GetRay(u, v core.Real) core.Ray {
	ray := core.Ray{
		Origin:    c.Origin,
		Direction: c.Upper_left_corner.Add(c.Horizontal.Mul(u)).Add(c.Vertical.Mul(v))}

	return ray
}

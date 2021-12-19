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

type CameraSettings struct {
	VerticalFOV      core.Real
	ImagePixelWidth  int
	ImagePixelHeight int

	LookFrom core.Vec3
	LookAt   core.Vec3
	GlobalUp core.Vec3

	Antialiasing      int
	MaxRayReflections int

	//Verbosity verbosity{ Verbosity::none };
	//float lensRadius{0.0};
}

func DefaultCameraSettings() CameraSettings {
	return CameraSettings{
		VerticalFOV:       90,
		ImagePixelWidth:   800,
		ImagePixelHeight:  400,
		LookFrom:          core.Vec3{0., 0., 0.},
		LookAt:            core.Vec3{0., 0., -1.},
		GlobalUp:          core.Vec3{0., 1., 0.},
		Antialiasing:      4,
		MaxRayReflections: 10,
	}
}

type Camera struct {
	origin          core.Vec3
	upperLeftCorner core.Vec3
	horizontalSpan  core.Vec3
	verticalSpan    core.Vec3
	settings        CameraSettings
}

func (c *Camera) initialize() {
	c.origin = c.settings.LookFrom
	verticalFOV_rad := c.settings.VerticalFOV * math32.Pi / 180
	aspectRatio := core.Real(c.settings.ImagePixelWidth) / core.Real(c.settings.ImagePixelHeight)

	halfHeight := math32.Tan(verticalFOV_rad / 2)
	halfWidth := aspectRatio * halfHeight

	back := c.settings.LookFrom.Sub(c.settings.LookAt).Normalize()
	right := c.settings.GlobalUp.Cross(back).Normalize()
	up := back.Cross(right)

	focusDistance := c.settings.LookFrom.Sub(c.settings.LookAt).Len()

	c.horizontalSpan = right.Mul(2 * halfWidth * focusDistance)
	c.verticalSpan = up.Mul(-2 * halfHeight * focusDistance)

	originToCorner := up.Mul(halfHeight).Sub(right.Mul(halfWidth)).Sub(back).Mul(focusDistance)
	c.upperLeftCorner = c.origin.Add(originToCorner)
}

func NewCamera(settings CameraSettings) *Camera {
	camera := Camera{settings: settings}
	camera.initialize()
	return &camera
}

func (c *Camera) Render(scene *Scene) *image.RGBA {
	width := c.settings.ImagePixelWidth
	height := c.settings.ImagePixelHeight

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	bar := progressbar.Default(int64(width), "rendering")

	start := time.Now()
	for x := 0; x < width; x++ {
		bar.Add(1)
		for y := 0; y < height; y++ {
			var clr core.Color
			for s := 0; s < c.settings.Antialiasing; s++ {
				u := (core.Real(x) + rand.Float32()) / core.Real(width)
				v := (core.Real(y) + rand.Float32()) / core.Real(height)
				ray := c.GetRay(u, v)
				c := scene.TestRay(ray)
				clr = clr.Add(c)
			}
			clr = clr.Mul(1.0 / core.Real(c.settings.Antialiasing))
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
		Origin:    c.origin,
		Direction: c.upperLeftCorner.Add(c.horizontalSpan.Mul(u)).Add(c.verticalSpan.Mul(v))}

	return ray
}

package render

import (
	"image"
	"log"
	"time"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/chewxy/math32"
	"github.com/schollz/progressbar/v3"
)

type CameraSettings struct {
	VerticalFOV      core.Real
	AspectRatio      core.Real
	ImagePixelHeight int

	LookFrom core.Vec3
	LookAt   core.Vec3
	GlobalUp core.Vec3

	Antialiasing      int
	MaxRayReflections int

	//float lensRadius{0.0};
}

func DefaultCameraSettings() CameraSettings {
	return CameraSettings{
		VerticalFOV:       90,
		AspectRatio:       2.,
		ImagePixelHeight:  360,
		LookFrom:          core.NewVec3(0., 0., 0.),
		LookAt:            core.NewVec3(0., 0., -1.),
		GlobalUp:          core.NewVec3(0., 1., 0.),
		Antialiasing:      4,
		MaxRayReflections: 10,
	}
}

type Camera struct {
	Origin          core.Vec3
	upperLeftCorner core.Vec3
	horizontalSpan  core.Vec3
	verticalSpan    core.Vec3

	PixelWidth     int
	PixelHeight    int
	sampling       int
	maxReflections int
}

func NewCamera(settings *CameraSettings) *Camera {
	camera := Camera{}

	camera.PixelHeight = settings.ImagePixelHeight
	camera.PixelWidth = int(core.Real(settings.ImagePixelHeight) * settings.AspectRatio)
	camera.sampling = settings.Antialiasing
	camera.maxReflections = settings.MaxRayReflections

	camera.Origin = settings.LookFrom
	verticalFOV_rad := settings.VerticalFOV * math32.Pi / 180

	halfHeight := math32.Tan(verticalFOV_rad / 2)
	halfWidth := settings.AspectRatio * halfHeight

	back := settings.LookFrom.Sub(settings.LookAt).Normalize()
	right := settings.GlobalUp.Cross(back).Normalize()
	up := back.Cross(right)

	focusDistance := settings.LookFrom.Sub(settings.LookAt).Len()

	camera.horizontalSpan = right.Mul(2 * halfWidth * focusDistance)
	camera.verticalSpan = up.Mul(-2 * halfHeight * focusDistance)

	originToCorner := up.Mul(halfHeight).Sub(right.Mul(halfWidth)).Sub(back).Mul(focusDistance)
	camera.upperLeftCorner = camera.Origin.Add(originToCorner)

	return &camera
}

func (c *Camera) createImage() *image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{c.PixelWidth, c.PixelHeight}

	return image.NewRGBA(image.Rectangle{upLeft, lowRight})
}

func (c *Camera) createProgressBar() *progressbar.ProgressBar {
	return progressbar.Default(int64(c.PixelWidth), "rendering")
}

func (c *Camera) IndexToU(index int) core.Real {
	return (core.Real(index) + core.Random().From01()) / core.Real(c.PixelWidth)
}

func (c *Camera) IndexToV(index int) core.Real {
	return (core.Real(index) + core.Random().From01()) / core.Real(c.PixelHeight)
}

func (c *Camera) Render(scene *Scene) *image.RGBA {
	img := c.createImage()
	bar := c.createProgressBar()

	start := time.Now()
	for x := 0; x < c.PixelWidth; x++ {
		bar.Add(1)
		for y := 0; y < c.PixelHeight; y++ {
			var pixelColor color.Color
			for s := 0; s < c.sampling; s++ {
				u := c.IndexToU(x)
				v := c.IndexToV(y)
				ray := c.GetRay(u, v)
				rayColor := scene.TestRay(ray)
				pixelColor = pixelColor.Add(rayColor)
			}
			pixelColor = pixelColor.Div(core.Real(c.sampling))
			img.Set(x, y, pixelColor.ToRGBA())
		}
	}
	elapsed := time.Since(start)
	log.Printf("Rendering took %s", elapsed)

	return img
}

func (c *Camera) GetRay(u, v core.Real) core.Ray {
	ray := core.Ray{
		Origin:    c.Origin,
		Direction: c.upperLeftCorner.Add(c.horizontalSpan.Mul(u)).Add(c.verticalSpan.Mul(v)).Sub(c.Origin)}

	return ray
}

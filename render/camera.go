package render

import (
	"image"
	"image/color"
	"log"
	"time"

	"github.com/Shamanskiy/go-ray-tracer/core"
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

	pixelWidth     int
	pixelHeight    int
	sampling       int
	maxReflections int
}

func NewCamera(settings *CameraSettings) *Camera {
	camera := Camera{}

	camera.pixelHeight = settings.ImagePixelHeight
	camera.pixelWidth = int(core.Real(settings.ImagePixelHeight) * settings.AspectRatio)
	camera.sampling = settings.Antialiasing
	camera.maxReflections = settings.MaxRayReflections

	camera.origin = settings.LookFrom
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
	camera.upperLeftCorner = camera.origin.Add(originToCorner)

	return &camera
}

func (c *Camera) createImage() *image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{c.pixelWidth, c.pixelHeight}

	return image.NewRGBA(image.Rectangle{upLeft, lowRight})
}

func (c *Camera) createProgressBar() *progressbar.ProgressBar {
	return progressbar.Default(int64(c.pixelWidth), "rendering")
}

func (c *Camera) indexToU(index int) core.Real {
	return (core.Real(index) + core.Random().From01()) / core.Real(c.pixelWidth)
}

func (c *Camera) indexToV(index int) core.Real {
	return (core.Real(index) + core.Random().From01()) / core.Real(c.pixelHeight)
}

func (c *Camera) Render(scene *Scene) *image.RGBA {
	img := c.createImage()
	bar := c.createProgressBar()

	start := time.Now()
	for x := 0; x < c.pixelWidth; x++ {
		bar.Add(1)
		for y := 0; y < c.pixelHeight; y++ {
			var pixelColor core.Color
			for s := 0; s < c.sampling; s++ {
				u := c.indexToU(x)
				v := c.indexToV(y)
				ray := c.GetRay(u, v)
				rayColor := scene.TestRay(ray)
				pixelColor = pixelColor.Add(rayColor)
			}
			pixelColor = core.Div(pixelColor, core.Real(c.sampling))
			img.Set(x, y, toRGBA(pixelColor))
		}
	}
	elapsed := time.Since(start)
	log.Printf("Rendering took %s", elapsed)

	return img
}

func toRGBA(c core.Color) color.RGBA {
	return color.RGBA{toZero255(c.X()), toZero255(c.Y()), toZero255(c.Z()), 255}
}

func toZero255(x core.Real) uint8 {
	return uint8(math32.Floor(255.99 * gammaCorrection(x)))
}

func gammaCorrection(input core.Real) core.Real {
	return math32.Sqrt(input)
}

func (c *Camera) GetRay(u, v core.Real) core.Ray {
	ray := core.Ray{
		Origin:    c.origin,
		Direction: c.upperLeftCorner.Add(c.horizontalSpan.Mul(u)).Add(c.verticalSpan.Mul(v)).Sub(c.origin)}

	return ray
}

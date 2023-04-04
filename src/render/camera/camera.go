package camera

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/render/image"
	"github.com/Shamanskiy/go-ray-tracer/src/render/log"
	"github.com/Shamanskiy/go-ray-tracer/src/render/scene"
	"github.com/chewxy/math32"
)

type Camera struct {
	Origin          core.Vec3
	upperLeftCorner core.Vec3
	horizontalSpan  core.Vec3
	verticalSpan    core.Vec3

	PixelWidth  int
	PixelHeight int
	sampling    int

	randomizer   random.RandomGenerator
	progressChan chan<- int
}

func NewCamera(settings *CameraSettings, randomizer random.RandomGenerator) *Camera {
	camera := Camera{randomizer: randomizer}

	camera.PixelHeight = settings.ImagePixelHeight
	camera.PixelWidth = int(core.Real(settings.ImagePixelHeight) * settings.AspectRatio)
	camera.sampling = settings.Antialiasing

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

	camera.progressChan = settings.ProgressChan

	return &camera
}

func (c *Camera) IndexToU(index int) core.Real {
	return (core.Real(index) + c.randomizer.Real()) / core.Real(c.PixelWidth)
}

func (c *Camera) IndexToV(index int) core.Real {
	return (core.Real(index) + c.randomizer.Real()) / core.Real(c.PixelHeight)
}

func (c *Camera) Render(scene scene.Scene) *image.Image {
	defer log.TimeExecution("rendering")()
	img := image.NewImage(c.PixelWidth, c.PixelHeight)

	for x := 0; x < c.PixelWidth; x++ {
		c.reportProgress(x+1, c.PixelWidth)

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
			img.SetPixelColor(x, y, pixelColor)
		}
	}

	return img
}

func (c *Camera) GetRay(u, v core.Real) core.Ray {
	rayDirection := c.upperLeftCorner.Add(c.horizontalSpan.Mul(u)).Add(c.verticalSpan.Mul(v)).Sub(c.Origin)
	ray := core.NewRay(c.Origin, rayDirection)

	return ray
}

func (c *Camera) reportProgress(currentColumn, imageWith int) {
	if c.progressChan == nil {
		return
	}

	progress := int(float64(currentColumn) / float64(imageWith) * 100)
	c.progressChan <- progress

	if currentColumn == imageWith {
		close(c.progressChan)
	}
}

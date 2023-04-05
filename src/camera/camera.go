package camera

import (
	"github.com/Shamanskiy/go-ray-tracer/src/camera/image"
	"github.com/Shamanskiy/go-ray-tracer/src/camera/log"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/scene"
)

type Camera struct {
	rayGenerator *RayGenerator

	// image + uv generation
	PixelWidth  int
	PixelHeight int

	sampling int

	randomizer   random.RandomGenerator
	progressChan chan<- int
}

func NewCamera(settings *CameraSettings, randomizer random.RandomGenerator) *Camera {
	camera := Camera{randomizer: randomizer}
	camera.rayGenerator = NewRayGenerator(settings.LookFrom, settings.LookAt, settings.VerticalFOV, settings.AspectRatio)

	camera.PixelHeight = settings.ImagePixelHeight
	camera.PixelWidth = int(core.Real(settings.ImagePixelHeight) * settings.AspectRatio)
	camera.sampling = settings.Antialiasing

	camera.progressChan = settings.ProgressChan

	return &camera
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
				ray := c.rayGenerator.GenerateRay(u, v)
				rayColor := scene.TestRay(ray)
				pixelColor = pixelColor.Add(rayColor)
			}
			pixelColor = pixelColor.Div(core.Real(c.sampling))
			img.SetPixelColor(x, y, pixelColor)
		}
	}

	return img
}

func (c *Camera) IndexToU(index int) core.Real {
	return (core.Real(index) + c.randomizer.Real()) / core.Real(c.PixelWidth)
}

func (c *Camera) IndexToV(index int) core.Real {
	return (core.Real(index) + c.randomizer.Real()) / core.Real(c.PixelHeight)
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

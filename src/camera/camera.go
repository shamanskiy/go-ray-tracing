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
	image        *image.Image
	randomizer   random.RandomGenerator
	progressChan chan<- int
	sampling     int
}

func NewCamera(settings *CameraSettings, randomizer random.RandomGenerator) *Camera {
	imageWidth := int(core.Real(settings.ImagePixelHeight) * settings.AspectRatio)

	return &Camera{
		rayGenerator: NewRayGenerator(settings.LookFrom, settings.LookAt, settings.VerticalFOV, settings.AspectRatio),
		image:        image.NewImage(imageWidth, settings.ImagePixelHeight),
		randomizer:   randomizer,
		progressChan: settings.ProgressChan,
		sampling:     settings.Antialiasing,
	}
}

func (c *Camera) Render(scene scene.Scene) *image.Image {
	defer log.TimeExecution("rendering")()

	for x := 0; x < c.image.Width(); x++ {
		c.reportProgress(x+1, c.image.Width())
		for y := 0; y < c.image.Height(); y++ {
			pixelColor := c.samplePixel(x, y, scene)
			c.image.SetPixelColor(x, y, pixelColor)
		}
	}

	return c.image
}

func (c *Camera) samplePixel(x, y int, scene scene.Scene) color.Color {
	var pixelColor color.Color
	for s := 0; s < c.sampling; s++ {
		u := c.toParam(x, c.image.Width())
		v := c.toParam(y, c.image.Height())
		ray := c.rayGenerator.GenerateRay(u, v)
		rayColor := scene.TestRay(ray)
		pixelColor = pixelColor.Add(rayColor)
	}
	return pixelColor.Div(core.Real(c.sampling))
}

func (c *Camera) toParam(index, maxIndex int) core.Real {
	return (core.Real(index) + c.randomizer.Real()) / core.Real(maxIndex)
}

func (c *Camera) reportProgress(done, total int) {
	if c.progressChan == nil {
		return
	}

	progress := int(float64(done) / float64(total) * 100)
	c.progressChan <- progress

	if done == total {
		close(c.progressChan)
	}
}

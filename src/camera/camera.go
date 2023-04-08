package camera

import (
	"fmt"
	"sync"

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
	progressChan chan<- log.ProgressUpdate
	sampling     int
}

type CameraSettings struct {
	VerticalFOV      core.Real
	AspectRatio      core.Real
	ImagePixelHeight int

	LookFrom core.Vec3
	LookAt   core.Vec3

	Antialiasing int
	ProgressChan chan<- log.ProgressUpdate
}

func NewCamera(settings *CameraSettings, randomizer random.RandomGenerator) *Camera {
	validateSettings(settings)

	imageWidth := int(core.Real(settings.ImagePixelHeight) * settings.AspectRatio)
	if settings.ImagePixelHeight <= 0 {
		panic(fmt.Errorf("new camera: invalid image pixel width: %v", imageWidth))
	}

	return &Camera{
		rayGenerator: NewRayGenerator(settings.LookFrom, settings.LookAt, settings.VerticalFOV, settings.AspectRatio),
		image:        image.NewImage(imageWidth, settings.ImagePixelHeight),
		randomizer:   randomizer,
		progressChan: settings.ProgressChan,
		sampling:     settings.Antialiasing,
	}
}

func validateSettings(settings *CameraSettings) {
	if settings.VerticalFOV <= 0 {
		panic(fmt.Errorf("new camera: invalid vertical FOV: %v", settings.VerticalFOV))
	}
	if settings.AspectRatio <= 0 {
		panic(fmt.Errorf("new camera: invalid aspect ratio: %v", settings.AspectRatio))
	}
	if settings.ImagePixelHeight <= 0 {
		panic(fmt.Errorf("new camera: invalid image pixel height: %v", settings.ImagePixelHeight))
	}
	if settings.LookFrom == settings.LookAt {
		panic(fmt.Errorf("new camera: lookAt coincides with lookFrom: %v", settings.LookAt))
	}
	if settings.Antialiasing < 1 {
		panic(fmt.Errorf("new camera: invalid antialiasing: %d", settings.Antialiasing))
	}
}

func (c *Camera) Render(scene scene.Scene) *image.Image {
	defer log.TimeExecution("rendering")()

	numWorkers := 10
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(numWorkers)

	for worker := 0; worker < numWorkers; worker++ {
		go c.render(scene, worker, numWorkers, &waitGroup)
	}

	waitGroup.Wait()
	if c.progressChan != nil {
		close(c.progressChan)
	}

	return c.image
}

func (c *Camera) render(scene scene.Scene, worker, numWorkers int, waitGroup *sync.WaitGroup) {
	for x := worker; x < c.image.Width(); x += numWorkers {
		c.reportProgress(x)
		for y := 0; y < c.image.Height(); y++ {
			pixelColor := c.samplePixel(x, y, scene)
			c.image.SetPixelColor(x, y, pixelColor)
		}
	}
	waitGroup.Done()
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

func (c *Camera) reportProgress(currentImageColumn int) {
	if c.progressChan == nil {
		return
	}

	c.progressChan <- log.ProgressUpdate{Max: c.image.Width()}
}

package camera_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/camera"
	"github.com/Shamanskiy/go-ray-tracer/src/camera/image"
	"github.com/Shamanskiy/go-ray-tracer/src/camera/log"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/scene"
	"github.com/stretchr/testify/assert"
)

var cameraSettings = camera.CameraSettings{
	VerticalFOV:      90,
	AspectRatio:      2,
	ImagePixelHeight: 5,
	LookAt:           core.NewVec3(0, 0, -1),
	Antialiasing:     1,
	NumRenderThreads: 1,
}

var randomizer = random.NewRandomGenerator()

func TestCamera_ShouldColorAllPixelsSameColor_IfEmptySceneWithFlatColorBackground(t *testing.T) {
	camera := camera.NewCamera(&cameraSettings, randomizer)
	scene := scene.NewFakeScene(color.Red)

	image := camera.Render(scene)

	assert.Equal(t, cameraSettings.ImagePixelHeight, image.Height())
	assert.Equal(t, cameraSettings.ImagePixelHeight*2, image.Width())
	assertAllPixelsColor(t, image, color.Red)
}

func TestCamera_ShouldReportProgress(t *testing.T) {
	waitForConsumer := make(chan struct{})
	progressUpdates := &[]log.ProgressUpdate{}
	settings := cameraSettings
	settings.ProgressChan = makeProgressConsumer(progressUpdates, waitForConsumer)
	camera := camera.NewCamera(&settings, randomizer)
	scene := scene.NewFakeScene(color.Red)

	camera.Render(scene)
	<-waitForConsumer

	imageWidth := 5 * 2
	assert.Len(t, *progressUpdates, imageWidth)
}

func TestCamera_ShouldMultiSampleEachPixel(t *testing.T) {
	settings := cameraSettings
	settings.ImagePixelHeight = 1
	settings.AspectRatio = 1
	settings.Antialiasing = 10
	camera := camera.NewCamera(&settings, randomizer)
	scene := scene.NewFakeScene(color.Red)

	camera.Render(scene)

	assert.Len(t, scene.RecordedRays, 10)
}

func assertAllPixelsColor(t *testing.T, image *image.Image, color color.Color) {
	for x := 0; x < image.Width(); x++ {
		for y := 0; y < image.Height(); y++ {
			assert.Equal(t, color, image.PixelColor(x, y))
		}
	}
}

func assertCornerPixelsColor(t *testing.T, image *image.Image, color color.Color) {
	width := image.Width()
	height := image.Height()
	assert.Equal(t, color, image.PixelColor(0, 0))
	assert.Equal(t, color, image.PixelColor(width-1, 0))
	assert.Equal(t, color, image.PixelColor(0, height-1))
	assert.Equal(t, color, image.PixelColor(width-1, height-1))
}

func makeProgressConsumer(progressMessages *[]log.ProgressUpdate, done chan struct{}) chan log.ProgressUpdate {
	progressChan := make(chan log.ProgressUpdate)
	go func() {
		for progress := range progressChan {
			*progressMessages = append(*progressMessages, progress)
		}
		done <- struct{}{}
	}()
	return progressChan
}

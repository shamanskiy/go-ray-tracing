package camera_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/camera"
	"github.com/Shamanskiy/go-ray-tracer/src/camera/image"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/scene"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/background"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
	"github.com/stretchr/testify/assert"
)

var cameraSettings = camera.CameraSettings{
	VerticalFOV:      90,
	AspectRatio:      2,
	ImagePixelHeight: 5,
	LookAt:           core.NewVec3(0, 0, -1),
	Antialiasing:     1,
}

var randomizer = random.NewRandomGenerator()

func TestCamera_ShouldColorAllPixelsSameColor_IfEmptySceneWithFlatColorBackground(t *testing.T) {
	camera := camera.NewCamera(&cameraSettings, randomizer)
	scene := scene.New(background.NewFlatColor(color.Red))

	image := camera.Render(scene)

	assert.Equal(t, cameraSettings.ImagePixelHeight, image.Height())
	assert.Equal(t, cameraSettings.ImagePixelHeight*2, image.Width())
	assertAllPixelsColor(t, image, color.Red)
}

func TestCamera_ShouldColorCentralPixelWithObjectColor(t *testing.T) {
	camera := camera.NewCamera(&cameraSettings, randomizer)
	scene := scene.New(background.NewFlatColor(color.White))
	scene.Add(geometries.NewSphere(core.NewVec3(0, 0, -2), 1), materials.NewDiffusive(color.Red, randomizer))

	image := camera.Render(scene)

	assertCornerPixelsColor(t, image, color.White)
	assert.Equal(t, color.Red, image.PixelColor(5, 2))
}

func TestCamera_ShouldReportProgress(t *testing.T) {
	settings := cameraSettings
	progressRecording := &[]int{}
	settings.ProgressChan = makeProgressConsumer(progressRecording)
	camera := camera.NewCamera(&settings, randomizer)
	scene := scene.New(background.NewFlatColor(color.Red))

	camera.Render(scene)

	assert.NotEmpty(t, progressRecording)
	assertOrdered(t, *progressRecording)
	assertLastElem(t, *progressRecording, 100)
}

func TestCamera_ShouldMultiSampleEachPixel(t *testing.T) {
	settings := cameraSettings
	settings.ImagePixelHeight = 1
	settings.AspectRatio = 1
	settings.Antialiasing = 10
	camera := camera.NewCamera(&settings, randomizer)
	scene := scene.NewFakeScene()

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

func makeProgressConsumer(progressValues *[]int) chan int {
	progressChan := make(chan int)
	go func() {
		for progress := range progressChan {
			*progressValues = append(*progressValues, progress)
		}
	}()
	return progressChan
}

func assertOrdered(t *testing.T, progressRecordging []int) {
	if len(progressRecordging) < 2 {
		return
	}

	for i := 0; i < len(progressRecordging)-1; i++ {
		assert.LessOrEqual(t, progressRecordging[i], progressRecordging[i+1])
	}
}

func assertLastElem(t *testing.T, progressRecordging []int, expectedValue int) {
	lastValue := progressRecordging[len(progressRecordging)-1]
	assert.Equal(t, expectedValue, lastValue)
}

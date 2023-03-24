package render

import (
	"image/color"
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/render"
	"github.com/Shamanskiy/go-ray-tracer/test"
)

func TestCamera_Default(t *testing.T) {
	t.Log("Default camera without randomness")
	settings := render.DefaultCameraSettings()
	camera := render.NewCamera(&settings)
	core.Random().Disable()
	defer core.Random().Enable()

	test.CheckResult(t, "Camera origin", camera.Origin, settings.LookFrom)

	ray := camera.GetRay(0.5, 0.5)
	test.AssertInDeltaVec3(t, settings.LookFrom, ray.Origin, core.Tolerance)
	test.AssertInDeltaVec3(t, settings.LookAt, ray.Eval(1), core.Tolerance)

	ray = camera.GetRay(0.5, 0)
	test.AssertInDeltaVec3(t, core.Vec3{0, 1, -1}, ray.Eval(1), core.Tolerance)

	ray = camera.GetRay(0.5, 1)
	test.AssertInDeltaVec3(t, core.Vec3{0, -1, -1}, ray.Eval(1), core.Tolerance)

	ray = camera.GetRay(0, 0)
	test.AssertInDeltaVec3(t, core.Vec3{-2, 1, -1}, ray.Eval(1), core.Tolerance)

	ray = camera.GetRay(1, 1)
	test.AssertInDeltaVec3(t, core.Vec3{2, -1, -1}, ray.Eval(1), core.Tolerance)
}

func TestCamera_Custom(t *testing.T) {
	t.Log("Camera with custom settings without randomness")
	settings := render.DefaultCameraSettings()
	settings.LookAt = core.Vec3{3, 0, 4}
	settings.AspectRatio = 1
	camera := render.NewCamera(&settings)
	core.Random().Disable()
	defer core.Random().Enable()

	test.AssertInDeltaVec3(t, settings.LookFrom, camera.Origin, core.Tolerance)

	ray := camera.GetRay(0.5, 0.5)
	test.AssertInDeltaVec3(t, settings.LookFrom, ray.Origin, core.Tolerance)
	test.AssertInDeltaVec3(t, settings.LookAt, ray.Eval(1), core.Tolerance)

	ray = camera.GetRay(0.5, 0)
	test.AssertInDeltaVec3(t, core.Vec3{3, 5, 4}, ray.Eval(1), core.Tolerance)

	ray = camera.GetRay(0, 0.5)
	test.AssertInDeltaVec3(t, core.Vec3{7, 0, 1}, ray.Eval(1), core.Tolerance)
}

func TestCamera_indexToU(t *testing.T) {
	t.Log("Camera with 100 px height and 1:1 aspect ratio")
	settings := render.DefaultCameraSettings()
	settings.ImagePixelHeight = 100
	settings.AspectRatio = 2.0
	camera := render.NewCamera(&settings)

	core.Random().Disable()
	defer core.Random().Enable()

	test.CheckResult(t, "Image height", camera.PixelHeight, 100)
	test.CheckResult(t, "Image width", camera.PixelWidth, 200)

	t.Log("  Pixel 0 to u")
	test.CheckResult(t, "u param", camera.IndexToU(0), core.Real(0))
	t.Log("  Pixel 100 to u")
	test.CheckResult(t, "u param", camera.IndexToU(100), core.Real(0.5))
	t.Log("  Pixel 200 to u")
	test.CheckResult(t, "u param", camera.IndexToU(200), core.Real(1))

	t.Log("  Pixel 0 to v")
	test.CheckResult(t, "v param", camera.IndexToV(0), core.Real(0))
	t.Log("  Pixel 50 to v")
	test.CheckResult(t, "v param", camera.IndexToV(50), core.Real(0.5))
	t.Log("  Pixel 100 to v")
	test.CheckResult(t, "v param", camera.IndexToV(100), core.Real(1))
}

func TestCamera_toRGBA(t *testing.T) {
	t.Log("Black to RGB")
	colorIn := core.Vec3{0, 0, 0}
	colorOut := color.RGBA{0, 0, 0, 255}
	test.CheckResult(t, "RGBA color", render.ToRGBA(colorIn), colorOut)

	t.Log("White to RGB")
	colorIn = core.Vec3{1., 1., 1.}
	colorOut = color.RGBA{255, 255, 255, 255}
	test.CheckResult(t, "RGBA color", render.ToRGBA(colorIn), colorOut)

	t.Log("Gray to RGB with gamma correction")
	colorIn = core.Vec3{0.64, 0.64, 0.64}
	colorOut = color.RGBA{204, 204, 204, 255}
	test.CheckResult(t, "RGBA color", render.ToRGBA(colorIn), colorOut)
}

func TestCamera_RenderEmptyScene(t *testing.T) {
	t.Log("Given an empty scene with white background")
	scene := render.Scene{SkyColorTop: core.White, SkyColorBottom: core.White}

	imageSize := 2
	t.Logf("and a camera with %vx%v resolution,\n", imageSize, imageSize)
	settings := render.DefaultCameraSettings()
	settings.ImagePixelHeight = imageSize
	settings.AspectRatio = 1
	settings.Antialiasing = 1
	camera := render.NewCamera(&settings)

	t.Logf("  the rendered image should be a %vx%v white square:\n", imageSize, imageSize)
	renderedImage := camera.Render(&scene)

	test.CheckResult(t, "Image width", renderedImage.Bounds().Size().X, imageSize)
	test.CheckResult(t, "Image height", renderedImage.Bounds().Size().Y, imageSize)

	expectedColor := color.RGBA{255, 255, 255, 255}
	for x := 0; x < imageSize; x++ {
		for y := 0; y < imageSize; y++ {
			test.CheckResult(t, "Pixel color", renderedImage.At(x, y), expectedColor)
		}
	}

}

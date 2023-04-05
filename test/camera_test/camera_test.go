package camera_test

// func TestCamera_indexToU(t *testing.T) {
// 	t.Log("Camera with 100 px height and 1:1 aspect ratio")
// 	settings := camera.DefaultCameraSettings()
// 	settings.ImagePixelHeight = 100
// 	settings.AspectRatio = 2.0
// 	randomizer := random.NewFakeRandomGenerator()

// 	camera := camera.NewCamera(&settings, randomizer)

// 	// assert.Equal(t, 100, camera.PixelHeight)
// 	// assert.Equal(t, 200, camera.PixelWidth)

// 	// assert.EqualValues(t, 0, camera.IndexToU(0))
// 	// assert.EqualValues(t, 0.5, camera.IndexToU(100))
// 	// assert.EqualValues(t, 1, camera.IndexToU(200))

// 	// assert.EqualValues(t, 0, camera.IndexToV(0))
// 	// assert.EqualValues(t, 0.5, camera.IndexToV(50))
// 	// assert.EqualValues(t, 1, camera.IndexToV(100))
// }

// func TestCamera_RenderEmptyScene(t *testing.T) {
// 	t.Log("Given an empty scene with white background")
// 	scene := scene.New(background.NewFlatColor(color.White))

// 	imageSize := 2
// 	t.Logf("and a camera with %vx%v resolution,\n", imageSize, imageSize)
// 	settings := camera.DefaultCameraSettings()
// 	settings.ImagePixelHeight = imageSize
// 	settings.AspectRatio = 1
// 	settings.Antialiasing = 1
// 	randomizer := random.NewFakeRandomGenerator()
// 	camera := camera.NewCamera(&settings, randomizer)

// 	t.Logf("  the rendered image should be a %vx%v white square:\n", imageSize, imageSize)
// 	camera.Render(scene)

// 	// assert.Equal(t, imageSize, renderedImage.Bounds().Size().X)
// 	// assert.Equal(t, imageSize, renderedImage.Bounds().Size().Y)

// 	// for x := 0; x < imageSize; x++ {
// 	// 	for y := 0; y < imageSize; y++ {
// 	// 		assert.Equal(t, color.White, renderedImage.PixelColor(x, y))
// 	// 	}
// 	// }

// }

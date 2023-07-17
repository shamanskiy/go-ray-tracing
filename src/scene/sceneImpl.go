package scene

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/background"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
)

const (
	DEFAULT_MIN_HIT_PARAM       core.Real = 0.0001
	DEFAULT_MAX_RAY_REFLECTIONS int       = 10
)

type SceneImpl struct {
	background background.Background
	bvh        *geometries.BVHNode

	minHitParam       core.Real // prevents black acne
	maxRayReflections int       // prevents infinite ray bouncing between parallel walls
}

func New(objects []Object, background background.Background, settings ...SceneImplSetting) *SceneImpl {
	scene := &SceneImpl{
		background:        background,
		minHitParam:       DEFAULT_MIN_HIT_PARAM,
		maxRayReflections: DEFAULT_MAX_RAY_REFLECTIONS,
	}

	for _, setting := range settings {
		setting(scene)
	}

	hittables := make([]geometries.Hittable, 0, len(objects))
	for _, object := range objects {
		hittables = append(hittables, object)
	}
	scene.bvh = geometries.BuildBVH(hittables)

	return scene
}

func (s *SceneImpl) TestRay(ray core.Ray) color.Color {
	return s.testRay(ray, 0)
}

func (s *SceneImpl) testRay(ray core.Ray, reflectionDepth int) color.Color {
	optionalHit := s.bvh.TestRay(ray, core.NewInterval(s.minHitParam, core.Inf()))
	if optionalHit.Empty() {
		return s.background.ColorRay(ray)
	}

	if reflectionDepth >= s.maxRayReflections {
		return color.Black
	}

	hit := optionalHit.Value()
	reflection := hit.Material.Reflect(ray.Direction(), hit.Point, hit.Normal)
	switch reflection.Type {
	case materials.Scattered:
		reflectedRayColor := s.testRay(reflection.Ray, reflectionDepth+1)
		return reflectedRayColor.MulColor(reflection.Color)
	case materials.Emitted:
		return reflection.Color
	case materials.Absorbed:
		return color.Black
	default:
		panic("unknown reflection type")
	}
}

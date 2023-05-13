package scene

import (
	"github.com/Shamanskiy/go-ray-tracer/src/camera/log"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/optional"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/background"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
	"github.com/google/uuid"
)

const (
	DEFAULT_MIN_HIT_PARAM       core.Real = 0.0001
	DEFAULT_MAX_RAY_REFLECTIONS int       = 10
)

type SceneImpl struct {
	objects    []geometries.Geometry
	materials  map[uuid.UUID]materials.Material
	background background.Background
	bvh        *geometries.BVHNode

	minHitParam       core.Real // prevents black acne
	maxRayReflections int       // prevents infinite ray bouncing between parallel walls
}

func New(background background.Background, settings ...SceneImplSetting) *SceneImpl {
	scene := &SceneImpl{
		background:        background,
		minHitParam:       DEFAULT_MIN_HIT_PARAM,
		maxRayReflections: DEFAULT_MAX_RAY_REFLECTIONS,
		materials:         make(map[uuid.UUID]materials.Material),
	}

	for _, setting := range settings {
		setting(scene)
	}

	return scene
}

func (s *SceneImpl) Add(object geometries.Geometry, material materials.Material) {
	s.objects = append(s.objects, object)
	s.materials[object.Id()] = material
	s.buildBVH()
}

func (s *SceneImpl) buildBVH() {
	defer log.TimeExecution("build bvh")()
	s.bvh = geometries.BuildBVH(s.objects)
}

func (s *SceneImpl) TestRay(ray core.Ray) color.Color {
	if len(s.objects) == 0 {
		return s.background.ColorRay(ray)
	}

	return s.testRay(ray, 0)
}

func (s *SceneImpl) testRay(ray core.Ray, reflectionDepth int) color.Color {
	optionalHit := s.hitClosestObject(ray)
	if optionalHit.Empty() {
		return s.background.ColorRay(ray)
	}

	if reflectionDepth >= s.maxRayReflections {
		return color.Black
	}

	hit := optionalHit.Value()
	reflection := hit.material.Reflect(ray.Direction(), hit.location.Point, hit.location.Normal)
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

type objectHit struct {
	location geometries.HitPoint
	material materials.Material
}

func (s *SceneImpl) hitClosestObject(ray core.Ray) optional.Optional[objectHit] {
	optionalHit := s.bvh.TestRay(ray, core.NewInterval(s.minHitParam, core.Inf()))

	if optionalHit.Empty() {
		return optional.Empty[objectHit]()
	}

	hit := optionalHit.Value()
	closestHitPoint := hit.HitGeometry.EvaluateHit(ray, hit.Param)
	return optional.Of(objectHit{
		location: closestHitPoint,
		material: s.materials[hit.HitGeometry.Id()],
	})
}

package scene

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
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
}

func (s *SceneImpl) TestRay(ray core.Ray) color.Color {
	return s.testRay(ray, 0)
}

func (s *SceneImpl) testRay(ray core.Ray, reflectionDepth int) color.Color {
	objectHit := s.hitClosestObject(ray)
	if !objectHit.hasHit {
		return s.background.ColorRay(ray)
	}

	if reflectionDepth >= s.maxRayReflections {
		return color.Black
	}

	reflection := objectHit.material.Reflect(ray.Direction(), objectHit.location.Point, objectHit.location.Normal)
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
	hasHit   bool
	location geometries.HitPoint
	material materials.Material
}

func (s *SceneImpl) hitClosestObject(ray core.Ray) objectHit {
	closestHit := geometries.Hit{false, core.Inf(), nil}

	for currentObjectIndex := range s.objects {
		hit := s.objects[currentObjectIndex].TestRay(ray, core.NewInterval(s.minHitParam, core.Inf()))

		if hit.HasHit && hit.Param < closestHit.Param {
			closestHit = hit
		}
	}

	if !closestHit.HasHit {
		return objectHit{}
	}

	closestHitPoint := closestHit.HitGeometry.EvaluateHit(ray, closestHit.Param)
	return objectHit{
		hasHit:   true,
		location: closestHitPoint,
		material: s.materials[closestHit.HitGeometry.Id()],
	}
}

package scene

import (
	"github.com/Shamanskiy/go-ray-tracer/src/background"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/slices"
	"github.com/Shamanskiy/go-ray-tracer/src/core/slices/filters"
	"github.com/Shamanskiy/go-ray-tracer/src/geometries"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
)

const (
	DEFAULT_MIN_HIT_PARAM       core.Real = 0.0001
	DEFAULT_MAX_RAY_REFLECTIONS int       = 10
)

type SceneImpl struct {
	objects    []geometries.Geometry
	materials  []materials.Material
	background background.Background

	minHitParam       core.Real // prevents black acne
	maxRayReflections int       // prevents infinite ray bouncing between parallel walls
}

func New(background background.Background, settings ...SceneImplSetting) *SceneImpl {
	scene := &SceneImpl{
		background:        background,
		minHitParam:       DEFAULT_MIN_HIT_PARAM,
		maxRayReflections: DEFAULT_MAX_RAY_REFLECTIONS,
	}

	for _, setting := range settings {
		setting(scene)
	}

	return scene
}

func (s *SceneImpl) Add(object geometries.Geometry, material materials.Material) {
	s.objects = append(s.objects, object)
	s.materials = append(s.materials, material)
}

func (s *SceneImpl) TestRay(ray core.Ray) color.Color {
	return s.testRay(ray, 0)
}

func (s *SceneImpl) testRay(ray core.Ray, reflectionDepth int) color.Color {
	objectHit := s.hitClosestObject(ray)
	if objectHit.noHit() {
		return s.background.ColorRay(ray)
	}

	if reflectionDepth >= s.maxRayReflections {
		return color.Black
	}

	reflection := objectHit.material.Reflect(ray.Direction(), objectHit.location.Point, objectHit.location.Normal)
	if reflection == nil {
		return color.Black
	}

	reflectedRayColor := s.testRay(reflection.Ray, reflectionDepth+1)
	return reflectedRayColor.MulColor(reflection.Color)
}

type objectHit struct {
	location *geometries.HitPoint
	material materials.Material
}

func (hit objectHit) noHit() bool {
	return hit.location == nil
}

func (s *SceneImpl) hitClosestObject(ray core.Ray) objectHit {
	closestHit := core.Inf()
	var objectIndex int

	for currentObjectIndex := range s.objects {
		hits := s.objects[currentObjectIndex].TestRay(ray)
		firstHit := slices.FindFirst(hits, filters.GreaterOrEqualThan(s.minHitParam))

		if firstHit == nil {
			continue
		}

		if *firstHit < closestHit {
			closestHit = *firstHit
			objectIndex = currentObjectIndex
		}
	}

	if closestHit == core.Inf() {
		return objectHit{}
	} else {
		closestHitPoint := s.objects[objectIndex].EvaluateHit(ray, closestHit)
		return objectHit{
			location: &closestHitPoint,
			material: s.materials[objectIndex],
		}
	}
}

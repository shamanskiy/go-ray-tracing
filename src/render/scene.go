package render

import (
	"fmt"

	"github.com/Shamanskiy/go-ray-tracer/src/background"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/slices"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
)

const (
	DEFAULT_MIN_HIT_PARAM       core.Real = 0.0001
	DEFAULT_MAX_RAY_REFLECTIONS int       = 10
)

type Scene struct {
	objects    []objects.Object
	materials  []materials.Material
	background background.Background

	minHitParam       core.Real // prevents black acne
	maxRayReflections int       // prevents infinite ray bouncing between parallel walls
}

type SceneSetting func(*Scene)

func MinRayHitParameter(minHitParam core.Real) SceneSetting {
	if minHitParam < 0 {
		panic(fmt.Errorf("invalid min ray hit parameter: %v", minHitParam))
	}
	return func(scene *Scene) {
		scene.minHitParam = minHitParam
	}
}

func MaxRayReflections(maxReflections int) SceneSetting {
	if maxReflections < 0 {
		panic(fmt.Errorf("invalid max ray reflections: %d", maxReflections))
	}
	return func(scene *Scene) {
		scene.maxRayReflections = maxReflections
	}
}

func NewScene(background background.Background, settings ...SceneSetting) *Scene {
	scene := &Scene{
		background:        background,
		minHitParam:       DEFAULT_MIN_HIT_PARAM,
		maxRayReflections: DEFAULT_MAX_RAY_REFLECTIONS,
	}

	for _, setting := range settings {
		setting(scene)
	}

	return scene
}

func (s *Scene) Add(object objects.Object, material materials.Material) {
	s.objects = append(s.objects, object)
	s.materials = append(s.materials, material)
}

func (s *Scene) TestRay(ray core.Ray) color.Color {
	return s.testRay(ray, 0)
}

func (s *Scene) testRay(ray core.Ray, depth int) color.Color {
	hit, objectIndex := s.hitClosestObject(ray)
	if hit == nil {
		return s.background.ColorRay(ray)
	}

	if depth >= s.maxRayReflections {
		return color.Black
	}

	reflection := s.materials[objectIndex].Reflect(ray.Direction(), hit.Point, hit.Normal)
	if reflection == nil {
		return color.Black
	}

	reflectedRayColor := s.testRay(reflection.Ray, depth+1)
	return reflectedRayColor.MulColor(reflection.Attenuation)
}

func (s *Scene) hitClosestObject(ray core.Ray) (hit *objects.HitRecord, objectIndex int) {
	closestHit := core.Inf()

	for currentObjectIndex := range s.objects {
		hits := s.objects[currentObjectIndex].TestRay(ray)
		firstHit := slices.FindFirstLargerOrEqualThan(hits, s.minHitParam)

		if firstHit == nil {
			continue
		}

		if *firstHit < closestHit {
			closestHit = *firstHit
			objectIndex = currentObjectIndex
		}
	}

	if closestHit == core.Inf() {
		return nil, 0
	} else {
		closestHitRecord := s.objects[objectIndex].EvaluateHit(ray, closestHit)
		return &closestHitRecord, objectIndex
	}
}

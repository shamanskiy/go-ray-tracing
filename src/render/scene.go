package render

import (
	"github.com/Shamanskiy/go-ray-tracer/src/background"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
)

type Scene struct {
	objects    []objects.Object
	materials  []materials.Material
	background background.Background
}

func NewScene(background background.Background) *Scene {
	return &Scene{
		background: background,
	}
}

func (s *Scene) HitClosestObject(ray core.Ray, minParam core.Real) (hit *objects.HitRecord, objectIndex int) {
	noHitYet := true
	closestHitParam := core.Inf()

	for i := range s.objects {
		hits := s.objects[i].TestRay(ray)
		for _, hitParam := range hits {
			if hitParam >= minParam && hitParam < closestHitParam {
				closestHitParam = hitParam
				objectIndex = i
				noHitYet = false
				break
			}
		}
	}

	if noHitYet {
		return nil, 0
	} else {
		closestHitRecord := s.objects[objectIndex].EvaluateHit(ray, closestHitParam)
		return &closestHitRecord, objectIndex
	}
}

func (s *Scene) Add(object objects.Object, material materials.Material) {
	s.objects = append(s.objects, object)
	s.materials = append(s.materials, material)
}

func (s *Scene) testRay(ray core.Ray, depth int) color.Color {
	hit, objectIndex := s.HitClosestObject(ray, 0.0001)
	if hit != nil {
		reflection := s.materials[objectIndex].Reflect(ray.Direction(), hit.Point, hit.Normal)
		if reflection != nil && depth < 10 {
			reflectedRayColor := s.testRay(reflection.Ray, depth+1)
			return reflectedRayColor.MulColor(reflection.Attenuation)
		} else {
			return color.Black
		}
	}

	return s.background.ColorRay(ray)
}

func (s *Scene) TestRay(ray core.Ray) color.Color {
	return s.testRay(ray, 0)
}

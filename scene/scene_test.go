package scene

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/objects"
)

func TestScene_RightSphere(t *testing.T) {
	leftSphere := objects.Sphere{core.Vec3{-6.0, 0.0, 0.0}, 2.0}
	rightSphere := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a scene with two spheres %v and %v,\n", leftSphere, rightSphere)
	var scene Scene
	scene.Add(leftSphere, rightSphere)

	hitRay := core.Ray{core.Vec3{4.0, 0.0, 0.0}, core.Vec3{-1.0, 0.0, 0.0}}
	t.Logf("\twe can test if a ray with origin %v and direction %v hits anything:\n", hitRay.Origin, hitRay.Direction)
	hitRecord := scene.Hit(hitRay)
	expected := objects.HitRecord{Param: 2.0, Point: core.Vec3{2.0, 0.0, 0.0}, Normal: core.Vec3{1.0, 0.0, 0.0}}
	if *hitRecord == expected {
		t.Logf("\t\tPASSED: the ray hit the right sphere, hit record is %v, expected %v", hitRecord, expected)
	} else {
		t.Fatalf("\t\tFAILED: hit record is %v, expected %v", hitRecord, expected)
	}
}

func TestScene_LeftSphere(t *testing.T) {
	leftSphere := objects.Sphere{core.Vec3{-6.0, 0.0, 0.0}, 2.0}
	rightSphere := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a scene with two spheres %v and %v,\n", leftSphere, rightSphere)
	var scene Scene
	scene.Add(leftSphere, rightSphere)

	hitRay := core.Ray{core.Vec3{4.0, 0.0, 0.0}, core.Vec3{-1.0, 0.0, 0.0}}
	t.Logf("\twe can test if a ray with origin %v and direction %v hits anything with minimum parameter 7.0:\n",
		hitRay.Origin, hitRay.Direction)
	hitRecord := scene.HitWithMin(hitRay, 7.0)
	expected := objects.HitRecord{Param: 8.0, Point: core.Vec3{-4.0, 0.0, 0.0}, Normal: core.Vec3{1.0, 0.0, 0.0}}
	if *hitRecord == expected {
		t.Logf("\t\tPASSED: the ray hit the left sphere, hit record is %v, expected %v", hitRecord, expected)
	} else {
		t.Fatalf("\t\tFAILED: hit record is %v, expected %v", hitRecord, expected)
	}
}

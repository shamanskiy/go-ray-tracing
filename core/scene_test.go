package core

import (
	"testing"
)

func TestScene_RightSphere(t *testing.T) {
	leftSphere := Sphere{Vec3{-6.0, 0.0, 0.0}, 2.0}
	rightSphere := Sphere{Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a scene with two spheres %v and %v,\n", leftSphere, rightSphere)
	var scene Scene
	scene.Add(leftSphere, rightSphere)

	hitRay := Ray{Origin: Vec3{4.0, 0.0, 0.0}, Direction: Vec3{-1.0, 0.0, 0.0}}
	t.Logf("\twe can test if a ray with origin %v and direction %v hits anything:\n", hitRay.Origin, hitRay.Direction)
	hitRecord := scene.Hit(hitRay)
	expected := HitRecord{Hit: true, Param: 2.0, Point: Vec3{2.0, 0.0, 0.0}, Normal: Vec3{1.0, 0.0, 0.0}}
	if hitRecord == expected {
		t.Logf("\t\tPASSED: the ray hit the right sphere, hit record is %v, expected %v", hitRecord, expected)
	} else {
		t.Fatalf("\t\tFAILED: hit record is %v, expected %v", hitRecord, expected)
	}
}

func TestScene_LeftSphere(t *testing.T) {
	leftSphere := Sphere{Vec3{-6.0, 0.0, 0.0}, 2.0}
	rightSphere := Sphere{Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a scene with two spheres %v and %v,\n", leftSphere, rightSphere)
	var scene Scene
	scene.Add(leftSphere, rightSphere)

	hitRay := Ray{Origin: Vec3{4.0, 0.0, 0.0}, Direction: Vec3{-1.0, 0.0, 0.0}, MinParam: 7.0}
	t.Logf("\twe can test if a ray with origin %v, direction %v and minimum parameter %v hits anything:\n",
		hitRay.Origin, hitRay.Direction, hitRay.MinParam)
	hitRecord := scene.Hit(hitRay)
	expected := HitRecord{Hit: true, Param: 8.0, Point: Vec3{-4.0, 0.0, 0.0}, Normal: Vec3{1.0, 0.0, 0.0}}
	if hitRecord == expected {
		t.Logf("\t\tPASSED: the ray hit the left sphere, hit record is %v, expected %v", hitRecord, expected)
	} else {
		t.Fatalf("\t\tFAILED: hit record is %v, expected %v", hitRecord, expected)
	}
}

package core

import (
	"testing"
)

func TestSphere_FirstHit(t *testing.T) {
	sphere := Sphere{Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a sphere with center at %v and radius %v,\n", sphere.Center, sphere.Radius)

	hitRay := Ray{Origin: Vec3{4.0, 0.0, 0.0}, Direction: Vec3{-1.0, 0.0, 0.0}}
	t.Logf("\twe can test if a ray with origin %v and direction %v hits the sphere:\n", hitRay.Origin, hitRay.Direction)
	hitRecord := sphere.Hit(hitRay)
	expected := HitRecord{Hit: true, Param: 2.0, Point: Vec3{2.0, 0.0, 0.0}, Normal: Vec3{1.0, 0.0, 0.0}}
	if hitRecord == expected {
		t.Logf("\t\tPASSED: result is %v, expected %v", hitRecord, expected)
	} else {
		t.Fatalf("\t\tFAILED: result is %v, expected %v", hitRecord, expected)
	}
}

func TestSphere_SecondHit(t *testing.T) {
	sphere := Sphere{Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a sphere with center at %v and radius %v,\n", sphere.Center, sphere.Radius)

	hitRay := Ray{Origin: Vec3{4.0, 0.0, 0.0}, Direction: Vec3{-1.0, 0.0, 0.0}, MinParam: 3.0}
	t.Logf("\twe can test if a ray with origin %v, direction %v and minimum parameter %v hits the sphere:\n", hitRay.Origin, hitRay.Direction, hitRay.MinParam)
	hitRecord := sphere.Hit(hitRay)

	expected := HitRecord{Hit: true, Param: 6.0, Point: Vec3{-2.0, 0.0, 0.0}, Normal: Vec3{-1.0, 0.0, 0.0}}
	if hitRecord == expected {
		t.Logf("\t\tPASSED: result is %v, expected %v", hitRecord, expected)
	} else {
		t.Fatalf("\t\tFAILED: result is %v, expected %v", hitRecord, expected)
	}
}

func TestSphere_TangentHit(t *testing.T) {
	sphere := Sphere{Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a sphere with center at %v and radius %v,\n", sphere.Center, sphere.Radius)

	hitRay := Ray{Origin: Vec3{4.0, 2.0, 0.0}, Direction: Vec3{-1.0, 0.0, 0.0}}
	t.Logf("\twe can test if a ray with origin %v and direction %v hits the sphere:\n", hitRay.Origin, hitRay.Direction)
	hitRecord := sphere.Hit(hitRay)

	expected := HitRecord{Hit: true, Param: 4.0, Point: Vec3{0.0, 2.0, 0.0}, Normal: Vec3{0.0, 1.0, 0.0}}
	if hitRecord == expected {
		t.Logf("\t\tPASSED: result is %v, expected %v", hitRecord, expected)
	} else {
		t.Fatalf("\t\tFAILED: result is %v, expected %v", hitRecord, expected)
	}
}

func TestSphere_NoHit_RayParamIsTooLarge(t *testing.T) {
	sphere := Sphere{Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a sphere with center at %v and radius %v,\n", sphere.Center, sphere.Radius)

	hitRay := Ray{Origin: Vec3{4.0, 0.0, 0.0}, Direction: Vec3{1.0, 0.0, 0.0}, MinParam: 7.0}
	t.Logf("\twe can test if a ray with origin %v, direction %v and minimum parameter %v hits the sphere:\n", hitRay.Origin, hitRay.Direction, hitRay.MinParam)
	hitRecord := sphere.Hit(hitRay)

	expected := HitRecord{}
	if hitRecord == expected {
		t.Logf("\t\tPASSED: no hit, result is %v, expected %v", hitRecord, expected)
	} else {
		t.Fatalf("\t\tFAILED: result is %v, expected %v", hitRecord, expected)
	}
}

func TestSphere_NoHit_RayMisses(t *testing.T) {
	sphere := Sphere{Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a sphere with center at %v and radius %v,\n", sphere.Center, sphere.Radius)

	hitRay := Ray{Origin: Vec3{4.0, 0.0, 0.0}, Direction: Vec3{0.0, 1.0, 0.0}}
	t.Logf("\twe can test if a ray with origin %v and direction %v hits the sphere:\n", hitRay.Origin, hitRay.Direction)
	hitRecord := sphere.Hit(hitRay)

	expected := HitRecord{}
	if hitRecord == expected {
		t.Logf("\t\tPASSED: no hit, result is %v, expected %v", hitRecord, expected)
	} else {
		t.Fatalf("\t\tFAILED: result is %v, expected %v", hitRecord, expected)
	}
}

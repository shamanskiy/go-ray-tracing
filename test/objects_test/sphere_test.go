package objects

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/Shamanskiy/go-ray-tracer/test"
)

func TestSphere_FirstHit(t *testing.T) {
	sphere := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a sphere with center at %v and radius %v,\n", sphere.Center, sphere.Radius)

	hitRay := core.Ray{core.Vec3{4.0, 0.0, 0.0}, core.Vec3{-1.0, 0.0, 0.0}}
	t.Logf("  we can test if a ray with origin %v and direction %v hits the sphere:\n", hitRay.Origin, hitRay.Direction)
	hitRecord := sphere.Hit(hitRay)
	expected := objects.HitRecord{Param: 2.0, Point: core.Vec3{2.0, 0.0, 0.0}, Normal: core.Vec3{1.0, 0.0, 0.0}}

	test.CheckResult(t, "hit record", *hitRecord, expected)
}

func TestSphere_SecondHit(t *testing.T) {
	sphere := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a sphere with center at %v and radius %v,\n", sphere.Center, sphere.Radius)

	hitRay := core.Ray{core.Vec3{4.0, 0.0, 0.0}, core.Vec3{-1.0, 0.0, 0.0}}
	t.Logf("  we can test if a ray with origin %v and direction %v hits the sphere with minimum parameter 3.0:\n",
		hitRay.Origin, hitRay.Direction)
	hitRecord := sphere.HitWithMin(hitRay, 3.0)

	expected := objects.HitRecord{Param: 6.0, Point: core.Vec3{-2.0, 0.0, 0.0}, Normal: core.Vec3{-1.0, 0.0, 0.0}}

	test.CheckResult(t, "hit record", *hitRecord, expected)
}

func TestSphere_TangentHit(t *testing.T) {
	sphere := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a sphere with center at %v and radius %v,\n", sphere.Center, sphere.Radius)

	hitRay := core.Ray{core.Vec3{4.0, 2.0, 0.0}, core.Vec3{-1.0, 0.0, 0.0}}
	t.Logf("  we can test if a ray with origin %v and direction %v hits the sphere:\n", hitRay.Origin, hitRay.Direction)
	hitRecord := sphere.Hit(hitRay)

	expected := objects.HitRecord{Param: 4.0, Point: core.Vec3{0.0, 2.0, 0.0}, Normal: core.Vec3{0.0, 1.0, 0.0}}

	test.CheckResult(t, "hit record", *hitRecord, expected)
}

func TestSphere_NoHit_RayParamIsTooLarge(t *testing.T) {
	sphere := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a sphere with center at %v and radius %v,\n", sphere.Center, sphere.Radius)

	hitRay := core.Ray{core.Vec3{4.0, 0.0, 0.0}, core.Vec3{1.0, 0.0, 0.0}}
	t.Logf("  we can test if a ray with origin %v and direction %v hits the sphere with minimum parameter 7.0:\n",
		hitRay.Origin, hitRay.Direction)
	hitRecord := sphere.HitWithMin(hitRay, 7.0)

	test.CheckNil(t, "hit record", hitRecord)
}

func TestSphere_NoHit_RayMisses(t *testing.T) {
	sphere := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 2.0}
	t.Logf("Given a sphere with center at %v and radius %v,\n", sphere.Center, sphere.Radius)

	hitRay := core.Ray{Origin: core.Vec3{4.0, 0.0, 0.0}, Direction: core.Vec3{0.0, 1.0, 0.0}}
	t.Logf("  we can test if a ray with origin %v and direction %v hits the sphere:\n", hitRay.Origin, hitRay.Direction)
	hitRecord := sphere.Hit(hitRay)

	test.CheckNil(t, "hit record", hitRecord)
}

package render

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/core"
)

func TestScene_Default(t *testing.T) {
	t.Log("Given a default empty scene, any ray tests black:")
	scene := Scene{}
	expected := core.Black

	ray := core.Ray{core.Vec3{}, core.Vec3{}}
	t.Logf("\tray %v:\n", ray)
	color := scene.TestRay(ray)
	if color == expected {
		t.Logf("\t\tPASSED: received color %v, expected %v", color, expected)
	} else {
		t.Fatalf("\t\tFAILED: received color %v, expected %v", color, expected)
	}

	ray = core.Ray{core.Vec3{1, 2, 3}, core.Vec3{4, 5, 6}}
	t.Logf("\tray %v:\n", ray)
	color = scene.TestRay(ray)
	if color == expected {
		t.Logf("\t\tPASSED: received color %v, expected %v", color, expected)
	} else {
		t.Fatalf("\t\tFAILED: received color %v, expected %v", color, expected)
	}

}

func TestScene_Empty(t *testing.T) {
	skyBottom, skyTop := core.Red, core.Blue
	scene := Scene{SkyColorBottom: skyBottom, SkyColorTop: skyTop}
	t.Logf("Given an empty scene with bottom sky color %v and top sky color %v,\n",
		skyBottom, skyTop)

	ray := core.Ray{core.Vec3{0, 0, 0}, core.Vec3{0, -1, 0}}
	t.Logf("\tdown-facing ray %v should return the bottom color:\n", ray)
	color := scene.TestRay(ray)
	expected := skyBottom
	if color == expected {
		t.Logf("\t\tPASSED: received color %v, expected %v", color, expected)
	} else {
		t.Fatalf("\t\tFAILED: received color %v, expected %v", color, expected)
	}

	ray = core.Ray{core.Vec3{0, 0, 0}, core.Vec3{0, 1, 0}}
	t.Logf("\tup-facing ray %v should return the top color:\n", ray)
	color = scene.TestRay(ray)
	expected = skyTop
	if color == expected {
		t.Logf("\t\tPASSED: received color %v, expected %v", color, expected)
	} else {
		t.Fatalf("\t\tFAILED: received color %v, expected %v", color, expected)
	}

	ray = core.Ray{core.Vec3{0, 0, 0}, core.Vec3{1, 0, 0}}
	t.Logf("\thorizontal ray %v should return the blend color:\n", ray)
	color = scene.TestRay(ray)
	expected = skyTop.Add(skyBottom).Mul(0.5)
	if color == expected {
		t.Logf("\t\tPASSED: received color %v, expected %v", color, expected)
	} else {
		t.Fatalf("\t\tFAILED: received color %v, expected %v", color, expected)
	}

}

// test scene.hitClosestObject
// test scene.testRay with white skybox and single sphere
// test multiple reflections (how?)

/*func TestScene_RightSphere(t *testing.T) {
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
*/

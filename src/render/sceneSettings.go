package render

import (
	"fmt"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
)

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

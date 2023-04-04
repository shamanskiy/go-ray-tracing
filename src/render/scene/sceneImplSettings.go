package scene

import (
	"fmt"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
)

type SceneImplSetting func(*SceneImpl)

func MinRayHitParameter(minHitParam core.Real) SceneImplSetting {
	if minHitParam < 0 {
		panic(fmt.Errorf("invalid min ray hit parameter: %v", minHitParam))
	}
	return func(scene *SceneImpl) {
		scene.minHitParam = minHitParam
	}
}

func MaxRayReflections(maxReflections int) SceneImplSetting {
	if maxReflections < 0 {
		panic(fmt.Errorf("invalid max ray reflections: %d", maxReflections))
	}
	return func(scene *SceneImpl) {
		scene.maxRayReflections = maxReflections
	}
}

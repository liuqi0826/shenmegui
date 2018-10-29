package display

import (
	"github.com/liuqi0826/seven/geom"
	"github.com/liuqi0826/shenmegui/global"
)

//++++++++++++++++++++ Picker ++++++++++++++++++++

type Picker struct {
	ray geom.Ray
}

func (this *Picker) Picker() {
	origin := new(geom.Vector4)
	origin.Vector4()
	direction := new(geom.Vector4)
	direction.Vector4()
	direction.Z = 1.0
	this.ray = geom.Ray{}
	this.ray.Ray(origin, direction)
}
func (this *Picker) Update(screenPosition *geom.Vector3, camera *Camera) {
	transform := camera.GetTransformMatrix().Clone()
	if transform.Invert() {
		csDirec := new(geom.Vector4)
		csDirec.X = ((2 * screenPosition.X / float32(global.Config.WindowWidth)) - 1) / camera.GetProjectionMatrix().Raw[0]
		csDirec.Y = -((2 * screenPosition.Y / float32(global.Config.WindowHeight)) - 1) / camera.GetProjectionMatrix().Raw[5]
		csDirec.Z = 1
		wsStart := new(geom.Vector4)
		wsStart.X = transform.Raw[12]
		wsStart.Y = transform.Raw[13]
		wsStart.Z = transform.Raw[14]
		wsDirec := new(geom.Vector4)
		wsDirec.X = csDirec.X*transform.Raw[0] + csDirec.Y*transform.Raw[4] + csDirec.Z*transform.Raw[8]
		wsDirec.Y = csDirec.X*transform.Raw[1] + csDirec.Y*transform.Raw[5] + csDirec.Z*transform.Raw[9]
		wsDirec.Z = csDirec.X*transform.Raw[2] + csDirec.Y*transform.Raw[6] + csDirec.Z*transform.Raw[10]
		this.ray.Ray(wsStart, wsDirec)
	}
}
func (this *Picker) CollisionTest(tester ICollisionTest) bool {
	switch tester.GetType() {
	case PLANE:
		if obj, ok := tester.GetBoundingBox().(*geom.Plane); ok {
			return this.ray.IntersectPlane(obj)
		}
	case SPHERE:
		if obj, ok := tester.GetBoundingBox().(*geom.Sphere); ok {
			return this.ray.IntersectSphere(obj)
		}
	case TRIANGLE:
		if obj, ok := tester.GetBoundingBox().(*geom.Triangle); ok {
			return this.ray.IntersectTriangle(obj)
		}
	case AABB:
		if obj, ok := tester.GetBoundingBox().(*geom.Box); ok {
			return this.ray.IntersectAABB(obj)
		}
	case OBB:
	}
	return false
}
func (this *Picker) CollisionTestWithHitPoint(tester ICollisionTest) *geom.Vector4 {
	switch tester.GetType() {
	case PLANE:
		if obj, ok := tester.GetBoundingBox().(*geom.Plane); ok {
			return this.ray.IntersectPlaneHitPoint(obj)
		}
	case SPHERE:
		if obj, ok := tester.GetBoundingBox().(*geom.Sphere); ok {
			return this.ray.IntersectSphereWithHitPoint(obj)
		}
	case TRIANGLE:
		if obj, ok := tester.GetBoundingBox().(*geom.Triangle); ok {
			return this.ray.IntersectTriangleWithHitPoint(obj)
		}
	case AABB:
		if obj, ok := tester.GetBoundingBox().(*geom.Box); ok {
			return this.ray.IntersectAABBWithHitPoint(obj)
		}
	case OBB:
	}
	return nil
}

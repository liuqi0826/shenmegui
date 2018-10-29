package display

import "github.com/liuqi0826/seven/geom"

//++++++++++++++++++++ AxisAlignedBoundingBox ++++++++++++++++++++

type AxisAlignedBoundingBox struct {
	min *geom.Vector4
	max *geom.Vector4
	box *geom.Box
}

func (this *AxisAlignedBoundingBox) AxisAlignedBoundingBox(min *geom.Vector4, max *geom.Vector4) {
	this.min = &geom.Vector4{X: min.X, Y: min.Y, Z: min.Z, W: min.W}
	this.max = &geom.Vector4{X: max.X, Y: max.Y, Z: max.Z, W: max.W}
	this.box.Box(this.min, this.max)
}
func (this *AxisAlignedBoundingBox) GetType() string {
	return AABB
}
func (this *AxisAlignedBoundingBox) GetBoundingBox() *geom.Box {
	return this.box
}
func (this *AxisAlignedBoundingBox) Update(transform *geom.Matrix4x4) {
	if transform != nil {
		compose := transform.Decompose(geom.EULER_ANGLES)
		compose[1] = &geom.Vector4{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0}
		mtx := new(geom.Matrix4x4)
		mtx.Recompose(compose, geom.EULER_ANGLES)
		min := mtx.TransformVector(this.min)
		max := mtx.TransformVector(this.max)
		this.box.Box(min, max)
	}
}

//++++++++++++++++++++ OrientedBoundingBox ++++++++++++++++++++

type OrientedBoundingBox struct {
	min *geom.Vector4
	max *geom.Vector4
	box *geom.Box
}

func (this *OrientedBoundingBox) OrientedBoundingBox(min *geom.Vector4, max *geom.Vector4) {
	this.min = &geom.Vector4{X: min.X, Y: min.Y, Z: min.Z, W: min.W}
	this.max = &geom.Vector4{X: max.X, Y: max.Y, Z: max.Z, W: max.W}
	this.box.Box(this.min, this.max)
}
func (this *OrientedBoundingBox) GetType() string {
	return OBB
}
func (this *OrientedBoundingBox) GetBoundingBox() *geom.Box {
	return this.box
}
func (this *OrientedBoundingBox) Update(transform *geom.Matrix4x4) {
	if transform != nil {
		min := transform.TransformVector(this.min)
		max := transform.TransformVector(this.max)
		this.box.Box(min, max)
	}
}

//++++++++++++++++++++ BoundingSphere ++++++++++++++++++++

type BoundingSphere struct {
	min    *geom.Vector4
	max    *geom.Vector4
	center *geom.Vector4
	radius float32
	sphere *geom.Sphere
}

func (this *BoundingSphere) BoundingSphere(min *geom.Vector4, max *geom.Vector4) {
	this.min = &geom.Vector4{X: min.X, Y: min.Y, Z: min.Z, W: min.W}
	this.max = &geom.Vector4{X: max.X, Y: max.Y, Z: max.Z, W: max.W}
	this.center = &geom.Vector4{X: max.X - min.X, Y: max.Y - min.Y, Z: max.Z - min.Z, W: 1.0}
	this.radius = geom.Vector4Distance(this.min, this.max)
	this.sphere.Sphere(this.center, this.radius)
}
func (this *BoundingSphere) GetType() string {
	return SPHERE
}
func (this *BoundingSphere) GetBoundingBox() *geom.Sphere {
	return this.sphere
}
func (this *BoundingSphere) Update(transform *geom.Matrix4x4) {
	if transform != nil {
		min := transform.TransformVector(this.min)
		max := transform.TransformVector(this.max)
		this.center = &geom.Vector4{X: max.X - min.X, Y: max.Y - min.Y, Z: max.Z - min.Z, W: 1.0}
		this.radius = geom.Vector4Distance(this.min, this.max)
		this.sphere.Sphere(this.center, this.radius)
	}
}

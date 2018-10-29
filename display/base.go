package display

import (
	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/seven/geom"
	"github.com/liuqi0826/shenmegui/core"
)

//++++++++++++++++++++ Object ++++++++++++++++++++

type Object struct {
	core.Unit

	X         float32
	Y         float32
	Z         float32
	RotationX float32
	RotationY float32
	RotationZ float32
	ScaleX    float32
	ScaleY    float32
	ScaleZ    float32

	position *geom.Vector4
	rotation *geom.Vector4
	scale    *geom.Vector4

	transform *geom.Matrix4x4
	changed   bool
}

func (this *Object) Object() {
	this.Unit.Unit()

	this.ScaleX = 1.0
	this.ScaleY = 1.0
	this.ScaleZ = 1.0

	this.position = new(geom.Vector4)
	this.position.Vector4()
	this.rotation = new(geom.Vector4)
	this.rotation.Vector4()
	this.scale = new(geom.Vector4)
	this.scale.Vector4()
	this.scale.X = this.ScaleX
	this.scale.Y = this.ScaleY
	this.scale.Z = this.ScaleZ

	this.transform = new(geom.Matrix4x4)
	this.transform.Matrix4x4(nil)
}
func (this *Object) GetPosition() *geom.Vector4 {
	this.position.X = this.X
	this.position.Y = this.Y
	this.position.Z = this.Z
	return this.position
}
func (this *Object) GetRotation() *geom.Vector4 {
	this.rotation.X = this.RotationX
	this.rotation.Y = this.RotationY
	this.rotation.Z = this.RotationZ
	return this.rotation
}
func (this *Object) GetScale() *geom.Vector4 {
	this.scale.X = this.ScaleX
	this.scale.Y = this.ScaleY
	this.scale.Z = this.ScaleZ
	return this.scale
}
func (this *Object) GetTransform() *geom.Matrix4x4 {
	if this.changed {
		components := [3]*geom.Vector4{}
		components[0] = this.position
		components[1] = this.rotation
		components[2] = this.scale
		this.transform.Recompose(components, geom.EULER_ANGLES)
		this.changed = false
	}
	return this.transform
}
func (this *Object) Update() {
	if this.position.X != this.X {
		this.position.X = this.X
		this.changed = true
	}
	if this.position.Y != this.Y {
		this.position.Y = this.Y
		this.changed = true
	}
	if this.position.Z != this.Z {
		this.position.Z = this.Z
		this.changed = true
	}

	if this.rotation.X != this.RotationX {
		this.rotation.X = this.RotationX
		this.changed = true
	}
	if this.rotation.Y != this.RotationY {
		this.rotation.Y = this.RotationY
		this.changed = true
	}
	if this.rotation.Z != this.RotationZ {
		this.rotation.Z = this.RotationZ
		this.changed = true
	}

	if this.scale.X != this.ScaleX {
		if this.ScaleX <= 0 {
			this.ScaleX = 0.00000000001
		}
		this.scale.X = this.ScaleX
		this.changed = true
	}
	if this.scale.Y != this.ScaleY {
		if this.ScaleY <= 0 {
			this.ScaleY = 0.00000000001
		}
		this.scale.Y = this.ScaleY
		this.changed = true
	}
	if this.scale.Z != this.ScaleZ {
		if this.ScaleZ <= 0 {
			this.ScaleZ = 0.00000000001
		}
		this.scale.Z = this.ScaleZ
		this.changed = true
	}
}

//++++++++++++++++++++ DisplayObject ++++++++++++++++++++

type DisplayObject struct {
	Object

	root      IContainer
	parent    IContainer
	renderer  IRenderer
	camera    *Camera
	layerMask int32
}

func (this *DisplayObject) DisplayObject() {
	this.Object.Object()
}
func (this *DisplayObject) GetRoot() IContainer {
	return this.root
}
func (this *DisplayObject) SetRoot(root IContainer) {
	this.root = root
}
func (this *DisplayObject) GetParent() IContainer {
	return this.parent
}
func (this *DisplayObject) SetParent(parent IContainer) {
	this.parent = parent
}
func (this *DisplayObject) GetLayerMask() int32 {
	return this.layerMask
}
func (this *DisplayObject) SetLayerMask(mask int32) {
	this.layerMask = mask
}
func (this *DisplayObject) GetRenderer() IRenderer {
	return this.renderer
}
func (this *DisplayObject) SetRenderer(renderer IRenderer) {
	this.renderer = renderer
}
func (this *DisplayObject) GetCamera() *Camera {
	return this.camera
}
func (this *DisplayObject) SetCamera(camera *Camera) {
	this.camera = camera
}
func (this *DisplayObject) Update(transform *geom.Matrix4x4) {
	this.Object.Update()

	if transform != nil {
		this.GetTransform().Append(transform)
	}
}
func (this *DisplayObject) Render() {
}

//++++++++++++++++++++ DisplayObjectContainer ++++++++++++++++++++

type DisplayObjectContainer struct {
	DisplayObject

	displayList []IDisplayObject
}

func (this *DisplayObjectContainer) DisplayObjectContainer() {
	this.DisplayObject.DisplayObject()
	this.displayList = make([]IDisplayObject, 0)
}
func (this *DisplayObjectContainer) AddChild(displayObject IDisplayObject) {
	this.displayList = append(this.displayList, displayObject)

	displayObject.SetRoot(this.GetRoot())
	displayObject.SetParent(this)

	event := new(events.Event)
	event.Type = events.ADDED
	displayObject.DispatchEvent(event)
}
func (this *DisplayObjectContainer) RemoveChild(displayObject IDisplayObject) IDisplayObject {
	for i, c := range this.displayList {
		if c == displayObject {
			this.displayList = append(this.displayList[:i], this.displayList[i+1:]...)
			c.SetRoot(nil)
			c.SetParent(nil)
			event := new(events.Event)
			event.Type = events.REMOVE
			c.DispatchEvent(event)
			return c
		}
	}
	return nil
}
func (this *DisplayObjectContainer) RemoveAllChildren() {
	for _, c := range this.displayList {
		this.RemoveChild(c)
	}
	this.displayList = make([]IDisplayObject, 0)
}
func (this *DisplayObjectContainer) GetChildByName(name string) IDisplayObject {
	for _, c := range this.displayList {
		if c.GetName() == name {
			return c
		}
	}
	return nil
}
func (this *DisplayObjectContainer) SetRoot(root IContainer) {
	this.root = root
	for _, c := range this.displayList {
		c.SetRoot(this.root)
	}
}
func (this *DisplayObjectContainer) GetChildrenNumber() int32 {
	return int32(len(this.displayList))
}
func (this *DisplayObjectContainer) Render() {
	for _, c := range this.displayList {
		c.Render()
	}
}

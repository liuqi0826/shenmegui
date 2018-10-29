package display

import (
	"github.com/liuqi0826/seven/geom"
)

//++++++++++++++++++++ Mesh ++++++++++++++++++++

type Mesh struct {
	DisplayObject

	Max    *geom.Vector4
	Min    *geom.Vector4
	Center *geom.Vector4

	renderer IRenderer
	geometry []*SubGeometry
	material *Material
	shader   *Program3D
}

func (this *Mesh) Mesh(geometry []*SubGeometry, material *Material, shader *Program3D) {
	this.DisplayObject.DisplayObject()

	this.geometry = geometry
	this.material = material
	this.shader = shader

	this.Max = &geom.Vector4{X: geometry[0].Max.X, Y: geometry[0].Max.Y, Z: geometry[0].Max.Z, W: 1.0}
	this.Min = &geom.Vector4{X: geometry[0].Min.X, Y: geometry[0].Min.Y, Z: geometry[0].Min.Z, W: 1.0}
	for _, sg := range this.geometry {
		if this.Max.X < sg.Max.X {
			this.Max.X = sg.Max.X
		}
		if this.Max.Y < sg.Max.Y {
			this.Max.Y = sg.Max.Y
		}
		if this.Max.Z < sg.Max.Z {
			this.Max.Z = sg.Max.Z
		}
		if this.Min.X > sg.Min.X {
			this.Min.X = sg.Min.X
		}
		if this.Min.Y > sg.Min.Y {
			this.Min.Y = sg.Min.Y
		}
		if this.Min.Z > sg.Min.Z {
			this.Min.Z = sg.Min.Z
		}
	}
	this.Center = &geom.Vector4{X: this.Max.X - this.Min.X, Y: this.Max.Y - this.Min.Y, Z: this.Max.Z - this.Min.Z, W: 1.0}

	this.renderer = CreateRender("default")
	if this.renderer != nil {
		this.renderer.Setup(this.GetCamera(), this.shader)
	}
}
func (this *Mesh) SetCamera(camera *Camera) {
	this.DisplayObject.SetCamera(camera)
	if this.renderer != nil {
		this.renderer.SetCamera(camera)
	}
}
func (this *Mesh) Update(transform *geom.Matrix4x4) {
	this.DisplayObject.Update(transform)

	if this.renderer != nil {
		this.renderer.SetValue("transform", this.DisplayObject.GetTransform().GetRawSlice())
	}
}
func (this *Mesh) Render() {
	if this.renderer != nil {
		this.material.Bind()
		for _, sg := range this.geometry {
			this.renderer.Render(sg)
		}
	}
}

//++++++++++++++++++++ Sprite ++++++++++++++++++++

type Sprite struct {
	Mesh

	animation []IAnimation
}

func (this *Sprite) Sprite(geometry []*SubGeometry, material *Material, shader *Program3D) {
	this.DisplayObject.DisplayObject()

	this.geometry = geometry
	this.material = material
	this.shader = shader

	this.Max = &geom.Vector4{X: geometry[0].Max.X, Y: geometry[0].Max.Y, Z: geometry[0].Max.Z, W: 1.0}
	this.Min = &geom.Vector4{X: geometry[0].Min.X, Y: geometry[0].Min.Y, Z: geometry[0].Min.Z, W: 1.0}
	for _, sg := range this.geometry {
		if this.Max.X < sg.Max.X {
			this.Max.X = sg.Max.X
		}
		if this.Max.Y < sg.Max.Y {
			this.Max.Y = sg.Max.Y
		}
		if this.Max.Z < sg.Max.Z {
			this.Max.Z = sg.Max.Z
		}
		if this.Min.X > sg.Min.X {
			this.Min.X = sg.Min.X
		}
		if this.Min.Y > sg.Min.Y {
			this.Min.Y = sg.Min.Y
		}
		if this.Min.Z > sg.Min.Z {
			this.Min.Z = sg.Min.Z
		}
	}
	this.Center = &geom.Vector4{X: this.Max.X - this.Min.X, Y: this.Max.Y - this.Min.Y, Z: this.Max.Z - this.Min.Z, W: 1.0}

	this.renderer = CreateRender("animation")
	if this.renderer != nil {
		this.renderer.Setup(this.GetCamera(), this.shader)
	}

	this.animation = make([]IAnimation, 0)
}
func (this *Sprite) SetCamera(camera *Camera) {
	this.DisplayObject.SetCamera(camera)
	if this.renderer != nil {
		this.renderer.SetCamera(camera)
	}
}
func (this *Sprite) AddAnimation(anim IAnimation) {
	this.animation = append(this.animation, anim)
}
func (this *Sprite) Update(transform *geom.Matrix4x4) {
	this.Mesh.Update(transform)
	for _, anim := range this.animation {
		anim.Update()
		anim.Apply(this.renderer)
	}
}

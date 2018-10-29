package display

import "github.com/liuqi0826/seven/geom"

type SkyBox struct {
	DisplayObject

	renderer IRenderer
	geometry []*SubGeometry
	material *Material
	shader   *Program3D
}

func (this *SkyBox) SkyBox(material *Material) {
	this.DisplayObject.DisplayObject()

	this.geometry = make([]*SubGeometry, 0)
	instance := ResourceManagerInstance.CreateSubgeometrie("sky")
	this.geometry = append(this.geometry, instance)
	this.material = material
	this.shader = ResourceManagerInstance.CreateShaderProgram("default")

	this.renderer = CreateRender("skybox")
	if this.renderer != nil {
		this.renderer.Setup(this.GetCamera(), this.shader)
	}
}
func (this *SkyBox) SetCamera(camera *Camera) {
	this.DisplayObject.SetCamera(camera)
	if this.renderer != nil {
		this.renderer.SetCamera(camera)
	}
}
func (this *SkyBox) Update(transform *geom.Matrix4x4) {
	if transform != nil {
		compose := transform.Decompose(geom.EULER_ANGLES)
		compose[0] = new(geom.Vector4)
		compose[2] = new(geom.Vector4)
		mtx := new(geom.Matrix4x4)
		mtx.Recompose(compose, geom.EULER_ANGLES)
		this.DisplayObject.Update(mtx)
	} else {
		this.DisplayObject.Update(transform)
	}

	if this.renderer != nil {
		this.renderer.SetValue("transform", this.DisplayObject.GetTransform().GetRawSlice())
	}
}
func (this *SkyBox) Render() {
	if this.renderer != nil {
		this.material.Bind()
		for _, sg := range this.geometry {
			this.renderer.Render(sg)
		}
	}
}

package display

import (
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/liuqi0826/shenmegui/api"
)

func CreateRender(id string) IRenderer {
	switch id {
	case "default":
		renderer := new(DefaultRender)
		return renderer
	case "skybox":
		renderer := new(SkyboxRender)
		return renderer
	case "animation":
		renderer := new(AnimationRender)
		return renderer
	}
	return nil
}

//++++++++++++++++++++ ForwordRender ++++++++++++++++++++

func ForwordRender() {

}

//++++++++++++++++++++ DefaultRender ++++++++++++++++++++

type DefaultRender struct {
	ready   bool
	camera  *Camera
	program *Program3D

	positionLocation uint32
	texcoordLocation uint32
	normalLocation   uint32

	cameraLocation     js.Value
	transformLocation  js.Value
	projectionLocation js.Value

	sampler0Location js.Value

	valueTransform []float32
}

func (this *DefaultRender) Setup(camera *Camera, program *Program3D) {
	if camera != nil {
		this.camera = camera
	}
	this.program = program
	this.binding()
}
func (this *DefaultRender) SetCamera(camera *Camera) {
	if camera != nil {
		this.camera = camera
		this.binding()
	}
}
func (this *DefaultRender) SetProgram(program *Program3D) {
	this.program = program
	this.binding()
}
func (this *DefaultRender) binding() {
	if this.camera != nil && this.program != nil {
		gl.UseProgram(this.program.Program)

		this.positionLocation = gl.GetAttribLocation(this.program.Program, "position")
		gl.EnableVertexAttribArray(uint32(this.positionLocation))
		this.texcoordLocation = gl.GetAttribLocation(this.program.Program, "texcoord")
		gl.EnableVertexAttribArray(uint32(this.texcoordLocation))
		this.normalLocation = gl.GetAttribLocation(this.program.Program, "normal")
		gl.EnableVertexAttribArray(uint32(this.normalLocation))

		this.cameraLocation = gl.GetUniformLocation(this.program.Program, "camera")
		this.transformLocation = gl.GetUniformLocation(this.program.Program, "transform")
		this.projectionLocation = gl.GetUniformLocation(this.program.Program, "projection")

		this.sampler0Location = gl.GetUniformLocation(this.program.Program, "sampler0")
		gl.Uniform1i(this.sampler0Location, 0)

		this.ready = true
		fmt.Println("DefaultRender binding is done!")
	} else {
		this.ready = false
	}
}
func (this *DefaultRender) SetValue(title string, value []float32) {
	switch title {
	case TRANSFORM:
		this.valueTransform = value
	}
}
func (this *DefaultRender) Render(target IRenderable) {
	if this.ready && target != nil && target.IsReady() {
		gl.UseProgram(this.program.Program)

		carmera := this.camera.GetTransformMatrix().Raw[:16]
		projection := this.camera.GetProjectionMatrix().Raw[:16]
		transform := this.valueTransform[:16]

		gl.UniformMatrix4fv(this.cameraLocation, false, carmera)
		gl.UniformMatrix4fv(this.projectionLocation, false, projection)
		gl.UniformMatrix4fv(this.transformLocation, false, transform)

		vlist := target.GetVertexBuffer()
		for _, vertexBuffer := range *vlist {
			if vertexBuffer != nil {
				gl.BindBuffer(api.GL_ARRAY_BUFFER, vertexBuffer.Buffer)
				gl.VertexAttribPointer(this.positionLocation, 3, api.GL_FLOAT, false, 8*4, 0)
				gl.VertexAttribPointer(this.texcoordLocation, 2, api.GL_FLOAT, false, 8*4, 3*4)
				gl.VertexAttribPointer(this.normalLocation, 3, api.GL_FLOAT, false, 8*4, 5*4)
			}
		}

		indexBuffer := target.GetIndexBuffer()
		gl.BindBuffer(api.GL_ELEMENT_ARRAY_BUFFER, indexBuffer.Buffer)
		gl.DrawElements(api.GL_TRIANGLES, int32(indexBuffer.Length), api.GL_UNSIGNED_SHORT, 0)
	}
}

//++++++++++++++++++++ SkyboxRender ++++++++++++++++++++

type SkyboxRender struct {
	ready   bool
	camera  *Camera
	program *Program3D

	positionLocation uint32
	texcoordLocation uint32
	normalLocation   uint32

	cameraLocation     js.Value
	transformLocation  js.Value
	projectionLocation js.Value

	sampler0Location js.Value

	valueTransform []float32
}

func (this *SkyboxRender) Setup(camera *Camera, program *Program3D) {
	if camera != nil {
		this.camera = camera
	}
	this.program = program
	this.binding()
}
func (this *SkyboxRender) SetCamera(camera *Camera) {
	if camera != nil {
		this.camera = camera
		this.binding()
	}
}
func (this *SkyboxRender) SetProgram(program *Program3D) {
	this.program = program
	this.binding()
}
func (this *SkyboxRender) binding() {
	if this.camera != nil && this.program != nil {
		gl.UseProgram(this.program.Program)

		this.positionLocation = gl.GetAttribLocation(this.program.Program, "position")
		gl.EnableVertexAttribArray(uint32(this.positionLocation))
		this.texcoordLocation = gl.GetAttribLocation(this.program.Program, "texcoord")
		gl.EnableVertexAttribArray(uint32(this.texcoordLocation))
		this.normalLocation = gl.GetAttribLocation(this.program.Program, "normal")
		gl.EnableVertexAttribArray(uint32(this.normalLocation))

		this.cameraLocation = gl.GetUniformLocation(this.program.Program, "camera")
		this.transformLocation = gl.GetUniformLocation(this.program.Program, "transform")
		this.projectionLocation = gl.GetUniformLocation(this.program.Program, "projection")

		this.sampler0Location = gl.GetUniformLocation(this.program.Program, "sampler0")
		gl.Uniform1i(this.sampler0Location, 0)

		this.ready = true
		fmt.Println("SkyboxRender binding is done!")
	} else {
		this.ready = false
	}
}
func (this *SkyboxRender) SetValue(title string, value []float32) {
	switch title {
	case TRANSFORM:
		this.valueTransform = value
	}
}
func (this *SkyboxRender) Render(target IRenderable) {
	if this.ready && target != nil && target.IsReady() {
		gl.UseProgram(this.program.Program)

		cmtx := this.camera.GetTransformMatrix().Clone()
		carmera := cmtx.Raw[:16]
		carmera[12] = 0.0
		carmera[13] = 0.0
		carmera[14] = 0.0
		projection := this.camera.GetProjectionMatrix().Raw[:16]
		transform := this.valueTransform[:16]

		gl.UniformMatrix4fv(this.cameraLocation, false, carmera)
		gl.UniformMatrix4fv(this.projectionLocation, false, projection)
		gl.UniformMatrix4fv(this.transformLocation, false, transform)

		vlist := target.GetVertexBuffer()
		for _, vertexBuffer := range *vlist {
			if vertexBuffer != nil {
				gl.BindBuffer(api.GL_ARRAY_BUFFER, vertexBuffer.Buffer)
				gl.VertexAttribPointer(this.positionLocation, 3, api.GL_FLOAT, false, 8*4, 0)
				gl.VertexAttribPointer(this.texcoordLocation, 2, api.GL_FLOAT, false, 8*4, 3*4)
				gl.VertexAttribPointer(this.normalLocation, 3, api.GL_FLOAT, false, 8*4, 5*4)
			}
		}

		indexBuffer := target.GetIndexBuffer()
		gl.BindBuffer(api.GL_ELEMENT_ARRAY_BUFFER, indexBuffer.Buffer)
		gl.DrawElements(api.GL_TRIANGLES, int32(indexBuffer.Length), api.GL_UNSIGNED_SHORT, 0)
	}
}

//++++++++++++++++++++ AnimationRender ++++++++++++++++++++

type AnimationRender struct {
	ready   bool
	camera  *Camera
	program *Program3D

	positionLocation uint32
	texcoordLocation uint32
	normalLocation   uint32
	indexLocation    uint32
	weightLocation   uint32

	cameraLocation        js.Value
	transformLocation     js.Value
	projectionLocation    js.Value
	texcoordValueLocation js.Value
	skeletonValueLocation []js.Value

	sampler0Location js.Value

	valueTransform []float32
	valueTexcoord  []float32
	valueSkeleton  []float32
}

func (this *AnimationRender) Setup(camera *Camera, program *Program3D) {
	if camera != nil {
		this.camera = camera
	}
	this.program = program
	this.binding()

	this.valueTexcoord = []float32{0, 0}
}
func (this *AnimationRender) SetCamera(camera *Camera) {
	if camera != nil {
		this.camera = camera
		this.binding()
	}
}
func (this *AnimationRender) SetProgram(program *Program3D) {
	this.program = program
	this.binding()
}
func (this *AnimationRender) binding() {
	if this.camera != nil && this.program != nil {
		gl.UseProgram(this.program.Program)

		this.positionLocation = gl.GetAttribLocation(this.program.Program, "position")
		gl.EnableVertexAttribArray(uint32(this.positionLocation))
		this.texcoordLocation = gl.GetAttribLocation(this.program.Program, "texcoord")
		gl.EnableVertexAttribArray(uint32(this.texcoordLocation))
		this.normalLocation = gl.GetAttribLocation(this.program.Program, "normal")
		gl.EnableVertexAttribArray(uint32(this.normalLocation))
		this.indexLocation = gl.GetAttribLocation(this.program.Program, "index")
		gl.EnableVertexAttribArray(uint32(this.indexLocation))
		this.weightLocation = gl.GetAttribLocation(this.program.Program, "weight")
		gl.EnableVertexAttribArray(uint32(this.weightLocation))

		this.cameraLocation = gl.GetUniformLocation(this.program.Program, "camera")
		this.transformLocation = gl.GetUniformLocation(this.program.Program, "transform")
		this.projectionLocation = gl.GetUniformLocation(this.program.Program, "projection")
		this.texcoordValueLocation = gl.GetUniformLocation(this.program.Program, "texcoord_value")

		this.skeletonValueLocation = make([]js.Value, 480)
		for i := 0; i < 480; i++ {
			this.skeletonValueLocation[i] = gl.GetUniformLocation(this.program.Program, "skeleton_value["+strconv.Itoa(i)+"]")
		}

		this.sampler0Location = gl.GetUniformLocation(this.program.Program, "sampler0")
		gl.Uniform1i(this.sampler0Location, 0)

		this.ready = true
		fmt.Println("AnimationRender binding is done!")
	} else {
		this.ready = false
	}
}
func (this *AnimationRender) SetValue(title string, value []float32) {
	switch title {
	case TRANSFORM:
		this.valueTransform = value
	case TEXCOORD:
		this.valueTexcoord = value
	case SKELETON:
		this.valueSkeleton = value
	}
}

func (this *AnimationRender) Render(target IRenderable) {
	if this.ready && target != nil && target.IsReady() {
		gl.UseProgram(this.program.Program)

		carmera := this.camera.GetTransformMatrix().Raw[:16]
		projection := this.camera.GetProjectionMatrix().Raw[:16]
		transform := this.valueTransform[:16]

		gl.UniformMatrix4fv(this.cameraLocation, false, carmera)
		gl.UniformMatrix4fv(this.projectionLocation, false, projection)
		gl.UniformMatrix4fv(this.transformLocation, false, transform)
		gl.Uniform2fv(this.texcoordValueLocation, this.valueTexcoord)

		skeletonUniformLenght := len(this.valueSkeleton) / 4
		for i := 0; i < skeletonUniformLenght; i++ {
			gl.Uniform4f(this.skeletonValueLocation[i], this.valueSkeleton[i*4], this.valueSkeleton[i*4+1], this.valueSkeleton[i*4+2], this.valueSkeleton[i*4+3])
		}

		vlist := target.GetVertexBuffer()
		for idx, vertexBuffer := range *vlist {
			if vertexBuffer != nil {
				switch idx {
				case 0:
					gl.BindBuffer(api.GL_ARRAY_BUFFER, vertexBuffer.Buffer)
					gl.VertexAttribPointer(this.positionLocation, 3, api.GL_FLOAT, false, 8*4, 0)
					gl.VertexAttribPointer(this.texcoordLocation, 2, api.GL_FLOAT, false, 8*4, 3*4)
					gl.VertexAttribPointer(this.normalLocation, 3, api.GL_FLOAT, false, 8*4, 5*4)
				case 1:
					gl.BindBuffer(api.GL_ARRAY_BUFFER, vertexBuffer.Buffer)
					gl.VertexAttribPointer(this.indexLocation, 4, api.GL_FLOAT, false, 8*4, 0)
					gl.VertexAttribPointer(this.weightLocation, 4, api.GL_FLOAT, false, 8*4, 4*4)
				}
			}
		}

		indexBuffer := target.GetIndexBuffer()
		gl.BindBuffer(api.GL_ELEMENT_ARRAY_BUFFER, indexBuffer.Buffer)
		gl.DrawElements(api.GL_TRIANGLES, int32(indexBuffer.Length), api.GL_UNSIGNED_SHORT, 0)
	}
}

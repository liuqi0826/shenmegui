package display

import (
	"fmt"

	"github.com/liuqi0826/seven/geom"
)

type SubGeometry struct {
	ID string

	Max    *geom.Vector4
	Min    *geom.Vector4
	Center *geom.Vector4

	IndexBuffer  *IndexBuffer
	VertexBuffer [8]*VertexBuffer

	userCount int32
	uploaded  bool

	geometryResource *GeometryResource
}

func (this *SubGeometry) SubGeometry(geometryResource *GeometryResource) {
	if geometryResource != nil {
		this.geometryResource = geometryResource
		this.Max = &geom.Vector4{X: geometryResource.Max[0], Y: geometryResource.Max[1], Z: geometryResource.Max[2], W: 1.0}
		this.Min = &geom.Vector4{X: geometryResource.Min[0], Y: geometryResource.Min[1], Z: geometryResource.Min[2], W: 1.0}
		this.Center = &geom.Vector4{X: geometryResource.Center[0], Y: geometryResource.Center[1], Z: geometryResource.Center[2], W: 1.0}
	} else {
		fmt.Println("GeometryResource is nil.")
	}
}
func (this *SubGeometry) Upload(context *Context3D) error {
	var err error
	this.IndexBuffer = context.CreateIndexBuffer()
	if this.IndexBuffer != nil {
		err = this.IndexBuffer.Upload(this.geometryResource.Geometrie.Index)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("IndexBuffer is nil")
	}
	for i, v := range this.geometryResource.Geometrie.Vertex {
		if v != nil {
			this.VertexBuffer[i] = context.CreateVertexBuffer()
			if this.VertexBuffer[i] != nil {
				err = this.VertexBuffer[i].Upload(v)
				if err != nil {
					return err
				}
			}
		}
	}
	this.uploaded = true
	return err
}
func (this *SubGeometry) GetIndexBuffer() *IndexBuffer {
	return this.IndexBuffer
}
func (this *SubGeometry) GetVertexBuffer() *[8]*VertexBuffer {
	return &this.VertexBuffer
}
func (this *SubGeometry) IsReady() bool {
	return this.uploaded
}
func (this *SubGeometry) AddCount() {
	this.userCount++
}
func (this *SubGeometry) SubCount() {
	if this.userCount > 0 {
		this.userCount--
	}
}
func (this *SubGeometry) Dispose() {
	this.uploaded = false
}

package display

import (
	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/seven/geom"
)

type IController interface {
	Update()
}

type IRenderer interface {
	Setup(camera *Camera, program3D *Program3D)
	SetCamera(camera *Camera)
	SetProgram(program *Program3D)
	SetValue(title string, value []float32)
	Render(renderable IRenderable)
}

type IRenderable interface {
	IsReady() bool
	GetIndexBuffer() *IndexBuffer
	GetVertexBuffer() *[8]*VertexBuffer
}

type IEntity interface {
	GetBound(boundType string) ICollisionTest
}

type ICollisionTest interface {
	GetType() string
	GetBoundingBox() interface{}
}

type IDisplayObject interface {
	events.IEventDispatcher

	GetID() uint32
	GetName() string
	GetRoot() IContainer
	SetRoot(root IContainer)
	GetParent() IContainer
	SetParent(parent IContainer)
	GetCamera() *Camera
	SetCamera(camera *Camera)
	GetLayerMask() int32
	SetLayerMask(int32)
	GetRenderer() IRenderer
	SetRenderer(renderer IRenderer)
	Update(transform *geom.Matrix4x4)
	Render()
}

type IContainer interface {
	events.IEventDispatcher

	AddChild(displayObject IDisplayObject)
	RemoveChild(displayObject IDisplayObject) IDisplayObject
	RemoveAllChildren()
	GetChildByName(name string) IDisplayObject
	GetChildrenNumber() int32
}

type IAnimation interface {
	GetType() string
	Update()
	Apply(renderer IRenderer)
}

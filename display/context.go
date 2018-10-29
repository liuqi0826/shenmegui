package display

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"syscall/js"

	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/shenmegui/api"
	"github.com/liuqi0826/shenmegui/global"
)

var canvas js.Value
var context js.Value
var gl *api.WebGL

var MainContext3D *Context3D

//++++++++++++++++++++ Context3D ++++++++++++++++++++

type Context3D struct {
	events.EventDispatcher

	currentShaderProgram *Program3D

	ready bool

	clearRed   float32
	clearGreen float32
	clearBlue  float32
	clearAlpha float32

	depthMask       bool
	passCompareMode string

	close chan int
}

func (this *Context3D) Setup(canvasID string) error {
	var err error
	this.EventDispatcher.EventDispatcher(this)

	canvas = global.Document.Call("getElementById", canvasID)
	context = canvas.Call("getContext", "webgl")
	if context == js.Undefined() {
		context = canvas.Call("getContext", "experimental-webgl")
	}
	if context == js.Undefined() {
		err = errors.New("browser might not support webgl")
		return err
	}
	this.close = make(chan int)

	MainContext3D = this

	canvas.Set("width", global.Config.WindowWidth)
	canvas.Set("height", global.Config.WindowHeight)
	go func() {
		resiazeEvt := js.NewEventCallback(0, func(e js.Value) {
			global.Config.WindowWidth = int32(global.Document.Get("body").Get("clientWidth").Int())
			global.Config.WindowHeight = int32(global.Document.Get("documentElement").Get("clientHeight").Int())
			fmt.Println(global.Config.WindowWidth, global.Config.WindowHeight)

			evt := new(events.Event)
			evt.Type = events.RESIZE
			this.DispatchEvent(evt)
		})
		defer resiazeEvt.Release()
		global.Window.Call("addEventListener", "resize", resiazeEvt)
		<-this.close
	}()

	gl = new(api.WebGL)
	gl.Setup(context)
	gl.Enable(api.GL_DEPTH_TEST)

	this.ready = true

	return err
}
func (this *Context3D) Clear(color bool, depth bool, stencil bool) {
	var mask uint32
	if color {
		mask = mask | api.GL_COLOR_BUFFER_BIT
	}
	if depth {
		mask = mask | api.GL_DEPTH_BUFFER_BIT
	}
	if stencil {
		mask = mask | api.GL_STENCIL_BUFFER_BIT
	}
	if this.ready {
		gl.Clear(mask)
	}
}
func (this *Context3D) ConfigureBackBuffer() {
}
func (this *Context3D) CreateCubeTexture() {
}
func (this *Context3D) CreateProgram() *Program3D {
	return new(Program3D)
}
func (this *Context3D) CreateTexture() {
}
func (this *Context3D) CreateIndexBuffer() *IndexBuffer {
	indexBuffer := new(IndexBuffer)
	indexBuffer.Buffer = gl.CreateBuffer()
	return indexBuffer
}
func (this *Context3D) CreateVertexBuffer() *VertexBuffer {
	vertexBuffer := new(VertexBuffer)
	vertexBuffer.Buffer = gl.CreateBuffer()
	return vertexBuffer
}
func (this *Context3D) Dispose() {
	this.close <- 0
}
func (this *Context3D) Present() {
}
func (this *Context3D) DrawTriangles(indexBuffer *IndexBuffer, firstIndex int32, numTriangles int32) {
	gl.DrawElements(api.GL_TRIANGLES, numTriangles, api.GL_UNSIGNED_SHORT, 0)
}
func (this *Context3D) SetBlendFactors() {
}
func (this *Context3D) SetClearColor(red float32, green float32, blue float32, alpha float32) {
	this.clearRed = red
	this.clearGreen = green
	this.clearBlue = blue
	this.clearAlpha = alpha
	if this.ready {
		gl.ClearColor(this.clearRed, this.clearGreen, this.clearBlue, this.clearAlpha)
	} else {
		fmt.Println("GL is not ready.")
	}
}
func (this *Context3D) SetColorMask() {
}
func (this *Context3D) SetCulling() {
}
func (this *Context3D) SetDepthTest(depthMask bool, passCompareMode string) {
	this.depthMask = depthMask
	this.passCompareMode = passCompareMode
	if this.ready {
		if this.depthMask {
			gl.Enable(api.GL_DEPTH_TEST)
		} else {
		}
		switch this.passCompareMode {
		case ALWAYS:
			gl.DepthFunc(api.GL_ALWAYS)
		case EQUAL:
			gl.DepthFunc(api.GL_EQUAL)
		case GREATER:
			gl.DepthFunc(api.GL_GREATER)
		case GREATER_EQUAL:
		case LESS:
			gl.DepthFunc(api.GL_LESS)
		case LESS_EQUAL:
		case NEVER:
			gl.DepthFunc(api.GL_NEVER)
		case NOT_EQUAL:
		}
	}
}
func (this *Context3D) SetProgram(program *Program3D) {
	this.currentShaderProgram = program
	gl.UseProgram(program.Program)
}
func (this *Context3D) SetProgramConstantsFromByteArray() {
}
func (this *Context3D) SetProgramConstantsFromMatrix() {
}
func (this *Context3D) SetProgramConstantsFromVector() {
}
func (this *Context3D) SetRenderToBackBuffer() {
}
func (this *Context3D) SetRenderToTexture() {
}
func (this *Context3D) SetScissorRectangle() {
}
func (this *Context3D) SetStencilActions() {
}
func (this *Context3D) SetStencilReferenceValue() {
}
func (this *Context3D) SetTextureAt() {
}
func (this *Context3D) SetVertexBufferAt(value string, stride int32, bufferOffset int32, format string) {
	if this.currentShaderProgram != nil {
		var size int32
		var xtype uint32

		switch format {
		case FLOAT_1:
			size = 1
			xtype = api.GL_FLOAT
		case FLOAT_2:
			size = 2
			xtype = api.GL_FLOAT
		case FLOAT_3:
			size = 3
			xtype = api.GL_FLOAT
		case FLOAT_4:
			size = 4
			xtype = api.GL_FLOAT
		case BYTES_4:
			size = 4
			xtype = api.GL_BYTE
		}

		attrib := uint32(gl.GetAttribLocation(this.currentShaderProgram.Program, value))
		gl.EnableVertexAttribArray(attrib)
		gl.VertexAttribPointer(attrib, size, xtype, false, stride, bufferOffset)
	}
}
func (this *Context3D) CursorLock() {
	canvas.Call("requestPointerLock")
}
func (this *Context3D) CursorUnlock() {
	global.Document.Call("exitPointerLock")
}

//++++++++++++++++++++ IndexBuffer ++++++++++++++++++++

type IndexBuffer struct {
	Index  uint32
	Length int
	Buffer js.Value
}

func (this *IndexBuffer) Upload(data []uint16) error {
	var err error
	this.Length = len(data)
	if this.Length > 65535 {
		err = errors.New("Upload index data is large than 65535.")
	} else {
		gl.BindBuffer(api.GL_ELEMENT_ARRAY_BUFFER, this.Buffer)
		gl.BufferData(api.GL_ELEMENT_ARRAY_BUFFER, data, api.GL_STATIC_DRAW)
	}
	return err
}
func (this *IndexBuffer) Dispose() {
}

//++++++++++++++++++++ VertexBuffer ++++++++++++++++++++

type VertexBuffer struct {
	Index  uint32
	Buffer js.Value
}

func (this *VertexBuffer) Upload(data []float32) error {
	var err error
	gl.BindBuffer(api.GL_ARRAY_BUFFER, this.Buffer)
	gl.BufferData(api.GL_ARRAY_BUFFER, data, api.GL_STATIC_DRAW)
	return err
}
func (this *VertexBuffer) Dispose() {
}

//++++++++++++++++++++ Program3D ++++++++++++++++++++

type Program3D struct {
	VertexShader   js.Value
	FragmentShader js.Value
	Program        js.Value
	userCount      uint32
}

func (this *Program3D) Upload(vertexProgram string, fragmentProgram string) error {
	var err error
	this.VertexShader, err = this.compileShader(vertexProgram, api.GL_VERTEX_SHADER)
	if err != nil {
		fmt.Println(err)
		return err
	}
	this.FragmentShader, err = this.compileShader(fragmentProgram, api.GL_FRAGMENT_SHADER)
	if err != nil {
		fmt.Println(err)
		return err
	}

	this.Program = gl.CreateProgram()
	gl.AttachShader(this.Program, this.VertexShader)
	gl.AttachShader(this.Program, this.FragmentShader)
	gl.LinkProgram(this.Program)

	pname := gl.GetProgramParameter(this.Program, api.GL_LINK_STATUS)
	if pv, ok := pname.(js.Value); ok {
		if !pv.Bool() {
			log := gl.GetProgramInfoLog(this.Program)
			fmt.Println(log)
		}
	}
	fmt.Println("Program3D Upload", pname)

	return err
}
func (this *Program3D) Dispose() {
	gl.DeleteShader(this.Program)
}
func (this *Program3D) AddCount() {
	this.userCount++
}
func (this *Program3D) SubCount() {
	this.userCount--
}
func (this *Program3D) GetCount() uint32 {
	return this.userCount
}
func (this *Program3D) compileShader(source string, shaderType uint32) (js.Value, error) {
	shader := gl.CreateShader(shaderType)
	gl.ShaderSource(shader, source)
	gl.CompileShader(shader)

	pname := gl.GetShaderParameter(shader, api.GL_COMPILE_STATUS)
	if pv, ok := pname.(js.Value); ok {
		if !pv.Bool() {
			log := gl.GetShaderInfoLog(shader)
			fmt.Println(log)
		}
	}
	fmt.Println("Program3D compileShader", pname)

	return shader, nil
}

//++++++++++++++++++++ Texture ++++++++++++++++++++

type Texture struct {
	Type    string
	Texture js.Value
}

func (this *Texture) Upload(source []byte, sourceType string) error {
	var err error
	this.Type = sourceType

	buff := bytes.NewBuffer(source)
	img, err := png.Decode(buff)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		fmt.Println("Unsupported stride.")
		return err
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	this.Texture = gl.CreateTexture()
	gl.ActiveTexture(api.GL_TEXTURE0)
	gl.BindTexture(api.GL_TEXTURE_2D, this.Texture)
	gl.TexImage2D(api.GL_TEXTURE_2D, 0, api.GL_RGBA, int32(img.Bounds().Dx()), int32(img.Bounds().Dy()), 0, api.GL_RGBA, api.GL_UNSIGNED_BYTE, rgba.Pix)
	gl.TexParameteri(api.GL_TEXTURE_2D, api.GL_TEXTURE_MAG_FILTER, api.GL_LINEAR)
	gl.TexParameteri(api.GL_TEXTURE_2D, api.GL_TEXTURE_MIN_FILTER, api.GL_LINEAR)
	gl.TexParameteri(api.GL_TEXTURE_2D, api.GL_TEXTURE_WRAP_S, api.GL_CLAMP_TO_EDGE)
	gl.TexParameteri(api.GL_TEXTURE_2D, api.GL_TEXTURE_WRAP_T, api.GL_CLAMP_TO_EDGE)
	return err
}
func (this *Texture) UploadCompressedTexture(source []byte) error {
	var err error
	return err
}
func (this *Texture) SetSlot(index int32) {
	var solt uint32
	switch index {
	case 0:
		solt = api.GL_TEXTURE0
	case 1:
		solt = api.GL_TEXTURE1
	case 2:
		solt = api.GL_TEXTURE2
	case 3:
		solt = api.GL_TEXTURE3
	case 4:
		solt = api.GL_TEXTURE4
	case 5:
		solt = api.GL_TEXTURE5
	case 6:
		solt = api.GL_TEXTURE6
	case 7:
		solt = api.GL_TEXTURE7
	}
	gl.ActiveTexture(solt)
	gl.BindTexture(api.GL_TEXTURE_2D, this.Texture)
}
func (this *Texture) Dispose() {

}

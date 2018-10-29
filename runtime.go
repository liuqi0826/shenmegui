package shenmegui

import (
	"runtime"
	"sync"
	"syscall/js"

	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/shenmegui/display"
	"github.com/liuqi0826/shenmegui/global"
	"github.com/liuqi0826/shenmegui/io"
	"github.com/liuqi0826/shenmegui/sound"
)

var Runtime *Engine
var ResourceManager *Resource

func init() {
	runtime.LockOSThread()

	Runtime = new(Engine)
	ResourceManager = new(Resource)

	Runtime.Engine()
}

//++++++++++++++++++++ Engine ++++++++++++++++++++

type Engine struct {
	events.EventDispatcher
	sync.Mutex

	Alive     bool
	Ready     bool
	Stage     *display.Stage
	Context3D *display.Context3D

	KeyboardManeger *io.KeyboardManeger
	MouseManeger    *io.MouseManeger

	actionList    []func()
	actionChannel chan func()

	close chan int
}

func (this *Engine) Engine() {
	this.actionList = make([]func(), 0)
	this.actionChannel = make(chan func(), 256)
	this.close = make(chan int)
}
func (this *Engine) Setup() {
	this.Context3D = new(display.Context3D)
	this.Context3D.Setup("context")

	this.KeyboardManeger = io.KeyboardInstance
	this.MouseManeger = io.MouseInstance

	ResourceManager.Setup()
	ResourceManager.AddEventListener(global.RESOURCE_EVENT, this.onResourceEvent)

	this.Ready = true
}
func (this *Engine) Start(function func()) {
	if this.Ready {
		this.Alive = true
		this.Stage = new(display.Stage)
		this.Stage.Setup()

		go this.actionListion()

		if function != nil {
			go function()
		}

		go func() {
			var renderFrame js.Callback
			renderFrame = js.NewCallback(func(args []js.Value) {
				this.frame()
				js.Global().Call("requestAnimationFrame", renderFrame)
			})
			defer renderFrame.Release()
			js.Global().Call("requestAnimationFrame", renderFrame)
			<-this.close
		}()
	}
}
func (this *Engine) Stop() {
	this.Alive = false
}
func (this *Engine) AddActionOSThread(function func()) {
	this.actionChannel <- function
}
func (this *Engine) CursorLock() {
	this.Context3D.CursorLock()
}
func (this *Engine) CursorUnlock() {
	this.Context3D.CursorUnlock()
}

func (this *Engine) onResourceEvent(event *events.Event) {
	this.AddActionOSThread(ResourceManager.displayResource.Upload)
}
func (this *Engine) actionListion() {
	for act := range this.actionChannel {
		this.Lock()
		this.actionList = append(this.actionList, act)
		this.Unlock()
	}
}
func (this *Engine) action() {
	this.Lock()
	defer this.Unlock()
	for _, fun := range this.actionList {
		if fun != nil {
			fun()
		}
	}
	this.actionList = make([]func(), 0)
}
func (this *Engine) frame() {
	this.action()
	this.Stage.Frame()
}

//++++++++++++++++++++ ResourceManager ++++++++++++++++++++

type Resource struct {
	events.EventDispatcher

	displayResource *display.ResourceManager
	soundResource   *sound.ResourceManager
}

func (this *Resource) Setup() {
	this.displayResource = display.ResourceManagerInstance
	this.displayResource.Setup()
	this.displayResource.AddEventListener(global.RESOURCE_EVENT, this.onEvent)
	
	this.soundResource = sound.ResourceManagerInstance
	this.soundResource.Setup()
	this.soundResource.AddEventListener(global.RESOURCE_EVENT, this.onEvent)
}
func (this *Resource) LoadGeometrie(url string) error {
	return this.displayResource.LoadGeometrie(url)
}
func (this *Resource) LoadGeometrieCompress(url string) error {
	return this.displayResource.LoadGeometrieCompress(url)
}
func (this *Resource) LoadSkeleton(url string) error {
	return this.displayResource.LoadSkeleton(url)
}
func (this *Resource) LoadSkeletonCompress(url string) error {
	return this.displayResource.LoadSkeletonCompress(url)
}
func (this *Resource) LoadMaterial(url string) error {
	return this.displayResource.LoadMaterial(url)
}
func (this *Resource) LoadMaterialCompress(url string) error {
	return this.displayResource.LoadMaterialCompress(url)
}
func (this *Resource) LoadAnimationClip(url string) error {
	return this.displayResource.LoadAnimationClip(url)
}
func (this *Resource) LoadAnimationClipCompress(url string) error {
	return this.displayResource.LoadAnimationClipCompress(url)
}
func (this *Resource) LoadShader(url string) error {
	return this.displayResource.LoadShader(url)
}
func (this *Resource) LoadShaderCompress(url string) error {
	return this.displayResource.LoadShaderCompress(url)
}
func (this *Resource) ParserGeometrie(value []byte) string {
	return this.displayResource.ParserGeometrie(value)
}
func (this *Resource) ParserSkeleton(value []byte) string {
	return this.displayResource.ParserSkeleton(value)
}
func (this *Resource) ParserMaterial(value []byte) string {
	return this.displayResource.ParserMaterial(value)
}
func (this *Resource) ParserAnimationClip(value []byte) string {
	return this.displayResource.ParserAnimationClip(value)
}
func (this *Resource) ParserShader(value []byte) string {
	return this.displayResource.ParserShader(value)
}
func (this *Resource) AddGeometrie(value *display.GeometryResource) {
	this.displayResource.AddGeometrie(value)
}
func (this *Resource) GetGeometrie(id string) *display.GeometryResource {
	return this.displayResource.GetGeometrie(id)
}
func (this *Resource) AddSkeleton(value *display.SkeletonResource) {
	this.displayResource.AddSkeleton(value)
}
func (this *Resource) GetSkeleton(id string) *display.SkeletonResource {
	return this.displayResource.GetSkeleton(id)
}
func (this *Resource) AddMaterial(value *display.MaterialResource) {
	this.displayResource.AddMaterial(value)
}
func (this *Resource) GetMaterial(id string) *display.MaterialResource {
	return this.displayResource.GetMaterial(id)
}
func (this *Resource) AddAnimationClip(value *display.AnimationClipResource) {
	this.displayResource.AddAnimationClip(value)
}
func (this *Resource) GetAnimationClip(id string) *display.AnimationClipResource {
	return this.displayResource.GetAnimationClip(id)
}
func (this *Resource) AddShader(value *display.ShaderResource) {
	this.displayResource.AddShader(value)
}
func (this *Resource) GetShader(id string) *display.ShaderResource {
	return this.displayResource.GetShader(id)
}

func (this *Resource) CreateSubgeometrie(id string) *display.SubGeometry {
	return this.displayResource.CreateSubgeometrie(id)
}
func (this *Resource) CreateMaterial(id string) *display.Material {
	return this.displayResource.CreateMaterial(id)
}
func (this *Resource) CreateShaderProgram(id string) *display.Program3D {
	return this.displayResource.CreateShaderProgram(id)
}
func (this *Resource) Upload() {
	this.displayResource.Upload()
}
func (this *Resource) onEvent(event *events.Event) {
	this.DispatchEvent(event)
}

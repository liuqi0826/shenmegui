package display

import (
	"bytes"
	"compress/zlib"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/shenmegui/global"
)

var ResourceManagerInstance *ResourceManager
var GeometryList map[string]string

func init() {
	GeometryList = make(map[string]string)

	GeometryList["sky"] = `{
		"id":"sky",
		"geometrie":{
			"index":[0,1,2, 0,2,3, 4,5,6, 4,6,7, 8,9,10, 8,10,11, 12,13,14, 12,14,15, 16,17,18, 16,18,19, 20,21,22, 20,22,23],
			"vertex":[
				[
					-50,50,50,		0,0.5,		0,-1,0,
					50,50,50,		0.25,0.5,	0,-1,0,
					50,50,-50,		0.25,1,		0,-1,0,
					-50,50,-50,		0,1,		0,-1,0,

					-50,-50,50,		0.25,0.5,	0,1,0,
					50,-50,50,		0.5,0.5,	0,1,0,
					50,-50,-50,		0.5,1,		0,1,0,
					-50,-50,-50,	0.25,1,		0,1,0,

					-50,50,50,		0,0,		0,0,-1,
					50,50,50,		0.25,0,		0,0,-1,
					50,-50,50,		0.25,0.5,	0,0,-1,
					-50,-50,50,		0,0.5,		0,0,-1,

					-50,50,-50,		0.5,0,		1,0,0,
					-50,50,50,		0.25,0,		1,0,0,
					-50,-50,50,		0.25,0.5,	1,0,0,
					-50,-50,-50,	0.5,0.5,	1,0,0,

					50,50,-50,		0.75,0,		0,0,1,
					-50,50,-50,		0.5,0,		0,0,1,
					-50,-50,-50,	0.5,0.5,	0,0,1,
					50,-50,-50,		0.75,0.5,	0,0,1,

					50,50,50,		1,0,		-1,0,0,
					50,50,-50,		0.75,0,		-1,0,0,
					50,-50,-50,		0.75,0.5,	-1,0,0,
					50,-50,50,		1,0.5,		-1,0,0
				],
				null,
				null,
				null,
				null,
				null,
				null,
				null
				]
		},
		"center":[0,50,0],
		"material":"sky",
		"skined":false,
		"shader":"default",
		"max":[50,100,50],
		"radius":86.60254037844386,
		"min":[-50,0,-50]
		}`

	ResourceManagerInstance = new(ResourceManager)
}

//++++++++++++++++++++ ResourcePackage ++++++++++++++++++++

type ResourcePackage struct {
	ID       string
	Geometry []byte
	Material []byte
	Skeleton []byte
	Shader   string
}

//++++++++++++++++++++ AnimationClipResource ++++++++++++++++++++

type AnimationChannel struct {
	Transform []float32 `json:"transform"`
	Time      []int32   `json:"time"`
}
type AnimationClipResource struct {
	ID          string              `json:"id"`
	Duration    int32               `json:"duration"`
	UploadIndex []uint32            `json:"uploadIndex"`
	Channel     []*AnimationChannel `json:"channel"`
}

func (this *AnimationClipResource) Parser(value []byte) error {
	var err error
	err = json.Unmarshal(value, this)
	return err
}

//++++++++++++++++++++ GeometryResource ++++++++++++++++++++

type GeometryResource struct {
	ID        string              `json:"id"`
	Shader    string              `json:"shader"`
	Material  string              `json:"material"`
	Max       []float32           `json:"max"`
	Min       []float32           `json:"min"`
	Center    []float32           `json:"center"`
	Radius    float32             `json:"radius"`
	Skined    bool                `json:"skined"`
	Geometrie *GeometrieComponent `json:"geometrie"`
}
type GeometrieComponent struct {
	Index  []uint16     `json:"index"`
	Vertex [8][]float32 `json:"vertex"`
}

func (this *GeometryResource) Parser(value []byte) error {
	var err error
	err = json.Unmarshal(value, this)
	return err
}

//++++++++++++++++++++ MaterialResource ++++++++++++++++++++

type MaterialResource struct {
	ID          string
	TextureList [8]*TextureResource
}

func (this *MaterialResource) Parser(value []byte) error {
	var err error
	reader := bytes.NewBuffer(value)
	decoder := gob.NewDecoder(reader)
	err = decoder.Decode(this)
	return err
}

//++++++++++++++++++++ ShaderResource ++++++++++++++++++++

type ShaderResource struct {
	ID       string `json:"id"`
	Vertex   string `json:"vertex"`
	Fragment string `json:"fragment"`
}

func (this *ShaderResource) Parser(value []byte) error {
	var err error
	err = json.Unmarshal(value, this)
	return err
}

//++++++++++++++++++++ SkeletonResource ++++++++++++++++++++

type SkeletonResource struct {
	ID            string        `json:"id"`
	Struct        []interface{} `json:"struct"`
	Root          []float32     `json:"root"`
	Foot          []float32     `json:"foot"`
	BindShape     []float32     `json:"bindShape"`
	Invbind       [][]float32   `json:"invbind"`
	NodeNameList  []string      `json:"nodeNameList"`
	JointNameList []string      `json:"jointNameList"`
}

func (this *SkeletonResource) Parser(value []byte) error {
	var err error
	err = json.Unmarshal(value, this)
	if this.Root == nil || len(this.Root) == 0 {
		this.Root = []float32{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	}
	if this.Foot == nil || len(this.Foot) == 0 {
		this.Foot = []float32{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	}
	return err
}

//++++++++++++++++++++ TextureResource ++++++++++++++++++++

type TextureResource struct {
	ID      string
	Type    string
	Texture []byte
}

//++++++++++++++++++++ ResourceManager ++++++++++++++++++++

type ResourceManager struct {
	sync.Mutex
	events.EventDispatcher

	//网络加载原始资源
	protoResource map[string][]byte

	//静态资源库
	geometryResource  map[string]*GeometryResource
	skeletonResource  map[string]*SkeletonResource
	materialResource  map[string]*MaterialResource
	animationResource map[string]*AnimationClipResource
	shaderResource    map[string]*ShaderResource

	//运行时资源
	geometryRuntime map[string]*SubGeometry
	materialRuntime map[string]*Material
	shaderRuntime   map[string]*Program3D

	shaderUploadQueue    []*ShaderUploader
	geometrieUploadQueue []*GeometrieUploader
	materialUploadQueue  []*MaterialUploader
}

func (this *ResourceManager) Setup() error {
	var err error

	this.EventDispatcher.EventDispatcher(this)

	if MainContext3D == nil {
		err = errors.New("Context3D is nil.")
	}

	this.protoResource = make(map[string][]byte)

	this.geometryResource = make(map[string]*GeometryResource)
	this.skeletonResource = make(map[string]*SkeletonResource)
	this.materialResource = make(map[string]*MaterialResource)
	this.animationResource = make(map[string]*AnimationClipResource)
	this.shaderResource = make(map[string]*ShaderResource)

	this.geometryRuntime = make(map[string]*SubGeometry)
	this.materialRuntime = make(map[string]*Material)
	this.shaderRuntime = make(map[string]*Program3D)

	this.shaderUploadQueue = make([]*ShaderUploader, 0)
	this.geometrieUploadQueue = make([]*GeometrieUploader, 0)
	this.materialUploadQueue = make([]*MaterialUploader, 0)

	this.pretreatment()

	return err
}
func (this *ResourceManager) LoadGeometrie(url string) error {
	var err error
	res, err := this.load(url)
	if err != nil {
		return err
	}
	id := this.ParserGeometrie(res)
	if id != "" {
		fmt.Println("Load Geometrie " + id)
	}
	return err
}
func (this *ResourceManager) LoadGeometrieCompress(url string) error {
	var err error
	res, err := this.load(url)
	if err != nil {
		return err
	}

	buff := bytes.NewReader(res)
	reader, err := zlib.NewReader(buff)
	defer reader.Close()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	id := this.ParserGeometrie(data)
	if id != "" {
		fmt.Println("Load Geometrie " + id)
	}
	return err
}
func (this *ResourceManager) LoadSkeleton(url string) error {
	var err error
	res, err := this.load(url)
	if err != nil {
		return err
	}
	id := this.ParserSkeleton(res)
	if id != "" {
		fmt.Println("Load Skeleton " + id)
	}
	return err
}
func (this *ResourceManager) LoadSkeletonCompress(url string) error {
	var err error
	res, err := this.load(url)
	if err != nil {
		return err
	}

	buff := bytes.NewReader(res)
	reader, err := zlib.NewReader(buff)
	defer reader.Close()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	id := this.ParserSkeleton(data)
	if id != "" {
		fmt.Println("Load Skeleton " + id)
	}
	return err
}
func (this *ResourceManager) LoadMaterial(url string) error {
	var err error
	res, err := this.load(url)
	if err != nil {
		return err
	}
	id := this.ParserMaterial(res)
	if id != "" {
		fmt.Println("Load Material " + id)
	}
	return err
}
func (this *ResourceManager) LoadMaterialCompress(url string) error {
	var err error
	res, err := this.load(url)
	if err != nil {
		return err
	}

	buff := bytes.NewReader(res)
	reader, err := zlib.NewReader(buff)
	defer reader.Close()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	id := this.ParserMaterial(data)
	if id != "" {
		fmt.Println("Load Material " + id)
	}
	return err
}
func (this *ResourceManager) LoadAnimationClip(url string) error {
	var err error
	res, err := this.load(url)
	if err != nil {
		return err
	}
	id := this.ParserAnimationClip(res)
	if id != "" {
		fmt.Println("Load Animation " + id)
	}
	return err
}
func (this *ResourceManager) LoadAnimationClipCompress(url string) error {
	var err error
	res, err := this.load(url)
	if err != nil {
		return err
	}

	buff := bytes.NewReader(res)
	reader, err := zlib.NewReader(buff)
	defer reader.Close()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	id := this.ParserAnimationClip(data)
	if id != "" {
		fmt.Println("Load Animation " + id)
	}
	return err
}
func (this *ResourceManager) LoadShader(url string) error {
	var err error
	res, err := this.load(url)
	if err != nil {
		return err
	}
	id := this.ParserShader(res)
	if id != "" {
		fmt.Println("Load Shader " + id)
	}
	return err
}
func (this *ResourceManager) LoadShaderCompress(url string) error {
	var err error
	res, err := this.load(url)
	if err != nil {
		return err
	}

	buff := bytes.NewReader(res)
	reader, err := zlib.NewReader(buff)
	defer reader.Close()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	id := this.ParserShader(data)
	if id != "" {
		fmt.Println("Load Shader " + id)
	}
	return err
}
func (this *ResourceManager) ParserGeometrie(value []byte) string {
	gr := new(GeometryResource)
	err := gr.Parser(value)
	if err == nil {
		this.AddGeometrie(gr)
		return gr.ID
	} else {
		fmt.Println(err)
	}
	return ""
}
func (this *ResourceManager) ParserSkeleton(value []byte) string {
	sk := new(SkeletonResource)
	err := sk.Parser(value)
	if err == nil {
		this.AddSkeleton(sk)
		return sk.ID
	} else {
		fmt.Println(err)
	}
	return ""
}
func (this *ResourceManager) ParserMaterial(value []byte) string {
	mt := new(MaterialResource)
	err := mt.Parser(value)
	if err == nil {
		this.AddMaterial(mt)
		return mt.ID
	} else {
		fmt.Println(err)
	}
	return ""
}
func (this *ResourceManager) ParserAnimationClip(value []byte) string {
	am := new(AnimationClipResource)
	err := am.Parser(value)
	if err == nil {
		this.AddAnimationClip(am)
		return am.ID
	} else {
		fmt.Println(err)
	}
	return ""
}
func (this *ResourceManager) ParserShader(value []byte) string {
	sr := new(ShaderResource)
	err := sr.Parser(value)
	if err == nil {
		this.AddShader(sr)
		return sr.ID
	} else {
		fmt.Println(err)
	}
	return ""
}

func (this *ResourceManager) AddGeometrie(value *GeometryResource) {
	this.geometryResource[value.ID] = value
}
func (this *ResourceManager) GetGeometrie(id string) *GeometryResource {
	if gr, ok := this.geometryResource[id]; ok {
		return gr
	}
	return nil
}
func (this *ResourceManager) AddSkeleton(value *SkeletonResource) {
	this.skeletonResource[value.ID] = value
}
func (this *ResourceManager) GetSkeleton(id string) *SkeletonResource {
	if sk, ok := this.skeletonResource[id]; ok {
		return sk
	}
	return nil
}
func (this *ResourceManager) AddMaterial(value *MaterialResource) {
	this.materialResource[value.ID] = value
}
func (this *ResourceManager) GetMaterial(id string) *MaterialResource {
	if mt, ok := this.materialResource[id]; ok {
		return mt
	}
	return nil
}
func (this *ResourceManager) AddAnimationClip(value *AnimationClipResource) {
	this.animationResource[value.ID] = value
}
func (this *ResourceManager) GetAnimationClip(id string) *AnimationClipResource {
	if an, ok := this.animationResource[id]; ok {
		return an
	}
	return nil
}
func (this *ResourceManager) AddShader(value *ShaderResource) {
	this.shaderResource[value.ID] = value
}
func (this *ResourceManager) GetShader(id string) *ShaderResource {
	if resource, ok := this.shaderResource[id]; ok {
		return resource
	}
	return nil
}

func (this *ResourceManager) CreateSubgeometrie(id string) *SubGeometry {
	if geometry, ok := this.geometryRuntime[id]; ok {
		geometry.AddCount()
		return geometry
	} else {
		resource := this.GetGeometrie(id)
		if resource != nil {
			subGeometry := new(SubGeometry)
			subGeometry.SubGeometry(resource)
			subGeometry.AddCount()
			this.geometryRuntime[id] = subGeometry

			this.Lock()
			gu := new(GeometrieUploader)
			gu.Target = subGeometry
			gu.Resource = resource
			this.geometrieUploadQueue = append(this.geometrieUploadQueue, gu)
			this.Unlock()

			evt := new(events.Event)
			evt.Type = global.RESOURCE_EVENT
			this.DispatchEvent(evt)

			return subGeometry
		}
	}
	return nil
}
func (this *ResourceManager) CreateMaterial(id string) *Material {
	if material, ok := this.materialRuntime[id]; ok {
		return material
	} else {
		resource := this.GetMaterial(id)
		if resource != nil {
			material := new(Material)
			material.Material(resource)
			material.AddCount()
			this.materialRuntime[id] = material

			this.Lock()
			mu := new(MaterialUploader)
			mu.Target = material
			mu.Resource = resource
			this.materialUploadQueue = append(this.materialUploadQueue, mu)
			this.Unlock()

			evt := new(events.Event)
			evt.Type = global.RESOURCE_EVENT
			this.DispatchEvent(evt)

			return material
		} else {
			fmt.Println("no material res.")
		}
	}
	return nil
}
func (this *ResourceManager) CreateShaderProgram(id string) *Program3D {
	if programe, ok := this.shaderRuntime[id]; ok {
		programe.AddCount()
		return programe
	} else {
		resource := this.GetShader(id)
		if resource != nil {
		} else {
			fmt.Println("no shader res.")
		}
	}
	return nil
}
func (this *ResourceManager) Upload() {
	this.Lock()
	defer this.Unlock()

	var su, gu, mu bool
	for _, s := range this.shaderUploadQueue {
		s.Target = MainContext3D.CreateProgram()
		s.Target.Upload(s.Resource.Vertex, s.Resource.Fragment)
		su = true
	}
	for _, g := range this.geometrieUploadQueue {
		g.Target.Upload(MainContext3D)
		gu = true
	}
	for _, m := range this.materialUploadQueue {
		m.Target.Upload()
		mu = true
	}
	if su {
		this.shaderUploadQueue = make([]*ShaderUploader, 0)
	}
	if gu {
		this.geometrieUploadQueue = make([]*GeometrieUploader, 0)
	}
	if mu {
		this.materialUploadQueue = make([]*MaterialUploader, 0)
	}
}
func (this *ResourceManager) load(url string) ([]byte, error) {
	var target []byte
	if res, has := this.protoResource[url]; has {
		target = res
	} else {
		resp, err := http.Post(url, "application/x-www-form-urlencode", nil)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		target = body
	}
	return target, nil
}
func (this *ResourceManager) pretreatment() {
	for _, v := range ShaderList {
		this.AddShader(v)

		//现阶段所有shader在开始时统一创建上传
		shader := MainContext3D.CreateProgram()
		shader.Upload(v.Vertex, v.Fragment)
		this.shaderRuntime[v.ID] = shader
	}

	for _, g := range GeometryList {
		this.ParserGeometrie([]byte(g))
	}
}

type ShaderUploader struct {
	Target   *Program3D
	Resource *ShaderResource
}
type GeometrieUploader struct {
	Target   *SubGeometry
	Resource *GeometryResource
}
type MaterialUploader struct {
	Target   *Material
	Resource *MaterialResource
}

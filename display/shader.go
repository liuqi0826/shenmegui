package display

import "fmt"

var ShaderList map[string]*ShaderResource

func init() {
	ShaderList = make(map[string]*ShaderResource)

	ShaderList["default"] = new(ShaderResource)
	ShaderList["default"].ID = "default"
	ShaderList["default"].Vertex = `
	attribute vec3 position;
	attribute vec2 texcoord;
	attribute vec3 normal;

	uniform mat4 camera;
	uniform mat4 projection;
	uniform mat4 transform;

	varying vec2 vtc;
	varying vec3 vn;

	void main() {
		vtc = texcoord;
		vn = normal;
		gl_Position = projection * camera * transform * vec4(position, 1.0);
	}`
	ShaderList["default"].Fragment = `
	precision mediump float;

	uniform sampler2D sampler0;

	varying vec2 vtc;
	varying vec3 vn;

	void main() {
		gl_FragColor = texture2D(sampler0, vtc);
	}`

	ShaderList["skeleton_matrix4x4"] = new(ShaderResource)
	ShaderList["skeleton_matrix4x4"].ID = "skeleton_matrix4x4"
	ShaderList["skeleton_matrix4x4"].Vertex = `
	attribute vec3 position;
	attribute vec2 texcoord;
	attribute vec3 normal;

	attribute vec4 index;
	attribute vec4 weight;

	uniform mat4 camera;
	uniform mat4 projection;
	uniform mat4 transform;

	const int count = 480;
	uniform vec2 texcoord_value;
	uniform vec4 skeleton_value[count];

	varying vec2 vtc;
	varying vec3 vn;

	void main() {
		vtc = texcoord;
		vn = normal;
		
		vec4 g_position = vec4(position, 1.0);
		vec4 t_position = vec4(0.0, 0.0, 0.0, 1.0);

		mat4 mtx1 = mat4(
			skeleton_value[int(index.x)*4+0].x, skeleton_value[int(index.x)*4+0].y, skeleton_value[int(index.x)*4+0].z, skeleton_value[int(index.x)*4+0].w, 
			skeleton_value[int(index.x)*4+1].x, skeleton_value[int(index.x)*4+1].y, skeleton_value[int(index.x)*4+1].z, skeleton_value[int(index.x)*4+1].w,
			skeleton_value[int(index.x)*4+2].x, skeleton_value[int(index.x)*4+2].y, skeleton_value[int(index.x)*4+2].z, skeleton_value[int(index.x)*4+2].w,
			skeleton_value[int(index.x)*4+3].x, skeleton_value[int(index.x)*4+3].y, skeleton_value[int(index.x)*4+3].z, skeleton_value[int(index.x)*4+3].w
		);
		t_position = mtx1 * g_position * weight.x;
		mat4 mtx2 = mat4(
			skeleton_value[int(index.y)*4+0].x, skeleton_value[int(index.y)*4+0].y, skeleton_value[int(index.y)*4+0].z, skeleton_value[int(index.y)*4+0].w, 
			skeleton_value[int(index.y)*4+1].x, skeleton_value[int(index.y)*4+1].y, skeleton_value[int(index.y)*4+1].z, skeleton_value[int(index.y)*4+1].w,
			skeleton_value[int(index.y)*4+2].x, skeleton_value[int(index.y)*4+2].y, skeleton_value[int(index.y)*4+2].z, skeleton_value[int(index.y)*4+2].w,
			skeleton_value[int(index.y)*4+3].x, skeleton_value[int(index.y)*4+3].y, skeleton_value[int(index.y)*4+3].z, skeleton_value[int(index.y)*4+3].w
		);
		t_position += mtx2 * g_position * weight.y;
		mat4 mtx3 = mat4(
			skeleton_value[int(index.z)*4+0].x, skeleton_value[int(index.z)*4+0].y, skeleton_value[int(index.z)*4+0].z, skeleton_value[int(index.z)*4+0].w, 
			skeleton_value[int(index.z)*4+1].x, skeleton_value[int(index.z)*4+1].y, skeleton_value[int(index.z)*4+1].z, skeleton_value[int(index.z)*4+1].w,
			skeleton_value[int(index.z)*4+2].x, skeleton_value[int(index.z)*4+2].y, skeleton_value[int(index.z)*4+2].z, skeleton_value[int(index.z)*4+2].w,
			skeleton_value[int(index.z)*4+3].x, skeleton_value[int(index.z)*4+3].y, skeleton_value[int(index.z)*4+3].z, skeleton_value[int(index.z)*4+3].w
		);
		t_position += mtx3 * g_position * weight.z;
		mat4 mtx4 = mat4(
			skeleton_value[int(index.w)*4+0].x, skeleton_value[int(index.w)*4+0].y, skeleton_value[int(index.w)*4+0].z, skeleton_value[int(index.w)*4+0].w, 
			skeleton_value[int(index.w)*4+1].x, skeleton_value[int(index.w)*4+1].y, skeleton_value[int(index.w)*4+1].z, skeleton_value[int(index.w)*4+1].w,
			skeleton_value[int(index.w)*4+2].x, skeleton_value[int(index.w)*4+2].y, skeleton_value[int(index.w)*4+2].z, skeleton_value[int(index.w)*4+2].w,
			skeleton_value[int(index.w)*4+3].x, skeleton_value[int(index.w)*4+3].y, skeleton_value[int(index.w)*4+3].z, skeleton_value[int(index.w)*4+3].w
		);
		t_position += mtx4 * g_position * weight.w;

		gl_Position = projection * camera * transform * t_position;
	}`
	ShaderList["skeleton_matrix4x4"].Fragment = `
	precision mediump float;

	uniform sampler2D sampler0;

	varying vec2 vtc;
	varying vec3 vn;

	void main() {
		gl_FragColor = texture2D(sampler0, vtc);
	}`

	ShaderList["skeleton_matrix4x3"] = new(ShaderResource)
	ShaderList["skeleton_matrix4x3"].ID = "skeleton_matrix4x3"
	ShaderList["skeleton_matrix4x3"].Vertex = `
	attribute vec3 position;
	attribute vec2 texcoord;
	attribute vec3 normal;

	attribute vec4 index;
	attribute vec4 weight;

	uniform mat4 camera;
	uniform mat4 projection;
	uniform mat4 transform;

	const int count = 480;
	uniform vec2 texcoord_value;
	uniform vec4 skeleton_value[count];

	varying vec2 vtc;
	varying vec3 vn;

	void main() {
		vtc = texcoord;
		vn = normal;
		
		vec4 g_position = vec4(position, 1.0);
		vec4 t_position = vec4(0.0, 0.0, 0.0, 1.0);
		
		mat4 mtx1 = mat4(
			skeleton_value[int(index.x)*3].x, skeleton_value[int(index.x)*3].y, skeleton_value[int(index.x)*3].z, skeleton_value[int(index.x)*3].w, 
			skeleton_value[int(index.x)*3+1].x, skeleton_value[int(index.x)*3+1].y, skeleton_value[int(index.x)*3+1].z, skeleton_value[int(index.x)*3+1].w,
			skeleton_value[int(index.x)*3+2].x, skeleton_value[int(index.x)*3+2].y, skeleton_value[int(index.x)*3+2].z, skeleton_value[int(index.x)*3+2].w,
			0, 0, 0, 1
		);
		t_position = mtx1 * g_position * weight.s;
		mat4 mtx2 = mat4(
			skeleton_value[int(index.y)*3].x, skeleton_value[int(index.y)*3].y, skeleton_value[int(index.y)*3].z, skeleton_value[int(index.y)*3].w, 
			skeleton_value[int(index.y)*3+1].x, skeleton_value[int(index.y)*3+1].y, skeleton_value[int(index.y)*3+1].z, skeleton_value[int(index.y)*3+1].w,
			skeleton_value[int(index.y)*3+2].x, skeleton_value[int(index.y)*3+2].y, skeleton_value[int(index.y)*3+2].z, skeleton_value[int(index.y)*3+2].w,
			0, 0, 0, 1
		);
		t_position += mtx2 * g_position * weight.y;
		mat4 mtx3 = mat4(
			skeleton_value[int(index.z)*3].x, skeleton_value[int(index.z)*3].y, skeleton_value[int(index.z)*3].z, skeleton_value[int(index.z)*3].w, 
			skeleton_value[int(index.z)*3+1].x, skeleton_value[int(index.z)*3+1].y, skeleton_value[int(index.z)*3+1].z, skeleton_value[int(index.z)*3+1].w,
			skeleton_value[int(index.z)*3+2].x, skeleton_value[int(index.z)*3+2].y, skeleton_value[int(index.z)*3+2].z, skeleton_value[int(index.z)*3+2].w,
			0, 0, 0, 1
		);
		t_position += mtx3 * g_position * weight.z;
		mat4 mtx4 = mat4(
			skeleton_value[int(index.w)*3].x, skeleton_value[int(index.w)*3].y, skeleton_value[int(index.w)*3].z, skeleton_value[int(index.w)*3].w, 
			skeleton_value[int(index.w)*3+1].x, skeleton_value[int(index.w)*3+1].y, skeleton_value[int(index.w)*3+1].z, skeleton_value[int(index.w)*3+1].w,
			skeleton_value[int(index.w)*3+2].x, skeleton_value[int(index.w)*3+2].y, skeleton_value[int(index.w)*3+2].z, skeleton_value[int(index.w)*3+2].w,
			0, 0, 0, 1
		);
		t_position += mtx4 * g_position * weight.w;

		gl_Position = projection * camera * transform * t_position;
	}`
	ShaderList["skeleton_matrix4x3"].Fragment = `
	precision mediump float;

	uniform sampler2D sampler0;

	varying vec2 vtc;
	varying vec3 vn;

	void main() {
		gl_FragColor = texture2D(sampler0, vtc);
	}`

	ShaderList["skeleton_quaternion"] = new(ShaderResource)
	ShaderList["skeleton_quaternion"].ID = "skeleton_quaternion"
	ShaderList["skeleton_quaternion"].Vertex = `
	attribute vec3 position;
	attribute vec2 texcoord;
	attribute vec3 normal;

	attribute vec4 index;
	attribute vec4 weight;

	uniform mat4 camera;
	uniform mat4 projection;
	uniform mat4 transform;

	const int count = 480;
	uniform vec2 texcoord_value;
	uniform vec4 skeleton_value[count];

	varying vec2 vtc;
	varying vec3 vn;

	void main() {
		vtc = texcoord;
		vn = normal;
		
		vec4 g_position = vec4(position, 1.0);
		vec4 t_position = vec4(0.0, 0.0, 0.0, 1.0);
		vec4 tt_position = vec4(0.0, 0.0, 0.0, 1.0);
		float tx, ty, tz, tw;

		tw = - skeleton_value[int(index.x)*2].x * g_position.x - skeleton_value[int(index.x)*2].y * g_position.y - skeleton_value[int(index.x)*2].z * g_position.z;
		tx = skeleton_value[int(index.x)*2].w * g_position.x + skeleton_value[int(index.x)*2].y * g_position.z - skeleton_value[int(index.x)*2].z * g_position.y;
		ty = skeleton_value[int(index.x)*2].w * g_position.y - skeleton_value[int(index.x)*2].x * g_position.z + skeleton_value[int(index.x)*2].z * g_position.x;
		tz = skeleton_value[int(index.x)*2].w * g_position.z + skeleton_value[int(index.x)*2].x * g_position.y - skeleton_value[int(index.x)*2].y * g_position.x;
		t_position.x = - tw * skeleton_value[int(index.x)*2].x + tx * skeleton_value[int(index.x)*2].w - ty * skeleton_value[int(index.x)*2].z + tz * skeleton_value[int(index.x)*2].y;
		t_position.y = - tw * skeleton_value[int(index.x)*2].y + tx * skeleton_value[int(index.x)*2].z + ty * skeleton_value[int(index.x)*2].w - tz * skeleton_value[int(index.x)*2].x;
		t_position.z = - tw * skeleton_value[int(index.x)*2].z - tx * skeleton_value[int(index.x)*2].y + ty * skeleton_value[int(index.x)*2].x + tz * skeleton_value[int(index.x)*2].w;
		t_position.x += skeleton_value[int(index.x)*2+1].x;
		t_position.y += skeleton_value[int(index.x)*2+1].y;
		t_position.z += skeleton_value[int(index.x)*2+1].z;
		t_position.x *= weight.x;
		t_position.y *= weight.x;
		t_position.z *= weight.x;

		tw = - skeleton_value[int(index.y)*2].x * g_position.x - skeleton_value[int(index.y)*2].y * g_position.y - skeleton_value[int(index.y)*2].z * g_position.z;
		tx = skeleton_value[int(index.y)*2].w * g_position.x + skeleton_value[int(index.y)*2].y * g_position.z - skeleton_value[int(index.y)*2].z * g_position.y;
		ty = skeleton_value[int(index.y)*2].w * g_position.y - skeleton_value[int(index.y)*2].x * g_position.z + skeleton_value[int(index.y)*2].z * g_position.x;
		tz = skeleton_value[int(index.y)*2].w * g_position.z + skeleton_value[int(index.y)*2].x * g_position.y - skeleton_value[int(index.y)*2].y * g_position.x;
		tt_position.x = - tw * skeleton_value[int(index.y)*2].x + tx * skeleton_value[int(index.y)*2].w - ty * skeleton_value[int(index.y)*2].z + tz * skeleton_value[int(index.y)*2].y;
		tt_position.y = - tw * skeleton_value[int(index.y)*2].y + tx * skeleton_value[int(index.y)*2].z + ty * skeleton_value[int(index.y)*2].w - tz * skeleton_value[int(index.y)*2].x;
		tt_position.z = - tw * skeleton_value[int(index.y)*2].z - tx * skeleton_value[int(index.y)*2].y + ty * skeleton_value[int(index.y)*2].x + tz * skeleton_value[int(index.y)*2].w;
		tt_position.x += skeleton_value[int(index.y)*2+1].x;
		tt_position.y += skeleton_value[int(index.y)*2+1].y;
		tt_position.z += skeleton_value[int(index.y)*2+1].z;
		t_position.x += tt_position.x * weight.y;
		t_position.y += tt_position.y * weight.y;
		t_position.z += tt_position.z * weight.y;

		tw = - skeleton_value[int(index.z)*2].x * g_position.x - skeleton_value[int(index.z)*2].y * g_position.y - skeleton_value[int(index.z)*2].z * g_position.z;
		tx = skeleton_value[int(index.z)*2].w * g_position.x + skeleton_value[int(index.z)*2].y * g_position.z - skeleton_value[int(index.z)*2].z * g_position.y;
		ty = skeleton_value[int(index.z)*2].w * g_position.y - skeleton_value[int(index.z)*2].x * g_position.z + skeleton_value[int(index.z)*2].z * g_position.x;
		tz = skeleton_value[int(index.z)*2].w * g_position.z + skeleton_value[int(index.z)*2].x * g_position.y - skeleton_value[int(index.z)*2].y * g_position.x;
		tt_position.x = - tw * skeleton_value[int(index.z)*2].x + tx * skeleton_value[int(index.z)*2].w - ty * skeleton_value[int(index.z)*2].z + tz * skeleton_value[int(index.z)*2].y;
		tt_position.y = - tw * skeleton_value[int(index.z)*2].y + tx * skeleton_value[int(index.z)*2].z + ty * skeleton_value[int(index.z)*2].w - tz * skeleton_value[int(index.z)*2].x;
		tt_position.z = - tw * skeleton_value[int(index.z)*2].z - tx * skeleton_value[int(index.z)*2].y + ty * skeleton_value[int(index.z)*2].x + tz * skeleton_value[int(index.z)*2].w;
		tt_position.x += skeleton_value[int(index.z)*2+1].x;
		tt_position.y += skeleton_value[int(index.z)*2+1].y;
		tt_position.z += skeleton_value[int(index.z)*2+1].z;
		t_position.x += tt_position.x * weight.z;
		t_position.y += tt_position.y * weight.z;
		t_position.z += tt_position.z * weight.z;

		tw = - skeleton_value[int(index.w)*2].x * g_position.x - skeleton_value[int(index.w)*2].y * g_position.y - skeleton_value[int(index.w)*2].z * g_position.z;
		tx = skeleton_value[int(index.w)*2].w * g_position.x + skeleton_value[int(index.w)*2].y * g_position.z - skeleton_value[int(index.w)*2].z * g_position.y;
		ty = skeleton_value[int(index.w)*2].w * g_position.y - skeleton_value[int(index.w)*2].x * g_position.z + skeleton_value[int(index.w)*2].z * g_position.x;
		tz = skeleton_value[int(index.w)*2].w * g_position.z + skeleton_value[int(index.w)*2].x * g_position.y - skeleton_value[int(index.w)*2].y * g_position.x;
		tt_position.x = - tw * skeleton_value[int(index.w)*2].x + tx * skeleton_value[int(index.w)*2].w - ty * skeleton_value[int(index.w)*2].z + tz * skeleton_value[int(index.w)*2].y;
		tt_position.y = - tw * skeleton_value[int(index.w)*2].y + tx * skeleton_value[int(index.w)*2].z + ty * skeleton_value[int(index.w)*2].w - tz * skeleton_value[int(index.w)*2].x;
		tt_position.z = - tw * skeleton_value[int(index.w)*2].z - tx * skeleton_value[int(index.w)*2].y + ty * skeleton_value[int(index.w)*2].x + tz * skeleton_value[int(index.w)*2].w;
		tt_position.x += skeleton_value[int(index.w)*2+1].x;
		tt_position.y += skeleton_value[int(index.w)*2+1].y;
		tt_position.z += skeleton_value[int(index.w)*2+1].z;
		t_position.x += tt_position.x * weight.w;
		t_position.y += tt_position.y * weight.w;
		t_position.z += tt_position.z * weight.w;

		gl_Position = projection * camera * transform * t_position;
	}`
	ShaderList["skeleton_quaternion"].Fragment = `
	precision mediump float;

	uniform sampler2D sampler0;

	varying vec2 vtc;
	varying vec3 vn;

	void main() {
		gl_FragColor = texture2D(sampler0, vtc);
	}`
}

type ShaderProgram struct {
	ID        string
	UsedCount int32
	Uploaded  bool

	Program  *Program3D
	Resource *ShaderResource
}

func (this *ShaderProgram) ShaderProgram(id string, Resource *ShaderResource) {
	this.ID = id
	this.Resource = Resource
}
func (this *ShaderProgram) Upload(context *Context3D) {
	this.Program = context.CreateProgram()
	err := this.Program.Upload(this.Resource.Vertex, this.Resource.Fragment)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		this.Uploaded = true
	}
}
func (this *ShaderProgram) Dispose() {
	this.Program.Dispose()
	this.Uploaded = false
}

type ShaderValue struct {
	Shader *ShaderProgram

	VertexData   []byte
	FragmentData []byte
}

func (this *ShaderValue) ShaderValue() {
	this.VertexData = make([]byte, 0)
	this.FragmentData = make([]byte, 0)
}
func (this *ShaderValue) UploadValue() {
}

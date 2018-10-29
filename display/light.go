package display

import (
	"github.com/liuqi0826/seven/geom"
)

type Light struct {
	Red       float32
	Green     float32
	Blue      float32
	Intensity float32
}

func (this *Light) Light(red float32, green float32, blue float32, intensity float32) {
	this.Red = red
	this.Green = green
	this.Blue = blue
	this.Intensity = intensity
}

//环境光
type AmbientLight struct {
	Light
}

func (this *AmbientLight) AmbientLight(red float32, green float32, blue float32, intensity float32) {
	this.Light.Light(red, green, blue, intensity)
}

//方向光
type DirectionalLight struct {
	Light

	Direction *geom.Vector4
}

func (this *DirectionalLight) DirectionalLight(red float32, green float32, blue float32, intensity float32, direction *geom.Vector4) {
	this.Light.Light(red, green, blue, intensity)
	this.Direction = direction
	if this.Direction == nil {
		this.Direction = new(geom.Vector4)
		this.Direction.Vector4()
	}
	this.Direction.Normalize()
}

//点光源
type PointLight struct {
	Light
}

func (this *PointLight) PointLight() {

}

//聚光灯
type SpotLight struct {
	Light
}

func (this *SpotLight) SpotLight() {

}

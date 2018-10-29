package display

type Material struct {
	ID string

	texture [8]*Texture

	userCount          uint32
	uploadTextureCount int32
	uploaded           bool

	materialResource *MaterialResource
}

func (this *Material) Material(materialResource *MaterialResource) {
	if materialResource != nil {
		this.ID = materialResource.ID
		this.materialResource = materialResource
		this.texture = [8]*Texture{}
	}
}
func (this *Material) Upload() {
	if this.materialResource != nil {
		for i, t := range this.materialResource.TextureList {
			if i < 8 && len(t.Texture) > 0 {
				text := new(Texture)
				text.Upload(t.Texture, t.Type)
				this.texture[i] = text
			}
		}
		this.uploaded = true
	}
}
func (this *Material) Bind() {
	for i, t := range this.texture {
		if t != nil {
			idx := int32(i)
			t.SetSlot(idx)
		}
	}
}
func (this *Material) AddCount() {
	this.userCount++
}
func (this *Material) SubCount() {
	if this.userCount > 0 {
		this.userCount--
	}
}
func (this *Material) GetCount() uint32 {
	return this.userCount
}
func (this *Material) Dispose() {
	this.uploaded = false
}

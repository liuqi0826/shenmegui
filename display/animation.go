package display

import (
	"time"

	"github.com/liuqi0826/seven/geom"
)

//++++++++++++++++++++ SkeletonAnimationChannel ++++++++++++++++++++

type SkeletonAnimationChannel struct {
	transform []*geom.Matrix4x4
	timeline  []int32
	endTime   int32

	prevTransform *geom.Matrix4x4
	nextTransform *geom.Matrix4x4
	percent       float32
}

func (this *SkeletonAnimationChannel) SkeletonAnimationChannel(cha *AnimationChannel) {
	this.timeline = cha.Time
	this.endTime = this.timeline[len(this.timeline)-1]

	count := len(cha.Transform) / 16
	this.transform = make([]*geom.Matrix4x4, 0)
	for i := 0; i < count; i++ {
		p := i * 16
		raw := &[16]float32{cha.Transform[p], cha.Transform[p+1], cha.Transform[p+2], cha.Transform[p+3], cha.Transform[p+4], cha.Transform[p+5], cha.Transform[p+6], cha.Transform[p+7], cha.Transform[p+8], cha.Transform[p+9], cha.Transform[p+10], cha.Transform[p+11], cha.Transform[p+12], cha.Transform[p+13], cha.Transform[p+14], cha.Transform[p+15]}
		mtx := new(geom.Matrix4x4)
		mtx.Matrix4x4(raw)
		this.transform = append(this.transform, mtx)
	}
}
func (this *SkeletonAnimationChannel) GetCurrentTransform() *geom.Matrix4x4 {
	if this.prevTransform != nil && this.nextTransform != nil {
		return geom.InterpolateMatrix4x4(this.prevTransform, this.nextTransform, this.percent)
	}
	return this.transform[0]
}
func (this *SkeletonAnimationChannel) Update(time int32) {
	if len(this.timeline) > 1 {
		for time > this.endTime {
			time = time - this.endTime
		}
		for ti := 0; ti < len(this.timeline); ti++ {
			if time < this.timeline[ti] {
				this.percent = float32((time - this.timeline[ti-1])) / float32((this.timeline[ti] - this.timeline[ti-1]))
				if ti <= 0 {
					this.prevTransform = this.transform[0]
				} else {
					this.prevTransform = this.transform[ti-1]
				}
				if ti+1 >= len(this.transform) {
					this.nextTransform = this.transform[len(this.transform)-1]
				} else {
					this.nextTransform = this.transform[ti]
				}
				break
			}
		}
	} else {
		this.percent = 1.0
		this.prevTransform = this.transform[0]
		this.nextTransform = this.transform[0]
	}
}

//++++++++++++++++++++ SkeletonAnimationClip ++++++++++++++++++++

type SkeletonAnimationClip struct {
	ID          string
	Duration    int32
	UploadIndex []uint32

	startTime            int32
	currentTime          int32
	playing              bool
	animationChannelList []*SkeletonAnimationChannel
}

func (this *SkeletonAnimationClip) SkeletonAnimationClip(clip *AnimationClipResource) {
	if clip != nil {
		this.ID = clip.ID
		this.Duration = clip.Duration
		this.UploadIndex = clip.UploadIndex
		this.animationChannelList = make([]*SkeletonAnimationChannel, 0)
		for _, cha := range clip.Channel {
			c := new(SkeletonAnimationChannel)
			c.SkeletonAnimationChannel(cha)
			this.animationChannelList = append(this.animationChannelList, c)
		}
	}
}
func (this *SkeletonAnimationClip) GetAnimationChannelList() []*SkeletonAnimationChannel {
	return this.animationChannelList
}
func (this *SkeletonAnimationClip) GetCurrenTime() int32 {
	return this.currentTime
}
func (this *SkeletonAnimationClip) Play() {
	this.playing = true
	this.startTime = int32(time.Now().UnixNano() / 1000000)
}
func (this *SkeletonAnimationClip) Stop() {
	this.playing = false
}
func (this *SkeletonAnimationClip) Update() {
	if this.playing {
		this.currentTime = int32(time.Now().UnixNano()/1000000) - this.startTime
		for this.currentTime > this.Duration {
			this.currentTime = this.currentTime - this.Duration
		}
		for _, channel := range this.animationChannelList {
			channel.Update(this.currentTime)
		}
	}
}

//++++++++++++++++++++ SkeletonAnimation ++++++++++++++++++++

type SkeletonAnimation struct {
	ID string

	skeleton      *Skeleton
	animationClip map[string]*SkeletonAnimationClip
	currentClip   *SkeletonAnimationClip
}

func (this *SkeletonAnimation) SkeletonAnimation(skeleton *Skeleton) {
	this.skeleton = skeleton
	this.animationClip = make(map[string]*SkeletonAnimationClip)
}
func (this *SkeletonAnimation) GetCurrentClip() *SkeletonAnimationClip {
	return this.currentClip
}
func (this *SkeletonAnimation) AddClip(clip *SkeletonAnimationClip) {
	this.animationClip[clip.ID] = clip
}
func (this *SkeletonAnimation) SwitchClip(id string, interval int32) bool {
	if clip, ok := this.animationClip[id]; ok {
		if this.currentClip != nil {
			this.currentClip.Stop()
		}
		this.skeleton.jointTree.Binding(clip.animationChannelList)
		this.currentClip = clip
		this.currentClip.Play()
		return true
	}
	return false
}
func (this *SkeletonAnimation) Update() {
	if this.currentClip != nil {
		this.currentClip.Update()
		this.skeleton.Update()
	}
}
func (this *SkeletonAnimation) Apply(renderer IRenderer) {
	this.skeleton.pose.Encode(this.currentClip.UploadIndex)
	renderer.SetValue(SKELETON, this.skeleton.pose.Buffer)
}
func (this *SkeletonAnimation) GetType() string {
	return "skeleton"
}

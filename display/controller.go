package display

import (
	"github.com/liuqi0826/seven/geom"
)

//++++++++++++++++++++ FitstPersonController ++++++++++++++++++++

type FitstPersonController struct {
}

func (this *FitstPersonController) Update() {

}

//++++++++++++++++++++ HoverController ++++++++++++++++++++

type HoverController struct {
}

func (this *HoverController) Update() {

}

//++++++++++++++++++++ LockPerspectiveController ++++++++++++++++++++

type LockPerspectiveController struct {
}

func (this *LockPerspectiveController) Update() {

}

//++++++++++++++++++++ Joint ++++++++++++++++++++

type Joint struct {
	Index     uint32
	Parent    *Joint
	Children  []*Joint
	Transform *geom.Matrix4x4

	channel *SkeletonAnimationChannel
}

func (this *Joint) Joint(parent *Joint) {
	this.Parent = parent
	this.Children = make([]*Joint, 0)
	this.Transform = new(geom.Matrix4x4)
	this.Transform.Matrix4x4(nil)
}
func (this *Joint) Update() {
	if this.channel != nil {
		this.Transform = this.channel.GetCurrentTransform()
	} else {
		this.Transform.Identity()
	}
	for _, child := range this.Children {
		child.Update()
	}
}
func (this *Joint) Blend(invbind []*geom.Matrix4x4, root *geom.Matrix4x4) {
	if this.Parent != nil {
		this.Transform.Append(this.Parent.Transform)
	} else {
		if root != nil {
			this.Transform.Append(root)
		}
	}
	for _, child := range this.Children {
		child.Blend(invbind, root)
	}
	this.Transform.Prepend(invbind[this.Index])
}
func (this *Joint) Apply(pose *Pose) {
	pose.NodeList[this.Index] = this.Transform
	for _, child := range this.Children {
		child.Apply(pose)
	}
}
func (this *Joint) Binding(channelList []*SkeletonAnimationChannel) {
	this.channel = channelList[this.Index]
	for _, child := range this.Children {
		child.Binding(channelList)
	}
}

//++++++++++++++++++++ Pose ++++++++++++++++++++

type Pose struct {
	NodeList []*geom.Matrix4x4
	Buffer   []float32
}

func (this *Pose) Pose(length int) {
	this.NodeList = make([]*geom.Matrix4x4, length)
	for i := 0; i < length; i++ {
		this.NodeList[i] = new(geom.Matrix4x4)
		this.NodeList[i].Matrix4x4(nil)
	}
}
func (this *Pose) Encode(upload []uint32) {
	this.Buffer = make([]float32, 0)
	for _, up := range upload {
		node := this.NodeList[up]
		for _, v := range node.Raw {
			this.Buffer = append(this.Buffer, v)
		}
	}
}

//++++++++++++++++++++ Skeleton ++++++++++++++++++++

type Skeleton struct {
	ID string

	structure             []interface{}
	bindShape             *geom.Matrix4x4
	root                  *geom.Matrix4x4
	foot                  *geom.Matrix4x4
	invbind               []*geom.Matrix4x4
	nodeNameList          []string
	jointNameList         []string
	jointTree             *Joint
	pose                  *Pose
	skeletonAnimationClip *SkeletonAnimationClip
}

func (this *Skeleton) Skeleton(skeleton *SkeletonResource) {
	if skeleton != nil {
		this.structure = skeleton.Struct

		bindArray := &[16]float32{skeleton.BindShape[0], skeleton.BindShape[1], skeleton.BindShape[2], skeleton.BindShape[3], skeleton.BindShape[4], skeleton.BindShape[5], skeleton.BindShape[6], skeleton.BindShape[7], skeleton.BindShape[8], skeleton.BindShape[9], skeleton.BindShape[10], skeleton.BindShape[11], skeleton.BindShape[12], skeleton.BindShape[13], skeleton.BindShape[14], skeleton.BindShape[15]}
		bindShape := new(geom.Matrix4x4)
		bindShape.Matrix4x4(bindArray)
		this.bindShape = bindShape

		rootArray := &[16]float32{skeleton.Root[0], skeleton.Root[1], skeleton.Root[2], skeleton.Root[3], skeleton.Root[4], skeleton.Root[5], skeleton.Root[6], skeleton.Root[7], skeleton.Root[8], skeleton.Root[9], skeleton.Root[10], skeleton.Root[11], skeleton.Root[12], skeleton.Root[13], skeleton.Root[14], skeleton.Root[15]}
		root := new(geom.Matrix4x4)
		root.Matrix4x4(rootArray)
		this.root = root

		footArray := &[16]float32{skeleton.Foot[0], skeleton.Foot[1], skeleton.Foot[2], skeleton.Foot[3], skeleton.Foot[4], skeleton.Foot[5], skeleton.Foot[6], skeleton.Foot[7], skeleton.Foot[8], skeleton.Foot[9], skeleton.Foot[10], skeleton.Foot[11], skeleton.Foot[12], skeleton.Foot[13], skeleton.Foot[14], skeleton.Foot[15]}
		foot := new(geom.Matrix4x4)
		foot.Matrix4x4(footArray)
		this.foot = foot

		this.invbind = make([]*geom.Matrix4x4, 0)
		for _, inv := range skeleton.Invbind {
			invArray := &[16]float32{inv[0], inv[1], inv[2], inv[3], inv[4], inv[5], inv[6], inv[7], inv[8], inv[9], inv[10], inv[11], inv[12], inv[13], inv[14], inv[15]}
			invmtx := new(geom.Matrix4x4)
			invmtx.Matrix4x4(invArray)
			this.invbind = append(this.invbind, invmtx)
		}

		this.nodeNameList = skeleton.NodeNameList
		this.jointNameList = skeleton.JointNameList

		this.jointTree = createJoint(nil, this.structure)

		this.pose = new(Pose)
		this.pose.Pose(len(this.nodeNameList))
	}
}
func (this *Skeleton) Update() {
	this.jointTree.Update()
	this.jointTree.Blend(this.invbind, nil)
	this.jointTree.Apply(this.pose)
}

func createJoint(parent *Joint, array []interface{}) *Joint {
	joint := new(Joint)
	if idx, ok := array[0].(float64); ok {
		joint.Index = uint32(idx)
		joint.Parent = parent
	}
	if children, ok := array[1].([]interface{}); ok {
		joint.Children = createJointList(joint, children)
	}
	return joint
}
func createJointList(parent *Joint, children []interface{}) []*Joint {
	childrenList := make([]*Joint, 0)
	for _, c := range children {
		if arr, ok := c.([]interface{}); ok {
			joint := createJoint(parent, arr)
			childrenList = append(childrenList, joint)
		}
	}
	return childrenList
}

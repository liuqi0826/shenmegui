package core

var instanceIndex uint32

func GetNextInstanceID() uint32 {
	instanceIndex++
	return instanceIndex
}

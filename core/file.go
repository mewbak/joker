package core

import (
	"unsafe"
)

type (
	File struct {
		*JokerWriter
	}
)

func (f *File) ToString(escape bool) string {
	return "#object[File]"
}

func (f *File) Equals(other interface{}) bool {
	return f == other
}

func (f *File) GetInfo() *ObjectInfo {
	return nil
}

func (f *File) GetType() *Type {
	return TYPE.File
}

func (f *File) Hash() uint32 {
	return hashPtr(uintptr(unsafe.Pointer(f)))
}

func (f *File) WithInfo(info *ObjectInfo) Object {
	return f
}

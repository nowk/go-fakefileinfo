package fakefileinfo

import (
	"os"
	"time"
)

// A FileInfo represents struct that impelments the os.FileInfo interface
type FileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modtime time.Time
	isdir   bool
	sys     interface{}
}

func New(name string, size int64, mode os.FileMode, modtime time.Time, isdir bool,
	sys interface{}) (fi os.FileInfo) {

	fi = &FileInfo{
		name:    name,
		size:    size,
		mode:    mode,
		modtime: modtime,
		isdir:   isdir,
		sys:     sys,
	}
	return
}

func (fi FileInfo) Name() string {
	return fi.name
}

func (fi FileInfo) Size() int64 {
	return fi.size
}

func (fi FileInfo) Mode() os.FileMode {
	return fi.mode
}

func (fi FileInfo) ModTime() time.Time {
	return fi.modtime
}

func (fi FileInfo) IsDir() bool {
	return fi.isdir
}

func (fi FileInfo) Sys() interface{} {
	return fi.sys
}

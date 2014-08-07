package fakefileinfo_test

import (
	"os"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	. "github.com/nowk/go-fakefileinfo"
)

func TestFakeFileInfo(t *testing.T) {
	var someTime = time.Now()
	fi := New("file.txt", int64(123), os.ModeType, someTime, true, nil)

	assert.Equal(t, fi.Name(), "file.txt")
	assert.Equal(t, fi.Size(), int64(123))
	assert.Equal(t, fi.Mode(), os.ModeType)

	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, fi.ModTime(), someTime)

	assert.Equal(t, fi.IsDir(), true)
	assert.Equal(t, fi.Sys(), nil)
}

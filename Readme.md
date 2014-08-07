# fakefileinfo

[![Build Status](https://travis-ci.org/nowk/go-fakefileinfo.svg?branch=master)](https://travis-ci.org/nowk/go-fakefileinfo)

`os.FileInfo` interface mock

## Example

    fi := New("file.txt", int64(123), os.ModeType, time.Now(), true, nil)

### License

MIT

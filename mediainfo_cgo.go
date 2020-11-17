package mediainfo

// #cgo CFLAGS: -DUNICODE -D_UNICODE
// #include "mediainfo.h"
// #cgo linux LDFLAGS: -ldl
import "C"

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type streamType uint

const (
	streamGeneral streamType = iota
	streamVideo
	streamAudio
	streamText
	streamOther
	streamImage
	streamMenu
	streamMax
)

var loaded bool

func init() {
	C.GoSetLocale()
}

var (
	inform  = wideString("Inform")
	jsonStr = wideString("JSON")
)

// MediaInfo - represents MediaInfo class, all interaction with libmediainfo through it
type mediaInfo struct {
	handle unsafe.Pointer
}

// Load loads the DLL/shared object. If successful returns true.
func Load() bool {
	r := C.MediaInfoDLL_Load()
	if r == 1 {
		loaded = true
		return true
	}
	return false
}

// Unload unloads the DLL/shared object
func Unload() {
	C.MediaInfoDLL_UnLoad()
}

// newMediaInfo - constructs new MediaInfo
func newMediaInfo() (*mediaInfo, error) {
	if !loaded {
		return nil, ErrNotLoaded
	}

	result := &mediaInfo{handle: C.GoMediaInfo_New()}
	return result, nil
}

// OpenFile - opens file
func (mi *mediaInfo) OpenFile(path string) error {
	p := wideString(path)
	s := C.GoMediaInfo_OpenFile(mi.handle, (*C.wchar_t)(unsafe.Pointer(&p[0])))
	if s == 0 {
		return fmt.Errorf("MediaInfo can't open file: %s", path)
	}
	return nil
}

// Close - closes file
func (mi *mediaInfo) Close() {
	C.GoMediaInfo_Close(mi.handle)
}

func (mi *mediaInfo) Inform() (informStruct, error) {
	C.GoMediaInfoOption(mi.handle, (*C.wchar_t)(unsafe.Pointer(&inform[0])), (*C.wchar_t)(unsafe.Pointer(&jsonStr[0])))
	str := toString(C.GoMediaInfoInform(mi.handle))
	info := informStruct{}
	err := json.Unmarshal([]byte(str), &info)
	return info, err
}

package mediainfo

import (
	"unicode/utf16"
	"unsafe"

	"golang.org/x/sys/windows"
)

// #include <wchar.h>
import "C"

func wideString(s string) []uint16 {
	if len(s) < windows.MAX_PATH-1 {
		return utf16.Encode([]rune(s + "\x00"))
	}

	// Add to the long filename a prefix to cause the API to handle filenames longer than 260 characters.
	p := utf16.Encode([]rune(`\\?\` + s + "\x00"))
	b := p // GetShortPathName says we can reuse buffer
	for {
		n, err := windows.GetShortPathName(&p[0], &b[0], uint32(len(b)))
		if err != nil {
			panic(err)
		}

		if n <= uint32(len(b)) {
			b = b[:n]
			break
		}
		b = make([]uint16, n)
	}
	return b
}

func toString(r *C.wchar_t) string {
	bytes := make([]uint16, 0)
	ptr := (*uint16)(unsafe.Pointer(r))
	for *ptr != 0 {
		bytes = append(bytes, *ptr)
		ptr = (*uint16)(unsafe.Pointer((uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(bytes[0]))))
	}
	return string(utf16.Decode(bytes))
}

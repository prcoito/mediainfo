package mediainfo

import (
	"strings"
	"unsafe"
)

// #include <wchar.h>
import "C"

func wideString(s string) []rune {
	return append([]rune(s), 0)
}

func toString(r *C.wchar_t) string {
	str := strings.Builder{}
	ptr := (*rune)(unsafe.Pointer(r))
	for *ptr != 0 {
		str.WriteRune(*ptr)
		ptr = (*rune)(unsafe.Pointer((uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(*ptr))))
	}
	return str.String()
}

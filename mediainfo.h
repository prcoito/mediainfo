#include "MediaInfoDLL.h"

// Unix
#if defined(UNIX) || defined(_UNIX) || defined(__UNIX__)
#include <locale.h>

void GoSetLocale(void) {
    setlocale(LC_ALL, "en_US.UTF-8");
}
#endif

// windows
#if defined(WIN32) || defined(WIN64)
void GoSetLocale(void) {}
#endif


void *GoMediaInfo_New() {
    return MediaInfo_New();
}

void GoMediaInfo_Delete(void *handle) {
    MediaInfo_Delete(handle);
}

size_t GoMediaInfo_OpenFile(void *handle, wchar_t *name) {
    return MediaInfo_Open(handle, name);
}

void GoMediaInfo_Close(void *handle) {
    MediaInfo_Close(handle);
}

const wchar_t *GoMediaInfoGet(void *handle, MediaInfo_stream_C s, size_t index, wchar_t *name) {
    return MediaInfo_Get(handle, s, index,  name, MediaInfo_Info_Text, MediaInfo_Info_Name);
}

const wchar_t *GoMediaInfoOption(void *handle, wchar_t *name, wchar_t *value) {
    return MediaInfo_Option(handle, name, value);
}

const wchar_t *GoMediaInfoInform(void *handle) {
    return MediaInfo_Inform(handle, 0);
}

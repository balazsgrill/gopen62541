package gopen62541

/*
#include <open62541/types.h>
#include <open62541/server.h>
*/
import "C"
import "unsafe"

func newLocalizedText(locale string, text string) C.UA_LocalizedText {
	l := C.CString(locale)
	defer C.free(unsafe.Pointer(l))
	t := C.CString(text)
	defer C.free(unsafe.Pointer(t))
	return C.UA_LOCALIZEDTEXT_ALLOC(l, t)
}

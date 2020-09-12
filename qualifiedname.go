package gopen62541

/*
#include <open62541/types.h>
#include <open62541/types_generated_handling.h>
*/
import "C"
import "unsafe"

type QualifiedName C.UA_QualifiedName

func NewQualifiedName(ns uint16, name string) QualifiedName {
	str := C.CString(name)
	defer C.free(unsafe.Pointer(str))
	return QualifiedName(C.UA_QUALIFIEDNAME_ALLOC(C.ushort(ns), str))
}

func (qn *QualifiedName) Clear() {
	C.UA_QualifiedName_clear((*C.UA_QualifiedName)(qn))
}

package gopen62541

/*
#include <open62541/types.h>
#include <open62541/types_generated_handling.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type Variant C.UA_Variant

func (variant *Variant) SetValue(value interface{}) error {
	switch v := value.(type) {
	case uint32:
		var r C.UA_UInt32 = C.UA_UInt32(v)
		C.UA_Variant_setScalarCopy((*C.UA_Variant)(variant), unsafe.Pointer(&r), &C.UA_TYPES[C.UA_TYPES_UINT32])
		return nil
	case int32:
		var r C.UA_Int32 = C.UA_Int32(v)
		C.UA_Variant_setScalarCopy((*C.UA_Variant)(variant), unsafe.Pointer(&r), &C.UA_TYPES[C.UA_TYPES_INT32])
		return nil
	default:
		return errors.New("Unknown type")
	}
}

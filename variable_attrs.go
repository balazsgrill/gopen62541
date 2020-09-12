package gopen62541

/*
#include <open62541/types.h>
#include <open62541/server.h>
*/
import "C"

type VariableAttributes C.UA_VariableAttributes

func NewDefaultVariableAttributes() *VariableAttributes {
	var v VariableAttributes = VariableAttributes(C.UA_VariableAttributes_default)
	return &v
}

func (va *VariableAttributes) SetDescription(locale string, text string) {
	va.description = newLocalizedText(locale, text)
}

func (va *VariableAttributes) SetDisplayName(locale string, text string) {
	va.displayName = newLocalizedText(locale, text)
}

func (va *VariableAttributes) Value() *Variant {
	return (*Variant)(&va.value)
}

func (va *VariableAttributes) Clear() {
	C.UA_VariableAttributes_clear((*C.UA_VariableAttributes)(va))
}

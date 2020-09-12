package gopen62541

/*
#include <open62541/types.h>
#include <open62541/types_generated_handling.h>
#include <open62541/nodeids.h>
*/
import "C"
import "unsafe"

// Design API to be compatible with https://github.com/gopcua/opcua/blob/master/ua/node_id.go

type NodeID C.UA_NodeId

const (
	NodeIDObjectsFolder = C.UA_NS0ID_OBJECTSFOLDER
	NodeIDOrganizes     = C.UA_NS0ID_ORGANIZES
)

func NodeIDNull() NodeID {
	return NodeID(C.UA_NODEID_NULL)
}

func NewNumericNodeID(ns uint16, id uint32) NodeID {
	return NodeID(C.UA_NODEID_NUMERIC(C.ushort(ns), C.uint(id)))
}

func NewStringNodeID(ns uint16, id string) NodeID {
	str := C.CString(id)
	defer C.free(unsafe.Pointer(str))
	return NodeID(C.UA_NODEID_STRING_ALLOC(C.ushort(ns), str))
}

func (n *NodeID) Clear() {
	C.UA_NodeId_clear((*C.UA_NodeId)(n))
}

package gopen62541

/*
#cgo LDFLAGS: -lopen62541
#include <open62541/server.h>
#include <open62541/server_config.h>
#include <open62541/server_config_default.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Server struct {
	uaServer *C.UA_Server
	running  C.UA_Boolean
}

func statusCodeToError(statusCode C.UA_StatusCode) error {
	switch statusCode {
	case C.UA_STATUSCODE_GOOD:
		return nil
	default:
		return fmt.Errorf("Server run error: %x", statusCode)
	}
}

func NewServer() *Server {
	server := &Server{
		uaServer: C.UA_Server_new(),
		running:  C.UA_Boolean(true),
	}
	C.UA_ServerConfig_setDefault(C.UA_Server_getConfig(server.uaServer))

	return server
}

func (s *Server) Run() error {
	return statusCodeToError(C.UA_Server_run(s.uaServer, &s.running))
}

func (s *Server) Stop() {
	s.running = C.UA_Boolean(false)
}

func (s *Server) Delete() {
	C.UA_Server_delete(s.uaServer)
}

func (s *Server) AddVariableNode(requestedNewNodeID NodeID,
	parentNodeID NodeID,
	referenceTypeID NodeID,
	browseName QualifiedName,
	typeDefinition NodeID,
	attr *VariableAttributes,
	nodeContext interface{}) (NodeID, error) {
	result := C.UA_NodeId{}
	statusCode := C.UA_Server_addVariableNode(
		s.uaServer,
		(C.UA_NodeId)(requestedNewNodeID),
		(C.UA_NodeId)(parentNodeID),
		(C.UA_NodeId)(referenceTypeID),
		(C.UA_QualifiedName)(browseName),
		(C.UA_NodeId)(typeDefinition),
		(C.UA_VariableAttributes)(*attr),
		unsafe.Pointer(&nodeContext),
		&result)
	return NodeID(result), statusCodeToError(statusCode)
}

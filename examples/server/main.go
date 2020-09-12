package main

import ua "github.com/balazsgrill/gopen62541"

func main() {
	s := ua.NewServer()

	vattr := ua.NewDefaultVariableAttributes()
	defer vattr.Clear()
	vattr.SetDescription("en-US", "the answer")
	vattr.SetDisplayName("en-US", "the answer")
	myIntegerNodeID := ua.NewStringNodeID(1, "the.answer")
	defer myIntegerNodeID.Clear()
	myIntegerName := ua.NewQualifiedName(1, "the answer")
	defer myIntegerName.Clear()
	parentNodeId := ua.NewNumericNodeID(0, ua.NodeIDObjectsFolder)
	parentReferenceNodeId := ua.NewNumericNodeID(0, ua.NodeIDOrganizes)

	s.AddVariableNode(myIntegerNodeID, parentNodeId, parentReferenceNodeId, myIntegerName, ua.NodeIDNull(), vattr, nil)

	s.Run()

}

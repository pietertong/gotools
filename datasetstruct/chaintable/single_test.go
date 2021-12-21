package chaintable

import "testing"

var headLinkNode *HeadLinkNode

func TestNew(t *testing.T) {
	headLinkNode = New()
	if headLinkNode == nil {
		t.Errorf("New() error, %v", headLinkNode)
	}
}

func TestHeadLinkNode_Add(t *testing.T) {
	headLinkNode = New()
	headLinkNode.Add("test")
	if headLinkNode.Node.Data != "" {
		t.Errorf("headLinkNode.Node.Data error, %v", headLinkNode.Node.Data)
	}
	if headLinkNode.Node.Next.Data != "test" {
		t.Errorf("add(test) error, %v", headLinkNode.Node.Data)
	}
}
func TestHeadLinkNode_Remove(t *testing.T) {

}

func TestHeadLinkNode_Get(t *testing.T) {

}
func TestHeadLinkNode_Foreach(t *testing.T) {

}
func TestHeadLinkNode_Len(t *testing.T) {

}
func TestHeadLinkNode_Search(t *testing.T) {

}
func TestHeadLinkNode_Insert(t *testing.T) {

}

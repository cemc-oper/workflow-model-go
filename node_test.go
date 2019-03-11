package workflowmodel_test

import (
	"fmt"
	"github.com/perillaroc/workflow-model-go"
	"testing"
)

func TestNode_AddChild(t *testing.T) {
	t1 := &workflowmodel.Node{}
	t1.Name = "t1"

	f1 := &workflowmodel.Node{}
	f1.Name = "f1"
	f1.AddChild(t1)

	if t1.Parent != f1 {
		t.Errorf("t1.Parent != f1")
	}

	if f1.Children[0] != t1 {
		t.Errorf("f1.Children[0] != t1")
	}
}

func TestNode_NodePath(t *testing.T) {
	t1 := &workflowmodel.Node{}
	t1.Name = "t1"

	f1 := &workflowmodel.Node{}
	f1.Name = "f1"
	f1.AddChild(t1)

	s1 := &workflowmodel.Node{}
	s1.Name = "s1"
	s1.AddChild(f1)

	root := &workflowmodel.Node{}
	root.AddChild(s1)

	tests := []struct {
		node *workflowmodel.Node
		path string
	}{
		{t1, "/s1/f1/t1"},
		{f1, "/s1/f1"},
		{s1, "/s1"},
		{root, "/"},
	}

	for _, test := range tests {
		if test.node.NodePath() != test.path {
			t.Errorf("NodePath for %v = %v, should %v",
				test.node, test.node.NodePath(), test.path)
		}
	}
}

func TestNode_MarshalJSON(t *testing.T) {
	t1 := &workflowmodel.Node{}
	t1.Name = "t1"

	f1 := &workflowmodel.Node{}
	f1.Name = "f1"
	f1.AddChild(t1)

	s1 := &workflowmodel.Node{}
	s1.Name = "s1"
	s1.AddChild(f1)

	root := &workflowmodel.Node{}
	root.AddChild(s1)

	b, err := root.MarshalJSON()
	if err != nil {
		t.Errorf("Node.MarshalJSON error: %v", err)
	}
	s := fmt.Sprintf("%s\n", b)
	fmt.Println(s)

	var n workflowmodel.Node
	err = n.UnmarshalJSON(b)
	if err != nil {
		t.Errorf("Node.UnmarshalJSON error: %v", err)
	}
}

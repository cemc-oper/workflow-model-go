package workflowmodel_test

import (
	"github.com/perillaroc/workflow-model-go"
	"testing"
)

func TestBunch_FindNode(t *testing.T) {
	t1 := &workflowmodel.Node{}
	t1.Name = "t1"

	f1 := &workflowmodel.Node{}
	f1.Name = "f1"
	f1.AddChild(t1)

	s1 := &workflowmodel.Node{}
	s1.Name = "s1"
	s1.AddChild(f1)

	root := &workflowmodel.Bunch{}
	root.AddChild(s1)

	tests := []struct {
		path string
		node *workflowmodel.Node
	}{
		{"/s1/f1/t1", t1},
		{"/s1/f1", f1},
		{"/s1", s1},
		{"/", &root.Node},
		{"/f1/t1", nil},
		{"/f1", nil},
	}

	for _, test := range tests {
		node := root.FindNode(test.path)
		if node != test.node {
			t.Errorf("FindNode(%v) = %v, should %v",
				test.path, node, test.node)
		}
	}
}

func TestBunch_AddNode(t *testing.T) {
	root := &workflowmodel.Bunch{}
	root.AddNode("/s1/f1/t1")
	root.AddNode("/s1/f1")
	root.AddNode("/s1/f2/t2")
	root.AddNode("/s1/f2/t3")

	t1 := root.FindNode("/s1/f1/t1")
	if t1 == nil {
		t.Errorf("/s1/f1/t1 not found")
	}

	t3 := root.FindNode("/s1/f2/t3")
	if t3 == nil {
		t.Errorf("/s1/f2/t3 not found")
	}

	if len(root.Children[0].Children) != 2 {
		t.Errorf("size of /s1's children is not 2")
	}
}

func TestBunch_AddNodeStatus(t *testing.T) {
	root := &workflowmodel.Bunch{}
	root.AddNodeStatus(workflowmodel.NodeStatusRecord{Path: "/s1/f1/t1", Status: "complete"})
	root.AddNodeStatus(workflowmodel.NodeStatusRecord{Path: "/s1/f1", Status: "complete"})
	root.AddNodeStatus(workflowmodel.NodeStatusRecord{Path: "/s1/f2/t2", Status: "active"})
	root.AddNodeStatus(workflowmodel.NodeStatusRecord{Path: "/s1/f2/t3", Status: "queued"})

	t1 := root.FindNode("/s1/f1/t1")
	if t1.Status != workflowmodel.Complete {
		t.Errorf("/s1/f1/t1 is not complete")
	}

	t3 := root.FindNode("/s1/f2/t3")
	if t3.Status != workflowmodel.Queued {
		t.Errorf("/s1/f2/t3 not queued")
	}

	if len(root.Children[0].Children) != 2 {
		t.Errorf("size of /s1's children is not 2")
	}
}

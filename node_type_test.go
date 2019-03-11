package workflowmodel_test

import (
	"github.com/perillaroc/workflow-model-go"
	"testing"
)

func TestGetNodeType(t *testing.T) {
	tests := []struct {
		s        string
		nodeType workflowmodel.NodeType
	}{
		{
			"unknown",
			workflowmodel.UnknownNodeType,
		},
		{
			"root",
			workflowmodel.Root,
		},
		{
			"suite",
			workflowmodel.Suite,
		},
		{
			"family",
			workflowmodel.Family,
		},
		{
			"task",
			workflowmodel.Task,
		},
		{
			"non-task",
			workflowmodel.NonTaskNode,
		},
		{
			"alias",
			workflowmodel.Alias,
		},
		{
			"meter",
			workflowmodel.Meter,
		},
	}

	for _, test := range tests {
		if workflowmodel.GetNodeType(test.s) != test.nodeType {
			t.Errorf("GetNodeType(%v) = %v, should %v",
				test.s, workflowmodel.GetNodeType(test.s), test.nodeType)
		}
	}
}

func TestNodeType_String(t *testing.T) {
	tests := []struct {
		nodeType workflowmodel.NodeType
		s        string
	}{
		{
			workflowmodel.UnknownNodeType,
			"unknown",
		},
		{
			workflowmodel.Root,
			"root",
		},
		{
			workflowmodel.Suite,
			"suite",
		},
		{
			workflowmodel.Family,
			"family",
		},
		{
			workflowmodel.Task,
			"task",
		},
		{
			workflowmodel.Alias,
			"alias",
		},
		{
			workflowmodel.NonTaskNode,
			"non-task",
		},
		{
			workflowmodel.Meter,
			"meter",
		},
	}

	for _, test := range tests {
		if test.nodeType.String() != test.s {
			t.Errorf("NodeType.String(%v) = %v, should %v",
				test.nodeType, test.nodeType.String(), test.s)
		}
	}
}

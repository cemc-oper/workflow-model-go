package workflowmodel_test

import (
	"github.com/perillaroc/workflow-model-go"
	"testing"
)

func TestGetNodeStatus(t *testing.T) {
	tests := []struct {
		s      string
		status workflowmodel.NodeStatus
	}{
		{
			"unknown",
			workflowmodel.Unknown,
		},
		{
			"unk",
			workflowmodel.Unknown,
		},
		{
			"suspended",
			workflowmodel.Suspended,
		},
		{
			"sus",
			workflowmodel.Suspended,
		},
		{
			"complete",
			workflowmodel.Complete,
		},
		{
			"com",
			workflowmodel.Complete,
		},
		{
			"queued",
			workflowmodel.Queued,
		},
		{
			"que",
			workflowmodel.Queued,
		},
		{
			"submitted",
			workflowmodel.Submitted,
		},
		{
			"sub",
			workflowmodel.Submitted,
		},
		{
			"active",
			workflowmodel.Active,
		},
		{
			"act",
			workflowmodel.Active,
		},
		{
			"aborted",
			workflowmodel.Aborted,
		},
		{
			"abo",
			workflowmodel.Aborted,
		},
		{
			"SHUTDOWN",
			workflowmodel.Shutdown,
		},
		{
			"shu",
			workflowmodel.Shutdown,
		},
		{
			"HALTED",
			workflowmodel.Halted,
		},
		{
			"hal",
			workflowmodel.Halted,
		},
		{
			"RUNNING",
			workflowmodel.Running,
		},
	}

	for _, test := range tests {
		if workflowmodel.GetNodeStatus(test.s) != test.status {
			t.Errorf("GetNodeStatus(%v) = %v, should %v",
				test.s, workflowmodel.GetNodeStatus(test.s), test.status)
		}
	}
}

func TestNodeStatus_String(t *testing.T) {
	tests := []struct {
		status workflowmodel.NodeStatus
		s      string
	}{
		{
			workflowmodel.Unknown,
			"unknown",
		},
		{
			workflowmodel.Suspended,
			"suspended",
		},
		{
			workflowmodel.Complete,
			"complete",
		},
		{
			workflowmodel.Queued,
			"queued",
		},
		{
			workflowmodel.Submitted,
			"submitted",
		},
		{
			workflowmodel.Active,
			"active",
		},
		{
			workflowmodel.Aborted,
			"aborted",
		},
		{
			workflowmodel.Shutdown,
			"SHUTDOWN",
		},
		{
			workflowmodel.Halted,
			"HALTED",
		},
		{
			workflowmodel.Running,
			"RUNNING",
		},
	}

	for _, test := range tests {
		if test.status.String() != test.s {
			t.Errorf("NodeStatus.String(%v) = %v, should %v",
				test.status, test.status.String(), test.s)
		}
	}
}

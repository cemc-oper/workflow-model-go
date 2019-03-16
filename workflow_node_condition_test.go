package workflowmodel_test

import (
	"github.com/perillaroc/workflow-model-go"
	"testing"
)

func TestWorkflowNodeStatusCondition_IsFit(t *testing.T) {
	condition := workflowmodel.WorkflowNodeStatusCondition{
		Checker: &workflowmodel.NodeStatusEqualValueChecker{
			ExpectedValue: workflowmodel.Active},
	}

	node := &workflowmodel.WorkflowNode{
		Status: workflowmodel.Active,
	}

	if !condition.IsFit(node) {
		t.Errorf("node status is not active")
	}
}

func TestWorkflowNodeVariableCondition_IsFit(t *testing.T) {
	condition := workflowmodel.WorkflowNodeVariableCondition{
		VariableName: "ECF_DATE",
		Checker: &workflowmodel.StringEqualValueChecker{
			ExpectedValue: "20190316"},
	}

	node1 := &workflowmodel.WorkflowNode{
		VariableList: workflowmodel.WorkflowNodeVariables{
			VariableList: []workflowmodel.NodeVariable{
				{
					Name:  "ECF_DATE",
					Value: "20190316",
					Type:  workflowmodel.Variable,
				},
			},
		},
	}

	if !condition.IsFit(node1) {
		t.Errorf("ECF_DATE is not 20190316")
	}

	node2 := &workflowmodel.WorkflowNode{
		VariableList: workflowmodel.WorkflowNodeVariables{
			VariableList: []workflowmodel.NodeVariable{
				{
					Name:  "ECF_DATE",
					Value: "20190315",
					Type:  workflowmodel.Variable,
				},
			},
		},
	}

	if condition.IsFit(node2) {
		t.Errorf("ECF_DATE is 20190316")
	}

}

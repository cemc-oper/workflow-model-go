package workflowmodel_test

import (
	"encoding/json"
	"fmt"
	"github.com/perillaroc/workflow-model-go"
	"testing"
)

func TestWorkflowNodeVariables_JSON(t *testing.T) {
	vars := workflowmodel.WorkflowNodeVariables{
		Path: "/meso_post/00/initial",
		VariableList: []workflowmodel.NodeVariable{
			{
				Name:  "BASE_DIR",
				Value: "/g1/u/nwp_pd/ecfworks/meso_post",
				Type:  workflowmodel.Variable,
			},
			{
				Name:  "BIN_DIR",
				Value: "/g1/u/nwp_pd/MESO_POST/bin",
				Type:  workflowmodel.Variable,
			},
		},
		GeneratedVariableList: []workflowmodel.NodeVariable{
			{
				Name:  "ECF_NAME",
				Value: "/meso_post/00/initial",
				Type:  workflowmodel.GeneratedVariable,
			},
			{
				Name:  "ECF_DATE",
				Value: "20190312",
				Type:  workflowmodel.GeneratedVariable,
			},
		},
	}

	varsJsonString, err := json.Marshal(vars)
	if err != nil {
		t.Errorf("json.Marshal(vars) error: %v", err)
	}

	fmt.Printf("json.Marshal(vars) = %s\n", varsJsonString)
}

func TestWorkflowNode_JSON(t *testing.T) {
	node := workflowmodel.WorkflowNode{
		Name:   "initial",
		Status: workflowmodel.Queued,
		Type:   workflowmodel.Task,
		VariableList: workflowmodel.WorkflowNodeVariables{
			Path: "/meso_post/00/initial",
			VariableList: []workflowmodel.NodeVariable{
				{
					Name:  "BASE_DIR",
					Value: "/g1/u/nwp_pd/ecfworks/meso_post",
					Type:  workflowmodel.Variable,
				},
				{
					Name:  "BIN_DIR",
					Value: "/g1/u/nwp_pd/MESO_POST/bin",
					Type:  workflowmodel.Variable,
				},
			},
			GeneratedVariableList: []workflowmodel.NodeVariable{
				{
					Name:  "ECF_NAME",
					Value: "/meso_post/00/initial",
					Type:  workflowmodel.GeneratedVariable,
				},
			},
		},
		InheritedVariableList: []workflowmodel.WorkflowNodeVariables{
			{
				Path: "/meso_post/00",
				VariableList: []workflowmodel.NodeVariable{
					{
						Name:  "HH",
						Value: "00",
						Type:  workflowmodel.Variable,
					},
				},
				GeneratedVariableList: []workflowmodel.NodeVariable{
					{
						Name:  "ECF_NAME",
						Value: "/meso_post/00",
						Type:  workflowmodel.GeneratedVariable,
					},
				},
			},
			{
				Path: "/meso_post",
				VariableList: []workflowmodel.NodeVariable{
					{
						Name:  "SUITE",
						Value: "meso_post",
						Type:  workflowmodel.Variable,
					},
				},
				GeneratedVariableList: []workflowmodel.NodeVariable{
					{
						Name:  "ECF_NAME",
						Value: "/meso_post",
						Type:  workflowmodel.GeneratedVariable,
					},
					{
						Name:  "ECF_DATE",
						Value: "20190312",
						Type:  workflowmodel.GeneratedVariable,
					},
				},
			},
		},
	}

	nodeJsonString, err := json.Marshal(node)
	if err != nil {
		t.Errorf("json.Marshal(node) error: %v", err)
	}

	fmt.Printf("json.Marshal(node) = %s\n", nodeJsonString)

}

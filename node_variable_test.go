package workflowmodel_test

import (
	"encoding/json"
	"fmt"
	"github.com/perillaroc/workflow-model-go"
	"testing"
)

func TestNodeVariable_JSON(t *testing.T) {
	v1 := workflowmodel.NodeVariable{
		Name:  "ECF_DATE",
		Value: "20190312",
		Type:  workflowmodel.GeneratedVariable,
	}

	v1JsonString, err := json.Marshal(v1)
	if err != nil {
		t.Errorf("json.Marshal(v1) error: %v", err)
	}

	fmt.Printf("json(v1) = %s\n", v1JsonString)

}

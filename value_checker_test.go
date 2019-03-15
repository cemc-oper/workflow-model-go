package workflowmodel_test

import (
	"github.com/perillaroc/workflow-model-go"
	"testing"
)

func TestStringEqualValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValue string
		data          string
		result        bool
	}{
		{"20190315", "20190315", true},
		{"20190315", "20190314", false},
	}

	for _, test := range tests {
		checker := workflowmodel.StringEqualValueChecker{
			ExpectedValue: test.expectedValue,
		}
		if checker.CheckValue(test.data) != test.result {
			t.Errorf("StringEqualValueChecker expected value %s, value %s, result %t",
				test.expectedValue, test.data, test.result)
		}
	}
}

func TestStringInValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValues []string
		data           string
		result         bool
	}{
		{[]string{"active", "complete"}, "active", true},
		{[]string{"active", "complete"}, "aborted", false},
	}

	for _, test := range tests {
		checker := workflowmodel.StringInValueChecker{
			ExpectedValues: test.expectedValues,
		}
		if checker.CheckValue(test.data) != test.result {
			t.Errorf("StringInValueChecker expected value %s, value %s, result %t",
				test.expectedValues, test.data, test.result)
		}
	}
}

func TestNodeStatusEqualValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValue workflowmodel.NodeStatus
		data          workflowmodel.NodeStatus
		result        bool
	}{
		{workflowmodel.Active, workflowmodel.Active, true},
		{workflowmodel.Complete, workflowmodel.Aborted, false},
	}

	for _, test := range tests {
		checker := workflowmodel.NodeStatusEqualValueChecker{
			ExpectedValue: test.expectedValue,
		}
		if checker.CheckValue(test.data) != test.result {
			t.Errorf("NodeStatusEqualValueChecker expected value %s, value %s, result %t",
				test.expectedValue, test.data, test.result)
		}
	}
}

func TestNodeStatusInValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValues []workflowmodel.NodeStatus
		data           workflowmodel.NodeStatus
		result         bool
	}{
		{
			[]workflowmodel.NodeStatus{workflowmodel.Active, workflowmodel.Complete, workflowmodel.Submitted},
			workflowmodel.Active,
			true,
		},
		{
			[]workflowmodel.NodeStatus{workflowmodel.Complete, workflowmodel.Queued},
			workflowmodel.Aborted,
			false},
	}

	for _, test := range tests {
		checker := workflowmodel.NodeStatusInValueChecker{
			ExpectedValues: test.expectedValues,
		}
		if checker.CheckValue(test.data) != test.result {
			t.Errorf("NodeStatusInValueChecker expected values %s, value %s, result %t",
				test.expectedValues, test.data, test.result)
		}
	}
}

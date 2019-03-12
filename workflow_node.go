package workflowmodel

type WorkflowNodeVariables struct {
	Path                  string         `json:"path"`
	VariableList          []NodeVariable `json:"variable_list"`
	GeneratedVariableList []NodeVariable `json:"generated_variable_list"`
}

type WorkflowNode struct {
	Name                  string                  `json:"name"`
	Status                NodeStatus              `json:"status"`
	Type                  NodeType                `json:"node_type"`
	VariableList          WorkflowNodeVariables   `json:"variable_list"`
	InheritedVariableList []WorkflowNodeVariables `json:"inherited_variable_list"`
}

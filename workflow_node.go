package workflowmodel

type WorkflowNodeVariables struct {
	Path                  string         `json:"path"`
	VariableList          []NodeVariable `json:"variable_list"`
	GeneratedVariableList []NodeVariable `json:"generated_variable_list"`
}

func (vars *WorkflowNodeVariables) GetVariable(name string) *NodeVariable {
	for _, variable := range vars.VariableList {
		if variable.Name == name {
			return &variable
		}
	}

	for _, variable := range vars.GeneratedVariableList {
		if variable.Name == name {
			return &variable
		}
	}

	return nil
}

type WorkflowNode struct {
	Name                  string                  `json:"name"`
	Status                NodeStatus              `json:"status"`
	Type                  NodeType                `json:"node_type"`
	VariableList          WorkflowNodeVariables   `json:"variable_list"`
	InheritedVariableList []WorkflowNodeVariables `json:"inherited_variable_list"`
}

func (node *WorkflowNode) GetVariable(name string) *NodeVariable {
	var variable *NodeVariable
	variable = node.VariableList.GetVariable(name)
	if variable != nil {
		return variable
	}

	for _, variableList := range node.InheritedVariableList {
		variable = variableList.GetVariable(name)
		if variable != nil {
			return variable
		}
	}

	return nil
}

package workflowmodel

type WorkflowNodeCondition interface {
	IsFit(node *WorkflowNode) bool
}

type WorkflowNodeStatusCondition struct {
	Checker NodeStatusValueChecker
}

func (c *WorkflowNodeStatusCondition) IsFit(node *WorkflowNode) bool {
	return c.Checker.CheckValue(node.Status)
}

type WorkflowNodeVariableCondition struct {
	VariableName string
	Checker      StringValueChecker
}

func (c *WorkflowNodeVariableCondition) IsFit(node *WorkflowNode) bool {
	variable := node.GetVariable(c.VariableName)
	if variable == nil {
		return false
	}

	return c.Checker.CheckValue(variable.Value)
}

package workflowmodel

type StringValueChecker interface {
	CheckValue(s string) bool
}

type StringEqualValueChecker struct {
	ExpectedValue string
}

func (c *StringEqualValueChecker) CheckValue(s string) bool {
	return s == c.ExpectedValue
}

type StringInValueChecker struct {
	ExpectedValues []string
}

func (c *StringInValueChecker) CheckValue(s string) bool {
	for _, v := range c.ExpectedValues {
		if s == v {
			return true
		}
	}
	return false
}

type NodeStatusValueChecker interface {
	CheckValue(status NodeStatus) bool
}

type NodeStatusEqualValueChecker struct {
	ExpectedValue NodeStatus
}

func (c *NodeStatusEqualValueChecker) CheckValue(s NodeStatus) bool {
	return s == c.ExpectedValue
}

type NodeStatusInValueChecker struct {
	ExpectedValues []NodeStatus
}

func (c *NodeStatusInValueChecker) CheckValue(s NodeStatus) bool {
	for _, v := range c.ExpectedValues {
		if s == v {
			return true
		}
	}
	return false
}

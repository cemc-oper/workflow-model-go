package workflowmodel

type NodeType uint

const (
	UnknownNodeType NodeType = iota
	Root
	Suite
	Family
	Task
	Alias
	NonTaskNode
	Meter
)

func GetNodeType(s string) NodeType {
	switch s {
	case "root":
		return Root
	case "suite":
		return Suite
	case "family":
		return Family
	case "task":
		return Task
	case "alias":
		return Alias
	case "non-task":
		return NonTaskNode
	case "meter":
		return Meter
	default:
		return UnknownNodeType
	}
}

func (nodeType NodeType) String() string {
	switch nodeType {
	case Root:
		return "root"
	case Suite:
		return "suite"
	case Family:
		return "family"
	case Task:
		return "task"
	case Alias:
		return "alias"
	case NonTaskNode:
		return "non-task"
	case Meter:
		return "meter"
	default:
		return "unknown"
	}
}

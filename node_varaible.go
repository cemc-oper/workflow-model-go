package workflowmodel

type NodeVariableType uint

const (
	Variable NodeVariableType = iota
	GeneratedVariable
)

type NodeVariable struct {
	Name  string           `json:"name"`
	Value string           `json:"value"`
	Type  NodeVariableType `json:"type"`
}

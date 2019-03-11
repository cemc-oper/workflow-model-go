package workflowmodel

import (
	"encoding/json"
	"strings"
)

type Node struct {
	Name     string
	Parent   *Node
	Children []*Node
	Status   NodeStatus
}

func (node *Node) AddChild(child *Node) {
	node.Children = append(node.Children, child)
	child.Parent = node
}

func (node *Node) NodePath() string {
	var nodeList []string
	currentNode := node
	for currentNode != nil {
		nodeList = append([]string{currentNode.Name}, nodeList...)
		currentNode = currentNode.Parent
	}
	path := strings.Join(nodeList, "/")
	if path == "" {
		path = "/"
	}
	return path
}

type nodeDict struct {
	Name     string     `json:"name"`
	Status   int        `json:"status"`
	Children []nodeDict `json:"children"`
}

func (node *Node) toNodeDict() nodeDict {
	childrenDict := make([]nodeDict, 0)
	for _, child := range node.Children {
		childrenDict = append(childrenDict, child.toNodeDict())
	}

	return nodeDict{
		Name:     node.Name,
		Status:   int(node.Status),
		Children: childrenDict,
	}
}

func createFromNodeDict(d nodeDict) Node {
	var node Node
	node.Name = d.Name
	node.Status = NodeStatus(d.Status)
	for _, c := range d.Children {
		cNode := createFromNodeDict(c)
		node.Children = append(node.Children, &cNode)
	}
	return node
}

func (node Node) MarshalJSON() ([]byte, error) {
	d := node.toNodeDict()
	return json.Marshal(&d)
}

func (node *Node) UnmarshalJSON(b []byte) error {
	var d nodeDict
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}

	*node = createFromNodeDict(d)

	return nil
}

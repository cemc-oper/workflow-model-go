package workflowmodel

import "strings"

type NodeStatusRecord struct {
	Path   string
	Status string
}

type Bunch struct {
	Node
}

func (bunch *Bunch) FindNode(path string) *Node {
	if path == "" {
		return nil
	}

	if path == "/" {
		return &bunch.Node
	}

	tokens := strings.Split(path, "/")
	tokens = tokens[1:]

	node := &bunch.Node
	for _, token := range tokens {
		var tNode *Node
		tNode = nil
		for _, child := range node.Children {
			if child.Name == token {
				tNode = child
				break
			}
		}
		if tNode == nil {
			return nil
		}
		node = tNode
	}
	return node
}

func (bunch *Bunch) AddNode(path string) *Node {
	if path == "" {
		return nil
	}

	if path == "/" {
		return &bunch.Node
	}

	tokens := strings.Split(path, "/")
	tokens = tokens[1:]

	node := &bunch.Node
	for _, token := range tokens {
		var tNode *Node
		tNode = nil
		for _, child := range node.Children {
			if child.Name == token {
				tNode = child
				break
			}
		}
		if tNode == nil {
			tNode = &Node{}
			tNode.Name = token
			node.AddChild(tNode)
		}
		node = tNode
	}
	return node
}

func (bunch *Bunch) AddNodeStatus(record NodeStatusRecord) *Node {
	node := bunch.AddNode(record.Path)
	node.Status = GetNodeStatus(record.Status)
	return node
}

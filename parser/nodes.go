package parser

import (
	"fmt"
	"go/token"
)

type Node interface {
	Type() NodeType
	// Position() Pos
	String() string
	Copy() Node
}

type NodeType int

func (t NodeType) Type() NodeType {
	return t
}

const (
	NodeIdent NodeType = iota
	NodeString
	NodeNumber
	NodeCall
	NodeVector
)

type IdentNode struct {
	// Pos
	NodeType
	Ident string
}

func (node *IdentNode) Copy() Node {
	return NewIdentNode(node.Ident)
}

func (node *IdentNode) String() string {
	if node.Ident == "nil" {
		return "()"
	}

	return node.Ident
}

// STRING NODE ----------------------------------
type StringNode struct {
	// Pos
	NodeType
	Value string
}

func (node *StringNode) Copy() Node {
	return newStringNode(node.Value)
}

func (node *StringNode) String() string {
	return node.Value
}

// NUMBER NODE ----------------------------------
type NumberNode struct {
	// Pos
	NodeType
	Value      string
	NumberType token.Token
}

func (node *NumberNode) Copy() Node {
	return &NumberNode{NodeType: node.Type(), Value: node.Value, NumberType: node.NumberType}
}

func (node *NumberNode) String() string {
	return node.Value
}

// VECTOR NODE ----------------------------------
type VectorNode struct {
	// Pos
	NodeType
	Nodes []Node
}

func (node *VectorNode) Copy() Node {
	vect := &VectorNode{NodeType: node.Type(), Nodes: make([]Node, len(node.Nodes))}
	for i, v := range node.Nodes {
		vect.Nodes[i] = v.Copy()
	}
	return vect
}

func (node *VectorNode) String() string {
	return fmt.Sprint(node.Nodes)
}

// CALL NODE ----------------------------------
type CallNode struct {
	// Pos
	NodeType
	Callee Node
	Args   []Node
}

func (node *CallNode) Copy() Node {
	call := &CallNode{NodeType: node.Type(), Callee: node.Callee.Copy(), Args: make([]Node, len(node.Args))}
	for i, v := range node.Args {
		call.Args[i] = v.Copy()
	}
	return call
}

func (node *CallNode) String() string {
	args := fmt.Sprint(node.Args)
	return fmt.Sprintf("(%s %s)", node.Callee, args[1:len(args)-1])
}

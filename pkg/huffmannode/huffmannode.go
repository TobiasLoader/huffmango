package huffmannode

import (
  "fmt"
)

type HuffmanNode struct {
  Name string
  Weight int
  Child1 *HuffmanNode
  Child2 *HuffmanNode
}

func NewHuffmanNode(name string, weight int) *HuffmanNode {
  return &HuffmanNode{
    Name:   name,
    Weight: weight,
    Child1: nil,
    Child2: nil,
  }
}

func (n *HuffmanNode) GetName() string {
  if n != nil {
    return n.Name;
  }
  return "nil";
}

func (n *HuffmanNode) GetWeight() int {
  if n != nil {
    return n.Weight;
  }
  return -1;
}

func (n *HuffmanNode) SetChild1(child *HuffmanNode) {
  n.Child1 = child;
}

func (n *HuffmanNode) SetChild2(child *HuffmanNode) {
  n.Child2 = child;
}

func (n *HuffmanNode) GetChild1() *HuffmanNode {
  if n != nil {
    return n.Child1;
  }
  return nil;
}

func (n *HuffmanNode) GetChild2() *HuffmanNode {
  if n != nil {
    return n.Child2;
  }
  return nil;
}

func (n *HuffmanNode) Show() {
  if n != nil {
    fmt.Println(n.GetName()+" = "+strconv.Itoa(n.GetWeight())+", children: ["+n.GetChild1().GetName()+", "+n.GetChild2().GetName()+"]")
  } else {
    fmt.Println("nil")
  }
}

package priority

import (
  "huffmango/pkg/node"
  "fmt"
)

// Priority Queue

type PriorityQueue struct {
  Queue []node.HuffmanNode
  Size int
}

func NewPriorityQueue() *PriorityQueue {
  return &PriorityQueue{
    Queue: []node.HuffmanNode{},
    Size: 0,
  }
}

// private functions

func (pq *PriorityQueue) getParent(i int) int {return (i-1)/2}
func (pq *PriorityQueue) getLeft(i int) int {return 2*i+1}
func (pq *PriorityQueue) getRight(i int) int {return 2*i+2}

func (pq *PriorityQueue) append(n *node.HuffmanNode) {
  pq.Queue = append(pq.Queue, *n)
}

func (pq *PriorityQueue) write(i int, n node.HuffmanNode) {
  pq.Queue[i] = n
}
func (pq *PriorityQueue) truncate() {
  pq.Queue = pq.Queue[:pq.Count()]
}

func (pq *PriorityQueue) incSize() {
  pq.Size += 1
}
func (pq *PriorityQueue) decSize() {
  pq.Size -= 1
}

func (pq *PriorityQueue) getNode(i int) node.HuffmanNode {
  return pq.Queue[i];
}

func (pq *PriorityQueue) getWeight(i int) int {
  return pq.getNode(i).Weight;
}

func (pq *PriorityQueue) heapify(i int) {
  if pq.Count() > 1 {
    left := pq.getLeft(i)
    right := pq.getRight(i)
    smallest := i
    if left < pq.Count() && pq.getWeight(left) < pq.getWeight(i) {
      smallest = left
    }
    if right < pq.Count() && pq.getWeight(right) < pq.getWeight(smallest) {
      smallest = right
    }
    if smallest != i {
      temp := pq.getNode(i)
      pq.write(i,pq.getNode(smallest))
      pq.write(smallest,temp)
      pq.heapify(smallest)
    }
  }
}

// public functions

func (pq *PriorityQueue) Peek() node.HuffmanNode {
  return pq.Queue[0];
}

func (pq *PriorityQueue) Show() {
  fmt.Println(pq.Queue)
}

func (pq *PriorityQueue) Count() int {
  return pq.Size;
}

func (pq *PriorityQueue) Enqueue(n *node.HuffmanNode) {
  pq.append(n)
  pq.incSize()
  current := pq.Count()-1
  for current > 0 {
    parent := pq.getParent(current)
    if pq.getWeight(parent) > pq.getWeight(current) {
      temp := pq.getNode(parent)
      pq.write(parent,pq.getNode(current))
      pq.write(current,temp)
      current = parent
    } else {
      current = 0
    }
  }
}

func (pq *PriorityQueue) Dequeue() *node.HuffmanNode {
  if pq.Count() > 0 {
    priority := pq.Peek()
    last := pq.getNode(pq.Count()-1)
    pq.write(0,last)
    pq.decSize()
    pq.truncate()
    pq.heapify(0)
    return &priority;
  } else {
    return nil;
  }
}
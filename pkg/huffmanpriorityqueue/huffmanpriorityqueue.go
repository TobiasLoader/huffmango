package huffmanpriorityqueue

import (
  "huffmango/pkg/huffmannode"
  "fmt"
)

// Priority Queue

type PriorityQueue struct {
  Queue []HuffmanNode
  Size int
}

func NewPriorityQueue() *PriorityQueue {
  return &PriorityQueue{
    Queue: []HuffmanNode{},
    Size: 0,
  }
}

// private functions

func (pq *PriorityQueue) getParent(i int) int {return (i-1)/2}
func (pq *PriorityQueue) getLeft(i int) int {return 2*i+1}
func (pq *PriorityQueue) getRight(i int) int {return 2*i+2}

func (pq *PriorityQueue) append(n *HuffmanNode) {
  pq.Queue = append(pq.Queue, *n)
}

func (pq *PriorityQueue) write(i int, n HuffmanNode) {
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

func (pq *PriorityQueue) getNode(i int) HuffmanNode {
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

func (pq *PriorityQueue) Peek() HuffmanNode {
  return pq.Queue[0];
}

func (pq *PriorityQueue) Show() {
  fmt.Println(pq.Queue)
}

func (pq *PriorityQueue) Count() int {
  return pq.Size;
}

func (pq *PriorityQueue) Enqueue(n *HuffmanNode) {
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

func (pq *PriorityQueue) Dequeue() *HuffmanNode {
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
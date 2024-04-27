# huffmango

### Go package

Huffman coding implemented in Go, [Wikipedia](https://en.wikipedia.org/wiki/Huffman_coding).

![mangotree](https://github.com/TobiasLoader/huffmango/assets/39600827/be583272-dbf5-417f-8112-fa613db0df74)

From `/huffmango` directory:
```bash
go run cmd/main.go
```

The core of the algorithm – in [algo.go](pkg/algo/algo.go) from line `84`:
```go
// create the Huffman tree
for p.Count()>1 {
  n1 := p.Dequeue()
  n2 := p.Dequeue()
  n0_name := n1.GetName()+n2.GetName()
  n0_weight := n1.GetWeight()+n2.GetWeight()
  n0 := node.NewHuffmanNode(n0_name,n0_weight)
  n0.SetChild1(n1)
  n0.SetChild2(n2)
  p.Enqueue(n0)
}

// extract the root
root := p.Dequeue()

// traverse the tree (with DFS) to create encoding
encoding := NewEncoding(name,blocksize)
encoding.DFS("",root)
encoding.Show()

return encoding;
```

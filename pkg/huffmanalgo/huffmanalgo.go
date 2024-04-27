package huffmanalgo

import (
  "huffmango/pkg/huffmannode"
  "huffmango/pkg/huffmanpriorityqueue"
  "fmt"
  "strconv"
  "sort"
)

// type Encoding

type Encoding struct {
  Name string
  Code map[string]string
  BlockSize int
}

func NewEncoding(name string, blocksize int) *Encoding {
  return &Encoding{
    Name: name,
    Code: make(map[string]string),
    BlockSize: blocksize,
  }
}

func (enc *Encoding) Show() {
  fmt.Println(enc.Name+" - block size "+strconv.Itoa(enc.BlockSize))
  symbols := make([]string, 0, len(enc.Code))
  for s := range enc.Code {
    symbols = append(symbols, s)
  }
  sort.Strings(symbols)
  for _, s := range symbols {
    fmt.Println("  "+s+": "+enc.Code[s])
  }
}

func (enc *Encoding) SetCodeword(symbol string, codeword string) {
  if len(symbol) == enc.BlockSize {
    enc.Code[symbol] = codeword;
  }
}

func (enc *Encoding) DFS(path string, node *HuffmanNode) {
  if node != nil {
    if node.GetChild1() == nil && node.GetChild2() == nil {
      enc.SetCodeword(node.GetName(),path);
    } else {
      enc.DFS(path+"0",node.GetChild1())
      enc.DFS(path+"1",node.GetChild2())
    }
  }
}

func checkBlockSize(blocksize int, symbol string) int {
  if blocksize == -1 {
    blocksize = len(symbol)
  } else {
    if blocksize != len(symbol) {
      return -1;
    }
  }
  return blocksize;
}

func HuffmanAlgo(name string, data map[string]int) *Encoding {
  blocksize := -1

  // initialise and build priority queue
  p := NewPriorityQueue()
  for symbol, count := range data {
    // check blocksize
    blocksize = checkBlockSize(blocksize,symbol)
    if blocksize==-1 {
      return NewEncoding("BLOCKSIZE ERROR: "+name,-1)
    }
    // add new node to priority queue
    n := NewHuffmanNode(symbol,count)
    p.Enqueue(n)
  }

  // create the Huffman tree
  for p.Count()>1 {
    n1 := p.Dequeue()
    n2 := p.Dequeue()
    n0_name := n1.GetName()+n2.GetName()
    n0_weight := n1.GetWeight()+n2.GetWeight()
    n0 := NewHuffmanNode(n0_name,n0_weight)
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
}

func avgMsgLength(enc *Encoding, data map[string]int, total int, multiplier float64) float64 {
  codewordTotalLength := 0
  for block, count := range data {
    lengthOfCodeword := 0
    if len(block) == enc.BlockSize {
      // assumes encoding block size same as data block size
      lengthOfCodeword = len(enc.Code[block])
    } else if enc.BlockSize == 1 {
      // assumes encoding block size is 1
      for _, char := range block {
        lengthOfCodeword += len(enc.Code[string(char)])
      }
    }
    codewordTotalLength += count*lengthOfCodeword
  }
  avgLength := multiplier * float64(codewordTotalLength)/float64(total)
  fmt.Println(enc.Name+" Average Message Length: "+fmt.Sprintf("%.2f", avgLength))
  return avgLength;
}
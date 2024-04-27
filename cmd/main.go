package main

import (
  "huffmango/pkg/huffmannode"
  "huffmango/pkg/huffmanpriorityqueue"
  "huffmango/pkg/huffmanalgo"
  "encoding/csv"
  "fmt"
  "log"
  "os"
  "strconv"
  "sort"
  "math"
)

func main() {
  csvnames := []string{"single_counts.csv","double_counts.csv"}
  csvfiles := make(map[string]map[string]int)
  csvtotals := make(map[string]int)
  imccsv := make(map[string]string)
  pairsfromsinglecounts := make(map[string]int)

  // LOADING FILES into MAPS
  for index, filename := range csvnames {
    file, err := os.Open("data/"+csvnames[index])
    if err != nil {
        log.Fatalf("Failed to open file: %v", err)
    }
    defer file.Close()
    csvfiles[filename] = make(map[string]int)
    reader := csv.NewReader(file)
    _, head_err := reader.Read()
    if head_err != nil {
        log.Fatalf("Failed to read header: %v", err)
    }
    total := 0
    for {
        record, err := reader.Read()
        if err != nil {
            break
        }
        i, err := strconv.Atoi(record[1])
        if err != nil {
            fmt.Println(err)
            return
        }
        csvfiles[filename][record[0]] = i
        total += i
    }
    csvtotals[filename] = total
  }
  // LOADING IMC into MAP
  file, err := os.Open("data/IMC.csv")
  if err != nil {
      log.Fatalf("Failed to open file: %v", err)
  }
  defer file.Close()
  reader := csv.NewReader(file)
  _, head_err := reader.Read()
  if head_err != nil {
      log.Fatalf("Failed to read header: %v", err)
  }
  for {
      record, err := reader.Read()
      if err != nil {
        break
      }
      imccsv[record[0]] = record[1]
  }

  for c1, n1 := range csvfiles["single_counts.csv"] {
    for c2, n2 := range csvfiles["single_counts.csv"] {
      // we don't need to normalise since all pairs
      // will be divided by the same total
      pairsfromsinglecounts[c1+c2] = n1*n2; // by independence
    }
  }

  MorseCode := NewEncoding("MorseCode",1)
  for symbol, codeword := range imccsv {
    MorseCode.SetCodeword(symbol,codeword)
  }
  MorseCode.Show()
  fmt.Println(csvfiles["single_counts.csv"])
  Q5b := HuffmanAlgo("Q5 (b)",csvfiles["single_counts.csv"])
  Q5c := HuffmanAlgo("Q5 (c)",pairsfromsinglecounts)
  Q5d := HuffmanAlgo("Q5 (d)",csvfiles["double_counts.csv"])

  fmt.Println("\n~\n")

  computeAvgLengthGivenDouble := func(enc *Encoding, base int) float64 {
    return avgMsgLength(enc,csvfiles["double_counts.csv"],csvtotals["double_counts.csv"],math.Log2(float64(base)))
  }

  computeAvgLengthGivenDouble(MorseCode, 3)
  computeAvgLengthGivenDouble(Q5b, 2)
  computeAvgLengthGivenDouble(Q5c, 2)
  computeAvgLengthGivenDouble(Q5d, 2)

  fmt.Println("\n~\n")

  infotheory3 := make(map[string]int)
  infotheory3["AAA"] = 343
  infotheory3["AAB"] = 49
  infotheory3["ABA"] = 49
  infotheory3["BAA"] = 49
  infotheory3["ABB"] = 7
  infotheory3["BAB"] = 7
  infotheory3["BBA"] = 7
  infotheory3["BBB"] = 1

  Sheet3Q4c := HuffmanAlgo("infotheory3",infotheory3)
  avgMsgLength(Sheet3Q4c,infotheory3,512,1)

}

// Some of the output:

/*
Q5 (b) - block size 1
  A: 1010
  B: 100000
  C: 110011
  D: 10111
  E: 1111
  F: 110001
  G: 100010
  H: 0100
  I: 0110
  J: 1100001010
  K: 11000011
  L: 10110
  M: 111000
  N: 0111
  O: 1001
  P: 100001
  Q: 11000010111
  R: 11101
  S: 0101
  T: 1101
  U: 111001
  V: 1100000
  W: 110010
  X: 110000100
  Y: 100011
  Z: 11000010110
  _: 00

  ~

  MorseCode Average Message Length: 44.61
  Q5 (b) Average Message Length: 8.24
  Q5 (c) Average Message Length: 8.18
  Q5 (d) Average Message Length: 7.44
*/

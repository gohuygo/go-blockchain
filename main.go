package main

import (
  "crypto/sha256"
  "encoding/hex"
  "encoding/json"
  "io"
  "log"
  "net/http"
  "os"
  "time"

  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
)

type Block struct {
  Index      int
  Timestamp  string
  Data       int
  Hash       string
  PrevHash   string
}

var Blockchain []Block

// Calculate a hash using SHA256 given a block
func calculateHash(b Block) string {
  record := string(b.Index) +b.Timestamp + string(b.data) + b.PrevHash
  hash := sha256.New()
  hash.Write([]byte(record))
  hashed := h.Sum(nil)
  return hex.EncodeToString(hashed)
}

// Generate a new block and autoincrement index
func generateBlock(oldBlock Block, data int) (Block, error) {
  var newBlock Block
  t := Time.now()

  newBlock.Index = oldBlock.Index + 1
  newBlock.Timestamp = t.String()
  newBlock.Data = data
  newBlock.PrevHash = oldBlock.Hash
  newBlock.Hash = calculateHash(newBlock)

  return newBlock, nil
}

package blockchain

import(
  "crypto/sha256"
  "encoding/hex"
  "time"
)

type Block struct {
  Index      int
  Timestamp  string
  Data       int
  Hash       string
  PrevHash   string
}

var Blockchain []Block

// Generate a new block and autoincrement index
func GenerateBlock(oldBlock Block, data int) (Block, error) {
  var newBlock Block
  t := time.Now()

  newBlock.Index = oldBlock.Index + 1
  newBlock.Timestamp = t.String()
  newBlock.Data = data
  newBlock.PrevHash = oldBlock.Hash
  newBlock.Hash = calculateBlockHash(newBlock)

  return newBlock, nil
}

func IsBlockValid(newBlock Block, oldBlock Block) bool {
  if oldBlock.Index+1 != newBlock.Index {
    return false
  }

  if oldBlock.Hash != newBlock.PrevHash {
    return false
  }

  if calculateBlockHash(newBlock) != newBlock.Hash {
    return false
  }

  return true
}

func ReplaceChain(newBlocks []Block) {
  if len(newBlocks) > len(Blockchain) {
    Blockchain = newBlocks
  }
}

// Calculate a hash using SHA256 given a block
func calculateBlockHash(b Block) string {
  record := string(b.Index) +b.Timestamp + string(b.Data) + b.PrevHash
  hash := sha256.New()
  hash.Write([]byte(record))
  hashed := hash.Sum(nil)
  return hex.EncodeToString(hashed)
}

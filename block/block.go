package block

import(
  "crypto/sha256"
  "encoding/hex"
  "time"
  "log"
)

type Block struct {
  Index        uint
  Timestamp    string
  Transaction  string
  Hash         string
  PrevHash     string
  Nonce        uint
}

var Blockchain []Block

// Generate a genesis block - will log fatal if a block already exists and terminate
func GenerateGenesisBlock(){
  if len(Blockchain) > 0 {
    log.Fatal("A genesis block already exists.")
    return
  }

  // TODO: Setup Nonce!
  genesisBlock := Block{0, time.Now().String(), "reddit.com - 1540542759 - Uber driver hair formed a perfect 25.", "", "", 0}
  genesisBlock.Hash = calculateBlockHash(genesisBlock)

  Blockchain = append(Blockchain, genesisBlock)
}

// Generate a new block and autoincrement index
func GenerateBlock(oldBlock Block, transaction string) (Block, error) {
  var newBlock Block
  t := time.Now()

  newBlock.Index = oldBlock.Index + 1
  newBlock.Timestamp = t.String()
  newBlock.Transaction = transaction
  newBlock.PrevHash = oldBlock.Hash
  newBlock.Hash = calculateBlockHash(newBlock)

  return newBlock, nil
}

func IsBlockValid(newBlock Block, oldBlock Block) bool {
  // TODO: Validate each UTXO
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
  record := string(b.Index) +b.Timestamp + string(b.Transaction) + b.PrevHash
  hash := sha256.New()
  hash.Write([]byte(record))
  hashed := hash.Sum(nil)
  return hex.EncodeToString(hashed)
}

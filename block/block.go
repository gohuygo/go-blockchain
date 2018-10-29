package block

import(
  "time"
  "log"
  "errors"
  "strconv"
  "fmt"

  "github.com/google/go-cmp/cmp"

  "github.com/gohuygo/go-blockchain/crypto"
)

type Block struct {
  Index        uint
  Timestamp    string
  Transaction  string
  Hash         []byte
  PrevHash     []byte
  Nonce        uint
}

var Blockchain []Block

const startingNonce = 0
const genesisNonce  = 521049

func (b *Block) SetHash(transaction string) {
  oldBlock := Blockchain[len(Blockchain)-1]
  t := time.Now()

  b.Index         = oldBlock.Index + 1
  b.Timestamp     = t.String()
  b.Transaction   = transaction
  b.PrevHash      = oldBlock.Hash
  b.Hash, b.Nonce = calculateBlockHash(*b, startingNonce)
}

func (b *Block) Data() []byte {
    index :=  strconv.Itoa(int(b.Index))
    transaction  := string(b.Transaction)
    prevHash := b.PrevHash
    nonce := strconv.Itoa(int(b.Nonce))

    data := []byte(index + transaction + string(prevHash) + nonce)
    return data
}

// Generate a new block and autoincrement index
func GenerateBlock(transaction string) (Block, error) {
  newBlock := &Block{}
  newBlock.SetHash(transaction)

  log.Println("Created Block #" + strconv.Itoa(int(newBlock.Index)))
  return *newBlock, nil
}

func IsBlockValid(newBlock Block) bool {
  oldBlock := Blockchain[len(Blockchain)-1]
  // TODO: Validate each UTXO
  if oldBlock.Index+1 != newBlock.Index {
    return false
  }

  if !cmp.Equal(oldBlock.Hash, newBlock.PrevHash) {
    return false
  }

  hash := crypto.DoubleSha256(newBlock.Data())

  if !cmp.Equal(hash, newBlock.Hash){
    return false
  }

  return true
}

func ReplaceChain(newBlocks []Block) {
  if len(newBlocks) > len(Blockchain) {
    Blockchain = newBlocks
  }
}

// Generate a genesis block - will log fatal if a block already exists and terminate
func GenerateGenesisBlock(){
  if len(Blockchain) > 0 {
    errors.New("A genesis block already exists.")
    return
  }
  genesisBlock := Block{
    0,
    time.Now().String(),
    "reddit.com - 1540542759 - Uber driver hair formed a perfect 25.",
    []byte(""),
    []byte(""),
    genesisNonce,
  }
  genesisBlock.Hash, genesisBlock.Nonce = calculateBlockHash(genesisBlock, genesisNonce)

  log.Println("Genesis Block created.")
  Blockchain = append(Blockchain, genesisBlock)
}

// Calculate a hash using SHA256 given a block
func calculateBlockHash(b Block, nonce uint) ([]byte, uint) {
  var blockHash []byte

  for {
    blockHash = crypto.DoubleSha256(b.Data())

    startsWithTarget := cmp.Equal(blockHash[:3], []byte("000"))

    if(startsWithTarget){
      fmt.Printf("%c\n", blockHash)
      log.Println("Solved with nonce: " + strconv.Itoa(int(nonce)))
      break;
    }

    nonce++
  }

  return blockHash, nonce
}

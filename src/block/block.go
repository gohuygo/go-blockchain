package block

import(
  "bytes"
  "time"
  "log"
  "errors"
  "strconv"

  "github.com/google/go-cmp/cmp"

  "github.com/gohuygo/go-blockchain/src/crypto"
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
const genesisNonce  = 89

const difficulty = 1

func (b *Block) header() []byte {
    index       :=  strconv.Itoa(int(b.Index))
    transaction := string(b.Transaction)
    prevHash    := b.PrevHash
    nonce       := strconv.Itoa(int(b.Nonce))

    return []byte(index + transaction + string(prevHash) + nonce)
}

// Generate a new block and autoincrement index
func New(transaction string) *Block {
  newBlock := &Block{}

  oldBlock := Blockchain[len(Blockchain)-1]
  t := time.Now()

  newBlock.Index         = oldBlock.Index + 1
  newBlock.Timestamp     = t.String()
  newBlock.Transaction   = transaction
  newBlock.PrevHash      = oldBlock.Hash

  newBlock.Hash,newBlock.Nonce = mine(*newBlock, startingNonce)

  log.Println("Created Block #" + strconv.Itoa(int(newBlock.Index)))
  return newBlock
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

  hash := crypto.DoubleSha256(append(newBlock.header(), byte(newBlock.Nonce)))

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
func GenerateGenesis(){
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
    0,
  }
  genesisBlock.Hash, genesisBlock.Nonce = mine(genesisBlock, genesisNonce)

  log.Println("Genesis Block created.")
  Blockchain = append(Blockchain, genesisBlock)
}

// Perform proof of work on a mine and return the blockhash and its nonce
func mine(b Block, nonce uint) ([]byte, uint) {
  var blockHash []byte

  targetPrefix := bytes.Repeat([]byte("0"), difficulty)

  for {
    log.Println(byte(nonce))
    guessBytes := append(b.header(), byte(nonce))
    blockHash = crypto.DoubleSha256(guessBytes)

    log.Println(blockHash[:difficulty])

    startsWithTarget := cmp.Equal(blockHash[:difficulty], targetPrefix)

    if(startsWithTarget){
      log.Println("Solved with nonce: " + strconv.Itoa(int(nonce)))
      break;
    }

    nonce++
  }

  return blockHash, nonce
}

package block

import(
  "time"
  "log"
  "errors"
  "strconv"

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

// Generate a new block and autoincrement index
func GenerateBlock(transaction string) (Block, error) {
  var newBlock Block

  oldBlock := Blockchain[len(Blockchain)-1]
  t := time.Now()

  newBlock.Index = oldBlock.Index + 1
  newBlock.Timestamp = t.String()
  newBlock.Transaction = transaction
  newBlock.PrevHash = oldBlock.Hash
  newBlock.Hash, newBlock.Nonce = calculateBlockHash(newBlock, startingNonce)

  log.Println("Created Block #" + strconv.Itoa(int(newBlock.Index)))
  return newBlock, nil
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

  hash := crypto.DoubleSha256(
    strconv.Itoa(int(newBlock.Index)),
    string(newBlock.Transaction),
    newBlock.PrevHash,
    strconv.Itoa(int(newBlock.Nonce)),
  )

  if !cmp.Equal(hash,newBlock.Hash){
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
  genesisBlock := Block{0, time.Now().String(), "reddit.com - 1540542759 - Uber driver hair formed a perfect 25.", []byte(""), []byte(""), genesisNonce}
  genesisBlock.Hash, genesisBlock.Nonce = calculateBlockHash(genesisBlock, genesisNonce)

  log.Println("Genesis Block created.")
  Blockchain = append(Blockchain, genesisBlock)
}

// Calculate a hash using SHA256 given a block
func calculateBlockHash(b Block, nonce uint) ([]byte, uint) {
  var encodedString []byte

  for {
    encodedString = crypto.DoubleSha256(
      strconv.Itoa(int(b.Index)),
      string(b.Transaction),
      b.PrevHash,
      strconv.Itoa(int(nonce)),
    )

    startsWithTarget := cmp.Equal(encodedString[:3], []byte("000")) // 0 maps to 48 in ASCII

    if(startsWithTarget){
      log.Println("Solved with nonce: " + strconv.Itoa(int(nonce)))
      break;
    }

    nonce++
  }

  return encodedString, nonce
}

package block

import(
  "time"
  "log"
  "strings"
  "strconv"

  "github.com/gohuygo/go-blockchain/crypto"
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

// Block difficulty is number of leading 0s.
// Every additional 0 decreases space by half (i.e. puzzle requires 2x hashing power to solve).
const blockTarget = "0000"

const startingNonce = 0
const genesisNonce  = 170

// Generate a new block and autoincrement index
func GenerateBlock(oldBlock Block, transaction string) (Block, error) {
  var newBlock Block
  t := time.Now()

  newBlock.Index = oldBlock.Index + 1
  newBlock.Timestamp = t.String()
  newBlock.Transaction = transaction
  newBlock.PrevHash = oldBlock.Hash
  newBlock.Hash, newBlock.Nonce = calculateBlockHash(newBlock, startingNonce)

  log.Println("Created Block #" + strconv.Itoa(int(newBlock.Index)))
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

  hash := crypto.DoubleSha(
    strconv.Itoa(int(newBlock.Index)),
    string(newBlock.Transaction),
    newBlock.PrevHash,
    strconv.Itoa(int(newBlock.Nonce)),
  )

  if hash != newBlock.Hash {
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
    log.Fatal("A genesis block already exists.")
    return
  }

  genesisBlock := Block{0, time.Now().String(), "reddit.com - 1540542759 - Uber driver hair formed a perfect 25.", "", "", genesisNonce}
  genesisBlock.Hash, genesisBlock.Nonce = calculateBlockHash(genesisBlock, genesisNonce)

  log.Println("Genesis Block created.")
  Blockchain = append(Blockchain, genesisBlock)
}

// Calculate a hash using SHA256 given a block
func calculateBlockHash(b Block, nonce uint) (string, uint) {
  startsWith := false
  encodedString := ""

  for {

    encodedString = crypto.DoubleSha(
      strconv.Itoa(int(b.Index)),
      string(b.Transaction),
      b.PrevHash,
      strconv.Itoa(int(nonce)),
    )

    startsWith = strings.HasPrefix(encodedString, blockTarget)

    if(startsWith){
      log.Println("Solved with nonce: " + strconv.Itoa(int(nonce)))
      break;
    }

    nonce++
  }

  return encodedString, nonce
}

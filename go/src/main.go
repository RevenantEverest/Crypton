package main

import ("fmt"
        "bytes"
      	"crypto/sha256"
      	"strconv"
      	"time"
        "math"
        "math/big"
        "encoding/binary"
        "encoding/gob"
        "log"
        )
var (
  maxNonce = math.MaxInt64
)

/*======== Block ======== */
type Block struct {
  Timestamp     int64
  Data          []byte
  PreviousHash  []byte
  Hash          []byte
  Nonce         int
}

func (b *Block) SetHash() {
  timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
  headers := bytes.Join([][]byte{b.PreviousHash, b.Data, timestamp}, []byte{})
  hash := sha256.Sum256(headers)

  b.Hash = hash[:]
}

func NewBlock(data string, previousHash []byte) *Block {
  block := &Block{time.Now().Unix(), []byte(data), previousHash, []byte{}, 0}
  pow := NewProofOfWork(block)
  nonce, hash := pow.Run()

  block.Hash = hash[:]
  block.Nonce = nonce
  //block.SetHash()
  return block
}

func (b *Block) Serialize() []byte {
  var result bytes.Buffer
  encoder := gob.NewEncoder(&result)

  err := encoder.Encode(b)
  return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
  var block Block
  decoder := gob.NewDecoder(bytes.NewReader(d))
  err := decoder.Decode(&block)

  return &block
}

func CreateGenesisBlock() *Block {
  return NewBlock("Genesis Block", []byte{})
}
/*-------- END --------*/

/*======== Blockchain ========*/
type Blockchain struct {
  //blocks []*Block
  tip []byte
  db *bolt.DB
}

func (bc *Blockchain) AddBlock(data string) {
  var lastHash []byte
  err := bc.db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(blocksBucket))
    lastHash = b.Get([]byte("1"))

    return nil
  })

  newBlock := NewBlock(data, lastHash)

  err = bc.db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(blocksBucket))
    err := b.Put(newBlock.Hash, newBlock.Serialize())
    err = b.Put([]byte("1"), newBlock.Hash)
    bc.tip = newBlock.Hash

    return nil
  })
}

func NewBlockchain() *Blockchain {
  var tip []byte
  db, err := bolt.Open(dbFile, 0000, nil)

  err =db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(blocksBucket))

    if b == nil {
      genesis := NewGenesisBlock()
      b, err := tx.CreateBucket([]byte(blocksBucket))
      err = b.Put(genesis.Hash, genesis.Serialize())
      err = b.put([]byte("1"), genesis.Hash)
      tip = genesis.Hash
    } else {
      tip = b.Get([]byte("1"))
    }

    return nil
  })

  bc := Blockchain{tip, db}
  return &bc
}

type BlockChainIterator struct {
  currentHash []byte
  db *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
  bci := &BlockchainIterator{bc.tip, bc.db}

  return bci
}
/*-------- END --------*/

/*======== Proof of Work ========*/
const targetBits = 24

type ProofOfWork struct {
  block *Block
  target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
  target := big.NewInt(1)
  target.Lsh(target, uint(256-targetBits))

  pow:= &ProofOfWork{b, target}
  return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PreviousHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
  var hashInt big.Int
  var hash [32]byte
  nonce := 0

  fmt.Printf("Mining block ... \"%s\"\n", pow.block.Data)
  for nonce < maxNonce {
    data := pow.prepareData(nonce)

    hash = sha256.Sum256(data)
    fmt.Printf("\r%x", hash)
    hashInt.SetBytes(hash[:])

    if hashInt.Cmp(pow.target) == -1 {
      break
    } else {
      nonce ++
    }
  }
  fmt.Print("\n\n")
  return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
  var hashInt big.Int

  data := pow.prepareData(pow.block.Nonce)
  hash := sha256.Sum256(data)
  hashInt.SetBytes(hash[:])

  isValid := hashInt.Cmp(pow.target) == -1
  return isValid
}

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

/*-------- END --------*/

/*======== CLI ========*/


type CLI struct {
  bc *Blockchain
}

func (cli *CLI) run() {
  cli.validateArgs()

  addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
  printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

  addBlockData := addBlockCmd.String("data", "", "Block data")

  switch os.Args[1] {
  case "addblock":
    err := addBlockCmd.Parse(os.Args[2:])
  case "printchain":
    err := printChainCmd.Parse(os.Args[2:])
  default:
    cli.printUsage()
    os.Exit(1)
  }

  if addBlockCmd.Parsed() {
    if *addBlockData == "" {
      addBlockCmd.Usage()
      os.Exit(1)
    }
    cli.addBlock(*addBlockData)
  }

  if printChainCmd.Parsed() {
    cli.printchain()
  }
}

/*-------- END --------*/

/*======== MAIN ========*/
func main() {
	bc := NewBlockchain()

	bc.AddBlock("New Block: Sent 1 Crypton")
	bc.AddBlock("New Block: Sent 3 Crypton")

	for _, block := range bc.blocks {
		fmt.Printf("Data: %s\n", block.Data)
    fmt.Printf("Previous hash: %x\n", block.PreviousHash)
		fmt.Printf("Hash: %x\n", block.Hash)
    pow := NewProofOfWork(block)
    fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
/*-------- END --------*/

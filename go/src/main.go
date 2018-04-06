package main

import ("fmt"
        "bytes"
      	"crypto/sha256"
      	"strconv"
      	"time"
        "math"
        "math/big"
        "encoding/binary"
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

func CreateGenesisBlock() *Block {
  return NewBlock("Genesis Block", []byte{})
}
/*-------- END --------*/

/*======== Blockchain ========*/
type Blockchain struct {
  blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
  prevBlock := bc.blocks[len(bc.blocks)-1]
  newBlock := NewBlock(data, prevBlock.Hash)
  bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
  return &Blockchain{[]*Block{CreateGenesisBlock()}}
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

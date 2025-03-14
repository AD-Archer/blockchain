package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Block struct { // This is the type for the blocks in the blockchain
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct { // This is the type for the blockchain
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func (b Block) calculateHash() string { // This function calculates the hash of the block
	data, _ := json.Marshal(b.data)                                                         // This converts the data in the block to a JSON string
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow) // This is the data that will be hashed
	blockHash := sha256.Sum256([]byte(blockData))                                           // This hashes the block data
	return fmt.Sprintf("%x", blockHash)                                                     // This returns the hash as a string
}

func (b *Block) mine(difficulty int) { // This function mines the block
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) { // This checks if the hash of the block has the correct number of leading zeros
		b.pow++                    // This increments the proof of work
		b.hash = b.calculateHash() // This recalculates the hash of the block
	}
}

func CreateBlockchain(difficulty int) Blockchain { // This function creates the first block in the blockchain
	genesisBlock := Block{ // This creates the first block in the blockchain
		hash:      "0",        // This is the hash of the first block it has to be 0 cause it is the first block
		timestamp: time.Now(), // This is the timestamp of the first block
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) addBlock(from, to string, amount float64) { // This function adds a new block to the blockchain
	blockData := map[string]interface{}{ // This is the data that will be added to the block
		"from":   from,   // This is the sender of the block
		"to":     to,     // This is the receiver of the block
		"amount": amount, // This is the amount of money that is being sent
	}
	lastBlock := b.chain[len(b.chain)-1] // This is the last block in the blockchain
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)         // This mines the new block
	b.chain = append(b.chain, newBlock) // This adds the new block to the blockchain
}

func (b *Blockchain) validateBlock(block Block) bool { // This function validates the block
	lastBlock := b.chain[len(b.chain)-1]
	return block.hash == block.calculateHash() && block.previousHash == lastBlock.hash
}

func (b Blockchain) isValid() bool { // This function validates the blockchain
	for i := range b.chain[1:] { // This loops through the blockchain
		previousBlock := b.chain[i]                                                                               // This is the previous block in the blockchain
		currentBlock := b.chain[i+1]                                                                              // This is the current block in the blockchain
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash { // This checks if the hash of the current block is valid
			return false
		}
	}
	return true
}

func (b Blockchain) printBlockchain() { // This function prints the blockchain after each transaction
	for i, block := range b.chain {
		print("Block %d:\n", i)
		print("  Data: from: %s, to: %s, amount: $%.2f\n", block.data["from"], block.data["to"], block.data["amount"])
		print("  Hash: %s\n", block.hash)
		print("  Previous Hash: %s\n", block.previousHash)
		print("  Timestamp: %s\n", block.timestamp)
		print("  Proof of Work: %d\n", block.pow)
		print("\n")
	}
}

func main() {
	// create a new blockchain instance with a mining difficulty of 2
	blockchain := CreateBlockchain(2)

	reader := bufio.NewReader(os.Stdin)

	for {
		// Prompt the user for transaction details
		var from, to string
		var amountStr string
		var amount float64

		fmt.Print("Enter sender: ")
		from, _ = reader.ReadString('\n')
		from = strings.TrimSpace(from)

		fmt.Print("Enter receiver: ")
		to, _ = reader.ReadString('\n')
		to = strings.TrimSpace(to)

		fmt.Print("Enter amount: ")
		amountStr, _ = reader.ReadString('\n')
		amountStr = strings.TrimSpace(amountStr)

		// Validate and clean the amount input
		re := regexp.MustCompile(`[^\d.]`)
		cleanAmountStr := re.ReplaceAllString(amountStr, "")
		amount, err := strconv.ParseFloat(cleanAmountStr, 64)
		if err != nil {
			fmt.Println("Invalid amount. Please enter a valid number.")
			continue
		}

		// Add the transaction to the blockchain
		blockchain.addBlock(from, to, amount)

		// Print the entire blockchain
		blockchain.printBlockchain()

		// Check if the blockchain is valid; expecting true
		fmt.Println("Is the blockchain valid?", blockchain.isValid())
	}
}

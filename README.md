# Blockchain Implementation in Go

This project is a simple blockchain implementation written in Go. It demonstrates the basic concepts of blockchain technology, including block creation, mining, and validation.

## Features

- **Block Creation**: Add new blocks with transaction data.
- **Mining**: Proof-of-work mining to secure the blockchain.
- **Validation**: Ensure the integrity of the blockchain.
- **Interactive Console**: Continuously add transactions via the console.

## Prerequisites

- Go 1.16 or later

## Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/ad-archer/blockchain
   cd blockchain
   ```

2. **Run the application**:
   ```bash
   go run blockchain.go
   ```

## Usage

- The application will prompt you to enter transaction details (sender, receiver, and amount).
- Transactions are added to the blockchain, and the entire blockchain is printed after each transaction.
- The blockchain's validity is checked and displayed after each transaction.

## Code Overview

- `Block`: Represents a single block in the blockchain.
- `Blockchain`: Manages the chain of blocks and provides methods for adding and validating blocks.
- `main()`: Interactive loop for adding transactions and displaying the blockchain.

## License

This project is open-source and available under the MIT License.

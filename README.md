# EtherealVault: Verifiable, Permissionless, and Confidential State Machine Replication

EtherealVault is a decentralized, Merkle-rooted state machine replication system engineered for building verifiable, permissionless, and confidential blockchain applications. It leverages zk-SNARKs (zero-knowledge succinct non-interactive arguments of knowledge) to achieve provable state transitions while ensuring data privacy. The system uses WebAssembly (WASM) smart contracts, offering flexibility and interoperability across different platforms. EtherealVault aims to provide a robust and secure foundation for developing decentralized applications (dApps) that require strong guarantees of correctness and confidentiality.

The core purpose of EtherealVault is to address the limitations of traditional blockchain systems concerning scalability, privacy, and security. By employing Merkle trees for state representation and zk-SNARKs for verifiable computation, EtherealVault minimizes the amount of data that needs to be stored and verified on the blockchain, significantly improving efficiency. Privacy is enhanced by allowing computations to be performed on encrypted data, with zk-SNARKs providing proof of correctness without revealing the underlying data. This makes EtherealVault suitable for applications such as confidential voting systems, private decentralized exchanges, and secure supply chain management.

EtherealVault offers a unique combination of features that sets it apart from other blockchain platforms. The WASM-based smart contract environment enables developers to write smart contracts in multiple languages, increasing accessibility and fostering innovation. The system incorporates a novel consensus mechanism optimized for zk-SNARK verification, ensuring fast and efficient block creation. Furthermore, EtherealVault includes a comprehensive suite of developer tools, including a compiler, debugger, and testing framework, to streamline the development and deployment of dApps.

**Key Features**

*   **Decentralized State Machine Replication:** EtherealVault replicates the state across a network of nodes, ensuring high availability and fault tolerance. The state is represented as a Merkle tree, allowing for efficient state updates and verification. Each state transition is validated by a decentralized consensus mechanism.

*   **zk-SNARK Integration:** The system utilizes zk-SNARKs to enable verifiable computation on encrypted data. This allows for privacy-preserving smart contracts where the execution logic and state remain confidential. Specifically, we use Groth16 for its proven efficiency and security. The circuit definitions are written in Circom and proven with SnarkJS.

*   **WASM Smart Contracts:** Smart contracts are written in WebAssembly, enabling interoperability and support for multiple programming languages. This approach allows developers to leverage existing codebases and expertise, simplifying the development process. A custom WASM runtime has been developed, optimized for execution within the EtherealVault environment.

*   **Merkle-Rooted State:** The entire state of the system is represented as a Merkle tree. This allows for efficient verification of state transitions and minimizes the data that needs to be stored and verified on-chain. The Merkle tree implementation uses a sparse Merkle tree construction with a depth of 256.

*   **Custom Consensus Mechanism:** EtherealVault employs a Byzantine Fault Tolerant (BFT) consensus mechanism optimized for zk-SNARK verification. This ensures fast and efficient block creation while maintaining security and resilience against malicious actors. The consensus algorithm uses a multi-signature scheme for block attestation.

*   **Verifiable Computation:** All state transitions are verified using zk-SNARKs, ensuring the integrity of the system. This eliminates the need for full replication of computation, significantly improving scalability. The zk-SNARK proofs are verified on-chain using precompiled contracts for gas efficiency.

*   **Developer Tooling:** EtherealVault includes a comprehensive suite of developer tools, including a compiler, debugger, and testing framework. These tools are designed to streamline the development and deployment of dApps on the platform. The compiler supports multiple languages that compile to WASM, including Rust and C++.

**Technology Stack**

*   **Go:** The core EtherealVault platform is written in Go, leveraging its concurrency features and performance characteristics for efficient state machine replication.
*   **WebAssembly (WASM):** Smart contracts are executed within a WASM runtime, providing a secure and portable execution environment.
*   **zk-SNARKs (Groth16):** zk-SNARKs are used to enable verifiable computation on encrypted data, ensuring privacy and integrity. We use Circom and SnarkJS for circuit definition and proof generation.
*   **Merkle Trees:** Merkle trees are used to represent the state of the system, enabling efficient state updates and verification.
*   **gRPC:** gRPC is used for inter-process communication between nodes in the network, ensuring high performance and reliability.
*   **LevelDB:** LevelDB is used as a persistent storage engine for storing the state of the blockchain.

**Installation**

1.  Install Go: Ensure you have Go 1.20 or later installed. Verify with `go version`.
2.  Clone the repository: `git clone https://github.com/jjfhwang/EtherealVault.git`
3.  Navigate to the project directory: `cd EtherealVault`
4.  Install dependencies: `go mod download`
5.  Build the EtherealVault executable: `go build -o etherealvault`

**Configuration**

EtherealVault can be configured using environment variables. The following variables are supported:

*   `ETHEREALVAULT_PORT`: The port to listen on for gRPC connections (default: 50051).
*   `ETHEREALVAULT_DATA_DIR`: The directory to store blockchain data (default: ./data).
*   `ETHEREALVAULT_PRIVATE_KEY`: The private key of the node (required for participating in consensus). This should be a hexadecimal string representing a 32-byte private key.
*   `ETHEREALVAULT_BOOTSTRAP_NODES`: A comma-separated list of bootstrap node addresses (e.g., `127.0.0.1:50051,127.0.0.1:50052`).

Example:

ETHEREALVAULT_PORT=50052 ETHEREALVAULT_DATA_DIR=/tmp/evdata ./etherealvault

**Usage**

To start an EtherealVault node:

./etherealvault

After starting a node, you can interact with it using gRPC. API documentation will be published separately. Example Go code to interact with the network:

package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"etherealvault/proto" // Assuming you have a proto package
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewEtherealVaultClient(conn) // Assuming you have an EtherealVaultClient

	r, err := c.GetState(context.Background(), &proto.GetStateRequest{Key: "somekey"}) // Example call
	if err != nil {
		log.Fatalf("could not get state: %v", err)
	}
	log.Printf("State: %s", r.GetValue())
}

**Contributing**

We welcome contributions to EtherealVault! Please follow these guidelines:

*   Fork the repository.
*   Create a new branch for your changes.
*   Write clear and concise commit messages.
*   Submit a pull request with a detailed description of your changes.
*   Ensure that your code adheres to the Go coding style.
*   Include unit tests for any new functionality.

**License**

This project is licensed under the MIT License. See the [LICENSE](https://github.com/jjfhwang/EtherealVault/blob/main/LICENSE) file for details.

**Acknowledgements**

We would like to thank the following projects and individuals for their contributions to the field of zero-knowledge proofs and blockchain technology:

*   The Zcash team for their pioneering work on zk-SNARKs.
*   The Ethereum community for their contributions to smart contract development.
*   The authors of the Circom and SnarkJS libraries.
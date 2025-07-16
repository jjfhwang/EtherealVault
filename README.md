# EtherealVault - Secure and Ephemeral Data Storage in Go

EtherealVault provides a robust and secure solution for managing sensitive data within a Go application, emphasizing ephemeral storage to minimize long-term exposure. It allows developers to encrypt, store, and retrieve data with configurable expiration times, ensuring that sensitive information is automatically purged when no longer needed. This approach reduces the risk of data breaches and simplifies compliance requirements related to data retention policies.

EtherealVault addresses the critical need for secure data handling in modern applications, particularly those dealing with Personally Identifiable Information (PII), API keys, or temporary access tokens. Unlike traditional databases or key-value stores, EtherealVault prioritizes data destruction after a defined lifespan. This "fire-and-forget" approach minimizes the attack surface and helps adhere to principles of least privilege. The library leverages industry-standard encryption algorithms and configurable storage backends to provide a flexible and secure solution for a wide range of use cases. It is designed to be easily integrated into existing Go projects with minimal code changes.

This library offers significant benefits for applications requiring stringent security and compliance. By automatically removing data after its intended purpose, EtherealVault reduces the risk of data leaks due to compromised databases or forgotten storage locations. Furthermore, the configurable expiration times allow developers to tailor the storage duration to specific security requirements. The transparent encryption and decryption processes ensure that sensitive data remains protected both at rest and in transit. EtherealVault is designed with performance in mind, minimizing overhead and ensuring seamless integration into high-throughput applications.

## Key Features

*   **Ephemeral Data Storage:** Data is automatically deleted after a configurable expiration time, ensuring that sensitive information is not stored indefinitely. The expiration time is defined during data insertion and is tracked internally using the `time.AfterFunc` function.
*   **AES-256 Encryption:** Data is encrypted using the Advanced Encryption Standard (AES) with a 256-bit key, providing robust protection against unauthorized access. Encryption is performed using the `crypto/aes` package and `crypto/cipher` package in the Go standard library. Key generation uses `crypto/rand` to ensure cryptographic strength.
*   **Configurable Storage Backend:** EtherealVault supports multiple storage backends, including in-memory storage (for development and testing) and persistent storage options like Redis (for production environments). The storage interface is defined to allow for the easy addition of new backend implementations. Currently, Redis leverages the `github.com/go-redis/redis/v8` client.
*   **Key Rotation Support:** The library supports key rotation, allowing you to change the encryption key periodically without disrupting existing data. This is achieved by storing the key identifier along with the encrypted data and providing a mechanism to retrieve the correct key for decryption. The implementation uses a versioning system for keys.
*   **Concurrency-Safe Operations:** EtherealVault is designed to be thread-safe, allowing you to use it in concurrent applications without worrying about data corruption or race conditions. Internal locking mechanisms (using `sync.Mutex`) protect critical data structures.
*   **Error Handling:** The library provides comprehensive error handling, making it easy to identify and resolve issues. All functions return detailed error messages, allowing you to understand the root cause of problems. Custom error types are defined for specific failure scenarios.
*   **Data Integrity Verification:** HMAC (Hash-based Message Authentication Code) is used to verify the integrity of the data, ensuring that it has not been tampered with. HMAC is calculated using the SHA-256 hash function (`crypto/sha256`) and a secret key.

## Technology Stack

*   **Go:** The primary programming language used for developing the library. Go provides excellent performance, concurrency support, and a strong standard library.
*   **crypto/aes:** Go standard library package for Advanced Encryption Standard (AES) encryption. Used to encrypt and decrypt sensitive data.
*   **crypto/cipher:** Go standard library package for cryptographic cipher implementations. Used in conjunction with `crypto/aes` for encryption operations.
*   **crypto/rand:** Go standard library package for generating cryptographically secure random numbers. Used for generating encryption keys and initialization vectors (IVs).
*   **crypto/sha256:** Go standard library package for SHA-256 hashing. Used for creating HMACs to ensure data integrity.
*   **github.com/go-redis/redis/v8:** Go Redis client. Used for persistent storage when the Redis backend is selected.
*   **sync:** Go standard library package for synchronization primitives (e.g., Mutexes). Used to ensure thread safety.
*   **time:** Go standard library package for time-related functions. Used for managing data expiration times.

## Installation

1.  Ensure you have Go installed and configured correctly. You can download Go from [https://golang.org/dl/](https://golang.org/dl/).
2.  Clone the EtherealVault repository:
    `git clone https://github.com/jjfhwang/EtherealVault.git`
3.  Navigate to the repository directory:
    `cd EtherealVault`
4.  Install the required dependencies:
    `go mod tidy`
5.  Build the library:
    `go build`

## Configuration

EtherealVault uses environment variables for configuration. The following environment variables are supported:

*   `ETHEREALVAULT_STORAGE_BACKEND`: Specifies the storage backend to use. Valid values are `memory` (for in-memory storage) and `redis` (for Redis storage). Defaults to `memory`.
*   `ETHEREALVAULT_REDIS_ADDR`: The address of the Redis server (e.g., `localhost:6379`). Required if `ETHEREALVAULT_STORAGE_BACKEND` is set to `redis`.
*   `ETHEREALVAULT_REDIS_PASSWORD`: The password for the Redis server. Optional.
*   `ETHEREALVAULT_ENCRYPTION_KEY`: The encryption key to use for AES-256 encryption. It must be a 32-byte (256-bit) key. If not provided, a random key will be generated on startup (not recommended for production).

Set these environment variables before running your application. For example:

`export ETHEREALVAULT_STORAGE_BACKEND=redis`
`export ETHEREALVAULT_REDIS_ADDR=localhost:6379`
`export ETHEREALVAULT_ENCRYPTION_KEY=YOUR_SECURE_ENCRYPTION_KEY_HERE`

## Usage

Here's a basic example of how to use EtherealVault:

// Import the EtherealVault package
// import "github.com/jjfhwang/EtherealVault"

// Initialize EtherealVault with the desired configuration
// vault, err := EtherealVault.NewVault()
// if err != nil {
//  panic(err)
// }

// Store data with an expiration time of 1 hour
// data := []byte("sensitive data to be stored")
// expiration := time.Hour
// id, err := vault.Store(data, expiration)
// if err != nil {
//  panic(err)
// }

// Retrieve the data
// retrievedData, err := vault.Retrieve(id)
// if err != nil {
//  panic(err)
// }

// fmt.Printf("Retrieved data: %s\n", string(retrievedData))

// After 1 hour, the data will be automatically deleted

For more detailed API documentation and examples, refer to the code comments within the EtherealVault package.

## Contributing

We welcome contributions to EtherealVault! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Write clear and concise code with thorough comments.
4.  Include unit tests to verify your changes.
5.  Submit a pull request with a detailed description of your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/jjfhwang/EtherealVault/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the Go community for providing excellent tools and resources that made this project possible.
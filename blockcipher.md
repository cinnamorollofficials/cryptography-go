## Blockcipher

Blockciper is an algorithm that encrypts a fixed-size of chunk (or "block") of data at a time. It's a fundamental component of symetric-key cryptography, meaning it uses the same secret key for both encrypting and decrypting the information.

Think of it like a highly complex, reversible shredder.
- Encryption: You feed in a block of plaintext(e.g 128 bits of a document) and a secret key. The cipher scrambles and outputs a 128-bit block of ciphertext (unreadable data).
- Decryption: You feed in that exact block of ciphertext and the same secret key. The cipher "un-scrambles" it, returning your original 128-bit block of plaintext.

This process is deterministoc: the same plaintext block with the same key will always produce the same ciphertext block (in the simplest mode). 

### Key concepts
Here are the core components that define how a block cipher works:
- Block size: This is the fixed amount of data the cipher operates on at once. A very common and secure block size today is 128 bits (16 bytes), used by the Advanced Encryption Standart (AES).
- Key Size: This is the length of secret key, which determine the cipher's securoty agains brute-force attacks. Common key sizes for AES are 128, 192 or 256 bits, A longe key means exponentially more possible combinations an attacker would have to try.

### Modes of Operaction
You rarely use a block cipher on just one block. To encrypt a large file or data stream, you must use a mode of operation. This is a set of rules that defines how to repetedly apply the cipher to encrypt multiple blocks of data. 

Here are common modes:
- **Electronic Codebook (ECB)**:
    - **How its works**: Each block is encrypted independently.
    - **Problem**: This is insecure and should not be used. If two plaintext blocks are identical (e.g a block of all-white pixels in an image ), their ciphertext blocks will also be identical. This leaks patterns.
- **Cipher Block Chaining (CBC)**:
    - **How its works**: Before encrypting a block, its XORed with the previous ciphertext block. This "chains" them together, so each block's encryption depends on all previous blocks.
    - **Benefit**: Identical plaintext blocks will result in different ciphertext blocks, hiding patterns.
- **Counter (CTR)**:
    - **How its works**: This mode effectively turns the block cipher into a stream cipher. Its a "counter" value for each block, then XORs that result with plaintext
    - **Benefit**: It's fast parallelizable (you can encrypt/decrypt multiple blocks at once), and doesnt require padding.


### Popular Block Cipher Algorithms 
- **AES (Advanced Encryption Standard)**: The modern-dat, secure standard. it's fast, secure, and used worldwide in everything from wi-fi (WPA2/3) to VPNs and file encryption.
- **DES (Data Encryption Standard)**: An older standard from the 1970s. Its 56-bit key is now insecure and can be broken by modern computers.
- **3DES (Triple DES)**: A successor to DES that applies the DES algorithm three times to increase the key size. It's much slower than AES and is now considered deprecated in favor of AES.

### Minimalis Example
```go

```

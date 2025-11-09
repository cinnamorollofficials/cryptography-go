## Padding

Padding is the process of adding extra data to a message before encryption to ensure its length is a multiple of the blockcipher's block size.

Think of it it this way: a block cipher is like a factory machine that only accepts shipping boxes of a perfectly spesific size (e.g., exact 16 bytes).

If your message is 20 bytes long, you can't just feed into the machine.
- You fill the first 16-byte box.
- You are left with 4 bytes, This "partial box" will be rejected.

Padding is the "filler material" (like styrofoam peanuts) you add to those last 4 bytes to create complete 16-byte box.

### Why is Padding Necessary?

Block ciphers, like the AES, are "dumb" in this spesific way: they only know how to encrypt a fixed-size block of data (e.g., 128 bits/16 bytes). The cannot operate a complete 16-byte box.

Padding provides a standard, reversible way to solve this "last block problem".

A crucial feature of any padding scheme is that must be unambigous. The person decrypting the message must be able to confidently remove the exact amount of padding that was added, and not accidentally remove any of the original message.

### How Padding Works (Common Scheme)

The most common padding scheme is PKCS#7. Its the logic is very simple:
- You add N bytes of padding, and the value of each byte i N.

Let's see this with an 8-byte block cipher.

**Case 1: Your last block is incomplete.**
- Plaintext (5 bytes):`HELLO`
- Block Size: 8 bytes
- Bytes Needed: 8 - 5 = 3 bytes
- You will add 3 bytes of padding, each with the value 0x03
- Padded Block: `[H E L L O 0x03 0x03 0x03]`

**Case 2: Your last block is exactly full.**
- Plaintext (8 bytes): `PASSWORD`
- Block Size: 8 bytes
- Bytes Needed: 0
- Problem: If you don't add padding, how will the decryptor know if the last byte (`D`) is part of the message or padding?
- Solution: You always add padding. In this case, you add an entirely new bock of padding.
- Bytes Needed: 8
- Padded Block 1: `[P A S S W O R D]`
- Padded Block 2: `[0x08 0x08 0x08 0x08 0x08 0x08 0x08 0x08]`

When decrypting, the application looks at the very last byte.
- If the value is `0x03`, it removes 3 bytes
- If the value is `0x08`, it removes 8 bytes
- If the value is `0x01`, it removes 1 bytes

This makes the process perfectly reversible.

### Security Warning: Padding Oracle Attacks

While padding is a functional requirement, it can be also create a security vulnerability if not handled carefully.

A padding oracle attack is a clever attack where an attacker send slightly modified ciphertext to a server and observers the error messages it returns.

- If the server responds with a "Bad Padding" error, the attacker learns the decryption failed at the padding step.
- If is responds with a different error (e.g., "Invalid data"), the attacker learns the padding was correct, even if the rest of the message was garbage.

By repeatedly "poking" the server and obsering these different error message, the attacker can act like and "oracle" and slowly, byte-by-byte, deduce the contents of the original plaintext without ever knowing the encryption key. This is why modern systems must handle padding errors very carefully. 


# Encrypt-Algorithms

RSA is an asymmetric encryption algorithm that uses two keys: a public key and a private key. The public key can be widely shared and used to encrypt messages, while the private key is kept secret and used for decryption.

The RSA algorithm is based on the mathematical problem of factoring large integers and includes the following steps:

1. **Key generation**:
     - Two large prime numbers `p` and `q` are selected.
     - You take into account their product `n = p*q`, which is used as a modulus for both keys.
     - The value of the Euler function from `n` is calculated: `φ(n) = (p - 1) * (q - 1)`.
     - You form compounds of the number `e`, for example, `1 < e < φ(n)` and `e` is coprime with `φ(n)`. The number `e` becomes the public key along with `n`.
     - Calculate the number `d`, the inverse of `e` modulo `φ(n)`: `d * e ≡ 1 (mod φ(n))`. The number `d` becomes the private key.

2. **Encryption**:
     - To encrypt the message `M` (where `M` is a number less than `n`), the formula `C = M^e mod n` is used, where `C` is the encrypted message.

3. **Decoding**:
     - To decrypt messages, the formula `M = C^d mod n` is used, where `M` is the original message.
  

AES is a symmetric encryption algorithm that uses the same key to encrypt and decrypt data. AES processes data in fixed-length blocks (typically 128 bits) and can use keys of varying lengths (128, 192, or 256 bits).

The AES algorithm includes several rounds of data processing, each of which consists of four stages:

1. SubBytes: non-linear replacement of each byte of a block with another byte using a lookup table (S-box).
2. ShiftRows: cyclic shift of rows of a data block.
3. MixColumns: Mixing the data of each column of the block.
4. AddRoundKey: adding a block of data with a round key (derived from the main encryption key).

The number of rounds depends on the key length: 10 rounds for a 128-bit key, 12 rounds for a 192-bit key, and 14 rounds for a 256-bit key.

DES is a legacy symmetric encryption algorithm that also uses a single key for encryption and decryption. It processes data in 64-bit blocks and uses a 56-bit key (actually 64 bits, but 8 bits are used for parity checking and are not involved in encryption).

The DES algorithm includes the following steps:

1. Initial Permutation (IP): The block data undergoes an initial permutation.
2. 16 rounds of encryption: each round includes expansion of half a block of data, addition with the round key, substitution using S-box, and permutation.
3. Final Permutation (FP): After 16 rounds, the reverse of the initial permutation is performed.

DES is considered insecure for use in modern applications due to its short key length, which makes it vulnerable to brute-force attacks. It is recommended to use more secure algorithms such as AES instead of DES.

# DaVinci
Fastest and secure way to encrypt and decrypt large files.

## Installation
You can download binary files for each platform from [the latest releases](https://github.com/eminmuhammadi/davinci/releases).


## Usage
- Generate passphrase (stored in a file)
```bash
davinci new-passphrase --folder ./path-to-folder
```
Returns a passphrase in the file `passphrase.txt`

- Generate public and private key (stored in a file)
```bash
davinci new-keypair --size 2048 --passphrase ./passphrase.key --folder ./path-to-folder
```
Returns the public key in the file `publicKey.pem` and the private key in the file `privateKey.pem`.

- Generate Symmetric key (stored in a file)
```bash
davinci key --folder ./path-to-folder
```
Return the symmetric key in the file `key.txt`.

- Encrypt a file
```bash
davinci encrypt --input ./file.ext --output ./file-decrypted.ext --key ./key.txt --passphrase ./passphrase.key --public-key ./publicKey.pem
```
Encrypts the file `file.ext` using the key `key.txt` (encrypted using RSA) and stores the result in `file-decrypted.ext`.

- Decrypt a file
```bash
davinci decrypt --input ./file-decrypted.ext --output ./file.ext --passphrase ./passphrase.key --private-key ./privateKey.pem
```
Decrypts the file `file-decrypted.ext` using the key `key.txt` (encrypted using RSA) and stores the result in `file.ext`.

## Building (obfuscated version)
```bash
go install
```

```bash
bash build.sh

# or

chmod +x build.sh
./build.sh
```

## Algorithm Specification
Key Generation:
- Generate the public and private key pair, afterwards we call this `pubK` and `privK`.
- Generate the symmetric key, afterwards we call this `R`.

Encryption:
- Encrypt symmetric key using the public key. `Enc(R, pubK)`, afterwards we call this `R_enc`.
- Encrypt file using the encrypted symmetric key. `Enc(file, R_enc)`, afterwards we call this `file_enc`.

Decryption:
- Decrypt symmetric key using the private key. `Dec(R_enc, privK)=R`
- Decrypt file using the decrypted symmetric key. `Dec(file_enc, R)`

## Requirements
- Golang 1.16+
- Garble https://github.com/burrowers/garble (optional, for obfuscation binary build file)

## Dependencies
- CLI Framework
  - github.com/cpuguy83/go-md2man/v2
  - github.com/russross/blackfriday/v2
  - github.com/urfave/cli/v2
  - github.com/xrash/smetrics
- Source code
  - strconv
  - encoding/base64
  - crypto/aes
  - crypto/cipher
  - crypto/rand
  - crypto/rsa
  - crypto/sha256
  - crypto/x509
  - encoding/pem
  - fmt
  - io/ioutil
  - os
  - path/filepath
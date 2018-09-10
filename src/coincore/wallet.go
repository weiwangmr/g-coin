package coincore

import (
	"fmt"
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"utils"

	"github.com/btcsuite/golangcrypto/ripemd160"
)

const VERSION = byte(0x00)
const WALLET_FILE = "wallet.dat"
const ADDRESS_CHECKSUM_LEN = 4

// Wallet stores private and public keys
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey []byte
}

// NewWallet creates and returns a Wallet
func NewWallet() *Wallet {
	private, public := newKeyPair() 
	wallet := Wallet{private, public}

	return &wallet
}

// GetAddress returns  wallet address
func (w Wallet) GetAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)

	versionedPayload := append([]byte{VERSION}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := utils.Base58Encode(fullPayload)

	return address
}

// HashPubKey hases public key
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

// ValidateAddress check if address if valid
func ValidateAddress(address string) bool {
	pubKeyHash := utils.Base58Decode([]byte(address))
	fmt.Println(pubKeyHash)
	actualChecksum := pubKeyHash[len(pubKeyHash) - ADDRESS_CHECKSUM_LEN:]
	fmt.Println(actualChecksum)
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash) - ADDRESS_CHECKSUM_LEN]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Compare(actualChecksum, targetChecksum) == 0
}
 
// Checksum generates a checksum for a public key
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:ADDRESS_CHECKSUM_LEN]
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}
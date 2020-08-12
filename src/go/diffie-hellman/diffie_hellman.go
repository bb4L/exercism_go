package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// PrivateKey create correct private key
func PrivateKey(p *big.Int) *big.Int {
	key := new(big.Int)
	limit := new(big.Int).Sub(p, big.NewInt(2))
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return key.Rand(seed, limit).Add(key, big.NewInt(2))
}

// PublicKey return public key
func PublicKey(privateKey, p *big.Int, g int64) (A *big.Int) {
	return new(big.Int).Exp(big.NewInt(g), privateKey, p)
}

// NewPair returns a new pair of keis
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey returns the secret key
func SecretKey(privateKey, publicKey, p *big.Int) *big.Int {
	return new(big.Int).Exp(publicKey, privateKey, p)
}

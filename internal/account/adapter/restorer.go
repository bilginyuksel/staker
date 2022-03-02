package adapter

import (
	"crypto/ed25519"

	"github.com/algorand/go-algorand-sdk/mnemonic"
)

type Restorer struct {
}

func NewRestorer() *Restorer {
	return &Restorer{}
}

func (r *Restorer) FromMnemonic(m string) (ed25519.PrivateKey, error) {
	return mnemonic.ToPrivateKey(m)
}

func (r *Restorer) FromPrivateKey(privateKey ed25519.PrivateKey) (string, error) {
	return mnemonic.FromPrivateKey(privateKey)
}

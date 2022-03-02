package adapter

import (
	"crypto/ed25519"

	"github.com/algorand/go-algorand-sdk/client/kmd"
)

type KMD struct {
	client kmd.Client
}

func NewKMD(url, token string) (*KMD, error) {
	c, err := kmd.MakeClient(url, token)
	return &KMD{c}, err
}

// InitWalletHandle given id and password return the wallet handle token
func (k *KMD) InitWalletHandle(id, password string) (string, error) {
	res, err := k.client.InitWalletHandle(id, password)
	return res.WalletHandleToken, err
}

// ImportKey given wallet handle token and private key
// return the address if successful
func (k *KMD) ImportKey(walletToken string, key ed25519.PrivateKey) (string, error) {
	res, err := k.client.ImportKey(walletToken, key)
	return res.Address, err
}

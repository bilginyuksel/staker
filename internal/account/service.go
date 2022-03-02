package account

import (
	"context"
	"crypto/ed25519"
	"errors"
	"time"
)

var (
	ErrMnemoicInvalid = errors.New("mnemonic is invalid")
)

const (
	_walletID       = "test"
	_walletPassword = "test"
)

type Account struct {
	Address    string
	Balance    uint64
	ImportedAt time.Time
}

type Restorer interface {
	FromMnemonic(mnemonic string) (ed25519.PrivateKey, error)
	FromPrivateKey(privateKey ed25519.PrivateKey) (string, error)
}

type (
	AlgodClient interface {
		// AccountInformation retrive the account information in node
		AccountInformation(ctx context.Context, address string) (Account, error)
	}

	KmdClient interface {
		// InitWalletHandle given id and password return the wallet handle token
		InitWalletHandle(id, password string) (string, error)
		// ImportKey given wallet handle token and private key
		// return the address if successful
		ImportKey(walletToken string, key ed25519.PrivateKey) (string, error)
	}
)

type Service struct {
	restorer Restorer
	algod    AlgodClient
	kmd      KmdClient
}

func NewService(restorer Restorer, algod AlgodClient, kmd KmdClient) *Service {
	return &Service{
		restorer: restorer,
		algod:    algod,
		kmd:      kmd,
	}
}

// Import provide the mnemonic to import the account to wallet
// if successful return the public key otherwise return the error
func (s *Service) Import(ctx context.Context, mnemonic string) (Account, error) {
	address, err := s.importAccountFrom(mnemonic)
	if err != nil {
		return Account{}, err
	}

	return s.algod.AccountInformation(ctx, address)
}

func (s *Service) importAccountFrom(mnemonic string) (string, error) {
	privateKey, err := s.restorer.FromMnemonic(mnemonic)
	if err != nil {
		return "", ErrMnemoicInvalid
	}

	walletToken, err := s.kmd.InitWalletHandle(_walletID, _walletPassword)
	if err != nil {
		return "", err
	}

	return s.kmd.ImportKey(walletToken, privateKey)
}

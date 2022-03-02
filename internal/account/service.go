package account

import "context"

type Account struct {
	PubKey  string
	Balance int64
}

type NodeClient interface {
	ImportAccount(ctx context.Context, mnemonic string) (Account, error)
}

type Service struct {
	nodeClient NodeClient
}

// Import provide the mnemonic to import the account to wallet
// if successful return the public key otherwise return the error
func (s *Service) Import(ctx context.Context, mnemonic string) (Account, error) {
	return s.nodeClient.ImportAccount(ctx, mnemonic)
}

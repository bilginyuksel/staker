package adapter

import (
	"context"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/bilginyuksel/staker/internal/account"
)

type Algod struct {
	client *algod.Client
}

func NewAlgod(url, token string) (*Algod, error) {
	c, err := algod.MakeClient(url, token)
	return &Algod{c}, err
}

func (a *Algod) AccountInformation(ctx context.Context, address string) (account.Account, error) {
	res, err := a.client.AccountInformation(address).Do(ctx)
	return account.Account{
		Address: address,
		Balance: res.Amount,
	}, err
}

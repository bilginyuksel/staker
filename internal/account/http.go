package account

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTPService interface {
	Import(ctx context.Context, mnemonic string) (Account, error)
	Get(ctx context.Context) ([]Account, error)
}

type HTTP struct {
	service HTTPService
}

func (h *HTTP) RegisterRoutes(router *echo.Echo) {
	router.POST("/accounts/import", h.ImportAccount)
	router.GET("/accounts", h.GetAccounts)
}

func NewHTTP(s HTTPService) *HTTP {
	return &HTTP{s}
}

type (
	importAccountReq struct {
		Mnemonic       string `json:"mnemonic"`
		WalletID       string `json:"wallet_id"`
		WalletPassword string `json:"wallet_password"`
	}

	importAccountRes struct {
		Address string `json:"address"`
		Balance uint64 `json:"balance"`
	}
)

// ImportAccount provide the mnemonic to import the account to
// the node to be able to stake.
func (h *HTTP) ImportAccount(c echo.Context) error {
	var req importAccountReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	account, err := h.service.Import(c.Request().Context(), req.Mnemonic)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, importAccountRes{
		Address: account.Address,
		Balance: account.Balance,
	})
}

type (
	accountDetails struct {
		Address string `json:"address"`
		Balance uint64 `json:"balance"`
	}

	getAccountRes []accountDetails
)

func fromAccounts(accounts []Account) (res getAccountRes) {
	for _, account := range accounts {
		res = append(res, fromAccount(account))
	}

	return
}

func fromAccount(account Account) accountDetails {
	return accountDetails{
		Address: account.Address,
		Balance: account.Balance,
	}
}

// GetAccounts fetch all accounts with their public id and balance.
// In the default wallet or provide the wallet id to fetch the accounts
func (h *HTTP) GetAccounts(c echo.Context) error {
	accounts, err := h.service.Get(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, fromAccounts(accounts))
}

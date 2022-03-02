package account

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTPService interface {
	Import(ctx context.Context, mnemonic string) (string, error)
}

type HTTP struct {
	service HTTPService
}

func (h *HTTP) RegisterRoutes(router *echo.Echo) {
	router.POST("/accounts/import", h.ImportAccount)
	router.GET("/accounts", h.GetAccounts)
}

type (
	importAccountReq struct {
		Mnemonic string `json:"mnemonic"`
	}

	importAccountRes struct {
		PubKey  string `json:"pub_key"`
		Balance int64  `json:"balance"`
	}
)

// ImportAccount provide the mnemonic to import the account to
// the node to be able to stake.
func (h *HTTP) ImportAccount(c echo.Context) error {
	var req importAccountReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	pubKey, err := h.service.Import(c.Request().Context(), req.Mnemonic)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, importAccountRes{
		PubKey:  pubKey,
		Balance: 0,
	})
}

// GetAccounts fetch all accounts with their public id and balance.
// In the default wallet or provide the wallet id to fetch the accounts
func (h *HTTP) GetAccounts(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

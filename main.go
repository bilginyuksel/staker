package main

import (
	"log"

	"github.com/bilginyuksel/staker/internal/account"
	"github.com/bilginyuksel/staker/internal/account/adapter"
	"github.com/labstack/echo/v4"
)

const (
	_algodURL   = "http://localhost:8080"
	_algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	_kmdURL   = "http://localhost:8888"
	_kmdToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
)

func main() {
	restorer := adapter.NewRestorer()
	algodAdapter, err := adapter.NewAlgod(_algodURL, _algodToken)
	if err != nil {
		panic(err)
	}

	kmdAdapter, err := adapter.NewKMD(_kmdURL, _kmdToken)
	if err != nil {
		panic(err)
	}

	accountService := account.NewService(restorer, algodAdapter, kmdAdapter)
	httpHandler := account.NewHTTP(accountService)

	e := echo.New()
	httpHandler.RegisterRoutes(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

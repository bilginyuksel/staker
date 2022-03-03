package main

import (
	"log"
	"os"

	"github.com/bilginyuksel/staker/internal/account"
	"github.com/bilginyuksel/staker/internal/account/adapter"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := readConfig()

	restorer := adapter.NewRestorer()
	algodAdapter, err := adapter.NewAlgod(conf.AlgoD.URL, conf.AlgoD.Token)
	if err != nil {
		panic(err)
	}

	kmdAdapter, err := adapter.NewKMD(conf.KMD.URL, conf.KMD.Token)
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

type Config struct {
	AlgoD struct {
		URL   string
		Token string
	}

	KMD struct {
		URL   string
		Token string
	}
}

func readConfig() Config {
	return Config{
		AlgoD: struct {
			URL   string
			Token string
		}{
			URL:   os.Getenv("ALGOD_URL"),
			Token: os.Getenv("ALGOD_TOKEN"),
		},
		KMD: struct {
			URL   string
			Token string
		}{
			URL:   os.Getenv("KMD_URL"),
			Token: os.Getenv("KMD_TOKEN"),
		},
	}
}

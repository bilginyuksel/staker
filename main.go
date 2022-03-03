package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bilginyuksel/staker/internal/account"
	"github.com/bilginyuksel/staker/internal/account/adapter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.Logger())
	httpHandler.RegisterRoutes(e)

	if err := e.Start(fmt.Sprintf(":%d", conf.AppPort)); err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	AppPort int

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
	port, err := strconv.Atoi(getenv("APP_PORT", "8888"))
	if err != nil {
		panic(err)
	}

	return Config{
		AppPort: port,
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

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

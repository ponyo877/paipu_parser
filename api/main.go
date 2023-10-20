package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ponyo877/paipu_parser/api/handler"
	"github.com/ponyo877/paipu_parser/repository"
	"github.com/ponyo877/paipu_parser/usecase/paipu"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	repository := repository.NewPaipuMahjongSoul()
	paipuService := paipu.NewService(repository)

	handler.MakePaipuHandlers(e, paipuService)

	e.Logger.Fatal(e.Start(":8080"))
}

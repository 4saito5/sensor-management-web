package main

import (
	"./sign"
	"./getPort"
	"./setPort"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echoのインスタンスを作成
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	// e.Use(middleware.Gzip())


	// ルーティング
	e.POST("/signin", sign.SignIn)
	e.POST("/signup", sign.SignUp)
	e.POST("/getPort", getPort.getPort)
	e.POST("/setPort", setPort.setPort)

	// サーバー起動
	e.Start(":5555")
}

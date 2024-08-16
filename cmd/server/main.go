package main

import (
	"github.com/iarsham/fasthttp-crud/internal/db"
	"github.com/iarsham/fasthttp-crud/internal/routers"
	"github.com/iarsham/fasthttp-crud/pkg/logger"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func main() {
	logs, err := logger.NewZapLog()
	if err != nil {
		panic(err)
	}

	sqliteDB, err := db.OpenDB()
	if err != nil {
		logs.Error("Failed to open sqlite database", zap.Error(err))
	}
	defer sqliteDB.Close()
	logs.Info("Database opened successfully")

	logs.Info("Server started on: ", zap.String("host", "0.0.0.0"), zap.String("port", ":8080"))
	if err := fasthttp.ListenAndServe(":8080", routers.SetupRouter(sqliteDB, logs).Handler); err != nil {
		logs.Error("Failed to start server", zap.Error(err))
	}
}

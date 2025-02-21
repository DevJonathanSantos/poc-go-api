package main

import (
	"log/slog"

	logger "github.com/DevJonathanSantos/poc-go-api/confiig"
)

func main() {
	logger.InitLogger()

	slog.Info("starting api")
}

package main

import (
	"github.com/AlexMinsk2017/PetAutchTest/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//  инициализировать объект конфига
	cfg := config.MustLoad()

	//  инициализировать логер
	log := setupLogger(cfg.Env)

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		// envLocal: локальный запуск — используем TextHandler, удобный для консоли,
		// и уровень логирования Debug (т.е. будем выводить все сообщения)
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		// envDev: запуск на удалённом dev-сервере — уровень логирования тот же, но формат вывода — JSON,
		// удобный для систем сбора логов (Kibana, Grafana Loki и т.п.)
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		// envProd: запуск в продакшене: повышаем уровень логирования до Info — нам не нужны debug-логи в проде.
		// Т.е. мы будем получать сообщения только с уровнем Info или Error.
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log

}

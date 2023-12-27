package main

import (
	"log/slog"
	"os"

	"github.com/Oden333/gRPC_SSO-service/internal/config"
	"github.com/Oden333/gRPC_SSO-service/internal/lib/logger/handlers/slogpretty"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//Инициализируем конфиг
	cfg := config.MustLoad()

	// Инициализируем логгер
	// Кастомный логгер просто своровал с исходного репо гита
	log := setupLogger(cfg.Env)
	log.Info("starting application", slog.Any("config", cfg))

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения
}

// Логгер может зависит от окружения, поэтому используем отдельную функцию
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:

		//хендлер - объект, который представляет логи.
		//	Это интерфейс с методом Handle, он пишет данные у удобном формате. Используем встроенный хендлер NewTextHandler - текстовый вывод
		//назначаем параметры, писать в консоль, а уровень lvldebug - чтобы писал вообще всё что есть (смотреть уровни LevelDebug)
		/*
			log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)
		*/

		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

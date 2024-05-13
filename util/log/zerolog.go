package log

import (
	"os"

	"github.com/rs/zerolog"
	"goosechase.ai/email-scheduler/config"
)

var logger zerolog.Logger

func Initialize() {
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	isProd := config.Env("ENV", "development") == "production"
	if isProd {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Error() *zerolog.Event {
	return logger.Error()
}

func Fatal() *zerolog.Event {
	return logger.Fatal()
}

func Panic() *zerolog.Event {
	return logger.Panic()
}

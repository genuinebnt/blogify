package logs

import (
	"io"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func Init() {
	var writer io.Writer

	if isLocalEnv, _ := strconv.ParseBool(os.Getenv("LOCAL_ENV")); isLocalEnv {
		writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	} else {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		writer = os.Stdout
	}

	logger = zerolog.New(writer).With().Timestamp().Logger()
}

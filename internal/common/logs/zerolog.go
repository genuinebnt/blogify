package logs

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/genuinebnt/blogify/internal/common/config"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(cfg *config.Config) {
	var writer io.Writer

	if env := cfg.Env; env == "development" {
		writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	} else {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		writer = os.Stdout
	}

	log.Logger = zerolog.New(writer).With().Timestamp().Logger()
}

type ZeroLogLogger struct {
	Logger zerolog.Logger
}

func (l *ZeroLogLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := l.Logger.Info()

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		entry = entry.Str("req_id", reqID)
	}

	entry = entry.Str("http_method", r.Method).
		Str("remote_addr", r.RemoteAddr).
		Str("uri", r.RequestURI)

	return &ZeroLogLoggerEntry{
		event: entry,
	}
}

type ZeroLogLoggerEntry struct {
	event *zerolog.Event
}

func (e *ZeroLogLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	e.event.
		Int("resp_status", status).
		Int("resp_bytes", bytes).
		Dur("resp_duration", elapsed).
		Msg("request complete")
}

// Panic logs panic messages
func (e *ZeroLogLoggerEntry) Panic(v interface{}, stack []byte) {
	zerolog.DefaultContextLogger.Error().
		Interface("panic", v).
		Bytes("stack", stack).
		Msg("panic occurred")
}

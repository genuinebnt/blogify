package logs

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/middleware"
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

// GetLogger returns the global logger instance
func GetLogger() *zerolog.Logger {
	return &logger
}

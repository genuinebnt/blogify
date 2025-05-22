package logs

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

func NewStructuredLogger(logger *zerolog.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

type StructuredLogger struct {
	Logger *zerolog.Logger
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := l.Logger.Info()

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		entry = entry.Str("req_id", reqID)
	}

	entry = entry.Str("http_method", r.Method).
		Str("remote_addr", r.RemoteAddr).
		Str("uri", r.RequestURI)

	return &StructuredLoggerEntry{
		event: entry,
	}
}

type StructuredLoggerEntry struct {
	event *zerolog.Event
}

func (e *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	e.event.
		Int("resp_status", status).
		Int("resp_bytes", bytes).
		Dur("resp_duration", elapsed).
		Msg("request complete")
}

// Panic logs panic messages
func (e *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	zerolog.DefaultContextLogger.Error().
		Interface("panic", v).
		Bytes("stack", stack).
		Msg("panic occurred")
}

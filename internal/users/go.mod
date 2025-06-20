module github.com/genuinebnt/blogify/internal/users

go 1.24.3

require github.com/genuinebnt/blogify/internal/common v0.0.0

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/go-chi/render v1.0.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.5 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/text v0.25.0 // indirect
)

require (
	github.com/go-chi/chi v1.5.5 // indirect
	github.com/go-chi/chi/v5 v5.2.1
	github.com/google/uuid v1.6.0
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rs/zerolog v1.34.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
)

replace github.com/genuinebnt/blogify/internal/common => ../common

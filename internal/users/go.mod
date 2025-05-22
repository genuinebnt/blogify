module github.com/genuinebnt/blogify/internal/users

go 1.24.3

require github.com/genuinebnt/blogify/internal/common v0.0.0

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/go-chi/render v1.0.3 // indirect
)

require (
	github.com/go-chi/chi v1.5.5 // indirect
	github.com/go-chi/chi/v5 v5.2.1
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rs/zerolog v1.34.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
)

replace github.com/genuinebnt/blogify/internal/common => ../common

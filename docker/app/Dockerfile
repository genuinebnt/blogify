FROM golang:1.24

RUN go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest && go install github.com/cespare/reflex@latest

COPY reflex.conf /

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]

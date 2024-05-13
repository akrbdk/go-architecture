FROM golang:1.22-alpine AS builder
LABEL authors="alexander.karabudka"

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

#dependencies
COPY ["app/go.mod", "app/go.sum", "./"]
RUN go mod download

#build
COPY app ./
RUN go build -o ./bin/app cmd/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /
COPY configs/config.yaml /config.yaml

CMD ["/app"]

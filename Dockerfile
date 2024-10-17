FROM registry.suse.com/bci/golang:1.23 as builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./app cmd/sky-mayor/main.go

FROM registry.suse.com/bci/bci-micro:15.6

COPY --from=builder /build/app /app

ENTRYPOINT ["/app"]

FROM cgr.dev/chainguard/go:latest AS builder

WORKDIR /work

COPY cmd/kumodemo/main.go /work/
COPY go.mod /work/

RUN CGO_ENABLED=0 go build -o bin/kumodemo *.go

FROM cgr.dev/chainguard/static:latest

COPY --from=builder /work/bin/kumodemo /kumodemo
CMD ["/kumodemo"]
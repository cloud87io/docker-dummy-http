FROM golang:1.19 as builder
COPY . .
RUN go build \
  -ldflags "-linkmode external -extldflags -static" \
  -a server.go

FROM scratch
ENV PORT=8080
COPY --from=builder /go/server ./server
CMD ["./server"]
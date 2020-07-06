FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build
COPY go.mod .
COPY go.sum .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o server cmd/app.go

FROM scratch
COPY --from=builder /build/server /app/
WORKDIR /app

ENTRYPOINT ["./server"]
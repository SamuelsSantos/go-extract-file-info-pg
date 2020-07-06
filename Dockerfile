FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o import-data .

FROM scratch
COPY --from=builder /build/import-data /app/
WORKDIR /app

ENTRYPOINT ["./import-data"]
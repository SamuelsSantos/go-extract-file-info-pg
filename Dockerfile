FROM golang:1.14-alpine AS build

WORKDIR /src/
COPY import-data.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/import-data

FROM scratch
COPY --from=build /bin/import-data /bin/import-data
ENTRYPOINT ["/bin/import-data"]
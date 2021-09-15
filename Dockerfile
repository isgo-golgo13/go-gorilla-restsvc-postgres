# stage 1
FROM golang:1.15-alpine as stage

WORKDIR /go-gorilla-restsvc-postgres
COPY go.mod go.sum ./
RUN go mod download
# copy the source from the current directory to the Working Directory inside the container
COPY . .

ENV GO111MODULE=auto
RUN CGO_ENABLED=0 GOOS=linux go build github.com/isgo-golgo13/go-gorilla-restsvc-postgres/cmd/service

# stage 2
FROM golang:1.15-alpine

WORKDIR /root/
COPY --from=stage /go-gorilla-restsvc-postgres .
CMD ["./service"]
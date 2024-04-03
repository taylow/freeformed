FROM golang:1.22.2 AS build
WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOPROXY=direct
RUN go build -v -o bin/freeformed cmd/freeformed/main.go

FROM scratch
COPY --from=build /go/src/app/bin/freeformed /go/bin/freeformed
ENTRYPOINT ["/go/bin/freeformed"]

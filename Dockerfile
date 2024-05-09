FROM golang:1.22.1 AS build

WORKDIR /mock-server-go

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest AS final

WORKDIR /root/

COPY --from=build /mock-server-go/main .

EXPOSE 8080

CMD ["./main"]
